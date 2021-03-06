package reference_test

import (
	"github.com/yubing24/das/businesslogic/reference"
	"github.com/yubing24/das/mock/businesslogic/reference"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestState_GetCities(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mock_reference.NewMockICityRepository(mockCtrl)

	// behavior 1
	mockRepo.EXPECT().SearchCity(&reference.SearchCityCriteria{StateID: 1}).Return([]reference.City{
		{CityID: 1, Name: "City of ID 1", StateID: 1},
		{CityID: 2, Name: "City of ID 2", StateID: 1},
	}, nil)

	// behavior 2
	mockRepo.EXPECT().SearchCity(&reference.SearchCityCriteria{StateID: 2}).Return(nil,
		errors.New("state does not exist"))

	state_1 := reference.State{ID: 1}
	cities_1, err_1 := state_1.GetCities(mockRepo)
	assert.EqualValues(t, 2, len(cities_1))
	assert.Nil(t, err_1)

	state_2 := reference.State{ID: 2}
	cities_2, err_2 := state_2.GetCities(mockRepo)
	assert.Nil(t, cities_2)
	assert.NotNil(t, err_2)

	cities_3, err_3 := state_1.GetCities(nil)
	assert.Nil(t, cities_3)
	assert.NotNil(t, err_3)

}
