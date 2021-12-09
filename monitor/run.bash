#!/bin/bash

docker-compose -f /home/icecream/Develop/go_dev/assignment2/standalone/docker-compose.yaml up -d --scale go_server="$1"
