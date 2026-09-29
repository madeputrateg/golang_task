# ==========================
# Build stage
# ==========================

FROM golang:1.27.1-alpine AS builder

WORKDIR /app


# Copy dependency files first
COPY go.mod go.sum ./


# Download dependencies
RUN go mod download


# Copy source code
COPY . .


# Ensure dependencies are clean
RUN go mod tidy


# Build webservice
RUN go build -o webservice ./cmd/webservice



# ==========================
# Runtime stage
# ==========================

FROM alpine:latest


WORKDIR /app


COPY --from=builder /app/webservice .


# If you have templates/static files
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/static ./static


EXPOSE 8080


CMD ["./webservice"]