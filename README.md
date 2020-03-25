# Rabbitmq Golang Producer/Consumer

Test your Rabbitmq things with this

Based on: 

* https://github.com/rabbitmq/rabbitmq-tutorials/tree/master/go
* https://www.rabbitmq.com/tutorials/tutorial-one-go.html

## Install

```
# local
go get "github.com/streadway/amqp"

# docker
docker build -t rabbitmq_go .
```

## Run producer 

Note: all parameters are mandatory for now, even the routing key

```bash
# local
go run producer.go "hello worlllld" exchange routing_key

# docker
docker run --rm -it --network host rabbitmq_go ./producer message exchange routingkey
```

## Run consumer


```
# local
go run consumer.go queue

# docker
docker run --rm -it --network host rabbitmq_go ./consumer queue
```

