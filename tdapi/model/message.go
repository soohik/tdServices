package model

const (
	SOK             = 200
	AuthWaitCode    = 201
	AuthSendTimeout = -1
	AuthSenCodeErr  = -2
	AuthorizationStateClosed
	RegisterFailed = 409
	PhoneNOTFOUND  = 404
)

type Message struct {
	Code int                    `json:"code"`
	Err  string                 `json:"err"`
	Data map[string]interface{} `json:"data"`
}
