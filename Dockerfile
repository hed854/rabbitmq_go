FROM golang:alpine as build
WORKDIR /go/src/
RUN apk update && apk --no-cache add \
    git
COPY producer.go .
COPY consumer.go .
RUN go get "github.com/streadway/amqp"
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o producer producer.go 
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o consumer consumer.go
RUN chmod +x producer consumer

FROM scratch as run
COPY --from=build /go/src/producer .
COPY --from=build /go/src/consumer .
