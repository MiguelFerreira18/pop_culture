package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"pop_culture/api/router"
	"pop_culture/cmd/migration"
	c "pop_culture/config"
	"pop_culture/logger"
	"strconv"
	"syscall"

	"github.com/go-chi/jwtauth/v5"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const fmtdbString = "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=%s"

func main() {

	config := c.New()

	logger := logger.Logger(config.Server.Debug)

	privateKey, err := c.LoadPrivateKey(config.KeyPaths.PrivateKeyPath)
	if err != nil {
		logger.Error().Err(err).Msg("")
		return
	}
	publicKey, err := c.LoadPublicKey(config.KeyPaths.PublicKeyPath)
	if err != nil {
		logger.Error().Err(err).Msg("")
		return
	}

	tokenAuth := jwtauth.New(jwa.RS256.String(), privateKey, publicKey)

	port := strconv.Itoa(int(config.Database.UrlPort))

	dbstring := fmt.Sprintf(fmtdbString, config.Database.User, config.Database.Password, config.Database.UrlAdress, port, config.Database.DatabaseName, config.Database.Locality)
	db, err := gorm.Open(mysql.Open(dbstring), &gorm.Config{})
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to connect to Database")
		return
	}
	migration.Migrate(db)

	//CREATE ROUTER
	r := router.New(logger, db, tokenAuth)
	server := &http.Server{
		Addr:         ":" + config.Server.Port,
		Handler:      r,
		ReadTimeout:  config.Server.ReadTimeout,
		WriteTimeout: config.Server.WriteTimeout,
		IdleTimeout:  config.Server.IdleTimeout,
	}

	closed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)
		<-sigint

		logger.Info().Msgf("Shutting down server %v", server.Addr)

		ctx, cancel := context.WithTimeout(context.Background(), config.Server.IdleTimeout)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			logger.Error().Err(err).Msg("Server shutdown failure")
		}

		sqlDB, err := db.DB()
		if err == nil {
			if err = sqlDB.Close(); err != nil {
				logger.Error().Err(err).Msg("DB connection closing failure")
			}
		}

		close(closed)

	}()

	// http.ListenAndServe(":"+config.Server.Port, r)
	logger.Info().Msgf("Starting server %v", server.Addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatal().Err(err).Msg("Server startup failure")
	}

	<-closed
	logger.Info().Msgf("Server shutdown successfully")

}
