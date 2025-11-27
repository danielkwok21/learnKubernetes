# what is this
repo to learn k8s.  
all lessons are in PR tab.

# directories
- `src` source code for a simple go server. we'll use this to build the image to be used for all lessons

# handy commands
```bash
# build & run docker image
docker build -f Dockerfile -t helloserver:<tagnumber> ./src
docker run -p 8080:8080 helloserver -d
curl http://localhost:8080

```