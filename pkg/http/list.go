package http

import (
	"fmt"

	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/release"
)

func List(cfg *action.Configuration) ([]*release.Release, error) {
	list := action.NewList(cfg)

	list.SetStateMask()

	results, err := list.Run()
	if err != nil {
		return nil, err
	}
	for _, res := range results {
		fmt.Printf("res: %+v ns:%v\n", res.Name, res.Namespace)
	}
	return results, err
}
