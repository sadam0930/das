package account

import (
	"github.com/yubing24/das/businesslogic"
	"github.com/yubing24/das/dataaccess/common"
	"database/sql"
	"errors"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"time"
)

const (
	DAS_USER_ACCOUNT_TABLE                  = "DAS.ACCOUNT"
	DAS_USER_ACCOUNT_COL_USER_TYPE_ID       = "ACCOUNT_TYPE_ID"
	DAS_USER_ACCOUNT_COL_USER_STATUS_ID     = "ACCOUNT_STATUS_ID"
	DAS_USER_ACCOUNT_COL_USER_GENDER_ID     = "USER_GENDER_ID"
	DAS_USER_ACCOUNT_COL_LAST_NAME          = "LAST_NAME"
	DAS_USER_ACCOUNT_COL_MIDDLE_NAMES       = "MIDDLE_NAMES"
	DAS_USER_ACCOUNT_COL_FIRST_NAME         = "FIRST_NAME"
	DAS_USER_ACCOUNT_COL_DATE_OF_BIRTH      = "DATE_OF_BIRTH"
	DAS_USER_ACCOUNT_COL_EMAIL              = "EMAIL"
	DAS_USER_ACCOUNT_COL_PHONE              = "PHONE"
	DAS_USER_ACCOUNT_COL_EMAIL_VERIFIED     = "EMAIL_VERIFIED"
	DAS_USER_ACCOUNT_COL_PHONE_VERIFIED     = "PHONE_VERIFIED"
	DAS_USER_ACCOUNT_HASH_ALGORITHM         = "HASH_ALGORITHM"
	DAS_USER_ACCOUNT_COL_PASSWORD_SALT      = "PASSWORD_SALT"
	DAS_USER_ACCOUNT_COL_PASSWORD_HASH      = "PASSWORD_HASH"
	DAS_USER_ACCOUNT_COL_DATETIME_CREATED   = "DATETIME_CREATED"
	DAS_USER_ACCOUNT_COL_DATETIME_UPDATED   = "DATETIME_UPDATED"
	DAS_USER_ACCOUNT_COL_TOS_ACCEPTED       = "TOS_ACCEPTED"
	DAS_USER_ACCOUNT_COL_PP_ACCEPTED        = "PP_ACCEPTED"
	DAS_USER_ACCOUNT_COL_BY_GUARDIAN        = "BY_GUARDIAN"
	DAS_USER_ACCOUNT_COL_GUARDIAN_SIGNATURE = "GUARDIAN_SIGNATURE"
)

type PostgresAccountRepository struct {
	Database   *sql.DB
	SqlBuilder sq.StatementBuilderType
}

func (repo PostgresAccountRepository) CreateAccount(account *businesslogic.Account) error {
	stmt := repo.SqlBuilder.
		Insert("").
		Into(DAS_USER_ACCOUNT_TABLE).
		Columns(DAS_USER_ACCOUNT_COL_USER_TYPE_ID,
			common.COL_UUID,
			DAS_USER_ACCOUNT_COL_USER_STATUS_ID,
			DAS_USER_ACCOUNT_COL_USER_GENDER_ID,
			DAS_USER_ACCOUNT_COL_LAST_NAME,
			DAS_USER_ACCOUNT_COL_MIDDLE_NAMES,
			DAS_USER_ACCOUNT_COL_FIRST_NAME,
			DAS_USER_ACCOUNT_COL_DATE_OF_BIRTH,
			DAS_USER_ACCOUNT_COL_EMAIL,
			DAS_USER_ACCOUNT_COL_PHONE,
			DAS_USER_ACCOUNT_COL_EMAIL_VERIFIED,
			DAS_USER_ACCOUNT_COL_PHONE_VERIFIED,
			DAS_USER_ACCOUNT_HASH_ALGORITHM,
			DAS_USER_ACCOUNT_COL_PASSWORD_SALT,
			DAS_USER_ACCOUNT_COL_PASSWORD_HASH,
			common.COL_DATETIME_CREATED,
			common.COL_DATETIME_UPDATED,
			DAS_USER_ACCOUNT_COL_TOS_ACCEPTED,
			DAS_USER_ACCOUNT_COL_PP_ACCEPTED,
			DAS_USER_ACCOUNT_COL_BY_GUARDIAN,
			DAS_USER_ACCOUNT_COL_GUARDIAN_SIGNATURE).Values(
		account.AccountTypeID, account.UUID, account.AccountStatusID, account.UserGenderID, account.LastName,
		account.MiddleNames, account.FirstName, account.DateOfBirth, account.Email, account.Phone,
		account.EmailVerified, account.PhoneVerified, account.HashAlgorithm, account.PasswordSalt, account.PasswordHash, time.Now(), time.Now(),
		account.ToSAccepted, account.PrivacyPolicyAccepted, account.ByGuardian, account.Signature,
	).Suffix(fmt.Sprintf("RETURNING %s", common.PRIMARY_KEY))

	// parsing arguments to ... parameters: https://golang.org/ref/spec#Passing_arguments_to_..._parameters
	// PostgreSQL does not return LastInsertID automatically: https://github.com/lib/pq/issues/24
	clause, args, err := stmt.ToSql()
	row := repo.Database.QueryRow(clause, args...)
	row.Scan(&account.ID)
	return err
}

func (repo PostgresAccountRepository) SearchAccount(criteria *businesslogic.SearchAccountCriteria) ([]businesslogic.Account, error) {
	stmt := repo.SqlBuilder.
		Select(
			fmt.Sprintf(
				"%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s",
				common.PRIMARY_KEY,
				common.COL_UUID,
				DAS_USER_ACCOUNT_COL_USER_TYPE_ID,
				DAS_USER_ACCOUNT_COL_USER_STATUS_ID,
				DAS_USER_ACCOUNT_COL_USER_GENDER_ID,
				DAS_USER_ACCOUNT_COL_LAST_NAME,
				DAS_USER_ACCOUNT_COL_MIDDLE_NAMES,
				DAS_USER_ACCOUNT_COL_FIRST_NAME,
				DAS_USER_ACCOUNT_COL_DATE_OF_BIRTH,
				DAS_USER_ACCOUNT_COL_EMAIL,
				DAS_USER_ACCOUNT_COL_PHONE,
				DAS_USER_ACCOUNT_COL_EMAIL_VERIFIED,
				DAS_USER_ACCOUNT_COL_PHONE_VERIFIED,
				DAS_USER_ACCOUNT_HASH_ALGORITHM,
				DAS_USER_ACCOUNT_COL_PASSWORD_SALT,
				DAS_USER_ACCOUNT_COL_PASSWORD_HASH,
				DAS_USER_ACCOUNT_COL_DATETIME_CREATED,
				DAS_USER_ACCOUNT_COL_DATETIME_UPDATED,
				DAS_USER_ACCOUNT_COL_TOS_ACCEPTED,
				DAS_USER_ACCOUNT_COL_PP_ACCEPTED,
				DAS_USER_ACCOUNT_COL_BY_GUARDIAN,
				DAS_USER_ACCOUNT_COL_GUARDIAN_SIGNATURE,
			)).From(DAS_USER_ACCOUNT_TABLE)
	if criteria.AccountType > 0 {
		stmt = stmt.Where(sq.Eq{DAS_USER_ACCOUNT_COL_USER_TYPE_ID: criteria.AccountType})
	}
	if len(criteria.UUID) != 0 {
		stmt = stmt.Where(sq.Eq{common.COL_UUID: criteria.UUID})
	}
	if criteria.ID > 0 {
		stmt = stmt.Where(sq.Eq{common.PRIMARY_KEY: criteria.ID})
	}
	if criteria.AccountStatus > 0 {
		stmt = stmt.Where(sq.Eq{DAS_USER_ACCOUNT_COL_USER_STATUS_ID: criteria.AccountStatus})
	}
	if criteria.Gender > 0 {
		stmt = stmt.Where(sq.Eq{DAS_USER_ACCOUNT_COL_USER_GENDER_ID: criteria.Gender})
	}
	if len(criteria.Email) > 0 {
		stmt = stmt.Where(sq.Eq{DAS_USER_ACCOUNT_COL_EMAIL: criteria.Email})
	}
	if len(criteria.Phone) > 0 {
		stmt = stmt.Where(sq.Eq{DAS_USER_ACCOUNT_COL_PHONE: criteria.Phone})
	}
	if len(criteria.LastName) > 0 {
		stmt = stmt.Where(sq.Eq{DAS_USER_ACCOUNT_COL_LAST_NAME: criteria.LastName})
	}
	if len(criteria.FirstName) > 0 {
		stmt = stmt.Where(sq.Eq{DAS_USER_ACCOUNT_COL_FIRST_NAME: criteria.FirstName})
	}
	if &criteria.DateOfBirth != nil && (time.Now().Year()-criteria.DateOfBirth.Year()) < 120 {
		stmt = stmt.Where(sq.Eq{DAS_USER_ACCOUNT_COL_DATE_OF_BIRTH: criteria.DateOfBirth})
	}

	accounts := make([]businesslogic.Account, 0)
	rows, err := stmt.RunWith(repo.Database).Query()
	if err != nil {
		return accounts, err
	}
	for rows.Next() {
		each := businesslogic.Account{}
		rows.Scan(
			&each.ID,
			&each.UUID,
			&each.AccountTypeID,
			&each.AccountStatusID,
			&each.UserGenderID,
			&each.LastName,
			&each.MiddleNames,
			&each.FirstName,
			&each.DateOfBirth,
			&each.Email,
			&each.Phone,
			&each.EmailVerified,
			&each.PhoneVerified,
			&each.HashAlgorithm,
			&each.PasswordSalt,
			&each.PasswordHash,
			&each.DateTimeCreated,
			&each.DateTimeModified,
			&each.ToSAccepted,
			&each.PrivacyPolicyAccepted,
			&each.ByGuardian,
			&each.Signature,
		)
		accounts = append(accounts, each)
	}
	rows.Close()
	return accounts, err
}

func (repo PostgresAccountRepository) DeleteAccount(account businesslogic.Account) error {
	if account.ID > 0 {
		stmt := repo.SqlBuilder.Delete("").From(DAS_USER_ACCOUNT_TABLE).Where(sq.Eq{common.PRIMARY_KEY: account.ID})
		_, err := stmt.RunWith(repo.Database).Exec()
		return err
	}
	return errors.New("account ID was not specified")
}

func (repo PostgresAccountRepository) UpdateAccount(account businesslogic.Account) error {
	return errors.New("not implemented")
}
