FROM golang:latest AS build

WORKDIR /app

COPY go.mod ./

RUN go mod tidy

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o auth ./cmd/auth

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/auth .

ENTRYPOINT [ "./auth" ]