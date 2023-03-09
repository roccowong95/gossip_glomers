#!/usr/bin/env bash
set -x
go install .
maelstrom test -w unique-ids --bin $(which maelstrom-uid) --time-limit 30 --rate 1000 --node-count 3 --availability total --nemesis partition