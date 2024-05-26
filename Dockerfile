FROM golang:1.21.6-bullseye

WORKDIR /app

COPY . ./

ENV SERVER_PORT=8889
ENV DB_HOST=localhost
ENV PG_USERNAME=postgres
ENV PG_PASSWORD=qwerty
ENV PG_PORT=5432
ENV PG_NAME=postgres
ENV PG_SSLMODE=disable

RUN go mod download

RUN go build -o ./bin/postsGraphQL ./server.go

EXPOSE 8889

ENTRYPOINT ["./bin/postsGraphQL"]