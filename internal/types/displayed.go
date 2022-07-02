package types

type DisplayableDashboard struct {
	ID         string
	Title      string
	Categories []DisplayableCategory
}

type DisplayableCategory struct {
	ID, Title, Icon, IconURL string
	Entries                  []DisplayableEntry
}

type DisplayableEntry struct {
	ID, Name, Url, Icon, IconURL string
	UseNewTab, Health            bool
}
