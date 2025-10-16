// Copyright 2025 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package typeban

import (
	"fmt"

	"golang.org/x/tools/go/analysis"
)

// New returns a new set of Analyzers.
func New() []*analysis.Analyzer {
	return []*analysis.Analyzer{
		newFor("NO_SYNC_POOL", "sync.Pool"),
	}
}

func newFor(name string, bannedType string) *analysis.Analyzer {
	return &analysis.Analyzer{
		Name: name,
		Doc:  fmt.Sprintf("Verifies that %s is not used.", bannedType),
		Run: func(pass *analysis.Pass) (any, error) {
			if typesInfo := pass.TypesInfo; typesInfo != nil {
				for expr, typeAndValue := range pass.TypesInfo.Types {
					if t := typeAndValue.Type; t != nil {
						value := t.String()
						if value == bannedType {
							pass.Reportf(expr.Pos(), "%s cannot be used", value)
						}
					}
				}
			}
			return nil, nil
		},
	}
}
