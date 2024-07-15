echo "building the Server"
go build -o ./bin/popCulture ./cmd/pop_culture/main.go 

echo "Running the Server"
./bin/popCulture