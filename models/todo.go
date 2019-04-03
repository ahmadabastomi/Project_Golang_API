package models

import "time"

type Todos struct {
	Id        int       `json:"id" orm:"auto"`
	Todo      string    `json:"todo"`
	CreatedAt time.Time `json:"created_at" orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `json:"updated_at" orm:"auto_now;type(datetime)"`
}
