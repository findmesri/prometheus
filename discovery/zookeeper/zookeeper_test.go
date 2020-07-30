// Copyright 2018 The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package zookeeper

import (
	"testing"
	"time"

	"github.com/prometheus/common/model"
	"go.uber.org/goleak"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
}

func TestNewDiscoveryError(t *testing.T) {
	_, err := NewDiscovery(
		[]string{"unreachable.test"},
		time.Second, []string{"/"},
		nil,
		func(data []byte, path string) (model.LabelSet, error) { return nil, nil })
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}
