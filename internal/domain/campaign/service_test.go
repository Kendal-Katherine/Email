package campaign

import (
	"email/internal/contract"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	mock.Mock
}

func (r *RepositoryMock) Save(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

func Test_Create_Campaign(t *testing.T) {

	assert := assert.New(t)
	service := Service{}

	newCampaign := contract.NewCampaign{
		Name:    "Test x",
		Content: "Body",
		Emails:  []string{"teste@e.com"},
	}

	id, err := service.Create(newCampaign)
	assert.NotNil(id)
	assert.Nil(err)
}

func Test_Create_SaveCampaign(t *testing.T) {

	repositoryMock := new(RepositoryMock)

	newCampaign := contract.NewCampaign{
		Name:    "Test x",
		Content: "Body",
		Emails:  []string{"teste@e.com"},
	}

	repositoryMock.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {
		if campaign.Name != newCampaign.Name ||
			campaign.Content != newCampaign.Content ||
			len(campaign.Contacts) != len(newCampaign.Emails) {
			return false
		}

		return true
	})).Return(nil)
	service := Service{Repository: repositoryMock}
	service.Create(newCampaign)
	repositoryMock.AssertExpectations(t)
}
