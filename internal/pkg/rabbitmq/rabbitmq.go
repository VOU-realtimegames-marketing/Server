package rabbitmq

import (
	"errors"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog/log"
)

const (
	_retryTimes     = 5
	_backOffSeconds = 2
)

type RabbitMQConnStr string

var ErrCannotConnectRabbitMQ = errors.New("cannot connect to rabbit")

func NewRabbitMQConn(rabbitMqURL RabbitMQConnStr) (*amqp.Connection, error) {
	var (
		amqpConn *amqp.Connection
		counts   int64
	)

	for {
		connection, err := amqp.Dial(string(rabbitMqURL))
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to connect to RabbitMQ: " + string(rabbitMqURL))
			counts++
		} else {
			amqpConn = connection

			break
		}

		if counts > _retryTimes {
			log.Fatal().Err(err).Msg("failed to retry")

			return nil, ErrCannotConnectRabbitMQ
		}

		log.Info().Msg("Backing off for 2 seconds...")
		time.Sleep(_backOffSeconds * time.Second)

		continue
	}

	log.Info().Msg("📫 connected to rabbitmq 🎉")

	return amqpConn, nil
}
