package dynaml

import (
	"fmt"
	"strings"
	"github.com/cloudfoundry-incubator/spiff/yaml"
	"github.com/cloudfoundry-incubator/spiff/debug"
)

type MergeExpr struct {
	Path []string
	Redirect bool
	Replace  bool
	Required bool
	KeyName  string
}

func (e MergeExpr) Evaluate(binding Binding) (yaml.Node, EvaluationInfo, bool) {
	var info EvaluationInfo
	if e.Redirect {
		info.RedirectPath=e.Path
	}
	info.KeyName = e.KeyName
	debug.Debug("/// lookup %v\n",e.Path)
	node, ok := binding.FindInStubs(e.Path)
	if ok {
		info.Replace=e.Replace
		info.Merged=true
	} else {
		info.Issue=fmt.Sprintf("'%s' not found in any stub",strings.Join(e.Path,"."))
	}
	return node, info, ok
}

func (e MergeExpr) String() string {
	rep := ""
	if e.Replace {
		rep = " replace"
	}
	
	if e.KeyName != "" {
		rep += " on "+e.KeyName
	}
	if e.Required && !e.Redirect && rep!="" {
		rep = " required"
	}
	if e.Redirect {
		return "merge" + rep + " " + strings.Join(e.Path, ".")
	}
	return "merge"+rep
}
