package swagger

import (
	_ "embed"

	"github.com/swaggo/swag"
)

//go:embed person-service.yaml
var swDoc string

type s struct{}

func (s *s) ReadDoc() string {
	return swDoc
}

func init() {
	swag.Register(swag.Name, &s{})
}
