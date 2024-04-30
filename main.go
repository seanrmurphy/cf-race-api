package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/seanrmurphy/anycloud-backend/api"
	"github.com/seanrmurphy/anycloud-backend/dbqueries"
	"github.com/seanrmurphy/anycloud-backend/service"
	"github.com/syumai/workers"
	"github.com/syumai/workers/cloudflare"
	"github.com/syumai/workers/cloudflare/d1"
	_ "github.com/syumai/workers/cloudflare/d1" // register driver
)

var standardApiValidityInStore = 600

type SecurityHandlerImplementation struct {
	ApiKeyNamespace    string
	DatabaseEnvVarName string
	db                 *sql.DB
}

func (s SecurityHandlerImplementation) HandleApiKeyAuth(ctx context.Context, operation string, k api.ApiKeyAuth) (context.Context, error) {
	// TODO: add support for admin operations; this should be done
	// by making a db query which also obtains the isAdmin field from
	// the linked user record

	// first, we check if the key is in the KV store...
	kv, err := cloudflare.NewKVNamespace(ctx, s.ApiKeyNamespace)
	if err != nil {
		// TODO: add proper logic here; we can probably work for a while without the KV
		// store but it will result in greater latency and more costs...
		log.Printf("failed to init KV: %v", err)
	}

	// if key does not exist in KV store, no error is returned, but
	// the string is empty
	countStr, err := kv.GetString(k.GetAPIKey(), nil)
	if err != nil {
		log.Printf("Error querying KV store for key: %v\n", err)
	}
	if countStr != "<null>" {
		log.Printf("API key found in KV store: %s\n", countStr)
		return ctx, nil
	} else {
		log.Printf("API key not found in KV store\n")
	}

	if s.db == nil {
		c, err := d1.OpenConnector(ctx, s.DatabaseEnvVarName)
		if err != nil {
			return ctx, fmt.Errorf("Internal server error: service unavailable: %w", err)
		}
		// use sql.OpenDB instead of sql.Open.
		s.db = sql.OpenDB(c)
	}

	queries := dbqueries.New(s.db)
	_, err = queries.GetKey(ctx, k.APIKey)
	if err != nil {
		if err == sql.ErrNoRows {
			// no key found in db
			return nil, ogenerrors.ErrSecurityRequirementIsNotSatisfied
		}
		return nil, fmt.Errorf("Internal server error")
	}

	// so here we have found a key from the db which is not in the KV store; we
	// must add it to the KV store with a TTL of 10 minutes
	ttl := standardApiValidityInStore // specified in seconds
	putOptions := &cloudflare.KVNamespacePutOptions{
		ExpirationTTL: ttl,
	}
	// TODO: make the entry in the kv store map to some more useful user related info
	err = kv.PutString(k.GetAPIKey(), "1", putOptions)
	if err != nil {
		log.Printf("Error adding key to KV store: %v - ignoring\n", err)
	} else {
		log.Printf("API Key added KV store with 10 minute lifetime\n")
	}

	return ctx, nil
	// this is how we add some differential treatment for different operations
	// switch operation {
	// case "AdminStatusGet":
	// 	if k.APIKey != s.ApiKey {
	// 		return nil, ogenerrors.ErrSecurityRequirementIsNotSatisfied
	// 	}
	// ... etc ...
	// default:
	// 	return nil, ogenerrors.ErrSecurityRequirementIsNotSatisfied
	// }
}

func main() {
	sec := SecurityHandlerImplementation{
		ApiKeyNamespace:    "apikeys",
		DatabaseEnvVarName: "DB",
	}

	// Create service instance.
	s := &service.ServerImplementation{
		ServiceInvocations:      0,
		AdminServiceInvocations: 0,
		// CostRecords:             initialCosts,
		DatabaseEnvVarName: "DB",
	}

	// Create generated server.
	srv, err := api.NewServer(s, sec)
	if err != nil {
		log.Fatal(err)
	}
	workers.Serve(srv)
}
