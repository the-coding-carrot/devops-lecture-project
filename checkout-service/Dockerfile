FROM golang:1.25-alpine AS builder

# Praktisch mkdir und cd in einem Schritt
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o server ./cmd/main.go

FROM alpine AS final 
WORKDIR /app
COPY --from=builder /app/server .

EXPOSE 8080
CMD [ "./server" ]