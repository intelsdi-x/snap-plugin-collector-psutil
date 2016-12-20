#http://www.apache.org/licenses/LICENSE-2.0.txt
#
#
#Copyright 2016 Intel Corporation
#
#Licensed under the Apache License, Version 2.0 (the "License");
#you may not use this file except in compliance with the License.
#You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
#Unless required by applicable law or agreed to in writing, software
#distributed under the License is distributed on an "AS IS" BASIS,
#WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#See the License for the specific language governing permissions and
#limitations under the License.

import sys
import os
import unittest

from spytest import bins
from spytest import utils
from spytest.logger import log
from unittest import TextTestRunner


class PsutilCollectorLargeTest(unittest.TestCase):

    def setUp(self):
        plugins_dir = os.getenv("PLUGINS_DIR", "/etc/snap/plugins")
        snap_dir = os.getenv("SNAP_DIR", "/usr/local/bin")

        snapteld_url = "http://snap.ci.snap-telemetry.io/snap/latest_build/linux/x86_64/snapteld"
        snaptel_url = "http://snap.ci.snap-telemetry.io/snap/latest_build/linux/x86_64/snaptel"
        psutil_url = "http://snap.ci.snap-telemetry.io/plugins/snap-plugin-collector-psutil/latest_build/linux/x86_64/snap-plugin-collector-psutil"
        file_url = "http://snap.ci.snap-telemetry.io/plugins/snap-plugin-publisher-file/latest_build/linux/x86_64/snap-plugin-publisher-file"

        # set and download required binaries (snapteld, snaptel, plugins)
        self.binaries = bins.Binaries()
        self.binaries.snapteld = bins.Snapteld(snapteld_url, snap_dir)
        self.binaries.snaptel = bins.Snaptel(snaptel_url, snap_dir)
        self.binaries.collector = bins.Plugin(psutil_url, plugins_dir, "collector", 9)
        self.binaries.publisher = bins.Plugin(file_url, plugins_dir, "publisher", -1)

        utils.download_binaries(self.binaries)

        self.task_file = "{}/examples/tasks/task-psutil.json".format(os.getenv("PROJECT_DIR", "/snap-plugin-collector-psutil"))

        log.info("starting snapteld")
        self.binaries.snapteld.start()
        if not self.binaries.snapteld.isAlive():
            self.fail("snapteld thread died")

        log.debug("Waiting for snapteld to finish starting")
        if not self.binaries.snapteld.wait():
            log.error("snapteld errors: {}".format(self.binaries.snapteld.errors))
            self.binaries.snapteld.kill()
            self.fail("snapteld not ready, timeout!")

    def test_psutil_collector_plugin(self):
        # load plugins
        for plugin in self.binaries.get_all_plugins():
            log.info("snaptel plugin load {}".format(os.path.join(plugin.dir, plugin.name)))
            loaded = self.binaries.snaptel.load_plugin(plugin)
            self.assertTrue(loaded, "{} loaded".format(plugin.name))

        # check available metrics, plugins and tasks
        metrics = self.binaries.snaptel.list_metrics()
        plugins = self.binaries.snaptel.list_plugins()
        tasks = self.binaries.snaptel.list_tasks()
        self.assertGreater(len(metrics), 0, "Metrics available {} expected {}".format(len(metrics), 0))
        self.assertEqual(len(plugins), 2, "Plugins available {} expected {}".format(len(plugins), 2))
        self.assertEqual(len(tasks), 0, "Tasks available {} expected {}".format(len(tasks), 0))

        # check config policy for metric
        rules = self.binaries.snaptel.metric_get("/intel/psutil/vm/free")
        self.assertEqual(len(rules), 0, "Rules available {} expected {}".format(len(rules), 0))

        # create and list available task
        log.info("snaptel task create -t {}".format(self.task_file))
        task_id = self.binaries.snaptel.create_task(self.task_file)
        tasks = self.binaries.snaptel.list_tasks()
        self.assertEqual(len(tasks), 1, "Tasks available {} expected {}".format(len(tasks), 1))

        # check if task hits and fails
        hits = self.binaries.snaptel.task_hits_count(task_id)
        fails = self.binaries.snaptel.task_fails_count(task_id)
        self.assertGreater(hits, 0, "Task hits {} expected {}".format(hits, ">0"))
        self.assertEqual(fails, 0, "Task fails {} expected {}".format(fails, 0))

        # stop task and list available tasks
        log.info("snaptel task stop {}".format(task_id))
        stopped = self.binaries.snaptel.stop_task(task_id)
        self.assertTrue(stopped, "Task stopped")
        tasks = self.binaries.snaptel.list_tasks()
        self.assertEqual(len(tasks), 1, "Tasks available {} expected {}".format(len(tasks), 1))

        # unload plugin, list metrics and plugins
        log.info("snaptel plugin unload {}".format(self.binaries.collector))
        self.binaries.snaptel.unload_plugin(self.binaries.collector)
        metrics = self.binaries.snaptel.list_metrics()
        plugins = self.binaries.snaptel.list_plugins()
        self.assertEqual(len(metrics), 0, "Metrics available {} expected {}".format(len(metrics), 0))
        self.assertEqual(len(plugins), 1, "Plugins available {} expected {}".format(len(plugins), 1))

        # check for snapteld errors
        self.assertEqual(len(self.binaries.snapteld.errors), 0, "Errors found during snapteld execution:\n{}"
                         .format("\n".join(self.binaries.snapteld.errors)))

    def tearDown(self):
        log.info("stopping snapteld")
        self.binaries.snapteld.stop()
        if self.binaries.snapteld.isAlive():
            log.warn("snapteld thread did not die")

if __name__ == "__main__":
    test_suite = unittest.TestLoader().loadTestsFromTestCase(PsutilCollectorLargeTest)
    test_result = TextTestRunner().run(test_suite)
    # exit with return code equal to number of failures
    sys.exit(len(test_result.failures))


