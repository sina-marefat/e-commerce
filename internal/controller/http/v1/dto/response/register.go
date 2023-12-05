package response

type RegisterStatusCode int

const (
	InvalidRegisterEmailPassStatusCodes RegisterStatusCode = 2002
	DuplicateUserStatusCode             RegisterStatusCode = 2003
)

type Register struct {
	Message    string `json:"message"`
}
