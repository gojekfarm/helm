package main

import (
	"fmt"
	"log"
	"os"

	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/gates"
	"helm.sh/helm/v3/pkg/http"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

var (
	settings = cli.New()
)

func debug(format string, v ...interface{}) {
	format = fmt.Sprintf("[debug] %s\n", format)
	log.Output(2, fmt.Sprintf(format, v...))
}

const FeatureGateOCI = gates.Gate("HELM_EXPERIMENTAL_OCI")

func main() {
	actionConfig := new(action.Configuration)
	for k, v := range settings.EnvVars() {
		fmt.Println(k, v)
	}
	if err := actionConfig.Init(settings.RESTClientGetter(), settings.Namespace(), os.Getenv("HELM_DRIVER"), debug); err != nil {
		log.Fatalf("error getting configuration: %v", err)
		return
	}
	if _, err := actionConfig.KubernetesClientSet(); err != nil {
		fmt.Println("hereeee...........", err)
		return
	}

	//doInstall(actionConfig, "h3-redis")
	doList(actionConfig)
}

func doInstall(cfg *action.Configuration, relname string) {
	chart := "/Users/dineshkumar/src/github.com/helm/charts/stable/redis"
	_, err := http.Install(cfg, chart, relname)
	if err != nil {
		fmt.Println("error installing chart....", err)
		return
	}
	fmt.Printf("installed chart: %s release: %s successfully", chart, relname)
}

func doList(cfg *action.Configuration) {
	_, err := http.List(cfg)
	if err != nil {
		panic(err)
	}
}
