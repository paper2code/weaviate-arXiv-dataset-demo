# arXiv-demo-dataset

This repository will contain a demo using Weaviate with data and metadata from the arXiv dataset.

## Datasets

### arXiv 

#### requirements
`gsutil` package 

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
