package http

import (
	"fmt"

	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/release"
)

func Install(cfg *action.Configuration, chartName, relName string) (*release.Release, error) {
	opts := make(map[string]interface{})

	settings := cli.New()
	install := action.NewInstall(cfg)
	install.ReleaseName = relName
	install.Namespace = "default"

	fmt.Printf("%+v %#v\n", install, settings)
	cp, err := install.ChartPathOptions.LocateChart(chartName, settings)
	if err != nil {
		return nil, err
	}
	var requestedChart *chart.Chart
	if requestedChart, err = loader.Load(cp); err != nil {
		return nil, err
	}

	release, err := install.Run(requestedChart, opts)
	if err != nil {
		return nil, err
	}
	//	deal with dependent charts later
	return release, nil
}
