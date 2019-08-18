package main

import "fmt"
import "context"
import kafka "github.com/segmentio/kafka-go"

func main() {
	fmt.Println("This is the producer")

	w := configureWriter()
	err := writeMessages(w)
	w.Close()

	if err != nil {
		fmt.Printf("Failed to send messages: %v\n", err)
		return
	}

	fmt.Println("Messages successfully sent")
}

func configureWriter() *kafka.Writer {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:	[]string{"localhost:9092"},
		Topic:		"example-topic",
		Balancer:	&kafka.LeastBytes{},
	})

	return w
}

func writeMessages(w *kafka.Writer) error {
	err := w.WriteMessages(context.Background(),
		kafka.Message{
			Value: []byte("Hello World!"),
		},
		kafka.Message{
			Value: []byte("Another one"),
		},
		kafka.Message{
			Value: []byte("Third time lucky"),
		},
	)

	return err
}