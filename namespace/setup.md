
# How to set up namespace and Prometheus

```bash
    k describe po | egrep "^Name|^Node"

    k create -f namespace-prom.json

    k get namespaces --show-labels

    k config current-context

    k config set-context prom --namespace=prom --cluster=minikube --user=minikube

    k config use-context prom

    k config current-context

    helm repo add prometheus-community https://prometheus-community.github.io/helm-charts

    helm install prometheus prometheus-community/prometheus --namespace=prom 

    k get po -n=prom

    kubectl expose service prometheus-server --type=NodePort --target-port=9090 --name=prometheus-server-np

    helm uninstall prometheus

    k config use-context minikube

    minikube service prometheus-server-np --namespace=prom

```

Set up expose service. Allow for inbound traffic on assigned port for prometheus-server-np. Access. Profit. 

cpu usage query:

```promsql
sum(rate(container_cpu_usage_seconds_total{name!="",name !~ ".*POD.*", id !~".*docker.*",pod=~"review.*|rating.*|detail.*|productpage.*"}[2m]))

sum(container_memory_usage_bytes{name!="",name !~ ".*POD.*", id !~".*docker.*",pod=~"review.*|rating.*|detail.*|productpage.*"})
```
