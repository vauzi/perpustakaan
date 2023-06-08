FROM golang:1.18-alpine

WORKDIR /app

COPY . .

RUN go build -o perpus-api

EXPOSE 8080

# Set the command to run the binary executable when the container starts
CMD perpus-api
