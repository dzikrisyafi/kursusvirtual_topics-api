package activity

import (
	"html"
	"strings"

	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
)

type Activity struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	SectionID int    `json:"section_id"`
	Hide      bool   `json:"hide"`
}

type Activitys []Activity

func (activity Activity) Validate(isHide int) rest_errors.RestErr {
	activity.Name = html.EscapeString(strings.TrimSpace(activity.Name))
	if activity.Name == "" {
		return rest_errors.NewBadRequestError("invalid activity name")
	}

	if activity.SectionID <= 0 {
		return rest_errors.NewBadRequestError("invalid section id")
	}

	if isHide < 0 || isHide > 1 {
		return rest_errors.NewBadRequestError("invalid status activity")
	}

	return nil
}
