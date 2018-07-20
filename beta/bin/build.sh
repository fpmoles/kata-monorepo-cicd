#!/usr/bin/env bash

result=$(docker build -t $1:$2 . | grep 'Successfully built' | cut -d " " -f 3)
echo result

