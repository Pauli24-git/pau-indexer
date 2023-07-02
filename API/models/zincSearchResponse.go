package models

import "time"

type ZSResponse struct {
	Timed_out bool    `json:"timed_out"`
	Shards    Shards  `json:"_shards"`
	Hits      Hits    `json:"hits"`
	Took      float64 `json:"took"`
}

type Shards struct {
	Total      float64 `json:"total"`
	Successful float64 `json:"successful"`
	Skipped    float64 `json:"skipped"`
	Failed     float64 `json:"failed"`
}

type Hits struct {
	Total     Total     `json:"total"`
	Max_score float64   `json:"max_score"`
	Hits      []Hits    `json:"hits"`
	Index     string    `json:"_index"`
	Type      string    `json:"_type"`
	Id        string    `json:"_id"`
	Score     float64   `json:"_score"`
	Timestamp time.Time `json:"@timestamp"`
	Source    Source    `json:"_source"`
}

type Total struct {
	Value float64 `json:"value"`
}

type Source struct {
	X_To                      string    `json:"X-To"`
	Bcc                       string    `json:"Bcc"`
	Date                      string    `json:"Date"`
	Subject                   string    `json:"Subject"`
	X_FileName                string    `json:"X-FileName"`
	X_Origin                  string    `json:"X-Origin"`
	From                      string    `json:"From"`
	Message_ID                string    `json:"Message-ID"`
	X_bcc                     string    `json:"X-bcc"`
	Content                   string    `json:"Content"`
	Content_Transfer_Encoding string    `json:"Content-Transfer-Encoding"`
	X_Folder                  string    `json:"X-Folder"`
	X_From                    string    `json:"X-From"`
	X_cc                      string    `json:"X-cc"`
	Timestamp                 time.Time `json:"@timestamp"`
	Cc                        string    `json:"Cc"`
	Mime_Version              string    `json:"Mime-Version"`
	To                        string    `json:"To"`
}
