#!/bin/csh
sudo docker rmi -f danding19931115/movie-micro:v1.0
sudo docker rmi -f danding19931115/movie-micro:latest

cd kubernetes
sudo docker build -t danding19931115/movie-micro:v1.0 .
sudo docker build -t danding19931115/movie-micro:latest .
