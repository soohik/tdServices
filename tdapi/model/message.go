package model

const (
	SOK             = 200
	AuthWaitCode    = 201
	AuthSendTimeout = -1
	AuthSenCodeErr  = -2
	AuthorizationStateClosed
	BadRequest     = 400
	RegisterFailed = 409
	PhoneNOTFOUND  = 404
	JoinLinkUrlErr = 410
)

type Message struct {
	Code int                    `json:"code"`
	Err  string                 `json:"err"`
	Data map[string]interface{} `json:"data"`
}
