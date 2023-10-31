#!/bin/sh
scripts=$0
cd $(dirname $(realpath $0)) || return
usage() {
    pwd
    echo "Scripts:"
    echo "$scripts tidy"
    echo "    Go mod tidy."
}

tidy() {
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
