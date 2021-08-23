package localpubsub

import (
	"sync"
)

type Message struct {
	Topic string
	Data  interface{}
}

type localPubSub struct {
	messageQueue chan Message
	mapTopic     map[string][]chan Message
	locker       *sync.Mutex
}

func NewMessage(data interface{}) *Message {
	msg := Message{
		Data: data,
	}
	return &msg
}

func NewPubSub() *localPubSub {
	pb := localPubSub{
		messageQueue: make(chan Message, 1000),
		mapTopic:     make(map[string][]chan Message),
		locker:       new(sync.Mutex),
	}
	return &pb
}

func (l *localPubSub) Publish(topic string, data interface{}) error {
	msg := Message{
		Topic: topic,
		Data:  data,
	}
	l.messageQueue <- msg
	return nil
}

func (l *localPubSub) Subscribe(topic string) chan Message {
	c := make(chan Message)

	l.locker.Lock()
	if _, exists := l.mapTopic[topic]; exists {
		l.mapTopic[topic] = append(l.mapTopic[topic], c)
	} else {
		l.mapTopic[topic] = []chan Message{c}
	}

	l.locker.Unlock()
	return c
}

func (l *localPubSub) Run() {
	go func() {
		for {
			msg := <-l.messageQueue
			topic := msg.Topic
			if _, exist := l.mapTopic[topic]; exist {
				for _, chanSub := range l.mapTopic[topic] {
					chanSub <- msg
				}
			}
		}
	}()
}

func (m *Message) GetTopic() string {
	return m.Topic
}

func (m *Message) SetTopic(topic string) error {
	m.Topic = topic
	return nil
}

func (m *Message) GetData() interface{} {
	return m.Data
}
