package service

import "github.com/donreno/dashboar/internal/types"

type DashboardRetriever func() (*types.Dashboard, error)

func GetDashboard() (*types.Dashboard, error) {
	return loadMockDashboard()
}

func loadMockDashboard() (*types.Dashboard, error) {
	dashboard := &types.Dashboard{
		Title: "My beautiful dashboard",
		Categories: []types.Category{
			{
				Name: "Networking",
				Icon: "bi-hdd-network",
				Entries: []types.Entry{
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
				Entries: []types.Entry{
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
				Entries: []types.Entry{
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
				Entries: []types.Entry{
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

	return dashboard, nil
}
