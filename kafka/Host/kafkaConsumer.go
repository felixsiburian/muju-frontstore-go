package Host

import (
	"fmt"
	"muju-frontstore-go/kafka/Host/Consumer"
)

func StartkafkaStore() {
	fmt.Println("Kafka Catalog is Starting")
	go Consumer.NewStoresConsumer()
	go Consumer.NewPackageConsumer()
	go Consumer.NewTemplateConsumer()
	go Consumer.NewTransactionConsumer()
}

