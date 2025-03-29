package model

type Token struct {
	SecretKey  string `json:"SecretKey,omitempty"`
	ExpiresAt  int64  `json:"ExpiresAt,omitempty"`
	MultiLogin bool   `json:"MultiLogin,omitempty"`
	Issuer     string `json:"Issuer,omitempty"`
}
