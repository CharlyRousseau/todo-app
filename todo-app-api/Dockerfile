FROM golang:1.22.2-alpine AS build

WORKDIR /app

COPY /todo-app-api/go.mod /todo-app-api/go.sum ./

RUN go mod download

COPY /todo-app-api/. .

RUN go build -o main .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=build /app/main .

EXPOSE 8080

CMD ["./main"]
