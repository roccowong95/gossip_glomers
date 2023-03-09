#!/usr/bin/env bash
set -x
go install .
maelstrom test -w broadcast --bin $(which ch3b) --node-count 5 --time-limit 20 --rate 10