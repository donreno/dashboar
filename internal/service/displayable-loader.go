package service

import (
	"github.com/donreno/dashboar/internal/types"
	"github.com/google/uuid"
)

type DisplayableBoarBuilder func(configuredBoard *types.Dashboard) (*types.DisplayableDashboard, error)

func MakeDisplayableBOARBuilder() DisplayableBoarBuilder {
	return func(configuredBoard *types.Dashboard) (dashboar *types.DisplayableDashboard, err error) {
		dashboar = new(types.DisplayableDashboard)

		dashboar.Title = configuredBoard.Title
		dashboar.Categories = []types.DisplayableCategory{}

		for _, category := range configuredBoard.Categories {
			mapCategory(category, dashboar)
		}

		return dashboar, nil
	}
}

func mapCategory(category types.Category, dashboar *types.DisplayableDashboard) {
	displayableCategory := types.DisplayableCategory{
		ID:      uuid.NewString(),
		Name:    category.Name,
		Icon:    category.Icon,
		Entries: []types.DisplayableEntry{},
	}

	for _, entry := range category.Entries {

		mapEntry(entry, &displayableCategory)
	}

	dashboar.Categories = append(dashboar.Categories, displayableCategory)
}

func mapEntry(entry types.Entry, displayableCategory *types.DisplayableCategory) {
	displayableEntry := types.DisplayableEntry{
		ID:        uuid.NewString(),
		Name:      entry.Name,
		Url:       entry.Url,
		UseNewTab: entry.UseNewTab,
		Health:    true, //Temporarily set to true
	}

	if entry.Icon == "favicon" {
		displayableEntry.IconURL = entry.Url + "/favicon.ico"
		displayableEntry.UseFavIcon = true
	} else {
		displayableEntry.Icon = entry.Icon
	}

	displayableCategory.Entries = append(displayableCategory.Entries, displayableEntry)
}
