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
	"github.com/shirou/gopsutil/mem"
)

func virtualMemory(ns []string) (*plugin.PluginMetricType, error) {
	mem, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}

	switch joinNamespace(ns) {
	case "/intel/psutil/vm/total":
		return &plugin.PluginMetricType{
			Namespace_: ns,
			Data_:      mem.Total,
		}, nil
	case "/intel/psutil/vm/available":
		return &plugin.PluginMetricType{
			Namespace_: ns,
			Data_:      mem.Available,
		}, nil
	case "/intel/psutil/vm/used":
		return &plugin.PluginMetricType{
			Namespace_: ns,
			Data_:      mem.Used,
		}, nil
	case "/intel/psutil/vm/used_percent":
		return &plugin.PluginMetricType{
			Namespace_: ns,
			Data_:      mem.UsedPercent,
		}, nil
	case "/intel/psutil/vm/free":
		return &plugin.PluginMetricType{
			Namespace_: ns,
			Data_:      mem.Free,
		}, nil
	case "/intel/psutil/vm/active":
		return &plugin.PluginMetricType{
			Namespace_: ns,
			Data_:      mem.Active,
		}, nil
	case "/intel/psutil/vm/inactive":
		return &plugin.PluginMetricType{
			Namespace_: ns,
			Data_:      mem.Inactive,
		}, nil
	case "/intel/psutil/vm/buffers":
		return &plugin.PluginMetricType{
			Namespace_: ns,
			Data_:      mem.Buffers,
		}, nil
	case "/intel/psutil/vm/cached":
		return &plugin.PluginMetricType{
			Namespace_: ns,
			Data_:      mem.Cached,
		}, nil
	case "/intel/psutil/vm/wired":
		return &plugin.PluginMetricType{
			Namespace_: ns,
			Data_:      mem.Wired,
		}, nil
	case "/intel/psutil/vm/shared":
		return &plugin.PluginMetricType{
			Namespace_: ns,
			Data_:      mem.Shared,
		}, nil
	}

	return nil, fmt.Errorf("Unknown error processing %v", ns)
}

func getVirtualMemoryMetricTypes() []plugin.PluginMetricType {
	t := []string{"total", "available", "used", "used_percent", "free", "active", "inactive", "buffers", "cached", "wired", "shared"}
	mts := make([]plugin.PluginMetricType, len(t))
	for i, te := range t {
		mts[i] = plugin.PluginMetricType{Namespace_: []string{"intel", "psutil", "vm", te}}
	}
	return mts
}
