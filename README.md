# weaviate-arXiv-dataset-demo

![Weaviate feat arXiv](./docs/weaviate_arxiv.png "Weaviate feat arXiv")

This repository will contain a demo using Weaviate with data and metadata from the arXiv dataset.

## Datasets

In this repository, we have created weaviate schemas for ArXiV and PapersWithCode datasets.
Below, we will detail how to fetch and prepare these datasets for ingestion into Weaviate.

### Weaviate 
Weaviate is a cloud-native, realtime vector search engine that allows you to bring your machine learning models to scale.
[Weaviate documentation](https://www.semi.technology/documentation/weaviate/current/)

#### The GraphQL-based Search Graph 
- Semantic Search engine
- Automatic Classification
- Knowledge Representation

#### Documentation

- [Documentation](https://semi.technology/documentation/weaviate/current/).
- [Getting Started Guide](https://www.semi.technology/documentation/weaviate/current/get-started/quick_start.html).

#### Support

- [Stackoverflow for questions](https://stackoverflow.com/questions/tagged/weaviate).
- [Github for issues](https://github.com/semi-technologies/weaviate/issues).

### ArXiV dataset

#### Requirements
`gsutil` package 

##### Installation
* Last updated: September 4, 2020
* Source: [Google Cloud Documentation](https://cloud.google.com/sdk/docs#deb)
```bash
# 1. Add Cloud SDK URI as a package source.

echo "deb [signed-by=/usr/share/keyrings/cloud.google.gpg] https://packages.cloud.google.com/apt cloud-sdk main" | sudo tee -a /etc/apt/sources.list.d/google-cloud-sdk.list

# 2. Install the necessary dependencies.
sudo apt-get install apt-transport-https ca-certificates gnupg

# 3. Import the Google Cloud public key.
curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key --keyring /usr/share/keyrings/cloud.google.gpg add -

# 4. Fetch the latest repository information, install Cloud SDK.
sudo apt-get update && sudo apt-get install google-cloud-sdk

# 5. Initialize Cloud SDK.
gcloud init
```

#### download the dataset

##### locally
```bash
gsutil cp -n gs://arxiv-dataset/metadata-v5/arxiv-metadata-oai.json .
```
or 
```bash
wget -nc https://paper2code.com/public/arxiv-metadata-oai.json.tar.gz
tar xvf arxiv-metadata-oai.json.tar.gz
```

#### Snapshot
```json
{
  "id": "0704.0001",
  "submitter": "Pavel Nadolsky",
  "authors": "C. Bal\\'azs, E. L. Berger, P. M. Nadolsky, C.-P. Yuan",
  "title": "Calculation of prompt diphoton production cross sections at Tevatron and\n  LHC energies",
  "comments": "37 pages, 15 figures; published version",
  "journal-ref": "Phys.Rev.D76:013009,2007",
  "doi": "10.1103/PhysRevD.76.013009",
  "abstract": "  A fully differential calculation in perturbative quantum chromodynamics is\npresented for the production of massive photon pairs at hadron colliders. All\nnext-to-leading order perturbative contributions from quark-antiquark,\ngluon-(anti)quark, and gluon-gluon subprocesses are included, as well as\nall-orders resummation of initial-state gluon radiation valid at\nnext-to-next-to-leading logarithmic accuracy. The region of phase space is\nspecified in which the calculation is most reliable. Good agreement is\ndemonstrated with data from the Fermilab Tevatron, and predictions are made for\nmore detailed tests with CDF and DO data. Predictions are shown for\ndistributions of diphoton pairs produced at the energy of the Large Hadron\nCollider (LHC). Distributions of the diphoton pairs from the decay of a Higgs\nboson are contrasted with those produced from QCD processes at the LHC, showing\nthat enhanced sensitivity to the signal can be obtained with judicious\nselection of events.\n",
  "report-no": "ANL-HEP-PR-07-12",
  "categories": [
    "hep-ph"
  ],
  "versions": [
    "v1",
    "v2"
  ]
}
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

### Snapshot
```json
  {
    "paper_url": "https://paperswithcode.com/paper/understanding-the-semantic-intent-of-natural",
    "arxiv_id": null,
    "title": "Understanding the Semantic Intent of Natural Language Query",
    "abstract": "",
    "url_abs": "https://www.aclweb.org/anthology/I13-1063/",
    "url_pdf": "https://www.aclweb.org/anthology/I13-1063",
    "proceeding": "IJCNLP 2013 10",
    "authors": [
      "Juan Xu",
      "Qi Zhang",
      "Xuanjing Huang"
    ],
    "tasks": [],
    "date": "2013-10-01"
  }
```

## Credits
