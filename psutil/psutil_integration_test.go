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

	. "github.com/smartystreets/goconvey/convey"

	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
)

func TestPsutilCollectMetrics(t *testing.T) {
	Convey("psutil collector", t, func() {
		p := &Psutil{}
		Convey("collect metrics", func() {
			config := plugin.Config{
				"mount_points": "/|/dev|/run",
			}
			mts := []plugin.Metric{
				plugin.Metric{
					Namespace: plugin.NewNamespace("intel", "psutil", "load", "load1"),
					Config:    config,
				},
				plugin.Metric{
					Namespace: plugin.NewNamespace("intel", "psutil", "load", "load5"),
					Config:    config,
				},
				plugin.Metric{
					Namespace: plugin.NewNamespace("intel", "psutil", "load", "load15"),
					Config:    config,
				},
				plugin.Metric{
					Namespace: plugin.NewNamespace("intel", "psutil", "vm", "total"),
					Config:    config,
				},
				plugin.Metric{
					Namespace: plugin.NewNamespace("intel", "psutil", "disk", "used"),
					Config:    config,
				},
			}
			if runtime.GOOS != "darwin" {
				mts = append(mts, plugin.Metric{
					Namespace: plugin.NewNamespace("intel", "psutil", "cpu", "cpu0", "user"),
				})
			}
			metrics, err := p.CollectMetrics(mts)
			So(err, ShouldBeNil)
			So(metrics, ShouldNotBeNil)
		})
		Convey("get metric types", func() {
			mts, err := p.GetMetricTypes(plugin.Config{})
			So(err, ShouldBeNil)
			So(mts, ShouldNotBeNil)
		})

	})
}
