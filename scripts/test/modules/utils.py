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

import urllib
import errno

from bins import *
from logger import log


def download_binaries(bins):
    for binary in bins.get_all_bins():
        log.debug("Downloading {} to {}".format(binary.url, binary.dir))
        _download_binary(binary)


def set_binaries():
    bins = Binaries()
    bins.snapd = Snapd(SNAP_URL, SNAP_DIR)
    bins.snapctl = Snapctl(SNAPCTL_URL, SNAP_DIR)
    bins.plugins = [Binary(PLUGIN_URL, PLUGIN_DIR)]
    return bins


def _ensure_dir(dirname):
    try:
        os.makedirs(dirname)
        log.debug("{} created".format(dirname))
    except OSError as e:
        if e.errno != errno.EEXIST:
            log.error(e.errno)
            raise


def _download_binary(binary):
    f = urllib.URLopener()
    try:
        if not os.path.isdir(binary.dir):
            _ensure_dir(binary.dir)
        fname, headers = f.retrieve(binary.url, os.path.join(binary.dir, os.path.basename(binary.url)))
        os.chmod(fname, 755)
        log.debug("chmod set to 755 for {}".format(fname))
    except IOError as e:
        log.error(e.args)

