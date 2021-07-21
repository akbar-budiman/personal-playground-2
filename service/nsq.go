package service

import (
	"encoding/json"
	"fmt"

	"github.com/akbar-budiman/personal-playground-2/entity"
	"github.com/nsqio/go-nsq"
)

var (
	nsqAddress        = "127.0.0.1:4150"
	nsqlookupdAddress = "127.0.0.1:4161"
	addUserTopic      = "NewUser"
)

func ProduceNewUserEvent(userData []byte) {
	config := nsq.NewConfig()
	p, err := nsq.NewProducer(nsqAddress, config)
	if err != nil {
		panic(err)
	}

	err = p.Publish(addUserTopic, userData)
	if err != nil {
		panic(err)
	}

	p.Stop()
}

type NewUserConsumer struct{}

func (h *NewUserConsumer) HandleMessage(m *nsq.Message) error {
	var newObj entity.User
	json.Unmarshal(m.Body, &newObj)

	AddOrReplaceUser(&newObj)
	return nil
}

func RegisterConsumer() {
	fmt.Println("Registering consumer")
	config := nsq.NewConfig()

	newUserConsumer, err := nsq.NewConsumer(addUserTopic, "channel1", config)
	if err != nil {
		panic(err)
	}
	newUserConsumer.AddHandler(&NewUserConsumer{})

	err = newUserConsumer.ConnectToNSQLookupd(nsqlookupdAddress)
	if err != nil {
		panic(err)
	}

	fmt.Println("consumer registered.")
}
