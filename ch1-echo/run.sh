#!/usr/bin/env bash

set -x
go install .
maelstrom test -w echo --bin $(which ch1) --node-count 1 --time-limit 10