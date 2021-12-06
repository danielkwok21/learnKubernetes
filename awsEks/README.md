# What is this
Playground for https://www.youtube.com/watch?v=QThadS3Soig

# Commands

## List all existing EKS clusters
```
aws eks list-clusters
```

## view clusters in detail
```
aws eks describe-cluster --name foo_cluster_name
```

## Create a new cluster, with vpc, subnets etc resources
```
aws eks create-cluster \
--name foo_new_cluster_name \
--role foo_aws_iam_role_arn \
--resources-vpc-config subnetIds=foo_subnetid1, foo_subnetid2,securityGroupIds=foo_securitygroupid1, endpointPublicAccess=true,endpointPrivateAccess=false
```

### (Example)
```
aws eks create-cluster \
--name my_eks_cluster \
--role arn:aws:iam::521197959066:role/myAmazonEKSClusterRole \
--resources-vpc-config subnetIds=subnet-0df0e7aa025e4a7c5,subnet-0b69e765dc93f5143,subnet-000445ca3b198a0c1,subnet-0030cdd711d03fb0b,securityGroupIds=sg-0ff5d7b1bf4b7eb71,endpointPublicAccess=true,endpointPrivateAccess=false
```

## Access cluster via kubeconfig
```
aws eks update-kubeconfig --name foo_created_cluster_name --region ap-southeast-2
```

## Creating EC2(s) and attaching to eks cluster
```
aws eks create-nodegroup \
--cluster-name foo_created_cluster_name \
--nodegroup-name foo_new_cluster_name \
--node-role foo_create_iam_arn_for_nodegroup \
--subnet foo_one_out_of_all_subnetid \
--disk-size 200 \
--scaling-config minSize=1,maxSize=2,desiredSize=1 \
--instance-types t2.small
```
### (Example)
```
aws eks create-nodegroup \
--cluster-name my_eks_cluster \
--nodegroup-name my_eks_nodegroup \
--node-role arn:aws:iam::521197959066:role/myAmazonEKSNodeRole \
--subnet subnet-0df0e7aa025e4a7c5 \
--disk-size 200 \
--scaling-config minSize=1,maxSize=2,desiredSize=1 \
--instance-types t2.small
```