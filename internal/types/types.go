package types

type Dashboard struct {
	Title  string
	Groups []DashboardGroup
}

type DashboardGroup struct {
	Name, Icon string
	Entries    []DashboardEntry
}

type DashboardEntry struct {
	Name, Url, Icon, IconURL, HealthCheckPath string
	UseNewTab                                 bool
}
