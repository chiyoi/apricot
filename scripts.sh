#!/bin/sh
usage() {
    dirname $(realpath $0)
    echo "Scripts:"
    echo "./scripts.sh tidy"
    echo "    Go mod tidy."
}

tidy() {
    go mod tidy
}

if test -z "$1" -o -n "$(echo "$1" | grep -Ex '\-{0,2}h(elp)?')"; then
usage
exit
fi

case "$1" in
tidy) ;;
*)
usage
exit 1
;;
esac

$@