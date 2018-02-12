# Technical Challenge - Lalamove
## Infrastructure Engineer

This exercise is designed to test a range of skills pertinent to being an infrastructure engineer here at Lalamove. These include, but are not limited to:
- Interacting with public APIs
- Packaging applications in Docker images
- Configuration in Kubernetes to run applications
- Writing complete, step-by-step instructions

Bonus points for:
- Providing appropriate RBAC privileges for this application
- Integration testing
- Full attack vector analysis on your developed application

# Preamble
There is a certificate API which allows you to provision certificates signed by the CA used internally in Kubernetes. To request one, you need to create a Certificate Signing Request, which then needs to be approved.

You can find a guide here: https://kubernetes.io/docs/tasks/tls/managing-tls-in-a-cluster/

There is already a project which creates these certificate signing requests for applications packaged as pods, that you can find here: https://github.com/kelseyhightower/certificate-init-container which also comes with a handy test project so you can make sure it works.
The deployment file, you can find here: https://raw.githubusercontent.com/kelseyhightower/certificate-init-container/master/deployments/tls-app.yaml

## The Challenge
Your challenge is to create a small application that listens to these requests and then approves them, *but only if the common name of the certificate resolves through DNS to the IP actually associated with the pod requesting the certificate*. How exactly you determine this is up to you.

This project should be written in Golang, and implemented using this client library:
https://github.com/kubernetes/client-go
In this repository, you'll find a skeleton main.go file with some of the imports you need to get started and a simple Dockerfile.

This is the API exposed by Kubernetes for doing so: 
https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.9/#certificatesigningrequest-v1beta1-certificates
You most likely need to interact directly with kubectl to approve the certificates in accordance with the guide `Managing TLS in a Cluster`.

We suggest using Minikube, though there is this issue: https://github.com/kubernetes/minikube/issues/1647
which means you need to start Minikube with the following command:
```
minikube start --extra-config=controller-manager.ClusterSigningCertFile="/var/lib/localkube/certs/ca.crt" --extra-config=controller-manager.ClusterSigningKeyFile="/var/lib/localkube/certs/ca.key"
```

Your project should be delivered as a compressed file containing a git repository with the code and Dockerfile that you are able to run in the root folder:
```
docker build . -t testproject:latest
```

and manifests to install the application in Kubernetes by running:
```
kubectl create -f manifest.yaml
```

To get your image into the minikube cluster, the easiest way is to connect your docker client to your minikube cluster by running:
```
eval $(minikube docker-env)
```

Then build it as normal, you can then reference your image by using `testproject:latest` in your manifest. Just make sure you use `pullPolicy: IfNotPresent`


Good luck! And ask us if you have any questions. :)
