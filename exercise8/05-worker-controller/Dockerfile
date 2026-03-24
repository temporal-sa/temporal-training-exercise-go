FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY versionA/ ./versionA/
# COPY versionB/ ./versionB/

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./versionA/worker
# RUN CGO_ENABLED=0 GOOS=linux go build -o app ./versionB/worker

FROM gcr.io/distroless/static-debian12

COPY --from=builder /app/app /app

# Run
CMD ["/app"]
