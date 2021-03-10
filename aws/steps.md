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

kubectl edit -n kube-system configmap/aws-auth
And put in members

For prometheus: 
See namespaces
helm install prometheus prometheus-community/prometheus --namespace=prom

Expose port using expose prometheus-server-np, then open up on security groups for inbound rules

NAME                            TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)        AGE     SELECTOR
prometheus-alertmanager         ClusterIP   10.100.10.102    <none>        80/TCP         4m33s   app=prometheus,component=alertmanager,release=prometheus
prometheus-kube-state-metrics   ClusterIP   10.100.78.86     <none>        8080/TCP       4m33s   app.kubernetes.io/instance=prometheus,app.kubernetes.io/name=kube-state-metrics
prometheus-node-exporter        ClusterIP   None             <none>        9100/TCP       4m33s   app=prometheus,component=node-exporter,release=prometheus
prometheus-pushgateway          ClusterIP   10.100.168.167   <none>        9091/TCP       4m33s   app=prometheus,component=pushgateway,release=prometheus
prometheus-server               ClusterIP   10.100.204.235   <none>        80/TCP         4m33s   app=prometheus,component=server,release=prometheus
prometheus-server-np            NodePort    10.100.197.96    <none>        80:__32454__/TCP   3m29s   app=prometheus,component=server,release=prometheus

Access prometheus through Public DNS on Instances on EC2. 

To delete the cluster: 
eksctl delete cluster --region=us-west-2 --name=istio-on-eks

Then you need to shut down: 
- NATs
- VPCs
- EC2 volumes

