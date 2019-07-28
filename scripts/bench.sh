#!/usr/bin/env bash

set -e

ab -n 10000 -c 1000 http://0.0:8080/
