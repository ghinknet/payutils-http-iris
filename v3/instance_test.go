package httpiris

import (
	"net/http"
	"testing"

	stderrors "errors"

	"github.com/kataras/iris/v12"

	"go.gh.ink/payutils/v3/errors"
)

func TestNewInstance_RejectsUnsupported(t *testing.T) {
	_, err := Driver{}.NewInstance(123)
	if !stderrors.Is(err, errors.ErrUnsupportedInstance) {
		t.Errorf("err = %v, want ErrUnsupportedInstance", err)
	}
}

func TestNewInstance_AcceptsParty(t *testing.T) {
	app := iris.New()

	inst, err := Driver{}.NewInstance(app)
	if err != nil {
		t.Fatalf("NewInstance(*iris.Application) error: %v", err)
	}

	// Registering handlers across every verb must not panic.
	noop := func(w http.ResponseWriter, _ *http.Request) {}
	inst.Get("/pay/get", noop)
	inst.Post("/pay/callback", noop)
	inst.Put("/pay/put", noop)
	inst.Patch("/pay/patch", noop)
	inst.Delete("/pay/delete", noop)
	inst.Head("/pay/head", noop)
	inst.Options("/pay/options", noop)
	inst.Any("/pay/any", noop)
}
