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
	"bytes"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/intelsdi-x/snap/control/plugin"
	"github.com/intelsdi-x/snap/control/plugin/cpolicy"
)

const (
	// Name of plugin
	name = "psutil"
	// Version of plugin
	version = 6
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
	metrics := make([]plugin.MetricType, len(mts))
	loadre := regexp.MustCompile(`^/intel/psutil/load/load[1,5,15]`)
	cpure := regexp.MustCompile(`^/intel/psutil/cpu.*/.*`)
	memre := regexp.MustCompile(`^/intel/psutil/vm/.*`)
	netre := regexp.MustCompile(`^/intel/psutil/net/.*`)

	for i, p := range mts {
		switch {
		case loadre.MatchString(p.Namespace().String()):
			metric, err := loadAvg(p.Namespace())
			if err != nil {
				return nil, err
			}
			metrics[i] = *metric
		case cpure.MatchString(p.Namespace().String()):
			metric, err := cpuTimes(p.Namespace())
			if err != nil {
				return nil, err
			}
			metrics[i] = *metric
		case memre.MatchString(p.Namespace().String()):
			metric, err := virtualMemory(p.Namespace())
			if err != nil {
				return nil, err
			}
			metrics[i] = *metric
		case netre.MatchString(p.Namespace().String()):
			metric, err := netIOCounters(p.Namespace())
			if err != nil {
				return nil, err
			}
			metrics[i] = *metric
		}
		metrics[i].Timestamp_ = time.Now()

	}
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

func joinNamespace(ns []string) string {
	return "/" + strings.Join(ns, "/")
}

func prettyPrint(mts []plugin.MetricType) error {
	var out bytes.Buffer
	mtsb, _, _ := plugin.MarshalMetricTypes(plugin.SnapJSONContentType, mts)
	if err := json.Indent(&out, mtsb, "", "  "); err != nil {
		return err
	}
	fmt.Println(out.String())
	return nil
}
