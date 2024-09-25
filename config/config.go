package config

import (
	gossiper "github.com/pieceowater-dev/lotof.lib.gossiper"
)

var GossiperConf = gossiper.Config{
	Env: gossiper.EnvConfig{
		Required: []string{"RABBITMQ_DSN"},
	},
	AMQPConsumer: gossiper.AMQPConsumerConfig{
		Queues: []gossiper.QueueConfig{
			{
				Name:    "template_queue",
				Durable: true,
			},
		},
		Consume: []gossiper.AMQPConsumeConfig{
			{
				Queue:    "template_queue",
				Consumer: "example_consumer",
				AutoAck:  true,
			},
		},
	},
}

func GetRabbitDSN() string {
	env := &gossiper.Env{}
	val, err := env.Get(GossiperConf.Env.Required[0])
	if err != nil {
		return ""
	}
	return val
}
