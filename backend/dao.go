package backend

import (
	"time"

	"golang.org/x/net/context"

	"github.com/TheAndruu/git-leaderboard/models"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
)

// SaveStats sets the time and saves stats to the datastore
func SaveStats(ctx context.Context, statsToSave *models.RepoStats) (string, error) {
	log.Infof(ctx, "Saving stats to db for %v", statsToSave.RepoName)

	// Update the date of the record
	statsToSave.DateUpated = time.Now()

	// Create a new key using the URL of the repo as the string name
	key := datastore.NewKey(ctx, "RepoStats", statsToSave.RepoURL, 0, nil)

	// Update the record
	_, err := datastore.Put(ctx, key, statsToSave)
	if err != nil {
		log.Errorf(ctx, "Issue saving RepoStats: %v", err)
		return "", err
	}

	return key.StringID(), nil
}

// GetRecentRepoStats returns up to the specified limit of most recently-updated RepoStats
func GetRecentRepoStats(ctx context.Context, limit int) *[]models.RepoStats {

	query := datastore.NewQuery("RepoStats").Order("-DateUpdated").Limit(limit)

	var results []models.RepoStats
	_, err := query.GetAll(ctx, &results)

	if err != nil {
		log.Errorf(ctx, "Issue querying most recent RepoStats: %v", err)
	}

	return &results
}
