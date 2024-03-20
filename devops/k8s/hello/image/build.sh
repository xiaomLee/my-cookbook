#!/bin/bash

if [ $# -eq 0 ];then
  echo "image tag is empty"
  exit 1
fi

docker build -t 850278995/web-demo:$1 .