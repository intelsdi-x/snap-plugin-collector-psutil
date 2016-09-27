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
	"time"

	"github.com/intelsdi-x/snap/control/plugin"
	"github.com/intelsdi-x/snap/core"
	"github.com/shirou/gopsutil/load"
)

func loadAvg(nss []core.Namespace) ([]plugin.MetricType, error) {
	load, err := load.Avg()
	if err != nil {
		return nil, err
	}

	results := make([]plugin.MetricType, len(nss))

	for i, ns := range nss {
		switch ns.Element(len(ns) - 1).Value {
		case "load1":
			results[i] = plugin.MetricType{
				Namespace_: ns,
				Data_:      load.Load1,
				Unit_:      "Load/1M",
				Timestamp_: time.Now(),
			}
		case "load5":
			results[i] = plugin.MetricType{
				Namespace_: ns,
				Data_:      load.Load5,
				Unit_:      "Load/5M",
				Timestamp_: time.Now(),
			}
		case "load15":
			results[i] = plugin.MetricType{
				Namespace_: ns,
				Data_:      load.Load15,
				Unit_:      "Load/15M",
				Timestamp_: time.Now(),
			}
		default:
			return nil, fmt.Errorf("Requested load statistic %s is not found", ns.Element(len(ns)-1).Value)
		}
	}

	return results, nil
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
