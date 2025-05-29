package campaign

import "time"

type Contacts struct {
	Emails string `json:"emails"`
}
type Campaign struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	CreatedOn time.Time  `json:"created_on"`
	Content   string     `json:"content"`
	Contacts  []Contacts `json:"contacts"`
}
