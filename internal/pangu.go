package internal

import (
	"github.com/pangum/pangu"
)

func init() {
	pangu.New().Dependencies(
		NewClient,
	)
}
