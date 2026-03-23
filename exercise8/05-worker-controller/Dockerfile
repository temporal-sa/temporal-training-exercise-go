FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY versionA/ ./versionA/

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./versionA/worker

FROM gcr.io/distroless/static-debian12

COPY --from=builder /app/app /app

# Run
CMD ["/app"]
