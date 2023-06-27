package contract

import (
	"errors"

	"github.com/hyperledger-labs/cckit/router"
)

func Graduate(c router.Context) (interface{}, error) {
	cc := NewContext(c)
	id := c.ParamString("id")

	student, err := cc.Repository.GetStudent(id)
	if err != nil {
		return nil, err
	}

	department, err := cc.Repository.GetDepartment(student.DepartmentId)
	if err != nil {
		return nil, err
	}

	transcript, err := cc.Repository.GetTranscript(id)
	if err != nil {
		return nil, err
	}

	haventDoneCourseIds := department.CourseIds
	for _, transcriptResult := range transcript.TranscriptResult {
		if !transcriptResult.Pass {
			continue
		}

		for j, courseId := range haventDoneCourseIds {
			if transcriptResult.CourseId == courseId {
				copy(haventDoneCourseIds[j:], haventDoneCourseIds[j+1:])               // Shift a[i+1:] left one index.
				haventDoneCourseIds[len(haventDoneCourseIds)-1] = ""                   // Erase last element (write zero value).
				haventDoneCourseIds = haventDoneCourseIds[:len(haventDoneCourseIds)-1] // Truncate slice.
			}
		}
	}

	if len(haventDoneCourseIds) > 0 {
		return nil, errors.New("haven't done all course")
	}

	student.Status = "graduated"

	return cc.Repository.UpdateStudent(student)
}

func GetTranscript(c router.Context) (interface{}, error) {
	cc := NewContext(c)
	id := c.ParamString("id")

	return cc.Repository.GetTranscript(id)
}
