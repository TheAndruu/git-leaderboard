package backend

import (
	"time"

	"golang.org/x/net/context"

	"github.com/TheAndruu/git-leaderboard/models"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
)

// SaveStats sets the time and saves stats to the datastore
func SaveStats(ctx context.Context, statsToSave *models.RepoStats) (int64, error) {
	log.Infof(ctx, "Saving stats to db for %v", statsToSave.RepoName)

	// Update the date of the record
	statsToSave.DateUpated = time.Now()

	// Query for URL to see if any already exist
	existingRemoteURLQuery := datastore.NewQuery("RepoStats").
		Filter("RepoURL =", statsToSave.RepoURL).
		KeysOnly().Limit(1)

	// Check if we already have a record with this remote URL
	var key *datastore.Key

	err := datastore.RunInTransaction(ctx, func(ctx context.Context) error {
		// This function's argument ctx shadows the variable ctx from the surrounding function.

		// last parameter is ignored because it's a keys-only query
		existingKeys, err := existingRemoteURLQuery.GetAll(ctx, new(models.RepoStats))
		isNewKey := true
		if len(existingKeys) > 0 {
			log.Infof(ctx, "Update existing record vice new key")
			// use existing key
			key = existingKeys[0]
			isNewKey = false
		} else {
			log.Infof(ctx, "No existing key found, use new key")
			key = datastore.NewIncompleteKey(ctx, "RepoStats", nil)
		}

		// Now have the key, put the new target record in place
		fullKey, err := datastore.Put(ctx, key, statsToSave)
		if err != nil {
			log.Errorf(ctx, "datastore.Put: %v", err)
			return err
		}
		if isNewKey {
			key = fullKey
		}
		return err
	}, nil)
	// end of transaction

	if err != nil {
		log.Errorf(ctx, "Issue saving RepoStats: %v", err)
		return 0, err
	}

	return key.IntID(), nil
}
