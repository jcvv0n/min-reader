docker build -t min-reader:v1 .
docker run --name min-reader-v1 -p 7676:7676 -d min-reader:v1
docker image prune