package backend

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/TheAndruu/git-leaderboard/models"
	"github.com/kr/pretty"
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

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	log.Infof(ctx, fmt.Sprintf("Accepting repo name: %v", target.RepoName))

	reMarshalled, err := json.Marshal(target)
	if err != nil {
		log.Errorf(ctx, "Issue marshalling json to string %v", err)
	}
	log.Infof(ctx, pretty.Sprintf(string(reMarshalled)))

	values := map[string]string{"message": fmt.Sprintf("thanks for the message from %s", target.RepoName)}
	asBytes, _ := json.Marshal(values)
	w.Write(asBytes)
}

// TODO Next:
// save the object in the db
//
