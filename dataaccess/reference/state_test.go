package reference

import (
	"github.com/yubing24/das/businesslogic/reference"
	"github.com/Masterminds/squirrel"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
	"time"
)

var stateRepository = PostgresStateRepository{

	Database:   nil,
	SqlBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

func TestDasStateRepository_SearchState(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	stateRepository.Database = db

	rows := sqlmock.NewRows(
		[]string{
			"ID", "NAME", "ADDRESS", "CITY_ID", "WEBSITE", "CREATE_USER_ID", "DATETIME_CREATED", "UPDATE_USER_ID", "DATETIME_UPDATED",
		},
	).AddRow(1, "Kanopy", "WI", 8, "www.example.com", 3, time.Now(), 4, time.Now())

	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	states, _ := stateRepository.SearchState(&reference.SearchStateCriteria{})

	assert.NotZero(t, len(states))
}
