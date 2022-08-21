package serializer

type ResponseUser struct {
	Status int    `json:"status" example:"200"`
	Data   User   `json:"data"`
	Msg    string `json:"msg" example:"ok"`
	Error  string `json:"error"`
}
