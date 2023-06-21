package state

type Lecturer struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Nik  string `json:"nik,omitempty"`
}

type Student struct {
	Id                   string `json:"id,omitempty"`
	Name                 string `json:"name,omitempty"`
	Nim                  string `json:"nim,omitempty"`
	DepartmentId         string `json:"department_id,omitempty"`
	EntryYear            int    `json:"entry_year,omitempty"`
	Status               string `json:"status,omitempty"`
	SupervisorLecturerId string `json:"supervisor_lecturer_id,omitempty"`
}
