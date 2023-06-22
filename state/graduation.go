package state

type Transcript struct {
	Id               string             `json:"id,omitempty"`
	StudentId        string             `json:"student_id,omitempty"`
	Score            float64            `json:"score,omitempty"`
	TranscriptResult []TranscriptResult `json:"array"`
}

type TranscriptResult struct {
	CourseResultId string  `json:"course_result_id,omitempty"`
	CourseId       string  `json:"course_id,omitempty"`
	Year           int     `json:"year,omitempty"`
	Semester       string  `json:"semester,omitempty"`
	Score          float64 `json:"score,omitempty"`
	Pass           bool    `json:"pass,omitempty"`
}
