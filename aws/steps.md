To create: 
```
eksctl create cluster \
    --region us-west-2 \
    --name istio-on-eks \
    --nodes 2 \
    --ssh-public-key "~/.ssh/id_rsa.pub"
```

IAM stuff: 
TODO


For prometheus: 
See namespaces
helm install prometheus prometheus-community/prometheus --namespace=prom

Expose port using expose prometheus-server-np, then open up on security groups for inbound rules

To delete the cluster: 
eksctl delete cluster --region=us-west-2 --name=istio-on-eks

Then you need to shut down: 
- VPCs
- 
