package psutil

import (
	"fmt"

	"github.com/intelsdi-x/pulse/control/plugin"
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
