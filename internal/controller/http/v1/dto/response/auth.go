package response

type LoginStatusCode int

const (
	InvalidLoginEmailPassStatusCodes LoginStatusCode = 2002
	UserPassNotMatchStatusCode       LoginStatusCode = 2003
)

type Authentication struct {
	AccessToken string `json:"access_token"`
}
