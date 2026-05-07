package hooks

import (
	dynatrace "github.com/Dynatrace/libbuildpack-dynatrace"
	"github.com/cloudfoundry/libbuildpack"
)

func init() {
	libbuildpack.AddHook(dynatrace.NewHook("dotnet", "process"))
}
