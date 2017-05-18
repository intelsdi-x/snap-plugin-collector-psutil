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
	"strings"

	"time"

	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"

	log "github.com/Sirupsen/logrus"
)

func NewPsutilCollector() *Psutil {
	return &Psutil{}
}

type Psutil struct {
}

// CollectMetrics returns metrics from gopsutil
func (p *Psutil) CollectMetrics(mts []plugin.Metric) ([]plugin.Metric, error) {
	loadReqs := []plugin.Namespace{}
	cpuReqs := []plugin.Namespace{}
	memReqs := []plugin.Namespace{}
	netReqs := []plugin.Namespace{}
	diskReqs := []plugin.Namespace{}

	for _, m := range mts {
		ns := m.Namespace
		switch ns[2].Value {
		case "load":
			loadReqs = append(loadReqs, ns)
		case "cpu":
			cpuReqs = append(cpuReqs, ns)
		case "vm":
			memReqs = append(memReqs, ns)
		case "net":
			netReqs = append(netReqs, ns)
		case "disk":
			diskReqs = append(diskReqs, ns)
		default:
			return nil, fmt.Errorf("Requested metric %s does not match any known psutil metric", m.Namespace.String())
		}
	}

	metrics := []plugin.Metric{}

	loadMts, err := loadAvg(loadReqs)
	if err != nil {
		return nil, err
	}
	metrics = append(metrics, loadMts...)

	cpuMts, err := cpuTimes(cpuReqs)
	if err != nil {
		return nil, err
	}
	metrics = append(metrics, cpuMts...)

	memMts, err := virtualMemory(memReqs)
	if err != nil {
		return nil, err
	}
	metrics = append(metrics, memMts...)

	netMts, err := netIOCounters(netReqs)
	if err != nil {
		return nil, err
	}
	metrics = append(metrics, netMts...)
	mounts := getMountpoints(mts[0].Config)
	diskMts, err := getDiskUsageMetrics(diskReqs, mounts)
	if err != nil {
		return nil, err
	}
	metrics = append(metrics, diskMts...)

	return metrics, nil
}

// GetMetricTypes returns the metric types exposed by gopsutil
func (p *Psutil) GetMetricTypes(_ plugin.Config) ([]plugin.Metric, error) {
	mts := []plugin.Metric{}

	mts = append(mts, getLoadAvgMetricTypes()...)
	mts_, err := getCPUTimesMetricTypes()
	if err != nil {
		return nil, err
	}
	mts = append(mts, mts_...)
	mts = append(mts, getVirtualMemoryMetricTypes()...)

	mts_, err = getNetIOCounterMetricTypes()
	if err != nil {
		return nil, err
	}
	mts = append(mts, mts_...)
	mts = append(mts, getDiskUsageMetricTypes()...)

	return mts, nil
}

//GetConfigPolicy returns a ConfigPolicy
func (p *Psutil) GetConfigPolicy() (plugin.ConfigPolicy, error) {
	c := plugin.NewConfigPolicy()
	c.AddNewStringRule([]string{"intel", "psutil", "disk"},
		"mount_points", false)
	return *c, nil
}

func getMountpoints(cfg plugin.Config) []string {
	if mp, err := cfg.GetString("mount_points"); err != nil {
		if mp == "*" {
			return []string{"all"}
		}
		mountPoints := strings.Split(mp, "|")
		return mountPoints
	}
	return []string{"physical"}
}

type label struct {
	description string
	unit        string
}

func timeSpent(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Debugf("%s took %s", name, elapsed)
}
