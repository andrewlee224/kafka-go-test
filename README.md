# kafka-go-test
Basic set up for running Kafka in an exposed container and producing/consuming messages using Go

1. Run Kafka in the container
```
$ docker-compose up -d
```

2. Build the producer and run it
```
$ go build producer/main.go
$ ./producer/producer
```

3. Build the consumer and run (you might have to Ctrl+C afterwards to stop it as it will wait for more messages)
```
$ go build consumer/main.go
$ ./consumer/consumer
```
