package campaign

import (
	"errors"
	"time"

	"github.com/rs/xid"
)

type Contacts struct {
	Email string `json:"emails"`
}
type Campaign struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	CreatedOn time.Time  `json:"created_on"`
	Content   string     `json:"content"`
	Contacts  []Contacts `json:"contacts"`
}

func NewCampaign(name string, content string, emails []string) (*Campaign, error) {
	if name == "" {
		return nil, errors.New("name is required")
	} else if content == "" {
		return nil, errors.New("content is required")
	} else if len(emails) == 0 {
		return nil, errors.New("contacts is required")
	}

	contacts := make([]Contacts, len(emails))
	for index, email := range emails {
		contacts[index].Email = email
	}

	return &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		CreatedOn: time.Now(),
		Content:   content,
		Contacts:  contacts,
	}, nil
}
