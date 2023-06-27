package contract

import (
	r "github.com/herbertabdillah/skripsi-contract-new/repository"
	s "github.com/herbertabdillah/skripsi-contract-new/service"
	"github.com/hyperledger-labs/cckit/router"
)

type Context struct {
	Service    s.Service
	Repository r.Repository
}

func NewContext(rc router.Context) Context {
	c := Context{}
	c.Repository = r.NewRepository(rc)
	c.Service = s.NewService(rc, c.Repository)

	return c
}
