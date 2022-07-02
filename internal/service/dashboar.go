package service

import "github.com/donreno/dashboar/internal/types"

type DashboarService interface {
	LoadDashboard() *types.Dashboard
}

type dashboarService struct{}

func New() DashboarService {
	return new(dashboarService)
}

func (d *dashboarService) LoadDashboard() *types.Dashboard {
	return loadMockDashboard()
}

func loadMockDashboard() *types.Dashboard {
	dashboard := &types.Dashboard{
		Title: "My beautiful dashboard",
		Groups: []types.DashboardGroup{
			{
				Name: "Networking",
				Icon: "bi-hdd-network",
				Entries: []types.DashboardEntry{
					{
						Name:      "PI-hole",
						Icon:      "bi-ethernet",
						Url:       "http://192.168.4.23:9000",
						UseNewTab: true,
					},
					{
						Name:      "NAS",
						Icon:      "bi-hdd-stack",
						Url:       "http://192.168.4.23:5000",
						UseNewTab: true,
					},
				},
			},
			{
				Name: "Coding",
				Icon: "bi-code-slash",
				Entries: []types.DashboardEntry{
					{
						Name:      "Github",
						Icon:      "bi-file-code",
						Url:       "http://github.com/donreno",
						UseNewTab: true,
					},
				},
			},
			{
				Name: "Networking",
				Icon: "bi-hdd-network",
				Entries: []types.DashboardEntry{
					{
						Name:      "PI-hole",
						Icon:      "bi-ethernet",
						Url:       "http://192.168.4.23:9000",
						UseNewTab: true,
					},
					{
						Name:      "NAS",
						Icon:      "bi-hdd-stack",
						Url:       "http://192.168.4.23:5000",
						UseNewTab: true,
					},
					{
						Name:      "Github",
						Icon:      "bi-file-code",
						Url:       "http://github.com/donreno",
						UseNewTab: true,
					},
				},
			},
			{
				Name: "Coding 2",
				Icon: "bi-code-slash",
				Entries: []types.DashboardEntry{
					{
						Name:      "Github",
						Icon:      "bi-file-code",
						Url:       "http://github.com/donreno",
						UseNewTab: true,
					},
				},
			},
		},
	}

	return dashboard
}
