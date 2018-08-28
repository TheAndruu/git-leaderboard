package backend

import (
	"context"
	"time"

	"github.com/TheAndruu/git-leaderboard/models"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
)

// SaveStats sets the time and saves stats to the datastore
func SaveStats(ctx context.Context, statsToSave *models.RepoStats) (string, error) {

	log.Infof(ctx, "Savinvg stats to db for %v", statsToSave.RepoName)
	partialKey := datastore.NewIncompleteKey(ctx, "RepoStats", nil)
	statsToSave.DateUpated = time.Now()

	fullKey, err := datastore.Put(ctx, partialKey, statsToSave)
	if err != nil {
		log.Errorf(ctx, "datastore.Put: %v", err)
		return "", err
	}
	idString := fullKey.StringID()
	return idString, nil
}
