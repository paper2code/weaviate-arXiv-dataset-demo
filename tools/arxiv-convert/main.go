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

func main() {
	start := time.Now()

	// create the export file
	e, err := os.Create("../shared/datasets/arXiv/arxiv-metadata-oai-weaviate.json")
	if err != nil {
		panic(err)
	}
	defer e.Close()

	// load arXiv categories
	var categories []*ArXivCategoryFlatten
	if err := json.Unmarshal([]byte(arXivSubjectsFlat), &categories); err != nil {
		panic(err)
	}
	pp.Println(categories)

	// Build a arXiv categories map
	categoriesMap := make(map[string]string, len(categories))
	for _, v := range categories {
		categoriesMap[v.Code] = v.Name
	}

	fileName := "../shared/datasets/arXiv/arxiv-metadata-oai.json"
	pp.Println("fileName:", fileName)

	f, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Error to read [file=%v]: %v", fileName, err.Error())
	}
	defer f.Close()

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

		arxiv.Title = strings.Replace(arxiv.Title, "\n", " ", -1)
		arxiv.Title = strings.Replace(arxiv.Title, "  ", " ", -1)

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
			for _, cat := range cats {
				if v, ok := categoriesMap[cat]; ok {
					arxiv.CategoriesFlatten = append(arxiv.CategoriesFlatten, &ArXivCategory{Code: cat, Name: v})
				}
			}
		}

		entry := &ArxivExport{
			Authors:    arxiv.Authors,
			Abstract:   arxiv.Abstract,
			Categories: arxiv.Categories,
			Comments:   arxiv.Comments,
			Doi:        arxiv.Doi,
			ID:         arxiv.ID,
			JournalRef: arxiv.JournalRef,
			ReportNo:   arxiv.ReportNo,
			Submitter:  arxiv.Submitter,
			Title:      arxiv.Title,
			Versions:   arxiv.Versions,
		}

		pp.Println("ID:", entry.ID, "Title:", entry.Title)
		entryBytes, _ := json.Marshal(entry)
		w := bufio.NewWriter(e)
		_, err := w.WriteString(string(entryBytes) + "\n")
		if err != nil {
			panic(err)
		}
		w.Flush()

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
