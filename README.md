# payutils-http-iris

[Iris](https://github.com/kataras/iris) HTTP driver for
[payutils](https://go.gh.ink/payutils/v3).

Adapts an Iris router so payutils can register provider callback routes on it.
Standard-library `http.HandlerFunc`s are bridged to Iris via `iris.FromStd`.

## Module

```
go.gh.ink/payutils/http/iris/v3
```

```bash
go get go.gh.ink/payutils/http/iris/v3
```

## Usage

Blank-import the driver (it self-registers under the name `iris`) and pass an
Iris router (`*iris.Application` or any party) in `Config.Instances`:

```go
import (
    "github.com/kataras/iris/v12"

    "go.gh.ink/payutils/v3/client"
    "go.gh.ink/payutils/v3/model"

	httpIris"go.gh.ink/payutils/http/iris/v3"
    _ "go.gh.ink/payutils/pay/alipay/v3"
)

app := iris.New()

c, err := client.NewClient(model.Config{
    Endpoint:    "https://api.example.com",
    Instances:   model.I{httpIris.Name: app}, // *iris.Application or a party
    Credentials: model.C{ /* ... */ },
    Contract:    myContract{},
})
```

payutils then registers `POST /{provider}/callback` on the router for each
configured provider.

## Accepted instance

Anything implementing `iris.Party` — `*iris.Application` and sub-parties
(`app.Party("/pay")`). Passing an unsupported value makes `NewInstance` return
`errors.ErrUnsupportedInstance`.

## Supported verbs

`Get` · `Post` · `Put` · `Patch` · `Delete` · `Head` · `Options` · `Any`.

## License

See [LICENSE](LICENSE).
