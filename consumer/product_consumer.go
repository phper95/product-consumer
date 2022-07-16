package consumer

import (
	"encoding/json"
	"gitee.com/phper95/pkg/es"
	"gitee.com/phper95/pkg/mq"
	"github.com/Shopify/sarama"
	"product-consumer/global"
	"strconv"
)

var productConsumer *mq.Consumer

func StartConsumer() {
	var err error
	productConsumer, err = mq.StartKafkaConsumer(global.CONFIG.Kafka.Hosts,
		[]string{global.Topic}, "product-consumer",
		nil, MsgHandler)
	if err != nil {
		panic(err)
	}
}

func MsgHandler(msg *sarama.ConsumerMessage) (bool, error) {
	mq.KafkaStdLogger.Println("partion : %d ; offset : %d ; msg %s",
		msg.Partition, msg.Offset, string(msg.Value))
	product := ProductMsg{}
	err := json.Unmarshal(msg.Value, &product)
	if err != nil {
		global.LOG.Error("Unmarshal", err, string(msg.Value))
		return true, nil
	}
	mq.KafkaStdLogger.Printf("product :%+v", product)
	productIndex := product.ProductIndex
	switch product.Operation {
	case global.OperationCreate, global.OperationUpdate:
	}
	es.GetClient(es.DefaultClient).BulkCreate(global.IndexName, strconv.FormatInt(productIndex.Id, 10),
		strconv.Itoa(productIndex.CateId), productIndex)

	return true, nil
}

func CloseProductConsumer() {
	if productConsumer != nil {
		productConsumer.Close()
	}
}
