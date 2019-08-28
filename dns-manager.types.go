package main

// App models the response we receive from couchbase when quering for a
// list of applications.
type App struct {
	ID string `json:"id" validate:"nonzero"`
}

type RBACTeamAssets struct {
	TeamName     string   `json:"teamname" validate:"nonzero"`
	DNSRecords   []string `json:"dnsRecords"`
	HealthChecks []string `json:"healthchecks"`
}
