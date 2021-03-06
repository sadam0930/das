package reference

import (
	"github.com/yubing24/das/businesslogic/reference"
	"github.com/yubing24/das/dataaccess/common"
	"database/sql"
	"fmt"
	sq "github.com/Masterminds/squirrel"
)

type PostgresGenderRepository struct {
	Database   *sql.DB
	SqlBuilder sq.StatementBuilderType
}

const (
	DAS_USER_GENDER_TABLE = "DAS.GENDER"
)

func (repo PostgresGenderRepository) GetAllGenders() ([]reference.Gender, error) {
	genders := make([]reference.Gender, 0)
	stmt := repo.SqlBuilder.Select(
		fmt.Sprintf(
			"%s, %s, %s, %s, %s, %s",
			common.PRIMARY_KEY,
			common.COL_NAME,
			common.COL_ABBREVIATION,
			common.COL_DESCRIPTION,
			common.COL_DATETIME_CREATED,
			common.COL_DATETIME_UPDATED,
		)).From(DAS_USER_GENDER_TABLE)

	rows, err := stmt.RunWith(repo.Database).Query()
	if err != nil {
		return genders, err
	}

	for rows.Next() {
		each := reference.Gender{}
		rows.Scan(
			&each.ID,
			&each.Name,
			&each.Abbreviation,
			&each.Description,
			&each.DateTimeCreated,
			&each.DateTimeUpdated,
		)
		genders = append(genders, each)
	}
	rows.Close()
	return genders, err
}
