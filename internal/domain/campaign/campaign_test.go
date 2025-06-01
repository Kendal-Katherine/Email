package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_NewCampaign_CreateCampaign(t *testing.T) {
	assert := assert.New(t)
	name := "Campaign x"
	content := "Body"
	contacts := []string{"email1@e.com", "email2@e.com"}

	campaign := NewCampaign(name, content, contacts)

	assert.Equal(name, campaign.Name)
	assert.Equal(content, campaign.Content)
	assert.Equal(len(campaign.Contacts), len(contacts))

}

func TestNewCampaign_IDIsNotNil(t *testing.T) {
	assert := assert.New(t)
	name := "Campaign x"
	content := "Body"
	contacts := []string{"email1@e.com", "email2@e.com"}

	campaign := NewCampaign(name, content, contacts)

	assert.NotNil(campaign.ID)

}

func TestNewCampaign_CreatedOnMustBeNow(t *testing.T) {
	assert := assert.New(t)
	name := "Campaign x"
	content := "Body"
	contacts := []string{"email1@e.com", "email2@e.com"}
	now := time.Now().Add(-time.Minute)

	campaign := NewCampaign(name, content, contacts)

	assert.Greater(campaign.CreatedOn, now)

}
