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

var cpuLabels = map[string]label{
	"user": label{
		description: "",
		unit:        "",
	},
	"system": label{
		description: "",
		unit:        "",
	},
	"idle": label{
		description: "",
		unit:        "",
	},
	"nice": label{
		description: "",
		unit:        "",
	},
	"iowait": label{
		description: "",
		unit:        "",
	},
	"irq": label{
		description: "",
		unit:        "",
	},
	"softirq": label{
		description: "",
		unit:        "",
	},
	"steal": label{
		description: "",
		unit:        "",
	},
	"guest": label{
		description: "",
		unit:        "",
	},
	"guest_nice": label{
		description: "",
		unit:        "",
	},
	"stolen": label{
		description: "",
		unit:        "",
	},
}

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
				Unit_:      cpuLabels["cpu"].unit,
			}, nil
		case regexp.MustCompile(`^/intel/psutil/cpu.*/system`).MatchString(ns.String()):
			return &plugin.MetricType{
				Namespace_: ns,
				Data_:      cpu.System,
				Unit_:      cpuLabels["system"].unit,
			}, nil
		case regexp.MustCompile(`^/intel/psutil/cpu.*/idle`).MatchString(ns.String()):
			return &plugin.MetricType{
				Namespace_: ns,
				Unit_:      cpuLabels["idle"].unit,
			}, nil
		case regexp.MustCompile(`^/intel/psutil/cpu.*/nice`).MatchString(ns.String()):
			return &plugin.MetricType{
				Namespace_: ns,
				Data_:      cpu.Nice,
				Unit_:      cpuLabels["nice"].unit,
			}, nil
		case regexp.MustCompile(`^/intel/psutil/cpu.*/iowait`).MatchString(ns.String()):
			return &plugin.MetricType{
				Namespace_: ns,
				Data_:      cpu.Iowait,
				Unit_:      cpuLabels["iowait"].unit,
			}, nil
		case regexp.MustCompile(`^/intel/psutil/cpu.*/irq`).MatchString(ns.String()):
			return &plugin.MetricType{
				Namespace_: ns,
				Data_:      cpu.Irq,
				Unit_:      cpuLabels["irq"].unit,
			}, nil
		case regexp.MustCompile(`^/intel/psutil/cpu.*/softirq`).MatchString(ns.String()):
			return &plugin.MetricType{
				Namespace_: ns,
				Data_:      cpu.Softirq,
				Unit_:      cpuLabels["softirq"].unit,
			}, nil
		case regexp.MustCompile(`^/intel/psutil/cpu.*/steal`).MatchString(ns.String()):
			return &plugin.MetricType{
				Namespace_: ns,
				Data_:      cpu.Steal,
				Unit_:      cpuLabels["steal"].unit,
			}, nil
		case regexp.MustCompile(`^/intel/psutil/cpu.*/guest`).MatchString(ns.String()):
			return &plugin.MetricType{
				Namespace_: ns,
				Data_:      cpu.Guest,
				Unit_:      cpuLabels["guest"].unit,
			}, nil
		case regexp.MustCompile(`^/intel/psutil/cpu.*/guest_nice`).MatchString(ns.String()):
			return &plugin.MetricType{
				Namespace_: ns,
				Data_:      cpu.GuestNice,
				Unit_:      cpuLabels["guest_nice"].unit,
			}, nil
		case regexp.MustCompile(`^/intel/psutil/cpu.*/stolen`).MatchString(ns.String()):
			return &plugin.MetricType{
				Namespace_: ns,
				Data_:      cpu.Stolen,
				Unit_:      cpuLabels["stolen"].unit,
			}, nil
		}

	}

	return nil, fmt.Errorf("Unknown error processing %v", ns)
}

func getCPUTimesMetricTypes() ([]plugin.MetricType, error) {
	//passing true to CPUTimes indicates per CPU
	//CPUTimes does not currently work on OSX https://github.com/shirou/gopsutil/issues/31
	mts := []plugin.MetricType{}
	switch runtime.GOOS {
	case "linux":
		c, err := cpu.Times(true)
		if err != nil {
			return nil, err
		}
		for _, i := range c {
			for k, label := range cpuLabels {
				mts = append(mts, plugin.MetricType{
					Namespace_:   core.NewNamespace("intel", "psutil", i.CPU, k),
					Description_: label.description,
					Unit_:        label.unit,
				})
			}
		}
	case "windows":
		_, err := cpu.Times(true)
		if err != nil {
			return nil, err
		}

		for _, label := range []string{"idle", "system", "user"} {
			mts = append(mts, plugin.MetricType{
				Namespace_:   core.NewNamespace("intel", "psutil", "cpu", label),
				Description_: cpuLabels[label].description,
				Unit_:        cpuLabels[label].unit,
			})
		}

	}
	return mts, nil
}
