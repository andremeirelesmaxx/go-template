FROM golang:1.25-alpine

WORKDIR /app

RUN apk add --no-cache git build-base curl

RUN go install github.com/air-verse/air@latest 
RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s -- -b $(go env GOPATH)/bin v2.5.0

RUN apk del --no-cache curl

# CMD ["tail", "-f", "/dev/null"]
CMD ["air", "-c", ".air.toml"]