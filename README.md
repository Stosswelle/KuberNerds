# KuberNerds
CSE561 Winter2021

# Common Commands

- Start Minikube
    - Start Docker
    - ```minikube start```

- Run Docker in Minikube
    - ```eval $(minikube docker-env)```

- Build Docker image
    - Edit Dockerfile
    - ```docker build -f Dockerfile -t <image_name>:<image_tag> .```
    <!-- docker build -f Dockerfile -t server:latest . -->

- Create K8s deployment
    - Edit deployment.yaml
    - ```kubectl create -f deployment.yaml```

- Port forwarding
    - ```kubectl port-forward service/<service_name> <external_port>:<internal_port>```
    <!-- kubectl port-forward service/service-server 12345:1453 -->

- Minikube tunneling
    - ```minikube service <service_name> --url```
    <!-- minikube service service-server --url -->

- ssh into k8s pod
    - ```kubectl exec -it <pod_name> -- bash```
    <!-- kubectl exec -it server-65987794c7-5kldn -- bash -->

- use Wireshark in specified pod
    - ```kubectl sniff <pod_name> <container_name>```
    <!-- kubectl sniff server-65987794c7-5kldn server -->