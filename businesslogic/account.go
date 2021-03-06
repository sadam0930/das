package businesslogic

import (
	"github.com/yubing24/das/businesslogic/reference"
	"errors"
	"github.com/bearbin/go-age"
	"github.com/google/uuid"
	"log"
	"time"
)

type Account struct {
	ID                    int    // userID will be account ID, too
	UUID                  string // uuid that will be used in communicating with client
	AccountTypeID         int
	AccountStatusID       int
	UserGenderID          int
	FirstName             string
	MiddleNames           string
	LastName              string
	DateOfBirth           time.Time
	Email                 string // can be used as login user name
	Phone                 string // for raw input
	EmailVerified         bool
	PhoneVerified         bool
	HashAlgorithm         string
	PasswordSalt          []byte
	PasswordHash          []byte
	DateTimeCreated       time.Time
	DateTimeModified      time.Time
	ToSAccepted           bool // users who do not accept ToS cannot proceed anything until accepted
	PrivacyPolicyAccepted bool
	ByGuardian            bool
	Signature             string
}

type IAccountRepository interface {
	SearchAccount(criteria *SearchAccountCriteria) ([]Account, error)
	CreateAccount(account *Account) error
	UpdateAccount(account Account) error
	DeleteAccount(account Account) error
}

type SearchAccountCriteria struct {
	ID            int
	UUID          string
	Email         string
	Phone         string
	FirstName     string
	LastName      string
	DateOfBirth   time.Time
	Gender        int
	AccountType   int
	AccountStatus int
}

func (account Account) GetName() string {
	return account.FirstName + " " + account.LastName
}

type ICreateAccountStrategy interface {
	CreateAccount(account Account, password string) error
}

type CreateAccountStrategy struct {
	AccountRepo IAccountRepository
}

func (strategy CreateAccountStrategy) CreateAccount(account Account, password string) error {
	if account.AccountTypeID == ACCOUNT_TYPE_ORGANIZER {
		return errors.New("creating an organizer account with the wrong strategy")
	}
	return createAccount(&account, password, strategy.AccountRepo)
}

type CreateOrganizerAccountStrategy struct {
	AccountRepo   IAccountRepository
	ProvisionRepo IOrganizerProvisionRepository
	HistoryRepo   IOrganizerProvisionHistoryRepository
}

func (strategy CreateOrganizerAccountStrategy) CreateAccount(account Account, password string) error {
	if strategy.AccountRepo == nil {
		return errors.New("account repository is null")
	}
	if strategy.HistoryRepo == nil {
		return errors.New("organizer history repository is null")
	}
	if strategy.ProvisionRepo == nil {
		return errors.New("organizer repository is null")
	}
	if account.AccountTypeID != ACCOUNT_TYPE_ORGANIZER {
		return errors.New("not an organizer account")
	}
	if err := createAccount(&account, password, strategy.AccountRepo); err != nil {
		return err
	}
	provision, history := initializeOrganizerProvision(account.ID)
	if err := strategy.ProvisionRepo.CreateOrganizerProvision(provision); err != nil {
		return err
	}
	if err := strategy.HistoryRepo.CreateOrganizerProvisionHistory(history); err != nil {
		return err
	}
	return nil
}

func createAccount(account *Account, password string, repo IAccountRepository) error {
	if err := validateAccountRegistration(account, repo); err != nil {
		return err
	}
	salt := GenerateSalt([]byte(password))
	hash := GenerateHash(salt, []byte(password))
	account.PasswordHash = hash
	account.PasswordSalt = salt
	account.UUID = uuid.New().String()

	// TODO: email and phone verification should be performed before account can be activated
	account.AccountStatusID = ACCOUNT_STATUS_ACTIVATED

	return repo.CreateAccount(account)
}

func GetAccountByEmail(email string, repo IAccountRepository) Account {
	accounts, _ := repo.SearchAccount(&SearchAccountCriteria{
		Email: email,
	})
	if len(accounts) != 1 {
		return Account{}
	}
	return accounts[0]
}

func GetAccountByID(accountID int, repo IAccountRepository) Account {
	accounts, _ := repo.SearchAccount(&SearchAccountCriteria{
		ID: accountID,
	})
	if len(accounts) != 1 {
		return Account{}
	}
	return accounts[0]
}

func GetAccountByUUID(uuid string, repo IAccountRepository) Account {
	accounts, _ := repo.SearchAccount(&SearchAccountCriteria{
		UUID: uuid,
	})

	if len(accounts) != 1 {
		return Account{} // if no account is find, a null account will be returned
	}
	return accounts[0]
}

func checkEmailUsed(email string, repo IAccountRepository) bool {
	accounts, err := repo.SearchAccount(&SearchAccountCriteria{
		Email: email,
	})
	if err != nil {
		log.Println(err)
	}
	if len(accounts) != 0 {
		return true
	}
	return false
}

func validateAccountRegistration(account *Account, accountRepo IAccountRepository) error {
	if account.AccountTypeID > ACCOUNT_TYPE_ADMINISTRATOR || account.AccountTypeID < ACCOUNT_TYPE_ATHLETE {
		return errors.New("invalid account type")
	}
	if len(account.FirstName) < 2 || len(account.LastName) < 2 {
		return errors.New("name is too short")
	}
	if len(account.FirstName) > 18 || len(account.LastName) > 18 {
		return errors.New("name is too long")
	}
	if len(account.Email) < 5 {
		return errors.New("invalid email address")
	}

	if len(account.Phone) < 3 {
		return errors.New("invalid phone number")
	}
	if checkEmailUsed(account.Email, accountRepo) {
		return errors.New("this email address is already used")
	}
	if account.UserGenderID != reference.GENDER_FEMALE && account.UserGenderID != reference.GENDER_MALE {
		return errors.New("invalid gender")
	}
	if (time.Now().Year() - account.DateOfBirth.Year()) > 120 {
		return errors.New("invalid date of birth")
	}
	if age.AgeAt(account.DateOfBirth, time.Now()) < 13 && !account.ByGuardian {
		return errors.New("must be 13 years old to register")
	}
	if age.AgeAt(account.DateOfBirth, time.Now()) < 13 && len(account.Signature) < 3 {
		return errors.New("must have consent from legal guardian")
	}
	if !account.ToSAccepted {
		return errors.New("Terms of Services must be accepted")
	}
	if !account.PrivacyPolicyAccepted {
		return errors.New("Privacy Policy must be accepted")
	}
	return nil
}
