package main

import (
	"k8s.io/client-go/kubernetes"
	certs "k8s.io/client-go/kubernetes/typed/certificates/v1beta1"
	core "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/rest"
	"log"
)

func approveCsrs(csr certs.CertificateSigningRequestInterface, pod core.PodInterface) bool {
	return true
}

func main() {
	// Initialise configuration
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// Set up the csr client
	certificateRestClient := clientset.Certificates().RESTClient()
	certificateClient := certs.New(certificateRestClient)
	csr := certificateClient.CertificateSigningRequests()

	// FIXME: only working in the default namespace
	pod := clientset.CoreV1().Pods("default")

	log.Fatalln(approveCsrs(csr, pod))
}
