package model

type Message struct {
	Code int                    `json:"code"`
	Err  string                 `json:"err"`
	Data map[string]interface{} `json:"data"`
}
