package storage

import (
	"github.com/pangum/pangu"
)

func init() {
	pangu.New().Dependency(newClient)
}