package reference

import (
	"github.com/yubing24/das/businesslogic/reference"
	"github.com/yubing24/das/dataaccess/common"
	"database/sql"
	"errors"
	"fmt"
	"github.com/Masterminds/squirrel"
	"log"
)

const (
	DAS_DANCE_TABLE = "DAS.DANCE"
)

type PostgresDanceRepository struct {
	Database   *sql.DB
	SqlBuilder squirrel.StatementBuilderType
}

func (repo PostgresDanceRepository) SearchDance(criteria *reference.SearchDanceCriteria) ([]reference.Dance, error) {
	stmt := repo.SqlBuilder.
		Select(fmt.Sprintf("%s, %s, %s, %s, %s, %s, %s, %s, %s",
			common.PRIMARY_KEY,
			common.COL_NAME,
			common.COL_ABBREVIATION,
			common.COL_DESCRIPTION,
			common.COL_STYLE_ID,
			common.COL_CREATE_USER_ID,
			common.COL_DATETIME_CREATED,
			common.COL_UPDATE_USER_ID,
			common.COL_DATETIME_UPDATED)).
		From(DAS_DANCE_TABLE).OrderBy(common.PRIMARY_KEY)
	if len(criteria.Name) > 0 {
		stmt = stmt.Where(squirrel.Eq{common.COL_NAME: criteria.Name})
	}
	if criteria.StyleID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.COL_STYLE_ID: criteria.StyleID})
	}
	if criteria.DanceID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.PRIMARY_KEY: criteria.DanceID})
	}
	rows, err := stmt.RunWith(repo.Database).Query()
	dances := make([]reference.Dance, 0)
	if err != nil {
		return dances, err
	}

	for rows.Next() {
		each := reference.Dance{}
		rows.Scan(
			&each.ID,
			&each.Name,
			&each.Abbreviation,
			&each.Description,
			&each.StyleID,
			&each.CreateUserID,
			&each.DateTimeCreated,
			&each.UpdateUserID,
			&each.DateTimeUpdated,
		)
		dances = append(dances, each)
	}
	rows.Close()
	return dances, err
}

func (repo PostgresDanceRepository) CreateDance(dance *reference.Dance) error {
	stmt := repo.SqlBuilder.Insert("").Into(DAS_DANCE_TABLE).Columns(
		common.COL_NAME,
		common.COL_ABBREVIATION,
		common.COL_DESCRIPTION,
		common.COL_STYLE_ID,
		common.COL_CREATE_USER_ID,
		common.COL_DATETIME_CREATED,
		common.COL_UPDATE_USER_ID,
		common.COL_DATETIME_UPDATED,
	).Values(
		dance.Name,
		dance.Abbreviation,
		dance.Description,
		dance.StyleID,
		dance.CreateUserID,
		dance.DateTimeCreated,
		dance.UpdateUserID,
		dance.DateTimeUpdated,
	).Suffix(
		fmt.Sprintf("RETURNING %s", common.PRIMARY_KEY),
	)

	clause, args, err := stmt.ToSql()
	if tx, txErr := repo.Database.Begin(); txErr != nil {
		return txErr
	} else {
		row := repo.Database.QueryRow(clause, args...)
		row.Scan(&dance.ID)
		tx.Commit()
	}
	return err
}

func (repo PostgresDanceRepository) UpdateDance(dance reference.Dance) error {
	stmt := repo.SqlBuilder.Update("").Table(DAS_DANCE_TABLE)
	if dance.ID > 0 {
		stmt = stmt.Set(common.COL_NAME, dance.Name).
			Set(common.COL_ABBREVIATION, dance.Abbreviation).
			Set(common.COL_DESCRIPTION, dance.Description).
			Set(common.COL_STYLE_ID, dance.StyleID).
			Set(common.COL_UPDATE_USER_ID, dance.UpdateUserID).
			Set(common.COL_DATETIME_UPDATED, dance.DateTimeUpdated)

		var err error
		if tx, txErr := repo.Database.Begin(); txErr != nil {
			return txErr
		} else {
			_, err = stmt.RunWith(repo.Database).Exec()
			tx.Commit()
		}
		return err
	}
	return errors.New("not implemented")
}

func (repo PostgresDanceRepository) DeleteDance(dance reference.Dance) error {
	if repo.Database == nil {
		log.Println(common.ERROR_NIL_DATABASE)
	}
	stmt := repo.SqlBuilder.Delete("").From(DAS_DANCE_TABLE).Where(
		squirrel.Eq{common.PRIMARY_KEY: dance.ID},
	)
	var err error
	if tx, txErr := repo.Database.Begin(); txErr != nil {
		return txErr
	} else {
		_, err = stmt.RunWith(repo.Database).Exec()
		tx.Commit()
	}
	return err
}
