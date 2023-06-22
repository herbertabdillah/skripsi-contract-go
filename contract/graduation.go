package contract

import (
	"errors"

	"github.com/herbertabdillah/skripsi-contract-new/state"
	"github.com/hyperledger-labs/cckit/router"
)

func Graduate(c router.Context) (interface{}, error) {
	id := c.ParamString("id")

	studentRes, err := c.State().Get("Student."+id, &state.Student{})
	if err != nil {
		return nil, err
	}
	student := studentRes.(state.Student)

	departmentRes, err := c.State().Get("Department."+student.DepartmentId, &state.Department{})
	if err != nil {
		return nil, err
	}
	department := departmentRes.(state.Department)
	courseIds := department.CourseIds

	transcriptRes, err := c.State().Get("Transcript."+id, &state.Transcript{})
	if err != nil {
		return nil, err
	}
	transcript := transcriptRes.(state.Transcript)

	for _, transcriptResult := range transcript.TranscriptResult {
		if transcriptResult.Pass {
			for j, courseId := range courseIds {
				if transcriptResult.CourseId == courseId {
					copy(courseIds[j:], courseIds[j+1:])     // Shift a[i+1:] left one index.
					courseIds[len(courseIds)-1] = ""         // Erase last element (write zero value).
					courseIds = courseIds[:len(courseIds)-1] // Truncate slice.
				}
			}

		}
	}

	if len(courseIds) == 0 {
		student.Status = "graduated"
		err = c.State().Put("Student."+id, student)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("haven't done all course")
	}
	return student, nil
}

func GetTranscript(c router.Context) (interface{}, error) {
	id := c.ParamString("id")

	return c.State().Get("Transcript."+id, &state.Transcript{})
}
