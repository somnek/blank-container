#!/bin/bash

set -e

container_name=empty-container
image_name=empty-container


if [ "$1" = "--run" ] || [ -z "$1" ]; then

  # container exist (exited), run it
  if [ "$(docker ps -aq --filter name=$container_name -f status=exited)" ]; then
    docker start $container_name
    echo "🚜  running in background ..."
  # container exist (running), ignore
  elif [ "$(docker ps -aq --filter name=$container_name)" ]; then
    echo "container is already running ..."
  else 
    # check if image already exist
    if [ "$(docker images -q $image_name)" ]; then 
      echo "running it again ..."
      docker run -d --name $container_name $image_name 
    else
      docker build -t $image_name -f Dockerfile . && \
      docker run -d --name $container_name $image_name 
    fi
      echo "🚜  running in background ..."
  fi

elif [ "$1" = "--stop" ]; then
  docker rm -f $container_name && echo "📮 stopped ..."

elif [ "$1" = "--clean" ]; then
  if [ "$(docker ps -aq --filter name=$container_name)" ]; then
    docker rm -f $container_name
  fi

  docker rmi $image_name && echo "🧹 removed ..."

else
  echo "️❓ unknown command"

fi