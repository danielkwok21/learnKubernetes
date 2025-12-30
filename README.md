# what is this
repo to learn k8s.  
all lessons are in PR tab.

# directories
- `src` source code for a simple go server. we'll use this to build the image to be used for all lessons

# handy commands
```bash
# build & run docker image
docker build -f Dockerfile -t demo:latest ./src
docker run -p 80:80 demo -d
curl http://localhost:80


# start nginx in ingress-nginx namespace
helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
helm repo update
helm upgrade --install ingress-nginx ingress-nginx/ingress-nginx \
  --namespace ingress-nginx \
  --create-namespace
  --version 4.11.0
kubectl get pods -n ingress-nginx
```