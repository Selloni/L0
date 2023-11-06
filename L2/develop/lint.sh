#!/bin/bash

for ((i = 1; i < 12; i++ ))
do
    echo "dev$i"   
    if (($i < 10)); then
        cd dev0$i && go vet && golint
        cd ..
    else 
        cd dev$i && go vet && golint
        cd ..        
    fi
done

#GO111MODULE=off go get -u golang.org/x/lint/golint
