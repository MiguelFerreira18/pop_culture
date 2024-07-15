FROM golang:1.22-alpine
WORKDIR /pop_culture

RUN apk add --no-cache gcc musl-dev

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -ldflags '-w -s' -a -o ./bin/popCulture ./cmd/pop_culture

CMD [ "/pop_culture/bin/popCulture" ]
EXPOSE 9090