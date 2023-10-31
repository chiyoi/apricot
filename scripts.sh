#!/bin/sh
cd $(dirname $(realpath $0)) || return
usage() {
    pwd
    echo "Scripts:"
    echo "./scripts.sh tidy"
    echo "    Tidy go modules."
}

tidy_all() {
    for m in kitsune logs neko sakana; do
    cd $m
    go mod tidy
    cd ..
    done
}

if test -z "$1" -o -n "$(echo "$1" | grep -Ex '\-{0,2}h(elp)?')"; then
usage
exit
fi

case "$1" in
tidy_all) ;;
*)
usage
exit 1
;;
esac

$@