Pico Microservices walkthrough
====

This project is a very small microservice project


## The foo service
The foo service is a simple phyton project that use flask as web server.

The snippet below allow to build Docker image and run the container

``` bash
docker build -t genocs/foo foo/.
docker run --name genocs-foo-container -p 5200:5000 genocs/foo
```


# Run foo service into kubernetes

Create the namespace
``` bash
kubectl create namespace genocs
kubectl create -f namespace-dev.json
```


``` bash
kubectl apply -f foo.yaml
```

