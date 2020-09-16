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

```bash
# local
go run producer.go -exchange myexchange -routingkey myroutingkey -body 'message body'

# docker
docker run --rm -it --network host rabbitmq_go ./producer -exchange myexchange -routingkey myroutingkey -body 'message body'
```

## Run consumer


```
# local
go run consumer.go queue

# docker
docker run --rm -it --network host rabbitmq_go ./consumer queue
```

