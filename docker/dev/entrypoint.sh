#!/bin/sh
set -x

reflex -s -r '(\.go$|go\.mod)' go run "$1" "$2"