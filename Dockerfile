# build stage
FROM golang:1.21-alpine3.18 AS build

LABEL maintainer="Javier Telio <jtelio118@gmail.com>"

# set working directory
WORKDIR /app

# copy source code
COPY . .

# install dependencies
RUN go mod download

# build binary
RUN go build -o ./bin/application ./cmd/api/main.go

# final stage
FROM alpine:3.18 AS final

# set working directory
WORKDIR /app

# copy binary
COPY --from=build /app/bin/application ./
COPY config.yaml .

ENV PORT=8080

EXPOSE $PORT

CMD ["./application"]
