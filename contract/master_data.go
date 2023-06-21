package contract

import (
	"github.com/herbertabdillah/skripsi-contract-new/state"
	"github.com/hyperledger-labs/cckit/router"
)

func GetFaculty(c router.Context) (interface{}, error) {
	var id = c.ParamString("id")

	return c.State().Get(id, &state.Faculty{})
}

func CreateFaculty(c router.Context) (interface{}, error) {
	var id = c.ParamString("id")
	var name = c.ParamString("name")
	var faculty = &state.Faculty{Name: name, Id: id}

	return faculty, c.State().Insert(id, faculty)
}
