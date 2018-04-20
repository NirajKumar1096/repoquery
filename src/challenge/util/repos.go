package util

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/tabwriter"
	"time"

	"github.com/gorilla/mux"
)

// Owner is the repository owner
type Owner struct {
	Login string
}

type Item struct {
	ID              int    `json:"repo_id"`
	Name            string `json:"name"`
	FullName        string `json:"full_name"`
	Owner           Owner
	Description     string
	CreatedAt       string `json:"created_at"`
	StargazersCount int    `json:"stargazers_count"`
}

// JSONData contains the GitHub API response and is primary structure stored in memory db
type JSONData struct {
	Count int `json:"total_count"`
	Items []Item
}

// This global variable servers as memory database container, holds retrieved repositories
var Repos JSONData

func QueryReposEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	query := params["query"]

	fmt.Println(query)
	data, err := Consumehub(query, 2, 10)
	if err != nil {
		log.Fatal(err)
	}
	Repos = data
	PrintData(Repos)
}

func GetReposEndpoint(w http.ResponseWriter, req *http.Request) {
	PrintData(Repos)
}

func PrintData(data JSONData) {
	log.Printf("Repositories found: %d", data.Count)
	const format = "%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Repository", "Stars", "Created at", "Description")
	fmt.Fprintf(tw, format, "----------", "-----", "----------", "----------")
	for _, i := range data.Items {
		desc := i.Description
		if len(desc) > 50 {
			desc = string(desc[:50]) + "..."
		}
		t, err := time.Parse(time.RFC3339, i.CreatedAt)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(tw, format, i.FullName, i.StargazersCount, t.Year(), desc)
	}
	tw.Flush()
}
