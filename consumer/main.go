package main

import "fmt"
import "context"
import kafka "github.com/segmentio/kafka-go"

func main() {
	fmt.Println("This is the consumer")

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:9092"},
		GroupID:   "test-group-1",
		Topic:     "example-topic",
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
	})

	defer r.Close()
	
	for i := 0; i < 5; i++ {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	}
	
	fmt.Println("Messages successfully received")
}
