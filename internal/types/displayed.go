package types

type DisplayableDashboard struct {
	Title      string
	Categories []DisplayableCategory
}

type DisplayableCategory struct {
	ID, Name, Icon, IconURL string
	Entries                 []DisplayableEntry
}

type DisplayableEntry struct {
	ID, Name, Url, Icon, IconURL  string
	UseNewTab, UseFavIcon, Health bool
}
