# What is this directory
This is an example configmap in action

configmap_mvp.yaml is where the config is stored, while deployment.yaml would be consuming it

# How to run
```
$ kubectl apply -f deployment.yaml #this deployment would consume the configmap
```