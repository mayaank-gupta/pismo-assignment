package pkg

type CreateAccountRequest struct {
	DocumentNumber string `json:"document_number" binding:"required"`
}

type CreateTransactionRequest struct {
	AccountID       uint    `json:"account_id" binding:"required"`
	OperationTypeID uint    `json:"operation_type_id" binding:"required"`
	Amount          float64 `json:"amount" binding:"required"`
}
