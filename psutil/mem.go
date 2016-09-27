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
	"github.com/shirou/gopsutil/mem"
)

func virtualMemory(nss []core.Namespace) ([]plugin.MetricType, error) {
	mem, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}

	results := make([]plugin.MetricType, len(nss))

	for i, ns := range nss {
		var data interface{}

		switch ns.Element(len(ns) - 1).Value {
		case "total":
			data = mem.Total
		case "available":
			data = mem.Available
		case "used":
			data = mem.Used
		case "used_percent":
			data = mem.UsedPercent
		case "free":
			data = mem.Free
		case "active":
			data = mem.Active
		case "inactive":
			data = mem.Inactive
		case "buffers":
			data = mem.Buffers
		case "cached":
			data = mem.Cached
		case "wired":
			data = mem.Wired
		default:
			return nil, fmt.Errorf("Requested memory statistic %s is not found", ns.String())
		}

		results[i] = plugin.MetricType{
			Namespace_: ns,
			Data_:      data,
			Unit_:      "B",
			Timestamp_: time.Now(),
		}
	}

	return results, nil
}

func getVirtualMemoryMetricTypes() []plugin.MetricType {
	return []plugin.MetricType{
		plugin.MetricType{
			Namespace_:   core.NewNamespace("intel", "psutil", "vm", "total"),
			Unit_:        "B",
			Description_: "total swap memory in bytes",
		},
		plugin.MetricType{
			Namespace_: core.NewNamespace("intel", "psutil", "vm", "available"),
			Unit_:      "B",
			Description_: `the actual amount of available memory that can be 
			given instantly to processes that request more memory in bytes; 
			this is calculated by summing different memory values depending 
			on the platform (e.g. free + buffers + cached on Linux) and it 
			is supposed to be used to monitor actual memory usage in a cross 
			platform fashion.`,
		},
		plugin.MetricType{
			Namespace_: core.NewNamespace("intel", "psutil", "vm", "used"),
			Unit_:      "B",
			Description_: `Memory used is calculated differently depending on 
			the platform and designed for informational purposes only.`,
		},
		plugin.MetricType{
			Namespace_: core.NewNamespace("intel", "psutil", "vm", "used_percent"),
			Unit_:      "B",
			Description_: `the percentage usage calculated as (total - available) 
			/ total * 100.`,
		},
		plugin.MetricType{
			Namespace_: core.NewNamespace("intel", "psutil", "vm", "free"),
			Description_: `memory not being used at all (zeroed) that is readily 
			available; note that this doesn’t reflect the actual memory available 
			(use ‘available’ instead).`,
			Unit_: "B",
		},
		plugin.MetricType{
			Namespace_: core.NewNamespace("intel", "psutil", "vm", "active"),
			Description_: `(UNIX): memory currently in use or very recently used, 
			and so it is in RAM.`,
			Unit_: "B",
		},
		plugin.MetricType{
			Namespace_:   core.NewNamespace("intel", "psutil", "vm", "inactive"),
			Description_: `(UNIX): memory that is marked as not used.`,
			Unit_:        "B",
		},
		plugin.MetricType{
			Namespace_:   core.NewNamespace("intel", "psutil", "vm", "buffers"),
			Description_: `(Linux, BSD): cache for things like file system metadata.`,
			Unit_:        "B",
		},
		plugin.MetricType{
			Namespace_:   core.NewNamespace("intel", "psutil", "vm", "cached"),
			Description_: `(Linux, BSD): cache for various things.`,
			Unit_:        "B",
		},
		plugin.MetricType{
			Namespace_: core.NewNamespace("intel", "psutil", "vm", "wired"),
			Description_: `(BSD, OSX): memory that is marked to always stay in RAM. 
			It is never moved to disk.`,
			Unit_: "B",
		},
	}
}
