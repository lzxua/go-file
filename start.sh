#!/bin/bash
docker rm -f go-file
docker rmi go-file
docker build -t go-file .
docker run -p 8080:8080 --name go-file -v /www/static/imgs:/www/static/imgs/ -d go-file