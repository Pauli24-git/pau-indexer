package models

import "time"

type ZSResponse struct {
	Took     int  `json:"took"`
	TimedOut bool `json:"timed_out"`
	Shards   struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Skipped    int `json:"skipped"`
		Failed     int `json:"failed"`
	} `json:"_shards"`
	Hits struct {
		Total struct {
			Value int `json:"value"`
		} `json:"total"`
		MaxScore float64 `json:"max_score"`
		Hits     []struct {
			Index     string    `json:"_index"`
			Type      string    `json:"_type"`
			ID        string    `json:"_id"`
			Score     float64   `json:"_score"`
			Timestamp time.Time `json:"@timestamp"`
			Source    Source    `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

type Source struct {
	Timestamp               time.Time `json:"@timestamp"`
	Bcc                     string    `json:"Bcc"`
	Cc                      string    `json:"Cc"`
	Content                 string    `json:"Content"`
	ContentTransferEncoding string    `json:"Content-Transfer-Encoding"`
	Date                    string    `json:"Date"`
	From                    string    `json:"From"`
	MessageID               string    `json:"Message-ID"`
	MimeVersion             string    `json:"Mime-Version"`
	Subject                 string    `json:"Subject"`
	To                      string    `json:"To"`
	XFileName               string    `json:"X-FileName"`
	XFolder                 string    `json:"X-Folder"`
	XFrom                   string    `json:"X-From"`
	XOrigin                 string    `json:"X-Origin"`
	XTo                     string    `json:"X-To"`
	XBcc                    string    `json:"X-bcc"`
	XCc                     string    `json:"X-cc"`
}
