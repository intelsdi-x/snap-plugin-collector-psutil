//
// +build medium

/*
http://www.apache.org/licenses/LICENSE-2.0.txt

Copyright 2015 Intel Corporation

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

package psutil

import (
	"runtime"
	"testing"

	"github.com/intelsdi-x/snap/control/plugin"
	"github.com/intelsdi-x/snap/core"
	. "github.com/smartystreets/goconvey/convey"
)

func TestPsutilCollectMetrics(t *testing.T) {
	Convey("psutil collector", t, func() {
		p := &Psutil{}
		Convey("collect metrics", func() {
			mts := []plugin.MetricType{
				plugin.MetricType{
					Namespace_: core.NewNamespace("intel", "psutil", "load", "load1"),
				},
				plugin.MetricType{
					Namespace_: core.NewNamespace("intel", "psutil", "load", "load5"),
				},
				plugin.MetricType{
					Namespace_: core.NewNamespace("intel", "psutil", "load", "load15"),
				},
				plugin.MetricType{
					Namespace_: core.NewNamespace("intel", "psutil", "vm", "total"),
				},
			}
			if runtime.GOOS != "darwin" {
				mts = append(mts, plugin.MetricType{
					Namespace_: core.NewNamespace("intel", "psutil", "cpu0", "user"),
				})
			}
			metrics, err := p.CollectMetrics(mts)
			So(err, ShouldBeNil)
			So(metrics, ShouldNotBeNil)
		})
		Convey("get metric types", func() {
			mts, err := p.GetMetricTypes(plugin.ConfigType{})
			So(err, ShouldBeNil)
			So(mts, ShouldNotBeNil)
		})

	})
}
