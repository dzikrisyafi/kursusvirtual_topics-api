package sections

type CourseSection struct {
	ID        int               `json:"id"`
	Name      string            `json:"name"`
	CourseID  int               `json:"course_id"`
	Activitys []SectionActivity `json:"activitys"`
}

type CourseSections []CourseSection

type SectionActivity struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Hide bool   `json:"hide"`
}
