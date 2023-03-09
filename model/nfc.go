package model

type NFCTag struct {
	ID         int    `sql:"nfc_id"`
	Name       string `sql:"name"`
	Activation bool   `sql:"activated"`
	CreatedAt  string `sql:"created_at"`
}
