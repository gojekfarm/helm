package api

import (
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/release"
)

type Upgrader struct {
	*action.Upgrade
}

type History struct {
	*action.History
}

type upgraderunner interface {
	Run(name string, chart *chart.Chart, vals map[string]interface{}) (*release.Release, error)
}

type historyrunner interface {
	Run(name string) ([]*release.Release, error)
}

func (u *Upgrader) SetConfig(cfg ReleaseConfig) {
	u.Namespace = cfg.Namespace
	u.ChartPathOptions.Version = cfg.Version
	u.Install = cfg.Install
}

func (u *Upgrader) GetInstall() bool {
	return u.Install
}

func (u *Upgrader) UpgradeLocateChart(name string, settings *cli.EnvSettings) (string, error) {
	return u.LocateChart(name, settings)
}

func (h *History) SetConfig() {
	h.Max = 1
}

func NewUpgrader(au *action.Upgrade) *Upgrader {
	return &Upgrader{au}
}

func NewHistory(ah *action.History) *History {
	return &History{ah}
}
