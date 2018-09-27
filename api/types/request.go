package types


/*
	key := r.FormValue("private_key")
	value := r.FormValue("value")
	to := r.FormValue("to")
	from := r.FormValue("from")

*/


type SendTransaction struct {
	PrivateKey string
	Value      string
	To         string
	From       string
	GasLimit   int64
	GasPrice   string
	Data  string
}

