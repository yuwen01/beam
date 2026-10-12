// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package harnessopts

import (
	"fmt"
	"strconv"

	"github.com/apache/beam/sdks/v2/go/pkg/beam/core/util/hooks"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/log"
)

const (
	minLogSeverityHook = "beam:go:hook:logging:minseverity"
)

// MinLogSeverity sets the lowest severity of log entries that SDK workers send
// to the runner. Entries below it are dropped on the worker, so runner logging
// backends such as Cloud Logging never receive them. By default every entry is
// sent.
func MinLogSeverity(sev log.Severity) error {
	if sev < log.SevDebug || sev > log.SevFatal {
		return fmt.Errorf("invalid log severity %v, must be between %v and %v", sev, log.SevDebug, log.SevFatal)
	}
	// The hook itself is defined in beam/core/runtime/harness/logging.go
	return hooks.EnableHook(minLogSeverityHook, strconv.Itoa(int(sev)))
}
