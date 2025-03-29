package model

type Token struct {
	SecretKey  []byte `json:"SecretKey,omitempty"`
	Expires    int64  `json:"Expires,omitempty"`
	MultiLogin bool   `json:"MultiLogin,omitempty"`
	Issuer     string `json:"Issuer,omitempty"`
}
