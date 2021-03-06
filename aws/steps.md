To create: 
```
eksctl create cluster \
    --region us-west-2 \
    --name istio-on-eks \
    --nodes 2 \
    --ssh-public-key "~/.ssh/id_rsa.pub"
```

IAM stuff: 
- Set up groups with names, save access keys somewhere
- Have people log into aws with IAM usernames 
- Set up cluster using eksctl
- Set up kubeconfig to add them as members to a group

For prometheus: 
See namespaces
helm install prometheus prometheus-community/prometheus --namespace=prom

Expose port using expose prometheus-server-np, then open up on security groups for inbound rules

To delete the cluster: 
eksctl delete cluster --region=us-west-2 --name=istio-on-eks

Then you need to shut down: 
- NATs
- VPCs
- EC2 volumes

