package Consumer

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/joho/godotenv"
	"log"
	"muju-frontstore-go/domain/model"
	"muju-frontstore-go/domain/repository"
	"muju-frontstore-go/kafka/Host/Config"
	"os"
	"os/signal"
	"strings"
)

func consumeStore(topics []string, master sarama.Consumer) (chan *sarama.ConsumerMessage, chan *sarama.ConsumerError) {
	consumers := make(chan *sarama.ConsumerMessage)
	errors := make(chan *sarama.ConsumerError)
	fmt.Println("Kafka Store is Ready")
	//	//fmt.Println(topics)
	for _, topic := range topics {
		if strings.Contains(topic, "__consumer_offsets") {
			continue
		}
		partitions, _ := master.Partitions(topic)
		// this only consumes partition no 1, you would probably want to consume all partitions
		consumer, err := master.ConsumePartition(topic, partitions[0], sarama.OffsetNewest)
		if nil != err {
			fmt.Println("error card : ", err.Error())
			fmt.Printf("Topic %v Partitions: %v", topic, partitions)
			panic(err)
		}
		//fmt.Println(" Start consuming topic ", topic)
		go func(topic string, consumer sarama.PartitionConsumer) {
			for {
				select {
				case consumerError := <-consumer.Errors():
					errors <- consumerError
					fmt.Println("consumerError: ", consumerError.Err)

				case msg := <-consumer.Messages():
					//*messageCountStart++
					//Deserialize
					store := model.Store{}
					switch msg.Topic {
					case "creates-store-topic":
						err := json.Unmarshal([]byte(msg.Value), &store)
						if err != nil {
							fmt.Println("Error : ",err.Error())
							os.Exit(1)
						}
						err = repository.CreateStores(&store)
						if err != nil {
							log.Fatal(err)
						}
						fmt.Println(string(msg.Value))
						fmt.Println("Store berhasil dibuat")
					case "updates-store-topic":
						fmt.Println("masuk ke topic Update")
						err := json.Unmarshal([]byte(msg.Value), &store)
						if err != nil {
							fmt.Println("Error : ", err.Error())
							log.Fatal(err)
						}
						err = repository.UpdateStores(&store)
						if err != nil {
							log.Fatal(err)
						}
						fmt.Println(string(msg.Value))
						fmt.Println("Store berhasil di update")
					case "deletes-store-topic":
						fmt.Println("masuk ke topic Delete")
						err := json.Unmarshal([]byte(msg.Value), &store)
						if err != nil {
							fmt.Println("Error Delete : ", err.Error())
							log.Fatal(err)
						}
						err = repository.DeleteStores(&store)
						if err != nil {
							log.Fatal(err)
						}
						fmt.Println("Store berhasil dihapus")
					}
				}
			}
		}(topic, consumer)
	}

	return consumers, errors
}

func NewStoresConsumer() {
	godotenv.Load(".env")
	brokers := []string{os.Getenv("KAFKA_IP")}

	kafkaConfig := Config.GetKafkaConfig("", "")

	master, err := sarama.NewConsumer(brokers, kafkaConfig)

	if err != nil {

		panic(err)

	}

	defer func() {

		if err := master.Close(); err != nil {

			panic(err)

		}

	}()

	//topic, err := master.Topics()
	if err != nil {
		panic(err)
	}
	topics, _ := master.Topics()
	//
	consumer, errors := consumeStore(topics, master)
	////consumer1, err := master.ConsumePartition(updateTopic, 0, sarama.OffsetNewest)
	//
	if errors != nil {
		fmt.Println(err)
		//panic(err)

	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// Count how many message processed
	msgCount := 0

	// Get signnal for finish
	doneCh := make(chan struct{})
	go func() {
		for {
			select {
			case msg := <-consumer:
				msgCount++
				fmt.Println("Received messages", string(msg.Key), string(msg.Value))
			case consumerError := <-errors:
				msgCount++
				fmt.Println("Received consumerError ", string(consumerError.Topic), string(consumerError.Partition), consumerError.Err)
				doneCh <- struct{}{}
			case <-signals:
				fmt.Println("Interrupt is detected")
				doneCh <- struct{}{}
			}
		}
	}()
	<-doneCh
	master.Close()
	fmt.Println("Processed", msgCount, "messages")

}
