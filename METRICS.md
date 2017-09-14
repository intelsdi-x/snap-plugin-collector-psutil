### Collected Metrics
This plugin has the ability to gather the following metrics:

Namespace | Data Type | Description (optional)
----------|-----------|------------
/intel/psutil/cpu/cpu-total/guest | float64 | time spent in guest mode accumulated over all cpus
/intel/psutil/cpu/cpu-total/guest_nice | float64 | time spent running a niced guest (virtual CPU for guest operating systems under the control of the Linux kernel) accumulated over all cpus
/intel/psutil/cpu/cpu-total/idle | float64 | time spent in the idle task accumulated over all cpus.  This value should be USER_HZ times the second entry in the /proc/uptime pseudo-file
/intel/psutil/cpu/cpu-total/iowait | float64 | time waiting for I/O to complete accumulated over all cpus
/intel/psutil/cpu/cpu-total/irq | float64 | time servicing interrupts accumulated over all cpus
/intel/psutil/cpu/cpu-total/nice | float64 | time spent in user mode with low priority (nice) accumulated over all cpus
/intel/psutil/cpu/cpu-total/softirq | float64 | time spent servicing softirqs accumulated over all cpus
/intel/psutil/cpu/cpu-total/steal | float64 | stolen time, which is the time spent in other operating systems when running in a virtualized environment accumulated over all cpus
/intel/psutil/cpu/cpu-total/stolen | float64 | stolen time, which is the time spent in other operating systems when running in a virtualized environment
/intel/psutil/cpu/cpu-total/system | float64 | time spent in system mode accumulated over all cpus
/intel/psutil/cpu/cpu-total/user | float64 | time spent in user mode accumulated over all cpus
/intel/psutil/cpu/[CPU]/guest | float64 | time spent in guest mode
/intel/psutil/cpu/[CPU]/guest_nice | float64 | time spent running a niced guest (virtual CPU for guest operating systems under the control of the Linux kernel)
/intel/psutil/cpu/[CPU]/idle | float64 | time spent in the idle task.  This value should be USER_HZ times the second entry in the /proc/uptime pseudo-file
/intel/psutil/cpu/[CPU]/iowait | float64 | time waiting for I/O to complete
/intel/psutil/cpu/[CPU]/irq | float64 | time servicing interrupts
/intel/psutil/cpu/[CPU]/nice | float64 | time spent in user mode with low priority (nice)
/intel/psutil/cpu/[CPU]/softirq | float64 | time spent servicing softirqs
/intel/psutil/cpu/[CPU]/steal | float64 | stolen time, which is the time spent in other operating systems when running in a virtualized environment
/intel/psutil/cpu/[CPU]/stolen | float64 | stolen time, which is the time spent in other operating systems when running in a virtualized environment
/intel/psutil/cpu/[CPU]/system | float64 | time spent in system mode
/intel/psutil/cpu/[CPU]/user | float64 | time spent in user mode
/intel/psutil/disk/[mount_point]/total | uint64 | total space which is available to root in mount point
/intel/psutil/disk/[mount_point]/used | uint64 | total space being used in general in mount point
/intel/psutil/disk/[mount_point]/free | uint64 | remaining free space usable by user mount point
/intel/psutil/disk/[mount_point]/percent | float64 | user usage percent compared to the total amount of space the user can use in mount point
/intel/psutil/load/load1 | float64 | load average over the last 1 minute
/intel/psutil/load/load15 | float64 | load average over the last 15 minutes
/intel/psutil/load/load5 | float64 | load average over the last 5 minutes
/intel/psutil/net/all/bytes_recv | uint64 | number of bytes sent
/intel/psutil/net/all/bytes_sent | uint64 | number of bytes received
/intel/psutil/net/all/dropin | uint64 | total number of incoming packets which were dropped
/intel/psutil/net/all/dropout | uint64 | total number of outgoing packets which were dropped (always 0 on OSX and BSD)
/intel/psutil/net/all/errin | uint64 | total number of errors while receiving
/intel/psutil/net/all/errout | uint64 | total number of errors while sending
/intel/psutil/net/all/packets_recv | uint64 | number of packets received
/intel/psutil/net/all/packets_sent | uint64 | number of packets sent
/intel/psutil/net/[INTERFACE]/bytes_recv | uint64 | number of bytes sent on given interface
/intel/psutil/net/[INTERFACE]/bytes_sent | uint64 | number of bytes received on given interface
/intel/psutil/net/[INTERFACE]/dropin | uint64 | total number of incoming packets which were dropped on given interface
/intel/psutil/net/[INTERFACE]/dropout | uint64 | total number of outgoing packets which were dropped (always 0 on OSX and BSD) o given interface
/intel/psutil/net/[INTERFACE]/errin | uint64 | total number of errors while receiving on given interface
/intel/psutil/net/[INTERFACE]/errout | uint64 | total number of errors while sending on given interface
/intel/psutil/net/[INTERFACE]/packets_recv | uint64 | number of packets received on given interface
/intel/psutil/net/[INTERFACE]/packets_sent | uint64 | number of packets sent on given interface
/intel/psutil/vm/active | uint64 | memory currently in use or very recently used, and so it is in RAM
/intel/psutil/vm/available | uint64 | the actual amount of available memory that can be given instantly to processes that request more memory in bytes; this is calculated by summing different memory values depending on the platform (e.g. free + buffers + cached on Linux) and it is supposed to be used to monitor actual memory usage in a cross platform fashion
/intel/psutil/vm/buffers | uint64 | cache for things like file system metadata
/intel/psutil/vm/cached | uint64 | cache for various things
/intel/psutil/vm/free | uint64 | memory not being used at all (zeroed) that is readily available; note that this doesn't reflect the actual memory available (use 'available' instead).
/intel/psutil/vm/inactive | uint64 | memory that is marked as not used
/intel/psutil/vm/total | uint64 | total physical memory available
/intel/psutil/vm/used | uint64 | memory used, calculated differently depending on the platform and designed for informational purposes only.
/intel/psutil/vm/used_percent | float64 | percent memory used
/intel/psutil/vm/wired | uint64 | memory that is marked to always stay in RAM. It is never moved to disk

*Please note that there is no possibility to request specific instance of dynamic disk metric passing it via requested metric in task manifest. I collect metrics based on configured mount points
All collected network counters contains information about the hardware address (tag -> hardware_address) and the MTU (tag -> mtu).
