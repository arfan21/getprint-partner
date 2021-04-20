package imgur

type ImgurResponse struct {
	Success bool      `json:"success"`
	Status  int       `json:"status"`
	Data    ImgurData `json:"data"`
}

type ImgurData struct {
	Id         string `json:"id,omitempty"`
	Title      string `json:"title,omitempty"`
	Link       string `json:"link,omitempty"`
	DeleteHash string `json:"deletehash,omitempty"`
	Error      string `json:"error,omitempty"`
	Request    string `json:"request,omitempty"`
	Method     string `json:"method,omitempty"`
}
