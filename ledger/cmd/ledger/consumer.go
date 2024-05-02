package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/IBM/sarama"
	"github.com/mmcferren/go-micro/internal/ledger"
	_ "github.com/go-sql-driver/mysql"
)

const (
	dbDriver = "mysql"
	dbName = "ledger"
	topic = "ledger"
)

var (
	db *sql.DB
	wg sync.WaitGroup
)

type LedgerMsg struct {
	OrderID string `json:"order_id"`
	UserID string `json:"user_id"`
	Amount int64`json:"amount"`
	Operation string `json:"operation"`
	Date string `json:"date"`
}

func main() {
	var err error

	dbUser := os.Getenv("MYSQL_USERNAME")
	dbPassword := os.Getenv("MYSQL_PASSWORD")

	// Open a database connection
	dsn := fmt.Sprintf("%s:%s@tcp(mysql-ledger:3306)/%s", dbUser, dbPassword, dbName)
	
	db, err = sql.Open(dbDriver, dsn)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = db.Close(); err != nil {
			log.Printf("error closing db: %s", err)
		}
	}()

	// Ping the database to ensure connection is valid
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan struct{})
	
	sarama.Logger = log.New(os.Stdout, "[sarama]", log.LstdFlags)

	consumer, err := sarama.NewConsumer([]string{"kafka-service:9092"}, sarama.NewConfig())
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		close(done)
		if err := consumer.Close(); err != nil {
			log.Println(err)
		}
	}()

	partitions, err := consumer.Partitions(topic)
	if err != nil {
		log.Fatal(err)
	}

	for _, partition := range partitions {
		partitionConsumer, err := consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)
		if err != nil {
			log.Fatal(err)
		}

		defer func() {
			if err := partitionConsumer.Close(); err != nil {
				log.Println(err)
			}
		}()

		wg.Add(1)
		go awaitMessages(partitionConsumer, partition, done)
	}

	wg.Wait()
}

func awaitMessages(partitionConsumer sarama.PartitionConsumer, partition int32, done chan struct{}) {
	defer wg.Done()

	for {
		select {
		case msg := <-partitionConsumer.Messages():
			fmt.Printf("Partition %d - Received message: %s\n", partition, string(msg.Value))
			handleMessage(msg)
		case <-done:
			fmt.Printf("Received done signal. Exiting...\n")
			return
		}
	}
}

func handleMessage(msg *sarama.ConsumerMessage) {
	var ledgerMsg LedgerMsg

	err := json.Unmarshal(msg.Value, &ledgerMsg)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = ledger.Insert(db, ledgerMsg.OrderID, ledgerMsg.UserID, ledgerMsg.Amount, ledgerMsg.Operation, ledgerMsg.Date)
	if err != nil {
		fmt.Println(err)
		return
	}
}