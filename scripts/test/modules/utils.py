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

