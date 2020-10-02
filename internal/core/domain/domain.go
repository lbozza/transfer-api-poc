package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type TransferCreateParams struct {
	PayeeID int     `json:"payeeId" binding:"required"`
	PayerID int     `json:"payerId" binding:"required"`
	Value   float32 `json:"value" binding:"required"`
}

type Transfer struct {
	ID    primitive.ObjectID `bson:"_id" json:"-"`
	Payee Customer
	Payer Customer
	Value float32
}

type Customer struct {
	ID       int
	Name     string
	Document string
	Balance  AccountBalance
}

type AccountBalance struct {
	Value float32
}

type TransferAuthorizationResponse struct {
	Authorized bool `json:"authorized"`
}

type HttpConfiguration struct {
	URL         string
	Method      string
	Headers     map[string]string
	ContentType string
	QueryParams map[string]string
	Data        interface{}
}

type HttpResponse struct {
	Status int
	Body   []byte
}
