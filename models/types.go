package models

//Datas comes from user
type ConvertRequest struct {
	From   string  `json:"from"`
	To     string  `json:"to"`
	Amount float64 `json:"amount"`
}

//Data comes to user
type ConvertResponse struct {
	Result float64 `json:"result"`
}
