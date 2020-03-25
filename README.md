# Rabbitmq Golang Producer/Consumer

Test your Rabbitmq things with this

Based on: 

* https://github.com/rabbitmq/rabbitmq-tutorials/tree/master/go
* https://www.rabbitmq.com/tutorials/tutorial-one-go.html

## Install

```
go get "github.com/streadway/amqp"
```

## Run producer 

Note: all parameters are mandatory for now, ie this works only for topics :D

```
go run producer.go "hello worlllld" exchange routing_key
```

## Run consumer

```
go run consumer.go queue
```

