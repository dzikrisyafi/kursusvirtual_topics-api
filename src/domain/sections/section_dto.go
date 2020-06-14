package sections

import (
	"html"
	"strings"

	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
)

type Section struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	CourseID int    `json:"course_id"`
}

type Sections []Section

func (section Section) Validate() rest_errors.RestErr {
	section.Name = html.EscapeString(strings.TrimSpace(section.Name))
	if section.Name == "" {
		return rest_errors.NewBadRequestError("invalid section name")
	}

	if section.CourseID <= 0 {
		return rest_errors.NewBadRequestError("invalid course id")
	}

	return nil
}
