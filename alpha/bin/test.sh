#!/usr/bin/env bash

result=$(go test | grep 'PASS\|FAIL')
if [ "$result" = "PASS" ]
then
    echo TRUE
else
    echo FALSE
fi

