#!/usr/bin/env bash
set -x
go install .
maelstrom test -w broadcast --bin $(which ch3d) --node-count 25 --time-limit 20 --rate 100 --latency 100
# maelstrom test -w broadcast --bin $(which ch3d) --node-count 25 --time-limit 20 --rate 100 --latency 100 --nemesis partition
