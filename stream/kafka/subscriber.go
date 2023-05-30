package kafka

//
// import (
// 	"context"
//
// 	confluentkafka "github.com/confluentinc/confluent-kafka-go/v2/kafka"
//
// 	"codeup.aliyun.com/qimao/leo/stream"
// )
//
// var _ stream.Subscriber = new(Subscriber)
//
// type Subscriber struct {
// 	topic          string
// 	consumer       *confluentkafka.Consumer
// 	chanBufferSize int
// 	rebalanceCb    confluentkafka.RebalanceCb
// 	messageDecoder func(ctx context.Context, kafkaMsg *confluentkafka.Message) (stream.Message, error)
// 	onNack         func(message stream.Message)
// }
//
// func (sub *Subscriber) Topic() string {
// 	return sub.topic
// }
//
// func (sub *Subscriber) Subscribe(ctx context.Context) (<-chan stream.Message, <-chan error) {
// 	msgC := make(chan stream.Message, sub.chanBufferSize)
// 	errC := make(chan error, sub.chanBufferSize)
// 	go sub.consumeMsg(ctx, msgC, errC)
// 	return msgC, errC
// }
//
// func (sub *Subscriber) Close(_ context.Context) error {
// 	return sub.consumer.Close()
// }
//
// func (sub *Subscriber) consumeMsg(ctx context.Context, msgC chan stream.Message, errC chan error) {
// 	defer close(msgC)
// 	defer close(errC)
// 	if err := sub.consumer.Subscribe(sub.topic, sub.rebalanceCb); err != nil {
// 		errC <- err
// 		return
// 	}
// 	defer func() {
// 		if err := sub.consumer.Unsubscribe(); err != nil {
// 			errC <- err
// 			return
// 		}
// 	}()
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			return
// 		default:
// 			kafkaMsg, err := sub.consumer.ReadMessage(-1)
// 			if err != nil {
// 				errC <- err
// 			} else {
// 				message, err := sub.messageDecoder(ctx, kafkaMsg)
// 				if err != nil {
// 					errC <- err
// 					continue
// 				}
// 				sub.handleMsg(ctx, message, kafkaMsg, msgC)
// 			}
//
// 		}
// 	}
// }
//
// func (sub *Subscriber) handleMsg(
// 	ctx context.Context,
// 	message stream.Message,
// 	kafkaMsg *confluentkafka.Message,
// 	msgC chan stream.Message,
// ) {
// 	commitFunc := func() (any, error) {
// 		return sub.consumer.CommitMessage(kafkaMsg)
// 	}
// 	ackedC := message.Acked(commitFunc)
// 	nackedC := message.Nacked(commitFunc)
// 	msgC <- message
// 	select {
// 	case <-ctx.Done():
// 		return
// 	case <-ackedC:
// 		return
// 	case <-nackedC:
// 		if sub.onNack != nil {
// 			sub.onNack(message)
// 		}
// 		return
// 	}
// }
//
// type SubscriberBuilder struct {
// 	topic          string
// 	consumer       *confluentkafka.Consumer
// 	chanBufferSize int
// 	rebalanceCb    confluentkafka.RebalanceCb
// 	messageDecoder func(ctx context.Context, kafkaMsg *confluentkafka.Message) (stream.Message, error)
// 	onNack         func(message stream.Message)
// }
//
// func (builder *SubscriberBuilder) SetTopic(topic string) *SubscriberBuilder {
// 	builder.topic = topic
// 	return builder
// }
//
// func (builder *SubscriberBuilder) SetConsumer(consumer *confluentkafka.Consumer) *SubscriberBuilder {
// 	builder.consumer = consumer
// 	return builder
// }
//
// func (builder *SubscriberBuilder) SetChanBufferSize(chanBufferSize int) *SubscriberBuilder {
// 	builder.chanBufferSize = chanBufferSize
// 	return builder
// }
//
// func (builder *SubscriberBuilder) SetRebalanceCb(rebalanceCb confluentkafka.RebalanceCb) *SubscriberBuilder {
// 	builder.rebalanceCb = rebalanceCb
// 	return builder
// }
//
// func (builder *SubscriberBuilder) SetMessageDecoder(messageDecoder func(ctx context.Context, kafkaMsg *confluentkafka.Message) (stream.Message, error)) *SubscriberBuilder {
// 	builder.messageDecoder = messageDecoder
// 	return builder
// }
//
// func (builder *SubscriberBuilder) SetOnNack(onNack func(message stream.Message)) *SubscriberBuilder {
// 	builder.onNack = onNack
// 	return builder
// }
//
// func (builder *SubscriberBuilder) Build() *Subscriber {
// 	if builder.messageDecoder == nil {
// 		builder.messageDecoder = DefaultMessageDecoder
// 	}
// 	return &Subscriber{
// 		topic:          builder.topic,
// 		consumer:       builder.consumer,
// 		chanBufferSize: builder.chanBufferSize,
// 		rebalanceCb:    builder.rebalanceCb,
// 		messageDecoder: builder.messageDecoder,
// 		onNack:         builder.onNack,
// 	}
// }
//
// func NewSubscriberBuilder() *SubscriberBuilder {
// 	return &SubscriberBuilder{}
// }
