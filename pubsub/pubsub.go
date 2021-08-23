package pubsub

type IMessage interface {
	GetTopic() string
	SetTopic(topic string) error
	GetData() interface{}
}

type IPubSub interface {
	Publish(topic string, data interface{}) error
	Subscribe(topic string) error
}
