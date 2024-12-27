package consumer

import (
	"context"

	"github.com/google/wire"
	"github.com/pkg/errors"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog/log"
)

const (
	_exchangeKind       = "direct"
	_exchangeDurable    = true
	_exchangeAutoDelete = false
	_exchangeInternal   = false
	_exchangeNoWait     = false

	_queueDurable    = true
	_queueAutoDelete = false
	_queueExclusive  = false
	_queueNoWait     = false

	_prefetchCount  = 5
	_prefetchSize   = 0
	_prefetchGlobal = false

	_consumeAutoAck   = false
	_consumeExclusive = false
	_consumeNoLocal   = false
	_consumeNoWait    = false

	_exchangeName   = "quiz-exchange"
	_queueName      = "quiz-queue"
	_bindingKey     = "quiz-routing-key"
	_consumerTag    = "quiz-consumer"
	_workerPoolSize = 24
)

type consumer struct {
	exchangeName, queueName, bindingKey, consumerTag string
	workerPoolSize                                   int
	amqpConn                                         *amqp.Connection
}

var _ EventConsumer = (*consumer)(nil)

var EventConsumerSet = wire.NewSet(NewConsumer)

func NewConsumer(amqpConn *amqp.Connection) (EventConsumer, error) {
	sub := &consumer{
		amqpConn:       amqpConn,
		exchangeName:   _exchangeName,
		queueName:      _queueName,
		bindingKey:     _bindingKey,
		consumerTag:    _consumerTag,
		workerPoolSize: _workerPoolSize,
	}

	return sub, nil
}

func (c *consumer) Configure(opts ...Option) EventConsumer {
	for _, opt := range opts {
		opt(c)
	}

	return c
}

// StartConsumer Start new rabbitmq consumer.
func (c *consumer) StartConsumer(fn worker) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch, err := c.createChannel()
	if err != nil {
		return errors.Wrap(err, "CreateChannel")
	}
	defer ch.Close()

	deliveries, err := ch.Consume(
		c.queueName,
		c.consumerTag,
		_consumeAutoAck,
		_consumeExclusive,
		_consumeNoLocal,
		_consumeNoWait,
		nil,
	)
	if err != nil {
		return errors.Wrap(err, "Consume")
	}

	forever := make(chan bool)

	for i := 0; i < c.workerPoolSize; i++ {
		go fn(ctx, deliveries)
	}

	chanErr := <-ch.NotifyClose(make(chan *amqp.Error))
	log.Fatal().Err(chanErr).Msg("ch.NotifyClose")

	<-forever

	return chanErr
}

// CreateChannel Consume messages.
func (c *consumer) createChannel() (*amqp.Channel, error) {
	ch, err := c.amqpConn.Channel()
	if err != nil {
		return nil, errors.Wrap(err, "Error amqpConn.Channel")
	}

	log.Info().Str("exchange_name", c.exchangeName).Msg("declaring exchange")

	err = ch.ExchangeDeclare(
		c.exchangeName,
		_exchangeKind,
		_exchangeDurable,
		_exchangeAutoDelete,
		_exchangeInternal,
		_exchangeNoWait,
		nil,
	)

	if err != nil {
		return nil, errors.Wrap(err, "Error ch.ExchangeDeclare")
	}

	queue, err := ch.QueueDeclare(
		c.queueName,
		_queueDurable,
		_queueAutoDelete,
		_queueExclusive,
		_queueNoWait,
		nil,
	)
	if err != nil {
		return nil, errors.Wrap(err, "Error ch.QueueDeclare")
	}

	log.Info().Str("queue", queue.Name).Int("messages_count", queue.Messages).
		Int("consumer_count", queue.Consumers).Str("exchange", c.exchangeName).
		Str("binding_key", c.bindingKey).
		Msg("declared queue, binding it to exchange")

	err = ch.QueueBind(
		queue.Name,
		c.bindingKey,
		c.exchangeName,
		_queueNoWait,
		nil,
	)
	if err != nil {
		return nil, errors.Wrap(err, "Error ch.QueueBind")
	}

	log.Info().Str("consumer_tag", c.consumerTag).Msg("queue bound to exchange, starting to consume from queue")

	err = ch.Qos(
		_prefetchCount,  // prefetch count
		_prefetchSize,   // prefetch size
		_prefetchGlobal, // global
	)
	if err != nil {
		return nil, errors.Wrap(err, "Error ch.Qos")
	}

	return ch, nil
}
