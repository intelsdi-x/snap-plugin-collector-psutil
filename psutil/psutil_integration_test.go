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

func TestPsutilGetMetricTypes(t *testing.T) {
	Convey("psutil collector", t, func() {
		p := NewPsutilCollector()
		Convey("get metric types", func() {
			config := plugin.Config{}

			metric_types, err := p.GetMetricTypes(config)
			So(err, ShouldBeNil)
			So(metric_types, ShouldNotBeNil)
			So(metric_types, ShouldNotBeEmpty)
			//55 collectable metrics
			So(len(metric_types), ShouldEqual, 55)
		})
	})

}

func TestPsutilCollectMetrics(t *testing.T) {
	Convey("psutil collector", t, func() {
		p := NewPsutilCollector()
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
					Namespace: plugin.NewNamespace("intel", "psutil", "vm", "available"),
					Config:    config,
				},
				plugin.Metric{
					Namespace: plugin.NewNamespace("intel", "psutil", "vm", "used"),
					Config:    config,
				},
				plugin.Metric{
					Namespace: plugin.NewNamespace("intel", "psutil", "vm", "used_percent"),
					Config:    config,
				},
				plugin.Metric{
					Namespace: plugin.NewNamespace("intel", "psutil", "vm", "free"),
					Config:    config,
				},
				plugin.Metric{
					Namespace: plugin.NewNamespace("intel", "psutil", "vm", "active"),
					Config:    config,
				},
				plugin.Metric{
					Namespace: plugin.NewNamespace("intel", "psutil", "vm", "inactive"),
					Config:    config,
				},
				plugin.Metric{
					Namespace: plugin.NewNamespace("intel", "psutil", "vm", "buffers"),
					Config:    config,
				},
				plugin.Metric{
					Namespace: plugin.NewNamespace("intel", "psutil", "vm", "cached"),
					Config:    config,
				},
				plugin.Metric{
					Namespace: plugin.NewNamespace("intel", "psutil", "vm", "wired"),
					Config:    config,
				},
				plugin.Metric{
					Namespace: plugin.NewNamespace("intel", "psutil", "disk", "used"),
					Config:    config,
				},
				plugin.Metric{
					Namespace: plugin.NewNamespace("intel", "psutil", "cpu", "*", "nice"),
					Config:    config,
				},
				plugin.Metric{
					Namespace: plugin.NewNamespace("intel", "psutil", "cpu", "*", "system"),
					Config:    config,
				},
				plugin.Metric{
					Namespace: plugin.NewNamespace("intel", "psutil", "cpu", "*", "iowait"),
					Config:    config,
				},
				plugin.Metric{
					Namespace: plugin.NewNamespace("intel", "psutil", "cpu", "*", "guest"),
					Config:    config,
				},
				plugin.Metric{
					Namespace: plugin.NewNamespace("intel", "psutil", "cpu", "*", "stolen"),
					Config:    config,
				},
				plugin.Metric{
					Namespace: plugin.NewNamespace("intel", "psutil", "cpu", "*", "idle"),
					Config:    config,
				},
				plugin.Metric{
					Namespace: plugin.NewNamespace("intel", "psutil", "cpu", "*", "guest_nice"),
					Config:    config,
				},
				plugin.Metric{
					Namespace: plugin.NewNamespace("intel", "psutil", "cpu", "*", "irq"),
					Config:    config,
				},
				plugin.Metric{
					Namespace: plugin.NewNamespace("intel", "psutil", "cpu", "*", "softirq"),
					Config:    config,
				},
				plugin.Metric{
					Namespace: plugin.NewNamespace("intel", "psutil", "cpu", "*", "steal"),
					Config:    config,
				},

				plugin.Metric{
					Namespace: plugin.NewNamespace("intel", "psutil", "net", "*", "bytes_sent"),
					Config:    config,
				},
				plugin.Metric{
					Namespace: plugin.NewNamespace("intel", "psutil", "net", "*", "bytes_recv"),
					Config:    config,
				},
				plugin.Metric{
					Namespace: plugin.NewNamespace("intel", "psutil", "disk", "*", "total"),
					Config:    config,
				},
				plugin.Metric{
					Namespace: plugin.NewNamespace("intel", "psutil", "disk", "*", "used"),
					Config:    config,
				},
				plugin.Metric{
					Namespace: plugin.NewNamespace("intel", "psutil", "disk", "*", "free"),
					Config:    config,
				},
				plugin.Metric{
					Namespace: plugin.NewNamespace("intel", "psutil", "disk", "*", "percent"),
					Config:    config,
				},
				plugin.Metric{
					Namespace: plugin.NewNamespace("intel", "psutil", "disk", "*", "used"),
					Config:    config,
				},
				plugin.Metric{
					Namespace: plugin.NewNamespace("intel", "psutil", "net", "*", "bytes_recv"),
					Config:    config,
				},

				plugin.Metric{
					Namespace: plugin.NewNamespace("intel", "psutil", "net", "*", "packets_sent"),
					Config:    config,
				},
				plugin.Metric{
					Namespace: plugin.NewNamespace("intel", "psutil", "net", "*", "packets_recv"),
					Config:    config,
				},
				plugin.Metric{
					Namespace: plugin.NewNamespace("intel", "psutil", "net", "*", "errin"),
					Config:    config,
				},
				plugin.Metric{
					Namespace: plugin.NewNamespace("intel", "psutil", "net", "*", "errout"),
					Config:    config,
				},
				plugin.Metric{
					Namespace: plugin.NewNamespace("intel", "psutil", "net", "lo", "dropin"),
					Config:    config,
				},
				plugin.Metric{
					Namespace: plugin.NewNamespace("intel", "psutil", "net", "all", "dropout"),
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

func TestPSUtilPlugin(t *testing.T) {
	Convey("Create PSUtil Collector", t, func() {
		psCol := NewPsutilCollector()
		Convey("So psCol should not be nil", func() {
			So(psCol, ShouldNotBeNil)
		})
		Convey("So psCol should be of Psutil type", func() {
			So(psCol, ShouldHaveSameTypeAs, &Psutil{})
		})
		Convey("psCol.GetConfigPolicy() should return a config policy", func() {
			configPolicy, _ := psCol.GetConfigPolicy()
			Convey("So config policy should not be nil", func() {
				So(configPolicy, ShouldNotBeNil)
			})
			Convey("So config policy should be a plugin.ConfigPolicy", func() {
				So(configPolicy, ShouldHaveSameTypeAs, plugin.ConfigPolicy{})
			})
		})
	})
}
