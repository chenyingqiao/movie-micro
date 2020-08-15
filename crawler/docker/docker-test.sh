
sudo docker build -t crawler .
sudo docker run --env-file=docker.env -it --rm --name crawler-test crawler bash