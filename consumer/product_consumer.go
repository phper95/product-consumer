package consumer

import (
	"context"
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
	mq.KafkaStdLogger.Printf("partion : %d ; offset : %d ; msg %s",
		msg.Partition, msg.Offset, string(msg.Value))
	product := ProductMsg{}
	err := json.Unmarshal(msg.Value, &product)
	if err != nil {
		global.LOG.Error("Unmarshal", err, string(msg.Value))
		//格式异常的数据，回到队列也不会解析成功
		return true, nil
	}
	mq.KafkaStdLogger.Printf("product :%+v", product)
	productIndex := product.ProductIndex
	esClient := es.GetClient(es.DefaultClient)
	switch product.Operation {
	case global.OperationCreate, global.OperationOnSale:
		mq.KafkaStdLogger.Printf("OperationOnSale product :%+v", product)
		esClient.BulkCreate(global.IndexName, strconv.FormatInt(productIndex.Id, 10),
			strconv.Itoa(productIndex.CateId), productIndex)
	case global.OperationUpdate:
		if product.IsShow == 0 {
			err = esClient.DeleteRefresh(context.Background(), global.IndexName, strconv.FormatInt(productIndex.Id, 10), strconv.Itoa(productIndex.CateId))
			if err != nil {
				global.LOG.Errorf("DeleteRefresh error %v ;productIndex : %+v", err, productIndex)
			}
		} else {
			esClient.BulkCreate(global.IndexName, strconv.FormatInt(productIndex.Id, 10),
				strconv.Itoa(productIndex.CateId), productIndex)
		}
	case global.OperationUnSale:
		mq.KafkaStdLogger.Printf("OperationUnSale product :%+v", product)
		err = esClient.DeleteRefresh(context.Background(), global.IndexName, strconv.FormatInt(productIndex.Id, 10), strconv.Itoa(productIndex.CateId))
		if err != nil {
			global.LOG.Errorf("DeleteRefresh error %v ;productIndex : %+v", err, productIndex)
		}
	}
	return true, nil
}

func CloseProductConsumer() {
	if productConsumer != nil {
		productConsumer.Close()
	}
}
