#!/bin/bash

set -e

cd $(dirname $0)

export COTENOER_SESSION=$(mktemp)

bin/cotenoer add self docker-compose.yml
bin/cotenoer add playground playground.docker-compose.yml
bin/cotenoer hud
