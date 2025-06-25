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
	ID        string     `json:"id" validate:"required"`
	Name      string     `json:"name" validate:"min=5,max+24"`
	CreatedOn time.Time  `json:"created_on" validate:"required"`
	Content   string     `json:"content" validate:"min=5,max=1024"`
	Contacts  []Contacts `json:"contacts" validate:"min=1"`
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
