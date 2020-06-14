package activity

import (
	"errors"

	"github.com/dzikrisyafi/kursusvirtual_topics-api/src/datasources/mysql/topics_db"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/logger"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
)

const (
	queryInsertActivity = `INSERT INTO activity(name, section_id, hide) VALUES(?, ?, ?);`
	queryGetActivity    = `SELECT name, section_id, hide FROM activity WHERE id=?;`
	queryUpdateActivity = `UPDATE activity SET name=?, section_id=?, hide=? WHERE id=?;`
	queryDeleteActivity = `DELETE FROM activity WHERE id=?;`
)

func (activity *Activity) Save(isHide int) rest_errors.RestErr {
	stmt, err := topics_db.DbConn().Prepare(queryInsertActivity)
	if err != nil {
		logger.Error("error when trying to prepare insert activity statement", err)
		return rest_errors.NewInternalServerError("error when trying to create activity", errors.New("database error"))
	}
	defer stmt.Close()

	result, saveErr := stmt.Exec(activity.Name, activity.SectionID, isHide)
	if saveErr != nil {
		logger.Error("error when trying to insert activity", err)
		return rest_errors.NewInternalServerError("error when trying to create activity", errors.New("database error"))
	}

	activityID, err := result.LastInsertId()
	if err != nil {
		logger.Error("error when trying to get last insert id after creating new activity", err)
		return rest_errors.NewInternalServerError("error when trying to create activity", errors.New("database error"))
	}
	activity.ID = int(activityID)

	return nil
}

func (activity *Activity) Get() rest_errors.RestErr {
	stmt, err := topics_db.DbConn().Prepare(queryGetActivity)
	if err != nil {
		logger.Error("error when trying to prepare get activity by id statement", err)
		return rest_errors.NewInternalServerError("error when trying to get activity", errors.New("database error"))
	}
	defer stmt.Close()

	result := stmt.QueryRow(activity.ID)
	if getErr := result.Scan(&activity.Name, &activity.SectionID, &activity.Hide); getErr != nil {
		logger.Error("error when trying to get activity by id", errors.New("database error"))
		return rest_errors.NewInternalServerError("error when trying to get activity", errors.New("database error"))
	}

	return nil
}

func (activity *Activity) Update(isHide int) rest_errors.RestErr {
	stmt, err := topics_db.DbConn().Prepare(queryUpdateActivity)
	if err != nil {
		logger.Error("error when trying to prepare update activity by id statement", err)
		return rest_errors.NewInternalServerError("error when trying to update activity", errors.New("database error"))
	}
	defer stmt.Close()

	if _, err := stmt.Exec(activity.Name, activity.SectionID, isHide, activity.ID); err != nil {
		logger.Error("error when trying to update activity by id", err)
		return rest_errors.NewInternalServerError("error when trying to update activity", errors.New("database error"))
	}

	return nil
}

func (activity *Activity) Delete() rest_errors.RestErr {
	stmt, err := topics_db.DbConn().Prepare(queryDeleteActivity)
	if err != nil {
		logger.Error("error when trying to prepare delete activity by id statement", err)
		return rest_errors.NewInternalServerError("error when trying to delete activity", errors.New("database error"))
	}
	defer stmt.Close()

	if _, err := stmt.Exec(activity.ID); err != nil {
		logger.Error("error when trying to delete activity by id", err)
		return rest_errors.NewInternalServerError("error when trying to delete activity", errors.New("database error"))
	}

	return nil
}
