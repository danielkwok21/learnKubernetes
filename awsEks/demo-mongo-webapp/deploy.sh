kubectl delete all --all
kubectl delete all --all
kubectl apply -n demo-mongo-webapp -f mongo-config.yaml
kubectl apply -n demo-mongo-webapp -f mongo-secret.yaml
kubectl apply -n demo-mongo-webapp -f mongo.yaml
kubectl apply -n demo-mongo-webapp -f webapp.yaml