//go:build !windows

/*
   Copyright The containerd Authors.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package platforms

import (
	specs "github.com/opencontainers/image-spec/specs-go/v1"
)

// NewMatcher returns the default Matcher for containerd
func newDefaultMatcher(platform specs.Platform) Matcher {
	m := &matcher{
		Platform: Normalize(platform),
	}

	p := mustReadConfig()
	m.Platform.Features = p.Features
	m.Platform.Compatibilities = p.Compatibilities

	if fs := m.Platform.Features; len(fs) > 0 {
		m.featuresSet = make(map[string]bool, len(fs))
		for _, f := range fs {
			m.featuresSet[f] = true
		}
	}

	return m
}

func GetWindowsOsVersion() string {
	return ""
}
