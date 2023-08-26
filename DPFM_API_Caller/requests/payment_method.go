package requests

type PaymentMethod struct {
	PaymentMethod       string `json:"PaymentMethod"`
	CreationDate        string `json:"CreationDate"`
	LastChangeDate      string `json:"LastChangeDate"`
	IsMarkedForDeletion *int   `json:"IsMarkedForDeletion"`
}
