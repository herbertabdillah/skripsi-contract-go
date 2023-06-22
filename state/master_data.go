package state

type Faculty struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Department struct {
	Id        string   `json:"id,omitempty"`
	FacultyId string   `json:"faculty_id,omitempty"`
	Name      string   `json:"name,omitempty"`
	CourseIds []string `json:"course_ids,omitempty"`
}

type Course struct {
	Id           string `json:"id,omitempty"`
	DepartmentId string `json:"department_id,omitempty"`
	Name         string `json:"name,omitempty"`
	Credit       int    `json:"credit,omitempty"`
	Kind         string `json:"kind,omitempty"`
}
