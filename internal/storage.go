package internal

type Entry struct {
	Account    string `json:"account"`
	Username   string `json:"username"`
	Ciphertext []byte `json:"ciphertext"`
	Nonce      []byte `json:"nonce"`
}

type Vault struct {
	Salt      []byte  `json:"salt"`
	CheckHash []byte  `json:"check_hash"`
	Entries   []Entry `json:"entries"`
}
