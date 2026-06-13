package httpiris

import (
	"net/http"

	"github.com/kataras/iris/v12"
)

func (i Instance) Get(path string, handler http.HandlerFunc) {
	i.Router.Get(path, iris.FromStd(handler))
}

func (i Instance) Post(path string, handler http.HandlerFunc) {
	i.Router.Post(path, iris.FromStd(handler))
}

func (i Instance) Put(path string, handler http.HandlerFunc) {
	i.Router.Put(path, iris.FromStd(handler))
}

func (i Instance) Patch(path string, handler http.HandlerFunc) {
	i.Router.Patch(path, iris.FromStd(handler))
}

func (i Instance) Delete(path string, handler http.HandlerFunc) {
	i.Router.Delete(path, iris.FromStd(handler))
}

func (i Instance) Head(path string, handler http.HandlerFunc) {
	i.Router.Head(path, iris.FromStd(handler))
}

func (i Instance) Options(path string, handler http.HandlerFunc) {
	i.Router.Options(path, iris.FromStd(handler))
}

func (i Instance) Any(path string, handler http.HandlerFunc) {
	i.Router.Any(path, iris.FromStd(handler))
}
