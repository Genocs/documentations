Pico Microservices walkthrough
====

This project is a very small microservice project


## The python-flask service
The python-flask service is a simple phyton project that use flask as web server.

``` bash
docker build -t genocs/python-flask python-flask/.
docker run --name genocs-python-flask-container -p 5200:5000 genocs/python-flask
curl http://localhost:5200
```

## The nodejs service
The nodejs service is a simple nodejs project that use a standard web server.

``` bash
docker build -t genocs/nodejs nodejs/.
docker run --name genocs-nodejs-container -p 5300:3000 genocs/nodejs
curl http://localhost:5300
```


## The golang service
The golang service is a simple go project that use a standard web server.
[Reference](https://www.callicoder.com/docker-golang-image-container-example/)

To run the project locally
``` bash
go mod init github.com/genocs/go-docker
go get
go run main.go

curl http://localhost:8080?name=Rajeev
```

``` bash
docker build -t genocs/go-docker go-docker/.
docker run --name genocs-go-docker-container -p 5400:8080 genocs/go-docker
curl http://localhost:5400?name=Giovanni
```

# Run services into kubernetes

Create the namespace
``` bash
kubectl create namespace genocs
kubectl create -f namespace-dev.json
```

Run the container
``` bash
kubectl apply -f python-flask.yaml
kubectl apply -f nodejs.yaml
kubectl apply -f golang.yaml
```

After the creation the services are 
python-flask:5300
nodejs:5400
golang:5500


