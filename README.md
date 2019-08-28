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

Flags:

| Name | Description | Default |
| --- | --- | --- |
| host: | The Couchbase address | localhost |
| bucket: | The Couchbase Bucket name  | test |
| password: | The password for your Couchbase | derp |
| kind: | The Datastore Kind you are using | test123 |
| project: | Your Project name   | test4321 |

Datastore credentials:   This one uses a local json file for authentication to Datastore 
by exporting the Datastore json file to an ENVIRONMENT variable. 
You get the json file from vault or your Google Project.

*export GOOGLE_APPLICATION_CREDENTIALS="PATH_TO_JSON"

* execute `go run *.go    -host=HOSTNAME -bucket=NAME -password=PASSWORD -kind=NAME -project=YOUR_PROJECT`

For example:

* execute `go run *.go -host=127.0.0.1 -bucket=dns-manager -password=ssshsecret -kind=dns-manager -project=test1-12345`



It should list out the categories it's importing
