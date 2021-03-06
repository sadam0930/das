package businesslogic_test

import (
	"github.com/yubing24/das/businesslogic"
	"github.com/yubing24/das/businesslogic/reference"
	"github.com/yubing24/das/mock/businesslogic"
	"github.com/go-errors/errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var testAccount = businesslogic.Account{
	FirstName:             "First Name",
	LastName:              "Last Name",
	UserGenderID:          reference.GENDER_MALE,
	DateOfBirth:           time.Date(2017, time.January, 1, 1, 1, 1, 1, time.UTC),
	ToSAccepted:           true,
	PrivacyPolicyAccepted: true,
	AccountTypeID:         businesslogic.ACCOUNT_TYPE_ATHLETE,
	Email:                 "test@test.com",
	Phone:                 "1232234442",
	Signature:             "I am a parent",
	ByGuardian:            true,
}

func TestGetAccountByEmail(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockedAccountRepo := mock_businesslogic.NewMockIAccountRepository(mockCtrl)
	mockedAccountRepo.EXPECT().SearchAccount(&businesslogic.SearchAccountCriteria{
		Email: "test@email.com",
	}).Return(nil, errors.New("should not return an account"))
	mockedAccountRepo.EXPECT().SearchAccount(&businesslogic.SearchAccountCriteria{
		Email: "newuser@email.com",
	}).Return([]businesslogic.Account{
		businesslogic.Account{
			ID: 1, Email: "newuser@email.com",
		},
	}, nil)

	result := businesslogic.GetAccountByEmail("test@email.com", mockedAccountRepo)
	assert.Equal(t, 0, result.ID)
	assert.Equal(t, "", result.Email)

	result = businesslogic.GetAccountByEmail("newuser@email.com", mockedAccountRepo)
	assert.NotNil(t, result)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "newuser@email.com", result.Email)

}

func TestGetAccountByID(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockedAccountRepo := mock_businesslogic.NewMockIAccountRepository(mockCtrl)
	mockedAccountRepo.EXPECT().SearchAccount(&businesslogic.SearchAccountCriteria{
		ID: 1,
	}).Return(nil, errors.New("should not return an account"))
	mockedAccountRepo.EXPECT().SearchAccount(&businesslogic.SearchAccountCriteria{
		ID: 2,
	}).Return([]businesslogic.Account{
		businesslogic.Account{
			ID: 2, Email: "newuser@email.com",
		},
	}, nil)

	result := businesslogic.GetAccountByID(1, mockedAccountRepo)
	assert.Equal(t, 0, result.ID)
	assert.Equal(t, "", result.Email)

	result = businesslogic.GetAccountByID(2, mockedAccountRepo)
	assert.NotNil(t, result)
	assert.Equal(t, 2, result.ID)
	assert.Equal(t, "newuser@email.com", result.Email)
}

func TestGetAccountByUUID(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockedAccountRepo := mock_businesslogic.NewMockIAccountRepository(mockCtrl)
	mockedAccountRepo.EXPECT().SearchAccount(&businesslogic.SearchAccountCriteria{
		UUID: "abc",
	}).Return(nil, errors.New("should not return an account"))
	mockedAccountRepo.EXPECT().SearchAccount(&businesslogic.SearchAccountCriteria{
		UUID: "123",
	}).Return([]businesslogic.Account{
		businesslogic.Account{
			ID: 2, Email: "newuser@email.com",
		},
	}, nil)

	result := businesslogic.GetAccountByUUID("abc", mockedAccountRepo)
	assert.Equal(t, 0, result.ID)
	assert.Equal(t, "", result.Email)

	result = businesslogic.GetAccountByUUID("123", mockedAccountRepo)
	assert.NotNil(t, result)
	assert.Equal(t, 2, result.ID)
	assert.Equal(t, "newuser@email.com", result.Email)
}

func TestCreateAccountStrategy_CreateAccount(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockedAccountRepo := mock_businesslogic.NewMockIAccountRepository(mockCtrl)
	mockedAccountRepo.EXPECT().SearchAccount(&businesslogic.SearchAccountCriteria{
		Email: "test@test.com",
	}).Return([]businesslogic.Account{}, errors.New("account does not exist"))
	mockedAccountRepo.EXPECT().CreateAccount(gomock.Any()).Return(nil)

	strategy := businesslogic.CreateAccountStrategy{
		AccountRepo: mockedAccountRepo,
	}

	err := strategy.CreateAccount(testAccount, "testpassword")
	assert.Nil(t, err, "should not throw an error when creating account of non-organizer")
}

func TestCreateOrganizerAccountStrategy_CreateAccount(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockedAccountRepo := mock_businesslogic.NewMockIAccountRepository(mockCtrl)
	mockedProvisionRepo := mock_businesslogic.NewMockIOrganizerProvisionRepository(mockCtrl)
	mockedHistoryRepo := mock_businesslogic.NewMockIOrganizerProvisionHistoryRepository(mockCtrl)

	mockedAccountRepo.EXPECT().SearchAccount(&businesslogic.SearchAccountCriteria{
		Email: "test@test.com",
	}).Return([]businesslogic.Account{}, errors.New("account does not exist"))
	mockedAccountRepo.EXPECT().CreateAccount(gomock.Any()).Return(nil)
	mockedProvisionRepo.EXPECT().CreateOrganizerProvision(gomock.Any()).Return(nil)
	mockedHistoryRepo.EXPECT().CreateOrganizerProvisionHistory(gomock.Any()).Return(nil)

	strategy := businesslogic.CreateOrganizerAccountStrategy{
		AccountRepo:   mockedAccountRepo,
		ProvisionRepo: mockedProvisionRepo,
		HistoryRepo:   mockedHistoryRepo,
	}

	// test behaviors
	err := strategy.CreateAccount(testAccount, "testpassword")
	assert.NotNil(t, err, "non-organizer account should not be created by CreateOrganizerAccountStrategy")

	testAccount.AccountTypeID = businesslogic.ACCOUNT_TYPE_ORGANIZER
	err = strategy.CreateAccount(testAccount, "testpassword")
	assert.Nil(t, err, "should create organizer account with CreateOrganizerAccountStrategy")
}
