//go:build tools

package tools

// Add any Go tools as _ imports and they will be installed as part of `make install-tools`
// https://marcofranssen.nl/manage-go-tools-via-go-modules

import (
	_ "github.com/boumenot/gocover-cobertura"
	_ "github.com/fullstorydev/grpcui/cmd/grpcui"
	_ "github.com/fzipp/gocyclo/cmd/gocyclo"
	_ "github.com/golang/mock/mockgen"
	_ "github.com/golang/protobuf/protoc-gen-go"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/securego/gosec/v2/cmd/gosec"
	_ "github.com/spf13/cobra/cobra"
	_ "golang.org/x/lint/golint"
	_ "golang.org/x/perf/cmd/benchstat"
	_ "golang.org/x/tools/cmd/stringer"
)
