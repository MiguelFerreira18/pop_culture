package middleware

import (
	"context"
	"net/http"

	"github.com/rs/xid"
)

const requestIDHeaderKey = "X-Request-ID"
const keyRequestId = "requestID"

func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		requestId := r.Header.Get(requestIDHeaderKey)
		if requestId == "" {
			requestId = xid.New().String()
		}
		ctx = context.WithValue(ctx, keyRequestId, requestId)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
