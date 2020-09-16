FROM golang:alpine as build
WORKDIR /go/src/
RUN apk update && apk --no-cache add \
    git
COPY producer/producer.go producer/producer.go
COPY consumer/consumer.go consumer/consumer.go
RUN go get "github.com/streadway/amqp"
RUN cd producer && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ../producer producer.go
RUN cd consumer && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ../consumer consumer.go
RUN chmod +x producer consumer

FROM scratch as run
COPY --from=build /go/src/producer .
COPY --from=build /go/src/consumer .
