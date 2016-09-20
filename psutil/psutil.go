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

	"github.com/intelsdi-x/snap/control/plugin"
	"github.com/intelsdi-x/snap/control/plugin/cpolicy"
	"github.com/intelsdi-x/snap/core"
)

const (
	// Name of plugin
	name = "psutil"
	// Version of plugin
	version = 7
	// Type of plugin
	pluginType = plugin.CollectorPluginType
)

func Meta() *plugin.PluginMeta {
	return plugin.NewPluginMeta(name, version, pluginType, []string{plugin.SnapGOBContentType}, []string{plugin.SnapGOBContentType})
}

func NewPsutilCollector() *Psutil {
	return &Psutil{}
}

type Psutil struct {
}

// CollectMetrics returns metrics from gopsutil
func (p *Psutil) CollectMetrics(mts []plugin.MetricType) ([]plugin.MetricType, error) {
	loadre := regexp.MustCompile(`^/intel/psutil/load/.*`)
	cpure := regexp.MustCompile(`^/intel/psutil/cpu/.*`)
	memre := regexp.MustCompile(`^/intel/psutil/vm/.*`)
	netre := regexp.MustCompile(`^/intel/psutil/net/.*`)

	loadReqs := []core.Namespace{}
	cpuReqs := []core.Namespace{}
	memReqs := []core.Namespace{}
	netReqs := []core.Namespace{}

	for _, p := range mts {
		switch {
		case loadre.MatchString(p.Namespace().String()):
			loadReqs = append(loadReqs, p.Namespace())
		case cpure.MatchString(p.Namespace().String()):
			cpuReqs = append(cpuReqs, p.Namespace())
		case memre.MatchString(p.Namespace().String()):
			memReqs = append(memReqs, p.Namespace())
		case netre.MatchString(p.Namespace().String()):
			netReqs = append(netReqs, p.Namespace())
		default:
			return nil, fmt.Errorf("Requested metric %s does not match any known psutil metric", p.Namespace().String())
		}
	}

	metrics := []plugin.MetricType{}

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

	return metrics, nil
}

// GetMetricTypes returns the metric types exposed by gopsutil
func (p *Psutil) GetMetricTypes(_ plugin.ConfigType) ([]plugin.MetricType, error) {
	mts := []plugin.MetricType{}

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

	return mts, nil
}

//GetConfigPolicy returns a ConfigPolicy
func (p *Psutil) GetConfigPolicy() (*cpolicy.ConfigPolicy, error) {
	c := cpolicy.New()
	return c, nil
}

type label struct {
	description string
	unit        string
}
