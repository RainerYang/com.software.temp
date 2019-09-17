#!/bin/bash

CGO_ENABLE=0 GOOS=linux go build -o com.software.temp

docker build -t temp_scratch ./
