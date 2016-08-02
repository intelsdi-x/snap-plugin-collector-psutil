import os
import threading
import subprocess
import fcntl

import time

from logger import log

PLUGIN_DIR = "/etc/snap/plugins"
SNAP_DIR = "/usr/local/bin"

PLUGIN_URL = "http://snap.ci.snap-telemetry.io/plugin/build/latest/snap-plugin-collector-psutil"
SNAP_URL = "http://snap.ci.snap-telemetry.io/snap/master/latest/snapd"
SNAPCTL_URL = "http://snap.ci.snap-telemetry.io/snap/master/latest/snapctl"


def _non_block_read(output):
    fd = output.fileno()
    fl = fcntl.fcntl(fd, fcntl.F_GETFL)
    fcntl.fcntl(fd, fcntl.F_SETFL, fl | os.O_NONBLOCK)
    try:
        return output.readline()
    except:
        return ""


class Binary(object):

    def __init__(self, url, location):
        self._url = url
        self._dir = location
        self._name = os.path.basename(url)

    @property
    def url(self):
        return self._url

    @url.setter
    def url(self, u):
        self._url = u

    @property
    def dir(self):
        return self._dir

    @dir.setter
    def dir(self, loc):
        self._dir = loc

    @property
    def name(self):
        return self._name

    @name.setter
    def name(self, n):
        self._name = n

    def __str__(self):
        return self._name


class Snapd(Binary, threading.Thread):

    def __init__(self, url, location):
        Binary.__init__(self, url, location)
        threading.Thread.__init__(self)
        self.stdout = None
        self.stderr = None
        self.errors = []
        self._stop = threading.Event()
        self._ready = threading.Event()
        self._process = None

    def run(self):
        cmd = '%s -t 0 -l 1 ' % (os.path.join(self.dir, self.name))
        log.debug("starting snapd thread: %s" % cmd)
        self._process = subprocess.Popen(cmd.split(), shell=False, stdout=subprocess.PIPE, stderr=subprocess.PIPE)
        while not self.stopped():
            out = _non_block_read(self._process.stderr)
            if "snapd started" in out:
                self._ready.set()
                log.debug("snapd is ready")
            if "error" in out:
                self.errors.append(out)
        self._process.kill()
        log.debug("exiting snapd thread")

    def stop(self):
        self._stop.set()
        self.join()

    def stopped(self):
        return self._stop.isSet()

    def wait(self, timeout=10):
        out = 0
        while not self._ready.isSet() or out > timeout:
            out += 0.5
            time.sleep(0.5)
        return out <= timeout


class Snapctl(Binary):

    def __init__(self, url, location):
        Binary.__init__(self, url, location)
        self.errors = []

    def load_plugin(self, plugin):
        cmd = '%s plugin load %s' % (os.path.join(self.dir, self.name), os.path.join(PLUGIN_DIR, plugin))
        log.debug("snapctl load plugin %s", cmd)
        out = self._start_process(cmd)
        log.debug("plugin loaded? %s", "Plugin loaded" in out)
        return "Plugin loaded" in out

    def unload_plugin(self, plugin_type, plugin_name, plugin_version):
        cmd = '%s plugin unload %s:%s:%s' % (os.path.join(self.dir, self.name), plugin_type, plugin_name, plugin_version)
        log.debug("snapctl unload plugin %s", cmd)
        out = self._start_process(cmd)
        log.debug("plugin unloaded? %s", "Plugin unloaded" in out)
        return "Plugin unloaded" in out

    def list_plugins(self):
        cmd = '%s plugin list' % (os.path.join(self.dir, self.name))
        log.debug("snapctl plugin list")
        plugins = self._start_process(cmd).split('\n')[1:-1]
        return len(plugins)

    def create_task(self, task):
        cmd = '%s task create -t %s' % (os.path.join(self.dir, self.name), task)
        log.debug("snapctl task create")
        out = self._start_process(cmd).split('\n')
        # sleeping for 10 seconds so the task can do some work
        time.sleep(10)
        if len(out) == 0:
            return False
        log.debug("task created? %s", out[1] == "Task created")
        task_id = out[2].split()
        if len(task_id):
            return task_id[1]
        return ""

    def list_tasks(self):
        return len(self._task_list())

    def stop_task(self, task_id):
        cmd = '%s task stop %s' % (os.path.join(self.dir, self.name), task_id)
        log.debug("snapctl task stop")
        out = self._start_process(cmd).split('\n')
        return "Task stopped" in out[0]

    def task_hits(self, task_id):
        tasks = self._task_list()
        hits = 0
        for task in tasks:
            if task.split()[0] == task_id:
                hits += int(task.split()[3])

        log.debug("task hits %s", hits)
        return hits

    def task_fails(self, task_id):
        tasks = self._task_list()
        fails = 0
        for task in tasks:
            if task.split()[0] == task_id:
                fails += int(task.split()[5])

        log.debug("task fails %s", fails)
        return fails

    def list_metrics(self):
        cmd = '%s metric list' % (os.path.join(self.dir, self.name))
        log.debug("snapctl metric list")
        metrics = self._start_process(cmd).split('\n')[1:-1]
        return len(metrics)

    def _task_list(self):
        cmd = '%s task list' % (os.path.join(self.dir, self.name))
        tasks = self._start_process(cmd).split('\n')[1:]
        tasks = filter(lambda t: t != '', tasks)
        if len(tasks):
            return tasks
        return []

    def _start_process(self, cmd):
        process = subprocess.Popen(cmd.split(), shell=False, stdout=subprocess.PIPE, stderr=subprocess.PIPE)
        out, err = process.communicate()
        if err:
            self.errors.append(err)
        return out


class Binaries(object):

    def __init__(self):
        self._snapd = None
        self._snapctl = None
        self._plugins = []

    @property
    def snapd(self):
        return self._snapd

    @snapd.setter
    def snapd(self, bin):
        self._snapd = bin

    @property
    def snapctl(self):
        return self._snapctl

    @snapctl.setter
    def snapctl(self, bin):
        self._snapctl = bin

    @property
    def plugins(self):
        return self._plugins

    @plugins.setter
    def plugins(self, bins):
        self._plugins = bins

    def get_all_bins(self):
        all_bins = [self.snapd, self.snapctl]
        all_bins.extend(self._plugins)
        return all_bins

    def __str__(self):
        return ";".join(map(lambda e: e.name, self.get_all_bins()))
