package main

import (
	"flag"
	"fmt"
	"github.com/mitchellh/go-homedir"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"strings"
)

var kubernetes_config string
var kpass_val bool
var verbose bool
var uri string
var passcount, passfail int

var help_text = `this is a validator for kubeapps:
	kaval 
		-c [Path to config file]
		-verbose
		-url uri to reach cluster, default is localhost

default path is .kube/config`

func parseFlags() bool {
	flag.StringVar(&kubernetes_config, "c", "~/.kube/config", "path to kubeconfig file")
	flag.BoolVar(&verbose, "verbose", false, "verbose on/off")
	flag.StringVar(&uri, "url", "http://localhost", "base uri to test")

	help := flag.Bool("help", false, "display help")

	flag.Parse()
	return *help
}

func kubernetesClient() (*kubernetes.Clientset, error) {

	cfg := kubernetes_config

	//replace home directory shortcut
	if strings.Contains(cfg, "~/") {
		d, err := homedir.Dir()
		if err != nil {
			return nil, err
		}
		cfg = strings.Replace(cfg, "~/", d+"/", 1)
	}

	config, err := clientcmd.BuildConfigFromFlags("", cfg)
	if err != nil {
		return nil, err
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return clientset, nil
}

func main() {
	defer Report()
	kpass_val = true
	help := parseFlags()
	if help {
		fmt.Println(help_text)
		os.Exit(1)
	}
	fmt.Println("kubeconfig path is:", kubernetes_config)
	client, err := kubernetesClient()
	if err != nil {
		fmt.Println("Error in kubernetes config:", err)
		os.Exit(2)
	}

	namespaces := []string{"kubeless", "kubeapps", "kube-system"}
	for _, n := range namespaces {
		CheckPods(n, client)
		CheckEndpoints(n, client)
	}
	paths := []string{"/", "/api/v1/repos", "/kubeless"}
	for _, p := range paths {
		PingPath(p)
	}
}
