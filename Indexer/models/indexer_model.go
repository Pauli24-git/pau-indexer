package models

type Mails struct {
	MessageID               string `json:"MessageID"`
	Date                    string `json:"Date"`
	From                    string `json:"From"`
	To                      string `json:"To"`
	Subject                 string `json:"Subject"`
	Cc                      string `json:"Cc"`
	MimeVersion             string `json:"Mime-Version"`
	ContentType             string `json:"Content-Type"`
	ContentTransferEncoding string `json:"Content-Transfer-Encoding"`
	Bcc                     string `json:"Bcc"`
	Xfrom                   string `json:"X-from"`
	Xto                     string `json:"X-to"`
	Xcc                     string `json:"X-cc"`
	Xbcc                    string `json:"X-bcc"`
	Xfolder                 string `json:"X-folder"`
	Xorigin                 string `json:"X-origin"`
	Xfilename               string `json:"X-filename"`
	Content                 string `json:"Content"`
}

type Records struct {
	Index   string              `json:"index"`
	Records []map[string]string `json:"records"`
}
