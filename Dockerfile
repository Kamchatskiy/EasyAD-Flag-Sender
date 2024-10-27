FROM golang AS build
WORKDIR /app
COPY ./src/go.mod .
COPY ./src/main.go .
RUN go build -o app .
EXPOSE 80
CMD ["./app"]