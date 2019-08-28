package main

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/datastore"
)

type CBInfo struct {
	Host           string
	BucketName     string
	BucketPassword string
}

type DSInfo struct {
	Creds   string
	Kind    string
	Project string
}

func (c *CBInfo) GetCouchBaseEntities() (*gocb.QueryResults, error) {

	clusterConnection, err := gocb.Connect(c.Host)
	if err != nil {
		return nil, err
	}

	fmt.Println("Reading Couchbase Bucket: ", c.BucketName)

	bucket, err := clusterConnection.OpenBucket(c.BucketName, c.BucketPassword)
	if err != nil {
		return nil, err
	}

	query := gocb.NewN1qlQuery(fmt.Sprintf("SELECT `%s`.* FROM `%s`", c.BucketName, c.BucketName))

	results, err := bucket.ExecuteN1qlQuery(query, nil)
	if err != nil {
		return nil, err
	}

	return &results, nil

}

func (d *DSInfo) PutDatastoreEntities(cbdata gocb.QueryResults) error {
	ctx := context.Background()
	// client, _ := ds.ConnectDatastore(d.Creds)
	client, err := datastore.NewClient(ctx, d.Project)
	if err != nil {
		fmt.Println(err)

	}

	row := RBACTeamAssets{} // Use source type here

	fmt.Println("Moving CB Bucket data to DS Kind: ", d.Kind, " in Project: ", d.Project)

	for cbdata.Next(&row) {
		key := datastore.NameKey(d.Kind, row.TeamName, nil)

		_, err := client.Put(ctx, key, &row)
		if err != nil {
			return err
		}

		fmt.Printf("%s added.\n", row.TeamName)
	}

	return nil
}

func main() {
	var stagingCouchbase = CBInfo{
		Host:           "couchbase://ADDRESS",
		BucketName:     "NAME",
		BucketPassword: "PASSWORD",
	}

	var stagingDatastore = DSInfo{
		Creds:   os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"),
		Kind:    "NAME",
		Project: "test1-12456",
	}

	results, err := stagingCouchbase.GetCouchBaseEntities()
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	err = stagingDatastore.PutDatastoreEntities(*results)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

}
