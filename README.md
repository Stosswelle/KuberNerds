# KuberNerds
CSE561 Winter2021

Common Commands

eval $(minikube docker-env)

docker build -f Dockerfile -t 461server:latest .

k create -f deployment.yaml

k port-forward service/service-461server 12345:1453

minikube service service-461server --url

kubectl exec -it 461server-85b7dbd976-rbrq9 bash

client -> 12345 tunnel -> 1453 K8s -> 12235 Docker