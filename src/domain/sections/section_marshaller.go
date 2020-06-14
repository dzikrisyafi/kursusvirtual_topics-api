package sections

type PublicSection struct {
	ID int `json:"id"`
}

func (section *Section) Marshall(isPublic bool) interface{} {
	if isPublic {
		return PublicSection{
			ID: section.ID,
		}
	}

	return section
}

func (sections CourseSections) Marshall(isPublic bool) []interface{} {
	result := make([]interface{}, len(sections))
	for index, section := range sections {
		result[index] = section.Marshall(isPublic)
	}

	return result
}

func (section CourseSection) Marshall(isPublic bool) interface{} {
	if isPublic {
		return PublicSection{
			ID: section.ID,
		}
	}

	return section
}
