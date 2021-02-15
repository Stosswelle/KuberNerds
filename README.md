# KuberNerds
CSE561 Winter2021

# Common Commands

- Run docker in minikube

```eval $(minikube docker-env)```

- Build docker image

```docker build -f Dockerfile -t server:latest .```

- Create k8s deployment

```k create -f deployment.yaml```

- Port forwarding

```k port-forward service/service-server 63383:1453```

- Minikube tunneling

```minikube service service-server --url```

- ssh into k8s pod

```kubectl exec -it <pod> -- bash```

- use Wireshark in specified pod

```k sniff <pod> server```