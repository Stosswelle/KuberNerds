
# How to set up namespace and Prometheus

```bash
    k create -f namespace-prom.json

    k get namespaces --show-labels

    k config current-context

    k config set-context prom --namespace=prom --cluster=minikube --user=minikube

    k config use-context prom

    k config current-context

    helm repo add prometheus-community https://prometheus-community.github.io/helm-charts

    helm install prometheus prometheus-community/prometheus

    kubectl expose service prometheus-server --type=NodePort --target-port=9090 --name=prometheus-server-np

    k config use-context minikube

    minikube service prometheus-server-np --namespace=prom

```

cpu usage query:

```promsql
rate(container_cpu_usage_seconds_total{name!="",name !~ ".*POD.*", id !~".*docker.*",pod=~"review.*|rating.*|detail.*|productpage.*"}[5m])
```
