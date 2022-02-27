package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"sort"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

var tmpl = template.Must(template.New("tracksTemplate").Parse(`<table>
<tr>
<th><a href="?sortby=title">Title</a></th>
<th><a href="?sortby=artist">Artist</a></th>
<th><a href="?sortby=album">Album</a></th>
<th><a href="?sortby=year">Year</a></th>
<th><a href="?sortby=length">Length</a></th>
</tr>
{{range .}}<tr>
<td>{{.Title}}</td>
<td>{{.Artist}}</td>
<td>{{.Album}}</td>
<td>{{.Year}}</td>
<td>{{.Length}}</td>
</tr>
{{end}}</table>
`))

func main() {
	http.HandleFunc("/", home)
	fmt.Println("Listening and serving on port 9000...")
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func home(w http.ResponseWriter, r *http.Request) {
	key := "title"
	for _, k := range r.URL.Query()["sortby"] {
		key = k
	}
	if cmp, ok := sortBy[key]; ok {
		if sort.IsSorted(customSort{tracks, cmp}) {
			sort.Stable(sort.Reverse(customSort{tracks, cmp}))
		} else {
			sort.Stable(customSort{tracks, cmp})
		}
	}

	err := tmpl.Execute(w, tracks)
	if err != nil {
		fmt.Fprintf(os.Stderr, "execute tracks template: %v\n", err)
	}
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (x customSort) Len() int           { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

var sortBy = map[string]func(*Track, *Track) bool{
	"title":  func(x, y *Track) bool { return x.Title < y.Title },
	"artist": func(x, y *Track) bool { return x.Artist < y.Artist },
	"album":  func(x, y *Track) bool { return x.Album < y.Album },
	"year":   func(x, y *Track) bool { return x.Year < y.Year },
	"length": func(x, y *Track) bool { return x.Length < y.Length },
}
