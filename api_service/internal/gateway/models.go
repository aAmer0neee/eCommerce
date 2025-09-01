package gateway

type ResponseMessage struct {
	Code    int
	Message string
}

type RegisterUserRequest struct {
	Email string `json:"email"  binding:"required"`
	Name  string `json:"name"  binding:"required"`
}

type RegisterUserResponse struct {
	Id      string
	Message ResponseMessage
}
