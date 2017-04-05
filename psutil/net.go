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
	"github.com/shirou/gopsutil/net"
)

var netIOCounterLabels = map[string]label{
	"bytes_sent": label{
		unit:        "",
		description: "",
	},
	"bytes_recv": label{
		unit:        "",
		description: "",
	},
	"packets_sent": label{
		unit:        "",
		description: "",
	},
	"packets_recv": label{
		unit:        "",
		description: "",
	},
	"errin": label{
		unit:        "",
		description: "",
	},
	"errout": label{
		unit:        "",
		description: "",
	},
	"dropin": label{
		unit:        "",
		description: "",
	},
	"dropout": label{
		unit:        "",
		description: "",
	},
}

var netConntrackCounterLabels = map[string]label{
	"conntrackcount": label{
		unit:        "",
		description: "",
	},
	"conntrackmax": label{
		unit:        "",
		description: "",
	},
}

func netIOCounters(nss []core.Namespace) ([]plugin.MetricType, error) {
	// gather accumulated metrics for all interfaces
	netsAll, err := net.IOCounters(false)
	if err != nil {
		return nil, err
	}

	// gather metrics per nic
	netsNic, err := net.IOCounters(true)
	if err != nil {
		return nil, err
	}

	// gather conntrack metric
	conntrack, err := net.FilterCounters()
	if err != nil {
		return nil, err
	}

	results := []plugin.MetricType{}

	for _, ns := range nss {
		// set requested metric name from last namespace element
		metricName := ns.Element(len(ns) - 1).Value
		// check if requested metric is dynamic (requesting metrics for all nics)

		switch ns[3].Value {
		case "*":
			for _, net := range netsNic {
				// prepare namespace copy to update value
				// this will allow to keep namespace as dynamic (name != "")
				dyn := make([]core.NamespaceElement, len(ns))
				copy(dyn, ns)
				dyn[3].Value = net.Name
				// get requested metric value
				val, err := getNetIOCounterValue(&net, metricName)
				if err != nil {
					return nil, err
				}

				metric := plugin.MetricType{
					Namespace_: dyn,
					Data_:      val,
					Timestamp_: time.Now(),
					Unit_:      netIOCounterLabels[metricName].unit,
				}
				results = append(results, metric)
			}
		case "conntrackcount":
			metric := plugin.MetricType{
				Namespace_: ns,
				Data_:      conntrack[0].ConnTrackCount,
				Timestamp_: time.Now(),
				Unit_:      netConntrackCounterLabels[metricName].unit,
			}
			results = append(results, metric)

		case "conntrackmax":
			metric := plugin.MetricType{
				Namespace_: ns,
				Data_:      conntrack[0].ConnTrackMax,
				Timestamp_: time.Now(),
				Unit_:      netConntrackCounterLabels[metricName].unit,
			}
			results = append(results, metric)

		default:
			stats := append(netsAll, netsNic...)
			// find stats for interface name or all nics
			stat := findNetIOStats(stats, ns[3].Value)
			if stat == nil {
				return nil, fmt.Errorf("Requested interface %s not found", ns[3].Value)
			}
			// get value for requested metric
			val, err := getNetIOCounterValue(stat, metricName)
			if err != nil {
				return nil, err
			}

			metric := plugin.MetricType{
				Namespace_: ns,
				Data_:      val,
				Timestamp_: time.Now(),
				Unit_:      netIOCounterLabels[metricName].unit,
			}
			results = append(results, metric)
		}
	}

	return results, nil
}

func findNetIOStats(nets []net.IOCountersStat, name string) *net.IOCountersStat {
	for _, net := range nets {
		if net.Name == name {
			return &net
		}
	}
	return nil
}

func getNetIOCounterValue(stat *net.IOCountersStat, name string) (uint64, error) {
	switch name {
	case "bytes_sent":
		return stat.BytesSent, nil
	case "bytes_recv":
		return stat.BytesRecv, nil
	case "packets_sent":
		return stat.PacketsSent, nil
	case "packets_recv":
		return stat.PacketsRecv, nil
	case "errin":
		return stat.Errin, nil
	case "errout":
		return stat.Errout, nil
	case "dropin":
		return stat.Dropin, nil
	case "dropout":
		return stat.Dropout, nil
	default:
		return 0, fmt.Errorf("Requested NetIOCounter statistic %s is not available", name)
	}
}

func getNetIOCounterMetricTypes() ([]plugin.MetricType, error) {
	mts := make([]plugin.MetricType, 0)

	for name, label := range netIOCounterLabels {
		//metrics which are the sum for all available nics
		mts = append(mts, plugin.MetricType{
			Namespace_:   core.NewNamespace("intel", "psutil", "net", "all", name),
			Description_: label.description,
			Unit_:        label.unit,
		})
		//dynamic metrics representing any nic
		mts = append(mts, plugin.MetricType{
			Namespace_: core.NewNamespace("intel", "psutil", "net").
				AddDynamicElement("nic_id", "network interface id").AddStaticElement(name),
			Description_: label.description,
			Unit_:        label.unit,
		})
	}

	for name, label := range netConntrackCounterLabels {
		mts = append(mts, plugin.MetricType{
			Namespace_:   core.NewNamespace("intel", "psutil", "net", name),
			Description_: label.description,
			Unit_:        label.unit,
		})
	}

	return mts, nil
}
