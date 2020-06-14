package sections

import (
	"errors"
	"fmt"

	"github.com/dzikrisyafi/kursusvirtual_topics-api/src/datasources/mysql/topics_db"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/logger"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
)

const (
	queryInsertSection         = `INSERT INTO sections(name, course_id) VALUES(?, ?);`
	queryGetSection            = `SELECT name, course_id FROM sections WHERE id=?;`
	queryGetSectionsByCourseID = `SELECT id, name, course_id FROM sections WHERE course_id=?;`
	queryGetAllActivity        = `SELECT id, name, hide FROM activity WHERE section_id=?;`
	queryUpdateSection         = `UPDATE sections SET name=?, course_id=? WHERE id=?;`
	queryDeleteSection         = `DELETE FROM sections WHERE id=?;`
)

func (section *Section) Save() rest_errors.RestErr {
	stmt, err := topics_db.DbConn().Prepare(queryInsertSection)
	if err != nil {
		logger.Error("error when trying to prepare insert section statement", err)
		return rest_errors.NewInternalServerError("error when trying to create section", errors.New("database error"))
	}
	defer stmt.Close()

	result, saveErr := stmt.Exec(section.Name, section.CourseID)
	if saveErr != nil {
		logger.Error("error when trying to insert section", err)
		return rest_errors.NewInternalServerError("error when trying to create section", errors.New("database error"))
	}

	sectionID, err := result.LastInsertId()
	if err != nil {
		logger.Error("error when trying to get last insert id after creating a section", err)
		return rest_errors.NewInternalServerError("error when trying to create section", errors.New("database error"))
	}
	section.ID = int(sectionID)

	return nil
}

func (section *Section) Get() rest_errors.RestErr {
	stmt, err := topics_db.DbConn().Prepare(queryGetSection)
	if err != nil {
		logger.Error("error when trying to prepare get section by id statement", err)
		return rest_errors.NewInternalServerError("error when trying to save section", errors.New("database error"))
	}
	defer stmt.Close()

	result := stmt.QueryRow(section.ID)
	if getErr := result.Scan(&section.Name, &section.CourseID); getErr != nil {
		logger.Error("error when trying to get section by id", getErr)
		return rest_errors.NewInternalServerError("error when trying to get section", errors.New("database error"))
	}

	return nil
}

func (section *CourseSection) GetAll() ([]CourseSection, rest_errors.RestErr) {
	stmt, err := topics_db.DbConn().Prepare(queryGetSectionsByCourseID)
	if err != nil {
		logger.Error("error when trying to prepare get all section by course id", err)
		return nil, rest_errors.NewInternalServerError("error when trying to get all section", errors.New("database error"))
	}
	defer stmt.Close()

	rows, err := stmt.Query(section.CourseID)
	if err != nil {
		logger.Error("error when trying to get all section by course id", err)
		return nil, rest_errors.NewInternalServerError("error when trying to get all section", errors.New("database error"))
	}
	defer rows.Close()

	result := make([]CourseSection, 0)
	for rows.Next() {
		if err := rows.Scan(&section.ID, &section.Name, &section.CourseID); err != nil {
			logger.Error("error when trying to scan section rows into section struct", err)
			return nil, rest_errors.NewInternalServerError("error when trying to get all section", errors.New("database error"))
		}

		result = append(result, *section)
	}

	if len(result) == 0 {
		return nil, rest_errors.NewNotFoundError(fmt.Sprintf("no section matching given course id %d", section.CourseID))
	}

	return result, nil
}

func (activity *SectionActivity) GetAllActivity(section *CourseSection) rest_errors.RestErr {
	stmt, err := topics_db.DbConn().Prepare(queryGetAllActivity)
	if err != nil {
		logger.Error("error when trying to prepapre get all activity by section id statement", err)
		return rest_errors.NewInternalServerError("error when trying to get all activity", errors.New("database error"))
	}
	defer stmt.Close()

	rows, err := stmt.Query(section.ID)
	if err != nil {
		logger.Error("error when trying to get all activity by section id", err)
		return rest_errors.NewInternalServerError("error when trying to get all activity", errors.New("database error"))
	}
	defer rows.Close()

	var isHide int
	for rows.Next() {
		if err := rows.Scan(&activity.ID, &activity.Name, &isHide); err != nil {
			logger.Error("error when trying to scan activity rows into activity struct", err)
			return rest_errors.NewInternalServerError("error when trying to get all activity", errors.New("database error"))
		}

		if isHide == 1 {
			activity.Hide = true
		} else {
			activity.Hide = false
		}

		section.Activitys = append(section.Activitys, *activity)
	}

	return nil
}

func (section *Section) Update() rest_errors.RestErr {
	stmt, err := topics_db.DbConn().Prepare(queryUpdateSection)
	if err != nil {
		logger.Error("error when trying to prepare update section statement", err)
		return rest_errors.NewInternalServerError("error when trying to update section", errors.New("database error"))
	}
	defer stmt.Close()

	if _, err := stmt.Exec(section.Name, section.CourseID, section.ID); err != nil {
		logger.Error("error when trying to update section", err)
		return rest_errors.NewInternalServerError("error when trying to update section", errors.New("database error"))
	}

	return nil
}

func (section *Section) Delete() rest_errors.RestErr {
	stmt, err := topics_db.DbConn().Prepare(queryDeleteSection)
	if err != nil {
		logger.Error("error when trying to prepare delete section by id statement", err)
		return rest_errors.NewInternalServerError("error when trying to delete section", errors.New("database error"))
	}
	defer stmt.Close()

	if _, err := stmt.Exec(section.ID); err != nil {
		logger.Error("error when trying to delete section by id", err)
		return rest_errors.NewInternalServerError("error when trying to delete section", errors.New("database error"))
	}

	return nil
}
