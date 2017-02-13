package kong

import (
	"service-cloud/utils/kong/models"
	"testing"
)

func TestAddConsumer(t *testing.T) {
	//test add consumer
	var consumer models.Consumer
	consumer.Username = "testconsumer"

	retConsumer, err := AddConsumer(&consumer)
	if err != nil {
		t.Error("AddConsumer Error, ", err)
	}
	if retConsumer.Username != consumer.Username {
		t.Errorf("AddConsumer error: %s.", retConsumer.Username)
		t.Error("AddConsumer Error, ", *retConsumer)
	}
	// test ListConsumers
	consumerlist, err := ListConsumers(5, "")
	if err != nil {
		t.Error("ListConsumers Error, ", err)
	}
	if consumerlist.Total < 1 {
		t.Error("ListConsumers Error, ", *consumerlist)
	}
	//test GetConsumer
	retConsumer, err = GetConsumer(consumer.Username)
	if err != nil {
		t.Error("GetConsumer Error, ", err)
	}
	if retConsumer.Username != consumer.Username {
		t.Error("GetConsumer Error, ", *retConsumer)
	}
	//test UpdateConsumer
	consumer.Username = "testconsumer1"
	retConsumer, err = UpdateConsumer(retConsumer.Username, &consumer)
	if err != nil {
		t.Error("UpdateConsumer Error, ", err)
	}
	if retConsumer.Username != "testconsumer1" {
		t.Error("UpdateConsumer Error, ", *retConsumer)
	}

	//test DeleteConsumer
	err = DeleteConsumer(retConsumer.ID)
	if err != nil {
		t.Error("DeleteConsumer Error, ", err)
	}
	retConsumer, err = GetConsumer(retConsumer.ID)
	if err == nil {
		t.Error("DeleteConsumer Error. process failure!")
	}
}
