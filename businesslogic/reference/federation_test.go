package reference_test

import (
	"github.com/yubing24/das/businesslogic/reference"
	"github.com/yubing24/das/mock/businesslogic/reference"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFederation_GetDivisions(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mock_reference.NewMockIDivisionRepository(mockCtrl)

	// behavior 1
	mockRepo.EXPECT().SearchDivision(&reference.SearchDivisionCriteria{FederationID: 1}).Return([]reference.Division{
		{ID: 1, Name: "Correct Division 1", FederationID: 1},
		{ID: 2, Name: "Correct Division 2", FederationID: 2},
	}, nil)

	// behavior 2
	mockRepo.EXPECT().SearchDivision(&reference.SearchDivisionCriteria{FederationID: 2}).Return(nil, errors.New("invalid search"))

	federation_1 := reference.Federation{ID: 1}
	federation_2 := reference.Federation{ID: 2}

	result_1, err_1 := federation_1.GetDivisions(mockRepo)
	assert.EqualValues(t, 2, len(result_1))
	assert.Nil(t, err_1)

	result_2, err_2 := federation_2.GetDivisions(mockRepo)
	assert.Nil(t, result_2)
	assert.NotNil(t, err_2)

	result_3, err_3 := federation_1.GetDivisions(nil)
	assert.Nil(t, result_3)
	assert.NotNil(t, err_3)
}
