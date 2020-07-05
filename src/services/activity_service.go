package services

import (
	"github.com/dzikrisyafi/kursusvirtual_topics-api/src/domain/activity"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
)

var (
	ActivityService activityServiceInterface = &activityService{}
)

type activityService struct{}

type activityServiceInterface interface {
	CreateActivity(activity.Activity) (*activity.Activity, rest_errors.RestErr)
	GetActivity(int) (*activity.Activity, rest_errors.RestErr)
	UpdateActivity(bool, activity.Activity) (*activity.Activity, rest_errors.RestErr)
	DeleteActivity(int) rest_errors.RestErr
}

func (s *activityService) CreateActivity(activity activity.Activity) (*activity.Activity, rest_errors.RestErr) {
	var isHide int
	if activity.Hide {
		isHide = 1
	} else {
		isHide = 0
	}

	if err := activity.Validate(isHide); err != nil {
		return nil, err
	}

	if err := activity.Save(isHide); err != nil {
		return nil, err
	}

	return &activity, nil
}

func (s *activityService) GetActivity(activityID int) (*activity.Activity, rest_errors.RestErr) {
	result := &activity.Activity{ID: activityID}
	if err := result.Get(); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *activityService) UpdateActivity(isPartial bool, activity activity.Activity) (*activity.Activity, rest_errors.RestErr) {
	current, err := s.GetActivity(activity.ID)
	if err != nil {
		return nil, err
	}

	var isHide int
	if activity.Hide {
		isHide = 1
	} else {
		isHide = 0
	}

	if isPartial {
		if activity.Name != "" {
			current.Name = activity.Name
		}

		if activity.SectionID > 0 {
			current.SectionID = activity.SectionID
		}

		if isHide == 0 || isHide == 1 {
			current.Hide = activity.Hide
		}
	} else {
		if err := activity.Validate(isHide); err != nil {
			return nil, err
		}

		current.Name = activity.Name
		current.SectionID = activity.SectionID
		current.Hide = activity.Hide
	}

	if err := current.Update(isHide); err != nil {
		return nil, err
	}

	return current, nil
}

func (s *activityService) DeleteActivity(activityID int) rest_errors.RestErr {
	dao := &activity.Activity{ID: activityID}
	return dao.Delete()
}
