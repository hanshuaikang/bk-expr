package core

import (
	"time"
)

type Function interface {
	Execute(args []string) (interface{}, error)
}

type Now struct {
}

func (n *Now) Execute(args []string) (interface{}, error) {
	return time.Now(), nil
}
