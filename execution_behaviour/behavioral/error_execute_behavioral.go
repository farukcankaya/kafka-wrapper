package behavioral

import (
	"context"
	"github.com/Shopify/sarama"
	"github.com/Trendyol/kafka-wrapper"
)

type errorBehaviour struct {
	executor kafka_wrapper.LogicOperator
}

func ErrorBehavioral(executor kafka_wrapper.LogicOperator) BehaviourExecutor {
	return &errorBehaviour{
		executor: executor,
	}
}

func (k *errorBehaviour) Process(ctx context.Context, message *sarama.ConsumerMessage) error {
	return k.executor.Operate(ctx, message)
}
