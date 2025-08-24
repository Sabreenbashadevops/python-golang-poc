# --- builder ---
FROM mirror.gcr.io/library/golang:1.22-alpine AS builder
RUN apk add --no-cache ca-certificates git && update-ca-certificates
WORKDIR /src
COPY go.mod ./
RUN go mod download
COPY . .
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags="-s -w" -o /out/server ./cmd/server

# --- runtime ---
FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /out/server /server
USER nonroot:nonroot
EXPOSE 8080
ENTRYPOINT ["/server"]
