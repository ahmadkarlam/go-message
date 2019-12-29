package request

type StoreMessageRequest struct {
	Body string `json:"body" binding:"required"`
}
