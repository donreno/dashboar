package service

import "github.com/donreno/dashboar/internal/types"

type DashboardRetriever func() (*types.Dashboard, error)

func MakeDashboarRetriever(dashboar *types.Dashboard) DashboardRetriever {
	return func() (*types.Dashboard, error) {
		return dashboar, nil
	}
}
