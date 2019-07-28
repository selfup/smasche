#!/usr/bin/env bash

set -e

ab -n 20000 -c 100 http://0.0:8080/ > .results

cat .results | grep '#/sec'
