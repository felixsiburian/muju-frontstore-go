package Store

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/labstack/echo"
	"log"
	"muju-frontstore-go/domain/model"
	"muju-frontstore-go/kafka/Host/Config"
	"net/http"
)

func PublishUpdateStore(c echo.Context) error {
	data := new(model.Store)
	err := json.NewDecoder(c.Request().Body).Decode(&data)
	//err := c.Bind(data)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
		panic(err)
	}
	fmt.Println(data)

	kafkaConfig := Config.GetKafkaConfig("", "")

	producer, err := sarama.NewSyncProducer([]string{"52.185.161.109:9092"}, kafkaConfig)

	if err != nil {

		panic(err)

	}

	defer func() {
		if err := producer.Close(); err != nil {
			panic(err)
		}
	}()

	var newTopic = "updates-store-topic"

	message, _ := json.Marshal(data)
	var data2 interface{}
	err = json.Unmarshal(message, &data2)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println("Data : ", data2)
	//message := `{
	//	"merchant_name":`+data.MerchantName+`,
	//	"merchant_email":"contact@jco.com"
	//}`
	msg := &sarama.ProducerMessage{
		Topic: newTopic,
		Value: sarama.StringEncoder(string(message)),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	fmt.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", newTopic, partition, offset)
	return c.JSON(http.StatusOK, &data2)
}
