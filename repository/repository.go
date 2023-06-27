package repository

import (
	"github.com/hyperledger-labs/cckit/router"
)

type Repository struct {
	context router.Context
}

func NewRepository(c router.Context) Repository {
	return Repository{context: c}
}
