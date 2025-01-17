#
# PROVIDE: tinyfilter
# REQUIRE: SERVERS
# KEYWORD: shutdown
#
# Add the following lines to /etc/rc.conf.local or /etc/rc.conf to enable tinyproxy:
# tinyfilter_enable (bool): Set to "NO" by default.
#                           Set it to "YES" to enable tinyproxy

. /etc/rc.subr

name="tinyfilter"
rcvar="tinyfilter_enable"

load_rc_config $name

# Make sure the pidfile matches what's in the config file.

: ${tinyfilter_enable="NO"}
: ${tinyfilter_pidfile="/var/run/${name}.pid"}

pidfile="${tinyfilter_pidfile}"
tinyfilter_log="/root/.tinyproxy/tinyfilter.log"

## Run

## See daemon man for details
## https://www.freebsd.org/cgi/man.cgi?query=daemon&sektion=8

command="/usr/sbin/daemon"
daemon_args="-P ${pidfile} -R2 -t \"${name} daemon\" -o ${tinyfilter_log} -m 3"
command_args="${daemon_args} /go/bin/tinyfilter web"

run_rc_command "${1}"
