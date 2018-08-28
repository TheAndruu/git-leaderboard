package backend

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/TheAndruu/git-leaderboard/models"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

func init() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/repostats", saveRepoPost)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello cruel, cruel world!")
}

func saveRepoPost(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	log.Infof(ctx, "Saving repo stats")
	target := models.RepoStats{}
	json.NewDecoder(r.Body).Decode(&target)
	defer r.Body.Close()

	// Save the data
	log.Infof(ctx, fmt.Sprintf("Accepting repo name: %v", target.RepoName))
	id, err := SaveStats(ctx, &target)
	if err != nil {
		log.Errorf(ctx, "Issue saving stats to datastore %v", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(420)
		values := map[string]string{"message": "Issue saving git stats to database"}
		asBytes, _ := json.Marshal(values)
		w.Write(asBytes)
		return
	}

	// Write content-type, statuscode, payload
	log.Infof(ctx, fmt.Sprintf("Created stats with id %d", id))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	values := map[string]string{"message": fmt.Sprintf("thanks for the message from %s", target.RepoName)}
	asBytes, _ := json.Marshal(values)
	w.Write(asBytes)
}
