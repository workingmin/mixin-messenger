package utils

type Keystore struct {
	Pin        string `json:"pin"`
	ClientId   string `json:"client_id"`
	SessionId  string `json:"session_id"`
	PinToken   string `json:"pin_token"`
	PrivateKey string `json:"private_key"`
}
