#!/usr/bin/env bash
#


#
# a db-connect role can be specified as the 1st argument, e.g.: "$0" localadm
# support dry mode by -n as first arg

DRY=
if [ "$1" = '-n' ]; then
    1>&2 echo DRY run mode
    DRY=echo
    shift
else
    DRY=
fi

role=
[ -n "$1" ] && role="-$1" && echo 1>&2 role set to $role
[ -z "$1" ] && role="-localadm" && echo 1>&2 role set to $role
# if a role is set, then an s-link with the role name is expected in the CWD, else use the default db-connect.sh
$DRY ./db-connect${role}.sh < "20-createDB.sql"
