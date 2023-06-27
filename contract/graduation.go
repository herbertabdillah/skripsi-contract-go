package contract

import (
	"github.com/hyperledger-labs/cckit/router"
)

func Graduate(c router.Context) (interface{}, error) {
	cc := NewContext(c)
	id := c.ParamString("id")

	return cc.Service.Graduate(id)
}

func GetTranscript(c router.Context) (interface{}, error) {
	cc := NewContext(c)
	id := c.ParamString("id")

	return cc.Repository.GetTranscript(id)
}
