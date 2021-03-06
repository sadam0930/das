package reference

import (
	"github.com/yubing24/das/businesslogic/reference"
	"github.com/yubing24/das/dataaccess/common"
	"database/sql"
	"errors"
	"fmt"
	"github.com/Masterminds/squirrel"
)

const (
	DAS_SCHOOL_TABLE = "DAS.SCHOOL"
)

type PostgresSchoolRepository struct {
	Database   *sql.DB
	SqlBuilder squirrel.StatementBuilderType
}

func (repo PostgresSchoolRepository) CreateSchool(school *reference.School) error {
	if repo.Database == nil {
		return errors.New("data source of PostgresSchoolRepository is not specified")
	}
	stmt := repo.SqlBuilder.Insert("").Into(DAS_SCHOOL_TABLE).Columns(
		common.COL_NAME,
		common.COL_CITY_ID,
		common.COL_CREATE_USER_ID,
		common.COL_DATETIME_CREATED,
		common.COL_UPDATE_USER_ID,
		common.COL_DATETIME_UPDATED,
	).Values(
		school.Name,
		school.CityID,
		school.CreateUserID,
		school.DateTimeCreated,
		school.UpdateUserID,
		school.DateTimeUpdated,
	).Suffix(
		fmt.Sprintf("RETURNING %s", common.PRIMARY_KEY),
	)

	clause, args, err := stmt.ToSql()
	if tx, txErr := repo.Database.Begin(); txErr != nil {
		return txErr
	} else {
		row := repo.Database.QueryRow(clause, args...)
		row.Scan(&school.ID)
		if err = tx.Commit(); err != nil {
			tx.Rollback()
		}
	}
	return err
}

func (repo PostgresSchoolRepository) UpdateSchool(school reference.School) error {
	if repo.Database == nil {
		return errors.New("data source of PostgresSchoolRepository is not specified")
	}
	stmt := repo.SqlBuilder.Update("").Table(DAS_SCHOOL_TABLE)
	if school.ID > 0 {
		stmt = stmt.Set(common.COL_NAME, school.Name).
			Set(common.COL_CITY_ID, school.CityID).
			Set(common.COL_UPDATE_USER_ID, school.UpdateUserID).
			Set(common.COL_DATETIME_UPDATED, school.DateTimeUpdated)
		var err error
		if tx, txErr := repo.Database.Begin(); txErr != nil {
			return txErr
		} else {
			_, err = stmt.RunWith(repo.Database).Exec()
			if err = tx.Commit(); err != nil {
				tx.Rollback()
			}
		}
		return err
	} else {
		return errors.New("school is not specified")
	}
}

func (repo PostgresSchoolRepository) DeleteSchool(school reference.School) error {
	if repo.Database == nil {
		return errors.New("data source of PostgresSchoolRepository is not specified")
	}
	stmt := repo.SqlBuilder.
		Delete("").
		From(DAS_SCHOOL_TABLE).
		Where(squirrel.Eq{common.PRIMARY_KEY: school.ID})
	var err error
	if tx, txErr := repo.Database.Begin(); txErr != nil {
		return txErr
	} else {
		_, err = stmt.RunWith(repo.Database).Exec()
		if err = tx.Commit(); err != nil {
			tx.Rollback()
		}
	}
	return err
}

func (repo PostgresSchoolRepository) SearchSchool(criteria *reference.SearchSchoolCriteria) ([]reference.School, error) {
	if repo.Database == nil {
		return nil, errors.New("data source of PostgresSchoolRepository is not specified")
	}
	stmt := repo.SqlBuilder.
		Select(fmt.Sprintf(
			`%s,%s, %s,%s, %s, %s, %s`,
			common.PRIMARY_KEY,
			common.COL_NAME,
			common.COL_CITY_ID,
			common.COL_CREATE_USER_ID,
			common.COL_DATETIME_CREATED,
			common.COL_UPDATE_USER_ID,
			common.COL_DATETIME_UPDATED)).
		From(DAS_SCHOOL_TABLE).
		OrderBy(`DAS.SCHOOL.ID`)
	if criteria.ID > 0 {
		stmt = stmt.Where(squirrel.Eq{`DAS.SCHOOL.ID`: criteria.ID})
	}
	if len(criteria.Name) > 0 {
		stmt = stmt.Where(squirrel.Eq{`DAS.SCHOOL.NAME`: criteria.Name})
	}
	if criteria.CityID > 0 {
		stmt = stmt.Where(squirrel.Eq{`DAS.SCHOOL.CITY_ID`: criteria.CityID})
	}
	if criteria.StateID > 0 {
		stmt = stmt.Join(`DAS.CITY C ON C.ID = DAS.SCHOOL.CITY_ID`).
			Where(squirrel.Eq{`C.STATE_ID`: criteria.StateID})
	}
	rows, err := stmt.RunWith(repo.Database).Query()
	schools := make([]reference.School, 0)
	if err != nil {
		return schools, err
	}
	for rows.Next() {
		each := reference.School{}
		rows.Scan(
			&each.ID,
			&each.Name,
			&each.CityID,
			&each.CreateUserID,
			&each.DateTimeCreated,
			&each.UpdateUserID,
			&each.DateTimeUpdated,
		)
		schools = append(schools, each)
	}
	rows.Close()
	return schools, err
}
