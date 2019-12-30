# kubernetes-assignment

This example creates two pods - One deploys a webpage page to take inputs from the user. The other pod had Mysql database running on it.
The input entered by user in webpage is added to the database running in different pod.

Create docker image using the available DockerFile :
`docker build -t go-webapp .`

Commands to create required deployments and services :
```
kubectl.exe apply -f "web-app-go.yaml"
kubectl.exe apply -f "mysql-deployment.yaml"
```

Get the URL for the hosted webpage :
`minikube.exe service web-service --url`