package main

import (
	"fmt"

	wh "github.com/cert-manager/cert-manager/pkg/acme/webhook"
	whapi "github.com/cert-manager/cert-manager/pkg/acme/webhook/apis/acme/v1alpha1"
	"k8s.io/client-go/kubernetes"
	restclient "k8s.io/client-go/rest"
)

type SolveConfig struct {
	ApiToken string

	SecretName string `json:"secret_name"`
	SecretKey  string `json:"secret_key"`

	// CreateRecord if set the solver will create a record after a successfull request.
	// If set to an IP address it'll create a new A record. if set to a hostname it'll
	// create a CNAME record.
	// it will create a record if there's no record already there. Set ForceCreateRecord
	// to true if you'd want to create a record regardless of there already being a
	// record with that name
	CreateRecord      string `json:"create_record"`
	ForceCreateRecord bool   `json:"force_create_record"`
}

// Solver is the ACME Hetzner DSN solver
type Solver struct {
	client *kubernetes.Clientset
	ApiKey string

	wh.Solver
}

// Name returns the name of the Solver
func (x Solver) Name() string {
	return `hetzner-dns`
}

// Present handles the DNS challenge with Hetzner DNS
func (x Solver) Present(cr *whapi.ChallengeRequest) error {
	if cr.Action == whapi.ChallengeActionPresent {
		return fmt.Errorf("error wrong action")
	}

	// Check if domain exists
	if err := checkDNSName(x.ApiKey, cr.DNSName); err != nil {
		return err
	}

	// present challenge

	return nil
}

// CleanUp cleans up the challenge record from Hetzner DNS
func (x Solver) CleanUp(cr *whapi.ChallengeRequest) error {
	if cr.Action == whapi.ChallengeActionCleanUp {
		return fmt.Errorf("error wrong action")
	}

	// Check if domain exists
	if err := checkDNSName(x.ApiKey, cr.DNSName); err != nil {
		return err
	}

	// cleanup challenge

	return nil
}

func (x *Solver) Initialize(kubeClientConfig *restclient.Config, stopCh <-chan struct{}) error {
	k8scli, err := kubernetes.NewForConfig(kubeClientConfig)
	if err != nil {
		return err
	}
	x.client = k8scli

	return nil
}

func checkDNSName(apiKey, dnsName string) error {

	return nil
}
