from modules import utils
from modules.logger import log
from unittest import TextTestRunner

import sys
import unittest


class LargeTest(unittest.TestCase):

    def setUp(self):
        self.binaries = utils.set_binaries()
        utils.download_binaries(self.binaries)
        log.debug("Starting snapd")
        self.binaries.snapd.start()
        if not self.binaries.snapd.isAlive():
            self.fail("snapd thread died")

        log.debug("Waiting for snapd to finish starting")
        self.binaries.snapd.wait()

    def test_psutil_collector_plugin(self):
        loaded = self.binaries.snapctl.load_plugin("snap-plugin-collector-psutil")
        self.assertTrue(loaded, "psutil collector loaded")

        metrics = self.binaries.snapctl.list_metrics()
        plugins = self.binaries.snapctl.list_plugins()
        tasks = self.binaries.snapctl.list_tasks()
        self.assertGreater(metrics, 0, "Metrics available %s" % metrics)
        self.assertEqual(plugins, 1, "Plugins available %s" % plugins)
        self.assertEqual(tasks, 0, "Tasks avaialble %s" % tasks)

        task_id = self.binaries.snapctl.create_task("/snap-plugin-collector-psutil/scripts/docker/large/psutil-task.yml")
        tasks = self.binaries.snapctl.list_tasks()
        self.assertEqual(tasks, 1, "Tasks avaialble %s" % tasks)

        hits = self.binaries.snapctl.task_hits(task_id)
        fails = self.binaries.snapctl.task_fails(task_id)
        self.assertGreater(hits, 0, "Task hits %s" % hits)
        self.assertEqual(fails, 0, "Task fails %s" % fails)

        stopped = self.binaries.snapctl.stop_task(task_id)
        self.assertTrue(stopped, "Task stopped")
        tasks = self.binaries.snapctl.list_tasks()
        self.assertEqual(tasks, 1, "Tasks avaialble %s" % tasks)

        self.binaries.snapctl.unload_plugin("collector", "psutil", "6")
        metrics = self.binaries.snapctl.list_metrics()
        plugins = self.binaries.snapctl.list_plugins()
        self.assertEqual(metrics, 0, "Metrics available %s" % metrics)
        self.assertEqual(plugins, 0, "Plugins available %s" % plugins)

        self.assertEqual(len(self.binaries.snapd.errors), 0, "Errors found during snapd execution")

    def tearDown(self):
        log.debug("Stopping snapd thread")
        self.binaries.snapd.stop()
        if self.binaries.snapd.isAlive():
            log.warn("snapd thread did not died")

if __name__ == "__main__":
    test_suite = unittest.TestLoader().loadTestsFromTestCase(LargeTest)
    test_result = TextTestRunner().run(test_suite)
    sys.exit(len(test_result.failures))


