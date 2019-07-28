#!/usr/bin/env bash

set -e

ab -n 1000 -c 50 http://0.0:8080/ > .results

cat .results | grep '#/sec'
