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

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	log.Infof(ctx, fmt.Sprintf("Accepting repo name: %v", target.RepoName))

	reMarshalled, err := json.Marshal(target)
	if err != nil {
		log.Errorf(ctx, "Issue marshalling json to string %v", err)
	}
	fmt.Fprintf(w, "%v", string(reMarshalled))
}
