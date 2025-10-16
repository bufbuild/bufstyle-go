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

package analyzers

import (
	"slices"

	"buf.build/go/bufstyle/internal/analyzers/america"
	"buf.build/go/bufstyle/internal/analyzers/casing"
	"buf.build/go/bufstyle/internal/analyzers/packagefilename"
	"buf.build/go/bufstyle/internal/analyzers/typeban"
	"golang.org/x/tools/go/analysis"
)

// New returns all Analyzers.
//
// We don't store this as a global because we modify these.
func New() []*analysis.Analyzer {
	return slices.Concat(
		america.New(),
		casing.New(),
		packagefilename.New(),
		typeban.New(),
	)
}
