package httpiris

import (
	"github.com/kataras/iris/v12"

	"go.gh.ink/payutils/v3/errors"
	"go.gh.ink/payutils/v3/model"
)

type Instance struct {
	Router iris.Party
}

type Driver struct{}

func (d Driver) NewInstance(instance any) (model.HttpInstance, error) {
	// Both *iris.Application and route parties implement iris.Party.
	router, ok := instance.(iris.Party)
	if !ok {
		return Instance{}, errors.ErrUnsupportedInstance
	}
	return Instance{Router: router}, nil
}
