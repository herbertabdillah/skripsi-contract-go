package state

type Transcript struct {
	Id               string   `json:"id,omitempty"`
	StudentId        string   `json:"student_id,omitempty"`
	Score            float32  `json:"score,omitempty"`
	TranscriptResult []string `json:"array"`
}

type TranscriptResult struct {
	CourseSemesterId string  `json:"course_semester_id,omitempty"`
	Score            float32 `json:"score,omitempty"`
	Pass             bool    `json:"pass,omitempty"`
}
