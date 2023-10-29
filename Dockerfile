FROM golang:1.19

WORKDIR /app

COPY . .
RUN go mod download

RUN go build -o /user-service

EXPOSE 8080

CMD [ "/user-service" ]