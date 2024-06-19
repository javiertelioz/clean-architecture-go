# build stage
FROM golang:1.22-alpine3.19 AS build

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
FROM alpine:3 AS final

# set working directory
WORKDIR /app

# copy binary
COPY --from=build /app/bin/application ./
COPY example.config.yaml config.yaml

ENV PORT=8080

EXPOSE $PORT

CMD ["./application"]
