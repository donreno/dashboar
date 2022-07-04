package service

import "github.com/donreno/dashboar/internal/types"

type DashboardRetriever func() (*types.DisplayableDashboard, error)

func MakeDashboarRetriever(dashboar *types.Dashboard, boarBuilder DisplayableBoarBuilder) DashboardRetriever {
	return func() (*types.DisplayableDashboard, error) {
		return boarBuilder(dashboar)
	}
}
