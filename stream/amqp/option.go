package amqp

import (
	"codeup.aliyun.com/qimao/leo/leo/log"
	"codeup.aliyun.com/qimao/leo/leo/stream"
	"github.com/rabbitmq/amqp091-go"
)

type options struct {
	Logger      log.Logger
	Marshaller  Marshaller
	NackHandler func(msg *stream.Message)

	ExchangeName func(topic string) string
	QueueName    func(topic string) string
	RoutingKeys  func(topic string) []string

	ExchangeOption  *ExchangeOption
	QueueOption     *QueueOption
	QueueBindOption *QueueBindOption
	ConfirmOption   *ConfirmOption
	QosOption       *QosOption

	PublishOption PublishOption
	ConsumeOption ConsumeOption

	PublishingDecorator func(*stream.Message, amqp091.Publishing) amqp091.Publishing

	AckOption AckOption
}

type Option func(o *options)

func (o *options) apply(opts ...Option) {
	for _, opt := range opts {
		opt(o)
	}
	if o.ExchangeName == nil {
		o.ExchangeName = func(topic string) string {
			return ""
		}
	}
	if o.QueueName == nil {
		o.QueueName = func(topic string) string {
			return ""
		}
	}
	if o.RoutingKeys == nil {
		o.RoutingKeys = func(topic string) []string {
			return []string{""}
		}
	}
}

func (o *options) init(topic string, channel *amqp091.Channel) error {
	if err := o.ExchangeOption.declare(o.ExchangeName(topic), channel); err != nil {
		return err
	}
	if err := o.QueueOption.declare(o.QueueName(topic), channel); err != nil {
		return err
	}
	if err := o.QueueBindOption.bind(o.ExchangeName(topic), o.QueueName(topic), o.RoutingKeys(topic), channel); err != nil {
		return err
	}
	if err := o.QosOption.qos(channel); err != nil {
		return err
	}
	if err := o.ConfirmOption.confirm(channel); err != nil {
		return err
	}
	if o.Marshaller == nil {
		o.Marshaller = DefaultMarshaller{}
	}
	return nil
}

func Logger(l log.Logger) Option {
	return func(o *options) {
		o.Logger = l
	}
}

func MessageMarshaller(m Marshaller) Option {
	return func(o *options) {
		o.Marshaller = m
	}
}

func NackHandler(h func(msg *stream.Message)) Option {
	return func(o *options) {
		o.NackHandler = h
	}
}

func ExchangeName(f func(topic string) string) Option {
	return func(o *options) {
		o.ExchangeName = f
	}
}

func QueueName(f func(topic string) string) Option {
	return func(o *options) {
		o.QueueName = f
	}
}

func RoutingKeys(f func(topic string) []string) Option {
	return func(o *options) {
		o.RoutingKeys = f
	}
}

func ExchangeOptions(eo *ExchangeOption) Option {
	return func(o *options) {
		o.ExchangeOption = eo
	}
}

func QueueOptions(qo *QueueOption) Option {
	return func(o *options) {
		o.QueueOption = qo
	}
}

func QueueBindOptions(qbo *QueueBindOption) Option {
	return func(o *options) {
		o.QueueBindOption = qbo
	}
}

func QosOptions(qo *QosOption) Option {
	return func(o *options) {
		o.QosOption = qo
	}
}

func PublishOptions(po PublishOption) Option {
	return func(o *options) {
		o.PublishOption = po
	}
}

func ConsumeOptions(co ConsumeOption) Option {
	return func(o *options) {
		o.ConsumeOption = co
	}
}

func ConfirmOptions(co *ConfirmOption) Option {
	return func(o *options) {
		o.ConfirmOption = co
	}
}

func AckOptions(ao AckOption) Option {
	return func(o *options) {
		o.AckOption = ao
	}
}

type ExchangeOption struct {
	Kind       string
	Durable    bool
	AutoDelete bool
	Internal   bool
	NoWait     bool
	Args       amqp091.Table
}

func (o *ExchangeOption) declare(exchangeName string, channel *amqp091.Channel) error {
	if o == nil {
		return nil
	}
	return channel.ExchangeDeclare(exchangeName, o.Kind, o.Durable, o.AutoDelete, o.Internal, o.NoWait, o.Args)
}

type QueueOption struct {
	Durable    bool
	AutoDelete bool
	Exclusive  bool
	NoWait     bool
	Args       amqp091.Table
}

func (o *QueueOption) declare(queueName string, channel *amqp091.Channel) error {
	if o == nil {
		return nil
	}
	_, err := channel.QueueDeclare(queueName, o.Durable, o.AutoDelete, o.Exclusive, o.NoWait, o.Args)
	return err
}

type QueueBindOption struct {
	NoWait bool
	Args   amqp091.Table
}

func (o *QueueBindOption) bind(exchangeName string, queueName string, routingKeys []string, channel *amqp091.Channel) error {
	if o == nil {
		return nil
	}
	for _, routingKey := range routingKeys {
		if err := channel.QueueBind(queueName, routingKey, exchangeName, o.NoWait, o.Args); err != nil {
			return err
		}
	}
	return nil
}

type ConfirmOption struct {
	NoWait bool
}

func (o *ConfirmOption) confirm(channel *amqp091.Channel) error {
	if o == nil {
		return nil
	}
	return channel.Confirm(o.NoWait)
}

type PublishOption struct {
	Mandatory    bool
	Immediate    bool
	DeliveryMode uint8
}

type QosOption struct {
	PrefetchSize  int
	PrefetchCount int
	Global        bool
}

func (o *QosOption) qos(channel *amqp091.Channel) error {
	if o == nil {
		return nil
	}
	return channel.Qos(o.PrefetchSize, o.PrefetchCount, o.Global)
}

type ConsumeOption struct {
	Tag       string
	AutoAck   bool
	Exclusive bool
	NoLocal   bool
	NoWait    bool
	Table     amqp091.Table
}

type AckOption struct {
	Multiple bool
	Requeue  bool
}
