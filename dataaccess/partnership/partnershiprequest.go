package partnership

import (
	"github.com/yubing24/das/businesslogic"
	"github.com/yubing24/das/dataaccess/common"
	"database/sql"
	"errors"
	"fmt"
	sq "github.com/Masterminds/squirrel"
)

const (
	DAS_PARTNERSHIP_REQUEST_TABLE              = "DAS.PARTNERSHIP_REQUEST"
	DAS_PARTNERSHIP_REQUEST_COL_SENDER_ID      = "SENDER_ID"
	DAS_PARTNERSHIP_REQUEST_COL_RECIPIEINT_ID  = "RECIPIENT_ID"
	DAS_PARTNERSHIP_REQUEST_COL_SENDER_ROLE    = "SENDER_ROLE"
	DAS_PARTNERSHIP_REQUEST_COL_RECIPIENT_ROLE = "RECIPIENT_ROLE"
	DAS_PARTNERSHIP_REQUEST_COL_MESSAGE        = "MESSAGE"
	DAS_PARTNERSHIP_REQUEST_COL_REQUEST_STATUS = "REQUEST_STATUS"
)

type PostgresPartnershipRequestRepository struct {
	Database   *sql.DB
	SqlBuilder sq.StatementBuilderType
}

func (repo PostgresPartnershipRequestRepository) SearchPartnershipRequest(criteria *businesslogic.SearchPartnershipRequestCriteria) ([]businesslogic.PartnershipRequest, error) {
	if repo.Database == nil {
		return nil, errors.New("data source of PostgresPartnershipRequestRepository is not specified")
	}
	requests := make([]businesslogic.PartnershipRequest, 0)
	stmt := repo.SqlBuilder.Select(fmt.Sprintf("%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s",
		common.PRIMARY_KEY,
		DAS_PARTNERSHIP_REQUEST_COL_SENDER_ID,
		DAS_PARTNERSHIP_REQUEST_COL_RECIPIEINT_ID,
		DAS_PARTNERSHIP_REQUEST_COL_SENDER_ROLE,
		DAS_PARTNERSHIP_REQUEST_COL_RECIPIENT_ROLE,
		DAS_PARTNERSHIP_REQUEST_COL_MESSAGE,
		DAS_PARTNERSHIP_REQUEST_COL_REQUEST_STATUS,
		common.COL_CREATE_USER_ID,
		common.COL_DATETIME_CREATED,
		common.COL_UPDATE_USER_ID,
		common.COL_DATETIME_UPDATED)).From(DAS_PARTNERSHIP_REQUEST_TABLE).OrderBy(common.PRIMARY_KEY)

	if criteria.Sender > 0 {
		stmt = stmt.Where(sq.Eq{DAS_PARTNERSHIP_REQUEST_COL_SENDER_ID: criteria.Sender})
	}
	if criteria.Recipient > 0 {
		stmt = stmt.Where(sq.Eq{DAS_PARTNERSHIP_REQUEST_COL_RECIPIEINT_ID: criteria.Recipient})
	}
	if criteria.Sender == 0 && criteria.Recipient == 0 {
		return requests, errors.New("either sender or recipient must be specified")
	}
	if criteria.RequestStatusID > 0 {
		stmt = stmt.Where(sq.Eq{DAS_PARTNERSHIP_REQUEST_COL_REQUEST_STATUS: criteria.RequestStatusID})
	}
	if criteria.RequestID > 0 {
		stmt = stmt.Where(sq.Eq{common.PRIMARY_KEY: criteria.RequestID})
	}

	rows, err := stmt.RunWith(repo.Database).Query()
	if err != nil {
		return requests, err
	}

	for rows.Next() {
		each := businesslogic.PartnershipRequest{}
		rows.Scan(
			&each.PartnershipRequestID,
			&each.SenderID,
			&each.RecipientID,
			&each.SenderRole,
			&each.RecipientRole,
			&each.Message,
			&each.Status,
			&each.CreateUserID,
			&each.DateTimeCreated,
			&each.UpdateUserID,
			&each.DateTimeUpdated,
		)
		requests = append(requests, each)
	}
	rows.Close()
	return requests, err
}

func (repo PostgresPartnershipRequestRepository) CreatePartnershipRequest(request *businesslogic.PartnershipRequest) error {
	stmt := repo.SqlBuilder.Insert("").Into(DAS_PARTNERSHIP_REQUEST_TABLE).Columns(
		DAS_PARTNERSHIP_REQUEST_COL_SENDER_ID,
		DAS_PARTNERSHIP_REQUEST_COL_RECIPIEINT_ID,
		DAS_PARTNERSHIP_REQUEST_COL_SENDER_ROLE,
		DAS_PARTNERSHIP_REQUEST_COL_RECIPIENT_ROLE,
		DAS_PARTNERSHIP_REQUEST_COL_MESSAGE,
		DAS_PARTNERSHIP_REQUEST_COL_REQUEST_STATUS,
		common.COL_CREATE_USER_ID,
		common.COL_DATETIME_CREATED,
		common.COL_UPDATE_USER_ID,
		common.COL_DATETIME_UPDATED,
	).Values(
		request.SenderID,
		request.RecipientID,
		request.SenderRole,
		request.RecipientRole,
		request.Message,
		request.Status,
		request.CreateUserID,
		request.DateTimeCreated,
		request.UpdateUserID,
		request.DateTimeUpdated,
	).Suffix(
		fmt.Sprintf("RETURNING %s", common.PRIMARY_KEY),
	)

	clause, args, err := stmt.ToSql()
	if tx, txErr := repo.Database.Begin(); txErr != nil {
		return txErr
	} else {
		row := repo.Database.QueryRow(clause, args...)
		row.Scan(&request.PartnershipRequestID)
		err = tx.Commit()
	}
	return err
}

func (repo PostgresPartnershipRequestRepository) UpdatePartnershipRequest(request businesslogic.PartnershipRequest) error {
	clause := repo.SqlBuilder.Update("").
		Table(DAS_PARTNERSHIP_REQUEST_TABLE).
		Set(DAS_PARTNERSHIP_REQUEST_COL_REQUEST_STATUS, request.Status).
		Set(common.COL_UPDATE_USER_ID, request.RecipientID).
		Set(common.COL_DATETIME_UPDATED, request.DateTimeUpdated).
		Where(sq.Eq{common.PRIMARY_KEY: request.PartnershipRequestID})

	_, err := clause.RunWith(repo.Database).Exec()
	return err
}

func (repo PostgresPartnershipRequestRepository) DeletePartnershipRequest(request businesslogic.PartnershipRequest) error {
	return errors.New("not implemented")
}
