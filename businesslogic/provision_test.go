package businesslogic_test

import (
	"github.com/yubing24/das/businesslogic"
	"github.com/yubing24/das/mock/businesslogic"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

var organizerProvision = businesslogic.OrganizerProvision{
	ID:          2,
	OrganizerID: 12,
	Available:   2,
	Hosted:      8,
}

func TestGetOrganizerProvision(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mock_businesslogic.NewMockIOrganizerProvisionRepository(mockCtrl)
	mockRepo.EXPECT().SearchOrganizerProvision(&businesslogic.SearchOrganizerProvisionCriteria{
		OrganizerID: 1,
	}).Return([]businesslogic.OrganizerProvision{
		{ID: 1, OrganizerID: 1, Available: 1, Hosted: 2},
	}, nil)
	mockRepo.EXPECT().SearchOrganizerProvision(&businesslogic.SearchOrganizerProvisionCriteria{
		OrganizerID: 0,
	}).Return(nil, errors.New("invalid search"))

	res_1, err_1 := mockRepo.SearchOrganizerProvision(&businesslogic.SearchOrganizerProvisionCriteria{
		OrganizerID: 1,
	})
	res_2, err_2 := mockRepo.SearchOrganizerProvision(&businesslogic.SearchOrganizerProvisionCriteria{
		OrganizerID: 0,
	})

	assert.Len(t, res_1, 1)
	assert.Nil(t, err_1)
	assert.Nil(t, res_2)
	assert.NotNil(t, err_2)
}
