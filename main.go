package main

import (
	"log"

	"github.com/simonswine/kube-lego/acme"
	client "k8s.io/kubernetes/pkg/client/unversioned"
)

type KubeLego struct {
	LegoClient     *acme.Client
	LegoURL        string
	LegoEmail      string
	LegoSecretName string
	legoUser       *LegoUser
	KubeClient     *client.Client
}

func NewKubeLego() *KubeLego {
	return &KubeLego{
		LegoURL: "https://acme-v01.api.letsencrypt.org/directory",
	}
}

func main() {
	log.Print("kube-lego starting")

	kl := NewKubeLego()

	err := kl.InitKube()
	if err != nil {
		log.Fatal(err)
	}

	err = kl.InitLego()
	if err != nil {
		log.Fatal(err)
	}

	err = kl.ProcessIngress()
	if err != nil {
		log.Fatal(err)

	}
}
