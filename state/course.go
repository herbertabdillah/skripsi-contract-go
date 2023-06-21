package state

type CoursePlan struct {
	Id                string   `json:"id,omitempty"`
	Year              int      `json:"year,omitempty"`
	Semester          string   `json:"semester,omitempty"`
	StudentId         string   `json:"student_id,omitempty"`
	Status            string   `json:"status,omitempty"`
	CourseSemesterIds []string `json:"array"`
}

type CourseSemeterResult struct {
	CourseSemesterId string  `json:"course_semester_id,omitempty"`
	Score            float32 `json:"score,omitempty"`
	Pass             bool    `json:"pass,omitempty"`
}

type CourseResult struct {
	Id           string                `json:"id,omitempty"`
	Year         int                   `json:"year,omitempty"`
	Semester     string                `json:"semester,omitempty"`
	StudentId    string                `json:"student_id,omitempty"`
	CoursePlanId string                `json:"course_plan_id,omitempty"`
	Result       []CourseSemeterResult `json:"array"`
}
