package dynaml

import (
	"github.com/cloudfoundry-incubator/spiff/yaml"
)

type Binding interface {
	FindFromRoot([]string) (yaml.Node, bool)
	FindReference([]string) (yaml.Node, bool)
	FindInStubs([]string) (yaml.Node, bool)
}

type Expression interface {
	Evaluate(Binding) (yaml.Node, bool)
}
