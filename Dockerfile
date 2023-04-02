# Use an official Golang runtime as a parent image
FROM golang:1.20-alpine3.16 AS build

# Set the working directory to /app
WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download
COPY . .

RUN CGO_ENABLED=0 go build -o microservice


FROM alpine:3.16.2
WORKDIR /bin

# Declare a build-time variable to hold the server port
ARG SERVER_PORT

RUN mkdir -p /usr/local/bin

COPY --from=build /app/microservice /bin/microservice
COPY --from=build /app/pkg/config/config.yml /bin/config.yml
COPY --from=build /app/docs/swagger.json /bin/doc.json

ENV GIN_MODE=release

# Set the environment variable for the microservice port
ENV PORT=${SERVER_PORT:-8080}

# Expose the microservice port
EXPOSE $PORT

# Run the microservice binary
CMD /bin/microservice
