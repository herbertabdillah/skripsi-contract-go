package state

type CourseYear struct {
	Year     int    `json:"id,omitempty"`
	Semester string `json:"semester,omitempty"`
	Status   string `json:"status,omitempty"`
}

type CourseSemester struct {
	Id         string `json:"id,omitempty"`
	Year       int    `json:"year,omitempty"`
	Semester   string `json:"semester,omitempty"`
	CourseId   string `json:"course_id,omitempty"`
	LecturerId string `json:"lecturer_id,omitempty"`
}
