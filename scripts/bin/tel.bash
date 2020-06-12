#!/usr/bin/env bash
# A port of plan9 'tel' program

for var in "$@"
do
    if test -f "$HOME/.tel"; then
        grep -i $1 $HOME/.tel
    fi

    grep -hi $1 /usr/lib/tel /usr/lib/areacodes
done

exit


