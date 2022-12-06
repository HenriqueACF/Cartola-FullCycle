package event

import (
	"context"

	"github.com/confluentinc/confluent-kafka-go/kafka"
    "github.com/cartola-backend/pkg/uow"
)

type ProcessEventStrategy interface {
    Process(ctx context.Context, msg *kafka.Message, uow uow.UowInterface) error
}