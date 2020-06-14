package services

import (
	"github.com/dzikrisyafi/kursusvirtual_topics-api/src/domain/sections"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
)

var (
	SectionsService sectionsServiceInterface = &sectionsService{}
)

type sectionsService struct{}

type sectionsServiceInterface interface {
	CreateSection(sections.Section) (*sections.Section, rest_errors.RestErr)
	GetSection(int) (*sections.Section, rest_errors.RestErr)
	GetAllSectionByCourseID(int) (sections.CourseSections, rest_errors.RestErr)
	GetAllActivity(*sections.CourseSection) rest_errors.RestErr
	UpdateSection(bool, sections.Section) (*sections.Section, rest_errors.RestErr)
	DeleteSection(int) rest_errors.RestErr
	DeleteSectionByCourseID(int) rest_errors.RestErr
}

func (s *sectionsService) CreateSection(section sections.Section) (*sections.Section, rest_errors.RestErr) {
	if err := section.Validate(); err != nil {
		return nil, err
	}

	if err := section.Save(); err != nil {
		return nil, err
	}

	return &section, nil
}

func (s *sectionsService) GetSection(sectionID int) (*sections.Section, rest_errors.RestErr) {
	result := &sections.Section{ID: sectionID}
	if err := result.Get(); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *sectionsService) GetAllSectionByCourseID(courseID int) (sections.CourseSections, rest_errors.RestErr) {
	dao := &sections.CourseSection{CourseID: courseID}
	allSection, err := dao.GetAll()
	if err != nil {
		return nil, err
	}

	results := make([]sections.CourseSection, 0)
	for _, section := range allSection {
		if err := s.GetAllActivity(&section); err != nil {
			return nil, err
		}

		results = append(results, section)
	}

	return results, nil
}

func (s *sectionsService) GetAllActivity(section *sections.CourseSection) rest_errors.RestErr {
	dao := &sections.SectionActivity{}
	return dao.GetAllActivity(section)
}

func (s *sectionsService) UpdateSection(isPartial bool, section sections.Section) (*sections.Section, rest_errors.RestErr) {
	current, err := s.GetSection(section.ID)
	if err != nil {
		return nil, err
	}

	if isPartial {
		if section.Name != "" {
			current.Name = section.Name
		}

		if section.CourseID > 0 {
			current.CourseID = section.CourseID
		}
	} else {
		if err := section.Validate(); err != nil {
			return nil, err
		}

		current.Name = section.Name
		current.CourseID = section.CourseID
	}

	if err := current.Update(); err != nil {
		return nil, err
	}

	return current, nil
}

func (s *sectionsService) DeleteSection(sectionID int) rest_errors.RestErr {
	dao := &sections.Section{ID: sectionID}
	return dao.Delete()
}

func (s *sectionsService) DeleteSectionByCourseID(courseID int) rest_errors.RestErr {
	dao := &sections.Section{CourseID: courseID}
	return dao.DeleteByCourseID()
}
