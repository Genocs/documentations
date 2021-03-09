Pico Microservices walkthrough
====

This project is a very small microservice project


## The python-flask service
The python-flask service is a simple phyton project that use flask as web server.

``` bash
docker build -t genocs/python-flask python-flask/.
docker run --name genocs-python-flask-container -p 5200:5000 genocs/python-flask
```

## The nodejs service
The nodejs service is a simple nodejs project that use a standard web server.

``` bash
docker build -t genocs/nodejs nodejs/.
docker run --name genocs-nodejs-container -p 5300:3000 genocs/nodejs
```


# Run python-flask service into kubernetes

Create the namespace
``` bash
kubectl create namespace genocs
kubectl create -f namespace-dev.json
```

Run the container
``` bash
kubectl apply -f python-flask.yaml
```

