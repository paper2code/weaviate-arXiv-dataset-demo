package main

import (
	"encoding/json"
)

// ArxivExport repesents...
type ArxivExport struct {
	Authors    []string `json:"authors"`
	Abstract   string   `json:"abstract"`
	Categories []string `json:"categories"`
	Comments   string   `json:"comments"`
	Doi        string   `json:"doi"`
	ID         string   `json:"id"`
	JournalRef string   `json:"journal-ref"`
	ReportNo   string   `json:"report-no"`
	Submitter  string   `json:"submitter"`
	Title      string   `json:"title"`
	Versions   []string `json:"versions"`
}

// ArXivCategory represents
type ArXivCategory struct {
	Code string `json:"Code"`
	Name string `json:"Name"`
}

// ArXivCategoryFlatten repesents...
type ArXivCategoryFlatten struct {
	Code       string `json:"Code"`
	Discipline string `json:"Discipline"`
	GroupCode  string `json:"GroupCode"`
	GroupName  string `json:"GroupName"`
	Name       string `json:"Name"`
}

// ArXivCategoryTree repesents...
type ArXivCategoryTree struct {
	Areas []ArXivCategoryTreeArea `json:"areas"`
	Name  string                  `json:"name"`
}

// ArXivCategoryTreeArea repesents...
type ArXivCategoryTreeArea struct {
	Code     string                     `json:"code"`
	Name     string                     `json:"name"`
	Subareas []ArXivCategoryTreeSubArea `json:"subareas"`
}

// ArXivCategoryTreeSubArea repesents...
type ArXivCategoryTreeSubArea struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

// ArxivMetadataJSON repesents...
type ArxivMetadataJSON struct {
	Authors           []string         `json:"-"`
	AuthorsStr        string           `json:"authors"`
	Abstract          string           `json:"abstract"`
	Categories        []string         `json:"-"`
	CategoriesSlice   []string         `json:"categories"`
	CategoriesFlatten []*ArXivCategory `json:"-"`
	Comments          string           `json:"comments"`
	Doi               string           `json:"doi"`
	ID                string           `json:"id"`
	JournalRef        string           `json:"journal-ref"`
	ReportNo          string           `json:"report-no"`
	Submitter         string           `json:"submitter"`
	Title             string           `json:"title"`
	Versions          []string         `json:"versions"`
}

func (a *ArxivMetadataJSON) Unmarshal(b []byte) error {
	return json.Unmarshal(b, a)
}
