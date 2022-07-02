package types

type Dashboard struct {
	Title      string
	Categories []Category
}

type Category struct {
	Name, Icon string
	Entries    []Entry
}

type Entry struct {
	Name            string
	Url             string
	Icon            string
	IconURL         string
	HealthCheckPath string
	UseNewTab       bool
}
