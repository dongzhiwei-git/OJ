package models

import "time"

type Contest struct {
	ContestId   int32     `json:"contestID"`
	Title       string    `json:"title"`
	StartTime   time.Time `json:`
	EndTime     time.Time `json:"default(null);type(datetime);null"`
	Defunct     string    `json:"type(char);size(1);default(N)"`
	Description string    `json:"type(text);null"`
	Private     uint8     `json:"type(4);default(0)"`
	Langmask    int       `json:"default(0);description:(bits for LANG to mask)"`
	Password    string    `json:"type(char);size(16);"`
	UserId      int32     `json:"null"`
}

func (Contest) TableName() string{

	return "contest"
}