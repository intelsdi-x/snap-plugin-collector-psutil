DISCONTINUATION OF PROJECT. 

This project will no longer be maintained by Intel.

This project has been identified as having known security escapes.

Intel has ceased development and contributions including, but not limited to, maintenance, bug fixes, new releases, or updates, to this project.  

Intel no longer accepts patches to this project.

# DISCONTINUATION OF PROJECT 

**This project will no longer be maintained by Intel.  Intel will not provide or guarantee development of or support for this project, including but not limited to, maintenance, bug fixes, new releases or updates.  Patches to this project are no longer accepted by Intel. If you have an ongoing need to use this project, are interested in independently developing it, or would like to maintain patches for the community, please create your own fork of the project.**


# Snap collector plugin - psutil
This plugin collects metrics from psutil which gathers information on running processes and system utilization (CPU, memory, disks, network). 

It's used in the [Snap framework](http://github.com:intelsdi-x/snap).

1. [Getting Started](#getting-started)
  * [System Requirements](#system-requirements)
  * [Installation](#installation)
  * [Configuration and Usage](#configuration-and-usage)
2. [Documentation](#documentation)
  * [Collected Metrics](#collected-metrics)
  * [Examples](#examples)
  * [Roadmap](#roadmap)
3. [Community Support](#community-support)
4. [Contributing](#contributing)
5. [License](#license-and-authors)
6. [Acknowledgements](#acknowledgements)

## Getting Started
### System Requirements 
* [golang 1.6+](https://golang.org/dl/) (needed only for building)

Note: This plugin does not require Python rather it depends on the go library [gopsutil](https://github.com/shirou/gopsutil).  

### Operating systems
All OSs currently supported by snap:
* Linux/amd64
* Darwin/amd64

### Installation
#### Download psutil plugin binary:
You can get the pre-built binaries for your OS and architecture under the plugin's [release](https://github.com/intelsdi-x/snap-plugin-collector-psutil/releases) page.  For Snap, check [here](https://github.com/intelsdi-x/snap/releases).


#### To build the plugin binary:
Fork https://github.com/intelsdi-x/snap-plugin-collector-psutil

Clone repo into `$GOPATH/src/github.com/intelsdi-x/`:

```
$ git clone https://github.com/<yourGithubID>/snap-plugin-collector-psutil.git
```

Build the plugin by running make within the cloned repo:
```
$ make
```
This builds the plugin in `./build/`

### Configuration and Usage
* Set up the [Snap framework](https://github.com/intelsdi-x/snap/blob/master/README.md#getting-started)

Some metrics are platform specific (see [gopsutil's current status](https://github.com/shirou/gopsutil/blob/master/README.rst#current-status)).

Available configuration option:
* mount_points - configuration of mount points to monitor, multiple paths should be separated with "|", e.g. "/|/dev|/run", default is set to collect only physical devices (hard disks, cd-rom, USB). Passing `*` enables collect data from all mount points.

## Documentation
There are a number of other resources you can review to learn to use this plugin:

* [gopsutil](https://github.com/shirou/gopsutil/) (go based implementation)
* [psutil](https://pythonhosted.org/psutil/) (python based implementation)
* [Snap psutil integration test](https://github.com/intelsdi-x/snap-plugin-collector-psutil/blob/master/psutil/psutil_integration_test.go)
* [Snap psutil unit test](https://github.com/intelsdi-x/snap-plugin-collector-psutil/blob/master/psutil/psutil_test.go)
* [Snap psutil examples](#examples)

### Collected metrics
List of metrics collected by this plugin can be found in [METRICS.md file](METRICS.md).

### Examples
This is an example running psutil and writing data to a file. It is assumed that you are using the latest Snap binary and plugins.

The example is run from a directory which includes snaptel, snapteld, along with the plugins and task file.

In one terminal window, open the Snap daemon (in this case with logging set to 1 and trust disabled):
```
$ snapteld -l 1 -t 0
```

In another terminal window:
Load psutil plugin
```
$ snaptel plugin load snap-plugin-collector-psutil
```
See available metrics for your system
```
$ snaptel metric list
```

Create a task manifest file (e.g. `task-psutil.json`):    
```json
{
    "version": 1,
    "schedule": {
        "type": "simple",
        "interval": "1s"
    },
    "workflow": {
        "collect": {
            "metrics": {
                "/intel/psutil/load/load1": {},
                "/intel/psutil/load/load5": {},
                "/intel/psutil/load/load15": {},
                "/intel/psutil/cpu/*/user": {},
                "/intel/psutil/net/*/bytes_sent": {},
                "/intel/psutil/vm/available": {},
                "/intel/psutil/vm/free": {},
                "/intel/psutil/vm/used": {}
            },
            "config": {
                "/intel/mock": {
                    "password": "secret",
                    "user": "root"
                }
            },
            "publish": [
                {                         
                    "plugin_name": "file",
                    "config": {
                        "file": "/tmp/published_psutil"
                    }
                }
            ]
        }
    }
}
```

Load file plugin for publishing:
```
$ snaptel plugin load snap-plugin-publisher-file
Plugin loaded
Name: file
Version: 3
Type: publisher
Signed: false
Loaded Time: Fri, 20 Nov 2015 11:41:39 PST
```

Create task:
```
$ snaptel task create -t task-psutil.json
Using task manifest to create task
Task created
ID: 02dd7ff4-8106-47e9-8b86-70067cd0a850
Name: Task-02dd7ff4-8106-47e9-8b86-70067cd0a850
State: Running
```

See file output (this is just part of the file):
```
2015-11-20 11:46:03.637390565 -0800 PST|[intel psutil load load1]|1.82|username-mac01.jf.intel.com
2015-11-20 11:46:03.641160359 -0800 PST|[intel psutil load load15]|2.09|username-mac01.jf.intel.com
2015-11-20 11:46:03.643858208 -0800 PST|[intel psutil load load5]|2.08|username-mac01.jf.intel.com
2015-11-20 11:46:03.661173851 -0800 PST|[intel psutil vm available]|168882176|username-mac01.jf.intel.com
2015-11-20 11:46:03.67167664 -0800 PST|[intel psutil vm free]|168943616|username-mac01.jf.intel.com
2015-11-20 11:46:03.681965105 -0800 PST|[intel psutil vm used]|17010798592|username-mac01.jf.intel.com
2015-11-20 11:46:04.641244629 -0800 PST|[intel psutil load load1]|1.82|username-mac01.jf.intel.com
2015-11-20 11:46:04.644420189 -0800 PST|[intel psutil load load15]|2.09|username-mac01.jf.intel.com
2015-11-20 11:46:04.647166418 -0800 PST|[intel psutil load load5]|2.08|username-mac01.jf.intel.com
2015-11-20 11:46:04.657065347 -0800 PST|[intel psutil vm available]|168984576|username-mac01.jf.intel.com
2015-11-20 11:46:04.666346721 -0800 PST|[intel psutil vm free]|169054208|username-mac01.jf.intel.com
2015-11-20 11:46:04.676683476 -0800 PST|[intel psutil vm used]|17010716672|username-mac01.jf.intel.com
```

Stop task:
```
$ snaptel task stop 02dd7ff4-8106-47e9-8b86-70067cd0a850
Task stopped:
ID: 02dd7ff4-8106-47e9-8b86-70067cd0a850
```

### Roadmap
There isn't a current roadmap for this plugin, but it is in active development. As we launch this plugin, we do not have any outstanding requirements for the next release. If you have a feature request, please add it as an [issue](https://github.com/intelsdi-x/snap-plugin-collector-psutil/issues/new) and/or submit a [pull request](https://github.com/intelsdi-x/snap-plugin-collector-psutil/pulls).

## Community Support
This repository is one of **many** plugins in **Snap**, a powerful telemetry framework. See the full project at http://github.com/intelsdi-x/snap To reach out to other users, head to the [main framework](https://github.com/intelsdi-x/snap#community-support)

## Contributing
We love contributions!

There's more than one way to give back, from examples to blogs to code updates. See our recommended process in [CONTRIBUTING.md](CONTRIBUTING.md).

## License
[Snap](http://github.com:intelsdi-x/snap), along with this plugin, is an Open Source software released under the Apache 2.0 [License](LICENSE).

## Acknowledgements
* Author: [@jcooklin](https://github.com/jcooklin/)

And **thank you!** Your contribution, through code and participation, is incredibly important to us.
