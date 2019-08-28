# CouchbaseToDatastore

Small app to move data from a Couchbase bucket into a Datastore project id in the default namespace.  This will search through a specified bucket and migrate the data based on the data types.  Mostly taken from https://github.com/GannettDigital/cb2ds

### Usage

* Add structs from your Couchbase  in notes_types.go. Example:

```
type RBACTeamAssets struct {
	TeamName     string   `json:"teamname" validate:"nonzero"`
	DNSRecords   []string `json:"dnsRecords"`
	HealthChecks []string `json:"healthchecks"`
}
```

* implement the Save/Load interfaces for any data structures/types Datastore doesn't support

* setup the Couchbase and Datastore info structs in the main.go.  You'll need a Couchbase address, bucket and password, along with Datastore credentials.  This one uses a local json file for authentication to Datastore by exporting the Datastore json file to an ENVIRONMENT. You get the json file from vault or your Google Project.

*export GOOGLE_APPLICATION_CREDENTIALS="PATH_TO_JSON"

* execute `go run *.go`

It should list out the categories it's importing