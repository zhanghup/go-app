// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package source

type NewUserRegister struct {
	RawData       string `json:"raw_data" xorm:"'raw_data'"`
	EncryptedData string `json:"encrypted_data" xorm:"'encrypted_data'"`
	Signature     string `json:"signature" xorm:"'signature'"`
	Iv            string `json:"iv" xorm:"'iv'"`
}

type NewUserRegisterMobile struct {
	EncryptedData string `json:"encrypted_data" xorm:"'encrypted_data'"`
	Iv            string `json:"iv" xorm:"'iv'"`
}
