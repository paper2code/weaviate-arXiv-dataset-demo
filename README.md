# arXiv-demo-dataset

This repository will contain a demo using Weaviate with data and metadata from the arXiv dataset.

## Datasets

### arXiv 

#### requirements
`gsutil` package

#### download the dataset
```bash
gsutil cp -n gs://arxiv-dataset/metadata-v5/arxiv-metadata-oai.json .
```

### Paperswithcode
You can download the full dataset behind [paperswithcode.com](https://paperswithcode.com) here:

Download links for the data dumps are:

- [All papers with abstracts](https://paperswithcode.com/media/about/papers-with-abstracts.json.gz)
- [Links between papers and code](https://paperswithcode.com/media/about/links-between-papers-and-code.json.gz)
- [Evaluation tables](https://paperswithcode.com/media/about/evaluation-tables.json.gz)

#### download the dataset
```bash
wget -nc https://paperswithcode.com/media/about/papers-with-abstracts.json.gz
gunzip papers-with-abstracts.json.gz
```

## Credits