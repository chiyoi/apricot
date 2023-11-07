#!/bin/sh
scripts=$0
cd $(dirname $(realpath $scripts)) || return
usage () {
    pwd
    echo "Scripts:"
    echo "$scripts tidy"
    echo "    Tidy go module."
}

tidy () {
    go mod tidy
}

case "$1" in
""|-h|-help|--help)
usage
exit
;;
tidy) ;;
*)
usage
exit 1
;;
esac

$@

