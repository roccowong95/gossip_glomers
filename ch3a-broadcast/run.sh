#!/usr/bin/env bash
set -x
go install .
maelstrom test -w broadcast --bin $(which ch3a) --node-count 1 --time-limit 20 --rate 10