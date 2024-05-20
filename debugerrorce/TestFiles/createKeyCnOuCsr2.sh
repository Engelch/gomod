#!/usr/bin/env bash

function err()
{
	echo $* > /dev/stderr
}

function usage() {
    1>&2 echo usage $(basename $0) '[ -4 ]' CN OU outputfile
    1>&2 echo '-4 for 4096bit certs, otherwise 2048bit be default'
    1>&2 echo version 0.6.0
}

if [ $# -lt 3 -o $# -gt 4 ] ; then
    err The number of arguments is $#
    usage
	exit 1
fi

outstr=
keys=2048
if [ "$1" = -4 ] ; then
    [ $# -ne 4 ] && 1>&2 echo wrong number of arguments. && usage && exit 1
    shift
    keys=4096
    outstr=.4k
fi

cn="$1"
ou="$2"
key="$(basename $3 .csr)"

if [ -f "$key"$outstr.key -o -f "$key$outstr.csr" ] ; then
	echo artifacts found. Not overwriting.
	exit 2
fi

openssl req -newkey rsa:$keys -nodes -keyout ${key}$outstr.key -subj "/C=CH/O=Schindler Digital/OU=$ou/CN=$cn" -out ${key}$outstr.csr
