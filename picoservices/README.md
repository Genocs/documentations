Pico Microservices walkthrough
====

This project is a very small microservice project containing different programming languages. They are:
1. nodejs
2. python
3. golang
4. dotnet


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

# Use the following command if u like to use a configuration file
kubectl create -f namespace-dev.json
```

Run the container
``` bash
kubectl apply -f python-flask.yaml
kubectl apply -f nodejs.yaml
kubectl apply -f golang.yaml
kubectl apply -f dotnet.yaml
```

After the creation the services are running on:
1. python-flask:5300
2. nodejs:5400
3. golang:5500
4. dotnet:5600

# Run with Dapr

How to run the service with Dapr 
``` bash
# Run the nodejs Dapr app
dapr run --app-id nodeapp --app-port 3000 --dapr-http-port 3500 node nodejs/hello-server.js

# Run the python Dapr app
dapr run --app-id pythonapp --app-port 5000 --dapr-http-port 3600 python python-flask/app.py

# Run the golang Dapr app
dapr run --app-id golangapp --app-port 8080 --dapr-http-port 3700 go run go-docker/main.go

# Check the app running with dapr
dapr list

# Call Get 
dapr invoke --app-id nodeapp --method ping -v GET

# Call Post
dapr invoke --app-id nodeapp --method post_my_data -d '{\"data\":{\"id\":\"Hello\"}}'
```


http://localhost:3500/v1.0/invoke/nodeapp/method/post_my_data