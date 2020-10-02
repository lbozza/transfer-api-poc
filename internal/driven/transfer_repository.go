package driven

import (
	"go.mongodb.org/mongo-driver/mongo"
	"main/internal/core/domain"
)

type transferRepository struct {
	db *mongo.Database
}

func (t transferRepository) Save(transfer domain.TransferCreateParams) (domain.Transfer, error) {
	return domain.Transfer{}, nil
}

func NewTransferRepository(db *mongo.Database) *transferRepository {
	return &transferRepository{db: db}
}
