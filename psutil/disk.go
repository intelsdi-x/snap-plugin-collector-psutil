///*
//http://www.apache.org/licenses/LICENSE-2.0.txt
//
//Copyright 2017 Intel Corporation
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//*/
//
package psutil

import (
	"strings"
	"time"

	"github.com/shirou/gopsutil/disk"

	"github.com/intelsdi-x/snap/control/plugin"
	"github.com/intelsdi-x/snap/core"
)

func getPSUtilDiskUsage(path string) (*disk.UsageStat, error) {
	disk_usage, err := disk.Usage(path)
	if err != nil {
		return nil, err
	}
	return disk_usage, nil
}

func getDiskUsageMetrics(nss []core.Namespace, mounts []string) ([]plugin.MetricType, error) {
	t := time.Now()
	var paths []disk.PartitionStat
	metrics := []plugin.MetricType{}
	namespaces := map[string][]string{}
	requested := map[string]core.Namespace{}
	for _, ns := range nss {
		namespaces[ns.Strings()[len(ns.Strings())-1]] = ns.Strings()
		requested[ns.Strings()[len(ns.Strings())-1]] = ns
	}
	if strings.Contains(mounts[0], "physical") {
		parts, err := disk.Partitions(false)
		if err != nil {
			return nil, err
		}
		paths = parts
	} else if strings.Contains(mounts[0], "all") {
		parts, err := disk.Partitions(true)
		if err != nil {
			return nil, err
		}
		paths = parts
	} else {
		parts, err := disk.Partitions(true)
		if err != nil {
			return nil, err
		}
		for _, part := range parts {
			for _, mtpoint := range mounts {
				if part.Mountpoint == mtpoint {
					paths = append(paths, part)
				}
			}
		}
	}

	for _, path := range paths {
		data, err := getPSUtilDiskUsage(path.Mountpoint)
		if err != nil {
			return nil, err
		}
		tags := map[string]string{}
		tags["device"] = path.Device
		for _, namespace := range namespaces {
			if strings.Contains(strings.Join(namespace, "|"), "total") {
				nspace := make([]core.NamespaceElement, len(requested["total"]))
				copy(nspace, requested["total"])
				nspace[3].Value = path.Mountpoint
				metrics = append(metrics, plugin.MetricType{
					Namespace_: nspace,
					Data_:      data.Total,
					Tags_:      tags,
					Timestamp_: t,
				})
			}
			if strings.Contains(strings.Join(namespace, "|"), "used") {
				nspace := make([]core.NamespaceElement, len(requested["used"]))
				copy(nspace, requested["used"])
				nspace[3].Value = path.Mountpoint
				metrics = append(metrics, plugin.MetricType{
					Namespace_: nspace,
					Data_:      data.Used,
					Tags_:      tags,
					Timestamp_: t,
				})
			}
			if strings.Contains(strings.Join(namespace, "|"), "free") {
				nspace := make([]core.NamespaceElement, len(requested["free"]))
				copy(nspace, requested["free"])
				nspace[3].Value = path.Mountpoint
				metrics = append(metrics, plugin.MetricType{
					Namespace_: nspace,
					Data_:      data.Free,
					Tags_:      tags,
					Timestamp_: t,
				})
			}
			if strings.Contains(strings.Join(namespace, "|"), "percent") {
				nspace := make([]core.NamespaceElement, len(requested["percent"]))
				copy(nspace, requested["percent"])
				nspace[3].Value = path.Mountpoint
				metrics = append(metrics, plugin.MetricType{
					Namespace_: nspace,
					Data_:      data.UsedPercent,
					Tags_:      tags,
					Timestamp_: t,
				})
			}
		}
	}
	return metrics, nil
}

func getDiskUsageMetricTypes() []plugin.MetricType {
	var mts []plugin.MetricType
	mts = append(mts, plugin.MetricType{
		Namespace_: core.NewNamespace("intel", "psutil", "disk").
			AddDynamicElement("mount_point", "Mount Point").
			AddStaticElement("total"),
	})
	mts = append(mts, plugin.MetricType{
		Namespace_: core.NewNamespace("intel", "psutil", "disk").
			AddDynamicElement("mount_point", "Mount Point").
			AddStaticElement("used"),
	})
	mts = append(mts, plugin.MetricType{
		Namespace_: core.NewNamespace("intel", "psutil", "disk").
			AddDynamicElement("mount_point", "Mount Point").
			AddStaticElement("free"),
	})
	mts = append(mts, plugin.MetricType{
		Namespace_: core.NewNamespace("intel", "psutil", "disk").
			AddDynamicElement("mount_point", "Mount Point").
			AddStaticElement("percent"),
	})
	return mts
}
