package producer

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:8080", "topic_stats", 0)
	if err != nil {
		log.Fatal("Cannot connect Kafka:", err)
	}

	conn.SetWriteDeadline(time.Now().Add(time.Second * 10))
}
