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
	"fmt"
	"regexp"
	"runtime"

	"github.com/intelsdi-x/snap/control/plugin"
	"github.com/intelsdi-x/snap/core"
	"github.com/shirou/gopsutil/cpu"
)

var cpuLabels = []string{"user", "system", "idle", "nice", "iowait",
	"irq", "softirq", "steal", "guest", "guest_nice", "stolen"}

func cpuTimes(ns core.Namespace) (*plugin.MetricType, error) {
	cpus, err := cpu.Times(true)
	if err != nil {
		return nil, err
	}

	for _, cpu := range cpus {
		switch {
		case regexp.MustCompile(`^/intel/psutil/cpu.*/user`).MatchString(ns.String()):
			return &plugin.MetricType{
				Namespace_: ns,
				Data_:      cpu.User,
			}, nil
		case regexp.MustCompile(`^/intel/psutil/cpu.*/system`).MatchString(ns.String()):
			return &plugin.MetricType{
				Namespace_: ns,
				Data_:      cpu.System,
			}, nil
		case regexp.MustCompile(`^/intel/psutil/cpu.*/idle`).MatchString(ns.String()):
			return &plugin.MetricType{
				Namespace_: ns,
				Data_:      cpu.Idle,
			}, nil
		case regexp.MustCompile(`^/intel/psutil/cpu.*/nice`).MatchString(ns.String()):
			return &plugin.MetricType{
				Namespace_: ns,
				Data_:      cpu.Nice,
			}, nil
		case regexp.MustCompile(`^/intel/psutil/cpu.*/iowait`).MatchString(ns.String()):
			return &plugin.MetricType{
				Namespace_: ns,
				Data_:      cpu.Iowait,
			}, nil
		case regexp.MustCompile(`^/intel/psutil/cpu.*/irq`).MatchString(ns.String()):
			return &plugin.MetricType{
				Namespace_: ns,
				Data_:      cpu.Irq,
			}, nil
		case regexp.MustCompile(`^/intel/psutil/cpu.*/softirq`).MatchString(ns.String()):
			return &plugin.MetricType{
				Namespace_: ns,
				Data_:      cpu.Softirq,
			}, nil
		case regexp.MustCompile(`^/intel/psutil/cpu.*/steal`).MatchString(ns.String()):
			return &plugin.MetricType{
				Namespace_: ns,
				Data_:      cpu.Steal,
			}, nil
		case regexp.MustCompile(`^/intel/psutil/cpu.*/guest`).MatchString(ns.String()):
			return &plugin.MetricType{
				Namespace_: ns,
				Data_:      cpu.Guest,
			}, nil
		case regexp.MustCompile(`^/intel/psutil/cpu.*/guest_nice`).MatchString(ns.String()):
			return &plugin.MetricType{
				Namespace_: ns,
				Data_:      cpu.GuestNice,
			}, nil
		case regexp.MustCompile(`^/intel/psutil/cpu.*/stolen`).MatchString(ns.String()):
			return &plugin.MetricType{
				Namespace_: ns,
				Data_:      cpu.Stolen,
			}, nil
		}

	}

	return nil, fmt.Errorf("Unknown error processing %v", ns)
}

func getCPUTimesMetricTypes() ([]plugin.MetricType, error) {
	//passing true to CPUTimes indicates per CPU
	//CPUTimes does not currently work on OSX https://github.com/shirou/gopsutil/issues/31
	mts := make([]plugin.PluginMetricType, 0)
	switch runtime.GOOS {
	case "linux":
		c, err := cpu.CPUTimes(true)
		if err != nil {
			return nil, err
		}
		for _, i := range c {
			for _, label := range cpuLabels {
				mts = append(mts, plugin.MetricType{Namespace_: core.NewNamespace([]string{"intel", "psutil", i.CPU, label})})
			}
		}
	case "windows":
		_, err := cpu.CPUTimes(true)
		if err != nil {
			return nil, err
		}

		for _, label := range []string{"idle", "system", "user"} {
			mts = append(mts, plugin.PluginMetricType{Namespace_: []string{"intel", "psutil", "cpu", label}})
		}

	}
	return mts, nil
}
