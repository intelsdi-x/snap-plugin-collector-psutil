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

	"github.com/intelsdi-x/snap/control/plugin"
	"github.com/intelsdi-x/snap/core"
	"github.com/shirou/gopsutil/load"
)

func loadAvg(ns core.Namespace) (*plugin.MetricType, error) {
	load, err := load.Avg()
	if err != nil {
		return nil, err
	}

	switch ns.String() {
	case "/intel/psutil/load/load1":
		return &plugin.MetricType{
			Namespace_: ns,
			Data_:      load.Load1,
			Unit_:      "Load/1M",
		}, nil
	case "/intel/psutil/load/load5":
		return &plugin.MetricType{
			Namespace_: ns,
			Data_:      load.Load5,
			Unit_:      "Load/5M",
		}, nil
	case "/intel/psutil/load/load15":
		return &plugin.MetricType{
			Namespace_: ns,
			Data_:      load.Load15,
			Unit_:      "Load/15M",
		}, nil
	}

	return nil, fmt.Errorf("Unknown error processing %v", ns)
}

func getLoadAvgMetricTypes() []plugin.MetricType {
	t := []int{1, 5, 15}
	mts := make([]plugin.MetricType, len(t))
	for i, te := range t {
		mts[i] = plugin.MetricType{
			Namespace_: core.NewNamespace("intel", "psutil", "load", fmt.Sprintf("load%d", te)),
			Unit_:      fmt.Sprintf("Load/%dM", te),
		}
	}
	return mts
}
