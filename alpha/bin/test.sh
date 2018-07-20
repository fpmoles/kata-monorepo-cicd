#!/usr/bin/env bash

result=$(go test | grep 'PASS\|FAIL')
if [ "$result" = "PASS" ]
then
    echo true
else
    echo false
fi

