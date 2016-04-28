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

func netIOCounters(ns core.Namespace) (*plugin.MetricType, error) {
	nets, err := net.IOCounters(true)
	if err != nil {
		return nil, err
	}

	for _, net := range nets {
		switch {
		case regexp.MustCompile(`^/intel/psutil/net/.*/bytes_sent$`).MatchString(ns.String()):
			return &plugin.MetricType{
				Namespace_: ns,
				Data_:      net.BytesSent,
			}, nil
		case regexp.MustCompile(`^/intel/psutil/net/.*/bytes_recv`).MatchString(ns.String()):
			return &plugin.MetricType{
				Namespace_: ns,
				Data_:      net.BytesRecv,
			}, nil
		case regexp.MustCompile(`^/intel/psutil/net/.*/packets_sent`).MatchString(ns.String()):
			return &plugin.MetricType{
				Namespace_: ns,
				Data_:      net.BytesSent,
			}, nil
		case regexp.MustCompile(`^/intel/psutil/net/.*/packets_recv`).MatchString(ns.String()):
			return &plugin.MetricType{
				Namespace_: ns,
				Data_:      net.BytesRecv,
			}, nil
		case regexp.MustCompile(`^/intel/psutil/net/.*/errin`).MatchString(ns.String()):
			return &plugin.MetricType{
				Namespace_: ns,
				Data_:      net.Errin,
			}, nil
		case regexp.MustCompile(`^/intel/psutil/net/.*/errout`).MatchString(ns.String()):
			return &plugin.MetricType{
				Namespace_: ns,
				Data_:      net.Errout,
			}, nil
		case regexp.MustCompile(`^/intel/psutil/net/.*/dropin`).MatchString(ns.String()):
			return &plugin.MetricType{
				Namespace_: ns,
				Data_:      net.Dropin,
			}, nil
		case regexp.MustCompile(`^/intel/psutil/net/.*/dropout`).MatchString(ns.String()):
			return &plugin.MetricType{
				Namespace_: ns,
				Data_:      net.Dropout,
			}, nil
		}
	}

	return nil, fmt.Errorf("Unknown error processing %v", ns)
}

func getNetIOCounterMetricTypes() ([]plugin.MetricType, error) {
	mts := make([]plugin.MetricType, 0)
	nets, err := net.IOCounters(false)
	if err != nil {
		return nil, err
	}
	//total for all nics
	for name, label := range netIOCounterLabels {
		mts = append(mts, plugin.MetricType{
			Namespace_:   core.NewNamespace("intel", "psutil", "net", nets[0].Name, name),
			Description_: label.description,
			Unit_:        label.unit,
		})
	}
	//per nic
	nets, err = net.IOCounters(true)
	if err != nil {
		return nil, err
	}
	for _, net := range nets {
		for name, label := range netIOCounterLabels {
			mts = append(mts, plugin.MetricType{
				Namespace_:   core.NewNamespace("intel", "psutil", "net", net.Name, name),
				Description_: label.description,
				Unit_:        label.unit,
			})
		}
	}
	return mts, nil
}
