#!/usr/bin/env bash

# This script builds this GO project for dev as well as prod.

rm -rfv ./bin/service

go build -o ./bin/service .

