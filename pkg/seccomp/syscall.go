package seccomp

import (
	"fmt"
)

// List of x86-x64 syscalls
// Source: https://raw.githubusercontent.com/torvalds/linux/master/arch/x86/entry/syscalls/syscall_64.tbl
var systemCalls = map[int]string{
	0:   "read",
	1:   "write",
	2:   "open",
	3:   "close",
	4:   "stat",
	5:   "fstat",
	6:   "lstat",
	7:   "poll",
	8:   "lseek",
	9:   "mmap",
	10:  "mprotect",
	11:  "munmap",
	12:  "brk",
	13:  "rt_sigaction",
	14:  "rt_sigprocmask",
	15:  "rt_sigreturn",
	16:  "ioctl",
	17:  "pread64",
	18:  "pwrite64",
	19:  "readv",
	20:  "writev",
	21:  "access",
	22:  "pipe",
	23:  "select",
	24:  "sched_yield",
	25:  "mremap",
	26:  "msync",
	27:  "mincore",
	28:  "madvise",
	29:  "shmget",
	30:  "shmat",
	31:  "shmctl",
	32:  "dup",
	33:  "dup2",
	34:  "pause",
	35:  "nanosleep",
	36:  "getitimer",
	37:  "alarm",
	38:  "setitimer",
	39:  "getpid",
	40:  "sendfile",
	41:  "socket",
	42:  "connect",
	43:  "accept",
	44:  "sendto",
	45:  "recvfrom",
	46:  "sendmsg",
	47:  "recvmsg",
	48:  "shutdown",
	49:  "bind",
	50:  "listen",
	51:  "getsockname",
	52:  "getpeername",
	53:  "socketpair",
	54:  "setsockopt",
	55:  "getsockopt",
	56:  "clone",
	57:  "fork",
	58:  "vfork",
	59:  "execve",
	60:  "exit",
	61:  "wait4",
	62:  "kill",
	63:  "uname",
	64:  "semget",
	65:  "semop",
	66:  "semctl",
	67:  "shmdt",
	68:  "msgget",
	69:  "msgsnd",
	70:  "msgrcv",
	71:  "msgctl",
	72:  "fcntl",
	73:  "flock",
	74:  "fsync",
	75:  "fdatasync",
	76:  "truncate",
	77:  "ftruncate",
	78:  "getdents",
	79:  "getcwd",
	80:  "chdir",
	81:  "fchdir",
	82:  "rename",
	83:  "mkdir",
	84:  "rmdir",
	85:  "creat",
	86:  "link",
	87:  "unlink",
	88:  "symlink",
	89:  "readlink",
	90:  "chmod",
	91:  "fchmod",
	92:  "chown",
	93:  "fchown",
	94:  "lchown",
	95:  "umask",
	96:  "gettimeofday",
	97:  "getrlimit",
	98:  "getrusage",
	99:  "sysinfo",
	100: "times",
	101: "ptrace",
	102: "getuid",
	103: "syslog",
	104: "getgid",
	105: "setuid",
	106: "setgid",
	107: "geteuid",
	108: "getegid",
	109: "setpgid",
	110: "getppid",
	111: "getpgrp",
	112: "setsid",
	113: "setreuid",
	114: "setregid",
	115: "getgroups",
	116: "setgroups",
	117: "setresuid",
	118: "getresuid",
	119: "setresgid",
	120: "getresgid",
	121: "getpgid",
	122: "setfsuid",
	123: "setfsgid",
	124: "getsid",
	125: "capget",
	126: "capset",
	127: "rt_sigpending",
	128: "rt_sigtimedwait",
	129: "rt_sigqueueinfo",
	130: "rt_sigsuspend",
	131: "sigaltstack",
	132: "utime",
	133: "mknod",
	134: "uselib",
	135: "personality",
	136: "ustat",
	137: "statfs",
	138: "fstatfs",
	139: "sysfs",
	140: "getpriority",
	141: "setpriority",
	142: "sched_setparam",
	143: "sched_getparam",
	144: "sched_setscheduler",
	145: "sched_getscheduler",
	146: "sched_get_priority_max",
	147: "sched_get_priority_min",
	148: "sched_rr_get_interval",
	149: "mlock",
	150: "munlock",
	151: "mlockall",
	152: "munlockall",
	153: "vhangup",
	154: "modify_ldt",
	155: "pivot_root",
	156: "_sysctl",
	157: "prctl",
	158: "arch_prctl",
	159: "adjtimex",
	160: "setrlimit",
	161: "chroot",
	162: "sync",
	163: "acct",
	164: "settimeofday",
	165: "mount",
	166: "umount2",
	167: "swapon",
	168: "swapoff",
	169: "reboot",
	170: "sethostname",
	171: "setdomainname",
	172: "iopl",
	173: "ioperm",
	174: "create_module",
	175: "init_module",
	176: "delete_module",
	177: "get_kernel_syms",
	178: "query_module",
	179: "quotactl",
	180: "nfsservctl",
	181: "getpmsg",
	182: "putpmsg",
	183: "afs_syscall",
	184: "tuxcall",
	185: "security",
	186: "gettid",
	187: "readahead",
	188: "setxattr",
	189: "lsetxattr",
	190: "fsetxattr",
	191: "getxattr",
	192: "lgetxattr",
	193: "fgetxattr",
	194: "listxattr",
	195: "llistxattr",
	196: "flistxattr",
	197: "removexattr",
	198: "lremovexattr",
	199: "fremovexattr",
	200: "tkill",
	201: "time",
	202: "futex",
	203: "sched_setaffinity",
	204: "sched_getaffinity",
	205: "set_thread_area",
	206: "io_setup",
	207: "io_destroy",
	208: "io_getevents",
	209: "io_submit",
	210: "io_cancel",
	211: "get_thread_area",
	212: "lookup_dcookie",
	213: "epoll_create",
	214: "epoll_ctl_old",
	215: "epoll_wait_old",
	216: "remap_file_pages",
	217: "getdents64",
	218: "set_tid_address",
	219: "restart_syscall",
	220: "semtimedop",
	221: "fadvise64",
	222: "timer_create",
	223: "timer_settime",
	224: "timer_gettime",
	225: "timer_getoverrun",
	226: "timer_delete",
	227: "clock_settime",
	228: "clock_gettime",
	229: "clock_getres",
	230: "clock_nanosleep",
	231: "exit_group",
	232: "epoll_wait",
	233: "epoll_ctl",
	234: "tgkill",
	235: "utimes",
	236: "vserver",
	237: "mbind",
	238: "set_mempolicy",
	239: "get_mempolicy",
	240: "mq_open",
	241: "mq_unlink",
	242: "mq_timedsend",
	243: "mq_timedreceive",
	244: "mq_notify",
	245: "mq_getsetattr",
	246: "kexec_load",
	247: "waitid",
	248: "add_key",
	249: "request_key",
	250: "keyctl",
	251: "ioprio_set",
	252: "ioprio_get",
	253: "inotify_init",
	254: "inotify_add_watch",
	255: "inotify_rm_watch",
	256: "migrate_pages",
	257: "openat",
	258: "mkdirat",
	259: "mknodat",
	260: "fchownat",
	261: "futimesat",
	262: "newfstatat",
	263: "unlinkat",
	264: "renameat",
	265: "linkat",
	266: "symlinkat",
	267: "readlinkat",
	268: "fchmodat",
	269: "faccessat",
	270: "pselect6",
	271: "ppoll",
	272: "unshare",
	273: "set_robust_list",
	274: "get_robust_list",
	275: "splice",
	276: "tee",
	277: "sync_file_range",
	278: "vmsplice",
	279: "move_pages",
	280: "utimensat",
	281: "epoll_pwait",
	282: "signalfd",
	283: "timerfd_create",
	284: "eventfd",
	285: "fallocate",
	286: "timerfd_settime",
	287: "timerfd_gettime",
	288: "accept4",
	289: "signalfd4",
	290: "eventfd2",
	291: "epoll_create1",
	292: "dup3",
	293: "pipe2",
	294: "inotify_init1",
	295: "preadv",
	296: "pwritev",
	297: "rt_tgsigqueueinfo",
	298: "perf_event_open",
	299: "recvmmsg",
	300: "fanotify_init",
	301: "fanotify_mark",
	302: "prlimit64",
	303: "name_to_handle_at",
	304: "open_by_handle_at",
	305: "clock_adjtime",
	306: "syncfs",
	307: "sendmmsg",
	308: "setns",
	309: "getcpu",
	310: "process_vm_readv",
	311: "process_vm_writev",
	312: "kcmp",
	313: "finit_module",
	314: "sched_setattr",
	315: "sched_getattr",
	316: "renameat2",
	317: "seccomp",
	318: "getrandom",
	319: "memfd_create",
	320: "kexec_file_load",
	321: "bpf",
	322: "execveat",
	323: "userfaultfd",
	324: "membarrier",
	325: "mlock2",
	326: "copy_file_range",
	327: "preadv2",
	328: "pwritev2",
	329: "pkey_mprotect",
	330: "pkey_alloc",
	331: "pkey_free",
	332: "statx",
	333: "io_pgetevents",
	334: "rseq",
	424: "pidfd_send_signal",
	425: "io_uring_setup",
	426: "io_uring_enter",
	427: "io_uring_register",
	428: "open_tree",
	429: "move_mount",
	430: "fsopen",
	431: "fsconfig",
	432: "fsmount",
	433: "fspick",
	434: "pidfd_open",
	435: "clone3",
	512: "rt_sigaction",
	513: "rt_sigreturn",
	514: "ioctl",
	515: "readv",
	516: "writev",
	517: "recvfrom",
	518: "sendmsg",
	519: "recvmsg",
	520: "execve",
	521: "ptrace",
	522: "rt_sigpending",
	523: "rt_sigtimedwait",
	524: "rt_sigqueueinfo",
	525: "sigaltstack",
	526: "timer_create",
	527: "mq_notify",
	528: "kexec_load",
	529: "waitid",
	530: "set_robust_list",
	531: "get_robust_list",
	532: "vmsplice",
	533: "move_pages",
	534: "preadv",
	535: "pwritev",
	536: "rt_tgsigqueueinfo",
	537: "recvmmsg",
	538: "sendmmsg",
	539: "process_vm_readv",
	540: "process_vm_writev",
	541: "setsockopt",
	542: "getsockopt",
	543: "io_setup",
	544: "io_submit",
	545: "execveat",
	546: "preadv2",
	547: "pwritev2",
}

// getSyscallName returns the system call name for a given syscallID
func getSyscallName(syscallID int) (string, error) {

	syscallName, exists := systemCalls[syscallID]
	if !exists {
		return "", fmt.Errorf("syscall id %d not supported", syscallID)
	}

	return syscallName, nil
}

func getAllSyscallNames() []string {
	unique := make(map[string]bool)
	syscalls := make([]string, 0, len(systemCalls))
	for _, n := range systemCalls {
		if _, exists := unique[n]; !exists {
			unique[n] = false
			syscalls = append(syscalls, n)
		}
	}
	return syscalls
}

func getMostFrequentSyscalls() []string {
	return []string{
		"arch_prctl",
		"bind",
		"clone", "clock_gettime", "close", "connect",
		"dup2",
		"execve", "exit", "exit_group", "epoll_pwait",
		"fcntl", "futex",
		"getpid", "getsockname", "getuid",
		"ioctl",
		"mprotect",
		"nanosleep",
		"open",
		"poll",
		"read", "recvfrom", "rt_sigaction", "rt_sigreturn", "rt_sigprocmask",
		"sendto", "setitimer", "socket", "set_tid_address", "setsockopt",
		"write", "writev",
	}
}

func getMostSyscalls() []string {
	return []string{
		"accept",
		"accept4",
		"access",
		"alarm",
		"arch_prctl",
		"bind",
		"brk",
		"capget",
		"capset",
		"chdir",
		"chmod",
		"chown",
		"chroot",
		"clock_getres",
		"clock_gettime",
		"clock_nanosleep",
		"close",
		"clone",
		"connect",
		"copy_file_range",
		"creat",
		"dup",
		"dup2",
		"dup3",
		"epoll_create",
		"epoll_create1",
		"epoll_ctl",
		"epoll_ctl_old",
		"epoll_pwait",
		"epoll_wait",
		"epoll_wait_old",
		"eventfd",
		"eventfd2",
		"execve",
		"execveat",
		"exit",
		"exit_group",
		"faccessat",
		"fadvise64",
		"fallocate",
		"fanotify_init",
		"fanotify_mark",
		"fchdir",
		"fchmod",
		"fchmodat",
		"fchown",
		"fchownat",
		"fcntl",
		"fdatasync",
		"fgetxattr",
		"flistxattr",
		"flock",
		"fork",
		"fremovexattr",
		"fsconfig",
		"fsetxattr",
		"fsmount",
		"fsopen",
		"fspick",
		"fstat",
		"fstatfs",
		"fsync",
		"ftruncate",
		"futex",
		"futimesat",
		"getcpu",
		"getcwd",
		"getdents",
		"getdents64",
		"getegid",
		"geteuid",
		"getgid",
		"getgroups",
		"getitimer",
		"getpeername",
		"getpgid",
		"getpgrp",
		"getpid",
		"getppid",
		"getpriority",
		"getrandom",
		"getresgid",
		"getresuid",
		"getrlimit",
		"get_robust_list",
		"getrusage",
		"getsid",
		"getsockname",
		"getsockopt",
		"get_thread_area",
		"gettid",
		"gettimeofday",
		"getuid",
		"getxattr",
		"inotify_add_watch",
		"inotify_init",
		"inotify_init1",
		"inotify_rm_watch",
		"io_cancel",
		"ioctl",
		"io_destroy",
		"io_getevents",
		"io_pgetevents",
		"ioprio_get",
		"ioprio_set",
		"io_setup",
		"io_submit",
		"io_uring_enter",
		"io_uring_register",
		"io_uring_setup",
		"kill",
		"lchown",
		"lgetxattr",
		"link",
		"linkat",
		"listen",
		"listxattr",
		"llistxattr",
		"lremovexattr",
		"lseek",
		"lsetxattr",
		"lstat",
		"madvise",
		"membarrier",
		"memfd_create",
		"migrate_pages",
		"mincore",
		"mkdir",
		"mkdirat",
		"mknod",
		"mknodat",
		"mlock",
		"mlock2",
		"mlockall",
		"mmap",
		"modify_ldt",
		"mprotect",
		"mq_getsetattr",
		"mq_notify",
		"mq_open",
		"mq_timedreceive",
		"mq_timedsend",
		"mq_unlink",
		"mremap",
		"msgctl",
		"msgget",
		"msgrcv",
		"msgsnd",
		"msync",
		"munlock",
		"munlockall",
		"munmap",
		"nanosleep",
		"newfstatat",
		"open",
		"openat",
		"open_tree",
		"pause",
		"pidfd_open",
		"pidfd_send_signal",
		"pipe",
		"pipe2",
		"poll",
		"ppoll",
		"prctl",
		"pread64",
		"preadv",
		"preadv2",
		"prlimit64",
		"pselect6",
		"pwrite64",
		"pwritev",
		"pwritev2",
		"read",
		"readahead",
		"readlink",
		"readlinkat",
		"readv",
		"recvfrom",
		"recvmmsg",
		"recvmsg",
		"remap_file_pages",
		"removexattr",
		"rename",
		"renameat",
		"renameat2",
		"rmdir",
		"rseq",
		"rt_sigaction",
		"rt_sigpending",
		"rt_sigprocmask",
		"rt_sigqueueinfo",
		"rt_sigreturn",
		"rt_sigsuspend",
		"rt_sigtimedwait",
		"rt_tgsigqueueinfo",
		"sched_getaffinity",
		"sched_getattr",
		"sched_getparam",
		"sched_get_priority_max",
		"sched_get_priority_min",
		"sched_getscheduler",
		"sched_rr_get_interval",
		"sched_setaffinity",
		"sched_setattr",
		"sched_setparam",
		"sched_setscheduler",
		"sched_yield",
		"select",
		"semctl",
		"semget",
		"semop",
		"semtimedop",
		"sendfile",
		"sendmmsg",
		"sendmsg",
		"sendto",
		"setfsgid",
		"setfsuid",
		"setgid",
		"setgroups",
		"setitimer",
		"setpgid",
		"setpriority",
		"setregid",
		"setresgid",
		"setresuid",
		"setreuid",
		"setrlimit",
		"set_robust_list",
		"setsid",
		"setsockopt",
		"set_thread_area",
		"set_tid_address",
		"setuid",
		"setxattr",
		"shmat",
		"shmctl",
		"shmdt",
		"shmget",
		"sigaltstack",
		"signalfd",
		"signalfd4",
		"socket",
		"socketpair",
		"splice",
		"stat",
		"statfs",
		"statx",
		"symlink",
		"symlinkat",
		"sync",
		"sync_file_range",
		"syncfs",
		"sysinfo",
		"syslog",
		"tee",
		"tgkill",
		"time",
		"timer_create",
		"timer_delete",
		"timerfd_create",
		"timerfd_gettime",
		"timerfd_settime",
		"timer_getoverrun",
		"timer_gettime",
		"timer_settime",
		"times",
		"tkill",
		"truncate",
		"umask",
		"uname",
		"unlink",
		"unlinkat",
		"utime",
		"utimensat",
		"utimes",
		"vfork",
		"vhangup",
		"vmsplice",
		"wait4",
		"waitid",
		"write",
		"writev",
	}
}

// Used docker's set of blocked system calls as starting point:
// https://docs.docker.com/engine/security/seccomp/
var highRiskSystemCalls = map[string]bool{
	"acct":              true,
	"add_key":           true,
	"bpf":               true,
	"clock_adjtime":     true,
	"clock_settime":     true,
	"create_module":     true,
	"delete_module":     true,
	"finit_module":      true,
	"get_kernel_syms":   true,
	"get_mempolicy":     true,
	"init_module":       true,
	"ioperm":            true,
	"iopl":              true,
	"kcmp":              true,
	"kexec_file_load":   true,
	"kexec_load":        true,
	"keyctl":            true,
	"lookup_dcookie":    true,
	"mbind":             true,
	"mount":             true,
	"move_pages":        true,
	"name_to_handle_at": true,
	"nfsservctl":        true,
	"open_by_handle_at": true,
	"perf_event_open":   true,
	"personality":       true,
	"pivot_root":        true,
	"process_vm_readv":  true,
	"process_vm_writev": true,
	"ptrace":            true,
	"query_module":      true,
	"quotactl":          true,
	"reboot":            true,
	"request_key":       true,
	"set_mempolicy":     true,
	"setns":             true,
	"settimeofday":      true,
	"stime":             true,
	"swapoff":           true,
	"swapon":            true,
	"_sysctl":           true,
	"sysfs":             true,
	"umount2":           true,
	"umount":            true,
	"unshare":           true,
	"uselib":            true,
	"userfaultfd":       true,
	"ustat":             true,
	"vm86old":           true,
	"vm86":              true,
}
