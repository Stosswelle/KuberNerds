# KuberNerds
CSE561 Winter2021

Common Commands

eval $(minikube docker-env)

docker build -f Dockerfile -t server:latest .

k create -f deployment.yaml

k port-forward service/service-server 63383:1453

minikube service service-server --url

kubectl exec -it 461server-85b7dbd976-rbrq9 bash

client -> 12345 tunnel -> 1453 K8s -> 12235 Docker