package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strings"
	"time"

	"github.com/k0kubun/pp"
)

type ArxivMetadataJSON struct {
	Authors         []string `json:"-"`
	AuthorsStr      string   `json:"authors"`
	Abstract        string   `json:"abstract"`
	Categories      []string `json:"-"`
	CategoriesSlice []string `json:"categories"`
	Comments        string   `json:"comments"`
	Doi             string   `json:"doi"`
	ID              string   `json:"id"`
	JournalRef      string   `json:"journal-ref"`
	ReportNo        string   `json:"report-no"`
	Submitter       string   `json:"submitter"`
	Title           string   `json:"title"`
	Versions        []string `json:"versions"`
}

func (a *ArxivMetadataJSON) Unmarshal(b []byte) error {
	return json.Unmarshal(b, a)
}

func main() {
	start := time.Now()

	fileName := "../shared/datasets/arXiv/arxiv-metadata-oai.json"
	pp.Println("fileName:", fileName)

	f, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Error to read [file=%v]: %v", fileName, err.Error())
	}

	fi, err := f.Stat()
	if err != nil {
		log.Fatalf("Could not obtain stat, handle error: %v", err.Error())
	}

	r := bufio.NewReader(f)
	d := json.NewDecoder(r)

	i := 0

	for {

		var arxiv ArxivMetadataJSON

		if err := d.Decode(&arxiv); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		arxiv.AuthorsStr = strings.Replace(arxiv.AuthorsStr, "\n", " ", -1)
		arxiv.AuthorsStr = strings.Replace(arxiv.AuthorsStr, " , ", ",", -1)
		arxiv.AuthorsStr = strings.Replace(arxiv.AuthorsStr, " and ", ",", -1)
		authors := strings.Split(arxiv.AuthorsStr, ",")
		for _, author := range authors {
			author = strings.Replace(author, "  ", " ", -1)
			author = strings.TrimSpace(author)
			if author != "" {
				arxiv.Authors = append(arxiv.Authors, author)
			}
		}

		arxiv.Abstract = strings.TrimSpace(strings.Replace(arxiv.Abstract, "\n", " ", -1))

		for _, category := range arxiv.CategoriesSlice {
			cats := strings.Split(category, " ")
			arxiv.Categories = append(arxiv.Categories, cats...)
		}

		pp.Println(arxiv)

	}

	elapsed := time.Since(start)

	fmt.Printf("Total of [%v] object created.\n", i)
	fmt.Printf("The [%s] is %s long\n", fileName, FileSize(fi.Size()))
	fmt.Printf("To parse the file took [%v]\n", elapsed)
}

func logn(n, b float64) float64 {
	return math.Log(n) / math.Log(b)
}

func humanateBytes(s uint64, base float64, sizes []string) string {
	if s < 10 {
		return fmt.Sprintf("%dB", s)
	}
	e := math.Floor(logn(float64(s), base))
	suffix := sizes[int(e)]
	val := float64(s) / math.Pow(base, math.Floor(e))
	f := "%.0f"
	if val < 10 {
		f = "%.1f"
	}

	return fmt.Sprintf(f+"%s", val, suffix)
}

// FileSize calculates the file size and generate user-friendly string.
func FileSize(s int64) string {
	sizes := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	return humanateBytes(uint64(s), 1024, sizes)
}
