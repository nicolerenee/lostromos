// Copyright 2017 the lostromos Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package version

import "go.uber.org/zap"

func ExamplePrint() {
	Version = "1"
	GitHash = "abc123"
	BuildTime = "Some point in time"

	Print(zap.NewExample().Sugar())
	// Output:
	// {"level":"info","msg":"version info","version":"1","gitCommitHash":"abc123","buildTime":"Some point in time"}
}
