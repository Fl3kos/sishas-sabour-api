#!/bin/bash

param=$1

run_api(){
    go run ./cmd/api/main.go
}
run_test(){
  echo "test completed"
}
case $param in
    "a" | "api" | "-a")
        run_api &
         ;;
    "b" | "build" | "-b")
        echo "Project compiled"
        echo "/ __| | | |/ __/ __/ _ \/ __/ __|"
        echo "\__ \ |_| | (_| (_|  __/\__ \__ \\"
        echo "|___/\__,_|\___\___\___||___/___/"
        ;;
    "cc" | "clear-cache" | "--clear-cache")
        go clean -testcache
        ;;
    "t" | "test" | "-t")
        go clean -testcache
        run_test
        ;;
    "p" | "pull" | "-p")
        git pull
        ;;
    "h" | "??" | "help" | "-q" | "--help")
        echo "Help:"
        echo "i: Inicialice the project, when download the project is the fist choice to use"
        echo "a: run the api system and swagger"
        echo "b: compile the project"
        echo "c: to clear files folder"
        echo "ca: to clear all project"
        echo "cc: to clear the cache"
        echo "l: run the layouts constructor sql file"
        echo "t: run the test"
        echo "p: pull repo, delete and create new folders"
        echo "h: help"
            ;;
esac