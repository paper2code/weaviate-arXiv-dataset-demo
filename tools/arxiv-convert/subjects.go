package main

const arXivSubjectsFlat = `[
 {
   "Code": "astro-ph.GA",
   "Name": "Astrophysics of Galaxies",
   "GroupCode": "Astrophysics",
   "GroupName": "astro-ph",
   "Discipline": "Physics"
 },
 {
   "Code": "astro-ph.CO",
   "Name": "Cosmology and Nongalactic Astrophysics",
   "GroupCode": "Astrophysics",
   "GroupName": "astro-ph",
   "Discipline": "Physics"
 },
 {
   "Code": "astro-ph.EP",
   "Name": "Earth and Planetary Astrophysics",
   "GroupCode": "Astrophysics",
   "GroupName": "astro-ph",
   "Discipline": "Physics"
 },
 {
   "Code": "astro-ph.HE",
   "Name": "High Energy Astrophysical Phenomena",
   "GroupCode": "Astrophysics",
   "GroupName": "astro-ph",
   "Discipline": "Physics"
 },
 {
   "Code": "astro-ph.IM",
   "Name": "Instrumentation and Methods for Astrophysics",
   "GroupCode": "Astrophysics",
   "GroupName": "astro-ph",
   "Discipline": "Physics"
 },
 {
   "Code": "astro-ph.SR",
   "Name": "Solar and Stellar Astrophysics",
   "GroupCode": "Astrophysics",
   "GroupName": "astro-ph",
   "Discipline": "Physics"
 },
 {
   "Code": "cond-mat.dis-nn",
   "Name": "Disordered Systems and Neural Networks",
   "GroupCode": "Condensed Matter",
   "GroupName": "cond-mat",
   "Discipline": "Physics"
 },
 {
   "Code": "cond-mat.mtrl-sci",
   "Name": "Materials Science",
   "GroupCode": "Condensed Matter",
   "GroupName": "cond-mat",
   "Discipline": "Physics"
 },
 {
   "Code": "cond-mat.mes-hall",
   "Name": "Mesoscale and Nanoscale Physics",
   "GroupCode": "Condensed Matter",
   "GroupName": "cond-mat",
   "Discipline": "Physics"
 },
 {
   "Code": "cond-mat.other",
   "Name": "Other Condensed Matter",
   "GroupCode": "Condensed Matter",
   "GroupName": "cond-mat",
   "Discipline": "Physics"
 },
 {
   "Code": "cond-mat.quant-gas",
   "Name": "Quantum Gases",
   "GroupCode": "Condensed Matter",
   "GroupName": "cond-mat",
   "Discipline": "Physics"
 },
 {
   "Code": "cond-mat.soft",
   "Name": "Soft Condensed Matter",
   "GroupCode": "Condensed Matter",
   "GroupName": "cond-mat",
   "Discipline": "Physics"
 },
 {
   "Code": "cond-mat.stat-mech",
   "Name": "Statistical Mechanics",
   "GroupCode": "Condensed Matter",
   "GroupName": "cond-mat",
   "Discipline": "Physics"
 },
 {
   "Code": "cond-mat.str-el",
   "Name": "Strongly Correlated Electrons",
   "GroupCode": "Condensed Matter",
   "GroupName": "cond-mat",
   "Discipline": "Physics"
 },
 {
   "Code": "cond-mat.supr-con",
   "Name": "Superconductivity",
   "GroupCode": "Condensed Matter",
   "GroupName": "cond-mat",
   "Discipline": "Physics"
 },
 {
   "Code": "gr-qc",
   "Name": "General Relativity and Quantum Cosmology",
   "GroupCode": "General Relativity and Quantum Cosmology",
   "GroupName": "gr-qc",
   "Discipline": "Physics"
 },
 {
   "Code": "hep-ex",
   "Name": "High Energy Physics - Experiment",
   "GroupCode": "High Energy Physics - Experiment",
   "GroupName": "hep-ex",
   "Discipline": "Physics"
 },
 {
   "Code": "hep-lat",
   "Name": "High Energy Physics - Lattice",
   "GroupCode": "High Energy Physics - Lattice",
   "GroupName": "hep-lat",
   "Discipline": "Physics"
 },
 {
   "Code": "hep-ph",
   "Name": "High Energy Physics - Phenomenology",
   "GroupCode": "High Energy Physics - Phenomenology",
   "GroupName": "hep-ph",
   "Discipline": "Physics"
 },
 {
   "Code": "hep-th",
   "Name": "High Energy Physics - Theory",
   "GroupCode": "High Energy Physics - Theory",
   "GroupName": "hep-th",
   "Discipline": "Physics"
 },
 {
   "Code": "math-ph",
   "Name": "Mathematical Physics",
   "GroupCode": "Mathematical Physics",
   "GroupName": "math-ph",
   "Discipline": "Physics"
 },
 {
   "Code": "nlin.AO",
   "Name": "Adaptation and Self-Organizing Systems",
   "GroupCode": "Nonlinear Sciences",
   "GroupName": "nlin",
   "Discipline": "Physics"
 },
 {
   "Code": "nlin.CG",
   "Name": "Cellular Automata and Lattice Gases",
   "GroupCode": "Nonlinear Sciences",
   "GroupName": "nlin",
   "Discipline": "Physics"
 },
 {
   "Code": "nlin.CD",
   "Name": "Chaotic Dynamics",
   "GroupCode": "Nonlinear Sciences",
   "GroupName": "nlin",
   "Discipline": "Physics"
 },
 {
   "Code": "nlin.SI",
   "Name": "Exactly Solvable and Integrable Systems",
   "GroupCode": "Nonlinear Sciences",
   "GroupName": "nlin",
   "Discipline": "Physics"
 },
 {
   "Code": "nlin.PS",
   "Name": "Pattern Formation and Solitons",
   "GroupCode": "Nonlinear Sciences",
   "GroupName": "nlin",
   "Discipline": "Physics"
 },
 {
   "Code": "nucl-ex",
   "Name": "Nuclear Experiment",
   "GroupCode": "Nuclear Experiment",
   "GroupName": "nucl-ex",
   "Discipline": "Physics"
 },
 {
   "Code": "nucl-th",
   "Name": "Nuclear Theory",
   "GroupCode": "Nuclear Theory",
   "GroupName": "nucl-th",
   "Discipline": "Physics"
 },
 {
   "Code": "physics.acc-ph",
   "Name": "Accelerator Physics",
   "GroupCode": "Physics",
   "GroupName": "physics",
   "Discipline": "Physics"
 },
 {
   "Code": "physics.app-ph",
   "Name": "Applied Physics",
   "GroupCode": "Physics",
   "GroupName": "physics",
   "Discipline": "Physics"
 },
 {
   "Code": "physics.ao-ph",
   "Name": "Atmospheric and Oceanic Physics",
   "GroupCode": "Physics",
   "GroupName": "physics",
   "Discipline": "Physics"
 },
 {
   "Code": "physics.atom-ph",
   "Name": "Atomic Physics",
   "GroupCode": "Physics",
   "GroupName": "physics",
   "Discipline": "Physics"
 },
 {
   "Code": "physics.atm-clus",
   "Name": "Atomic and Molecular Clusters",
   "GroupCode": "Physics",
   "GroupName": "physics",
   "Discipline": "Physics"
 },
 {
   "Code": "physics.bio-ph",
   "Name": "Biological Physics",
   "GroupCode": "Physics",
   "GroupName": "physics",
   "Discipline": "Physics"
 },
 {
   "Code": "physics.chem-ph",
   "Name": "Chemical Physics",
   "GroupCode": "Physics",
   "GroupName": "physics",
   "Discipline": "Physics"
 },
 {
   "Code": "physics.class-ph",
   "Name": "Classical Physics",
   "GroupCode": "Physics",
   "GroupName": "physics",
   "Discipline": "Physics"
 },
 {
   "Code": "physics.comp-ph",
   "Name": "Computational Physics",
   "GroupCode": "Physics",
   "GroupName": "physics",
   "Discipline": "Physics"
 },
 {
   "Code": "physics.data-an",
   "Name": "Data Analysis",
   "GroupCode": "Statistics and Probability Physics",
   "GroupName": "physics",
   "Discipline": "Physics"
 },
 {
   "Code": "physics.flu-dyn",
   "Name": "Fluid Dynamics",
   "GroupCode": "Physics",
   "GroupName": "physics",
   "Discipline": "Physics"
 },
 {
   "Code": "physics.gen-ph",
   "Name": "General Physics",
   "GroupCode": "Physics",
   "GroupName": "physics",
   "Discipline": "Physics"
 },
 {
   "Code": "physics.geo-ph",
   "Name": "GeoPhysics",
   "GroupCode": "Physics",
   "GroupName": "physics",
   "Discipline": "Physics"
 },
 {
   "Code": "physics.hist-ph",
   "Name": "History and Philosophy of Physics",
   "GroupCode": "Physics",
   "GroupName": "physics",
   "Discipline": "Physics"
 },
 {
   "Code": "physics.ins-det",
   "Name": "Instrumentation and Detectors",
   "GroupCode": "Physics",
   "GroupName": "physics",
   "Discipline": "Physics"
 },
 {
   "Code": "physics.med-ph",
   "Name": "Medical Physics",
   "GroupCode": "Physics",
   "GroupName": "physics",
   "Discipline": "Physics"
 },
 {
   "Code": "physics.optics",
   "Name": "Optics",
   "GroupCode": "Physics",
   "GroupName": "physics",
   "Discipline": "Physics"
 },
 {
   "Code": "physics.ed-ph",
   "Name": "Physics Education",
   "GroupCode": "Physics",
   "GroupName": "physics",
   "Discipline": "Physics"
 },
 {
   "Code": "physics.soc-ph",
   "Name": "Physics and Society",
   "GroupCode": "Physics",
   "GroupName": "physics",
   "Discipline": "Physics"
 },
 {
   "Code": "physics.plasm-ph",
   "Name": "Plasma Physics",
   "GroupCode": "Physics",
   "GroupName": "physics",
   "Discipline": "Physics"
 },
 {
   "Code": "physics.pop-ph",
   "Name": "Popular Physics",
   "GroupCode": "Physics",
   "GroupName": "physics",
   "Discipline": "Physics"
 },
 {
   "Code": "physics.space-ph",
   "Name": "Space Physics",
   "GroupCode": "Physics",
   "GroupName": "physics",
   "Discipline": "Physics"
 },
 {
   "Code": "quant-ph",
   "Name": "Quantum Physics",
   "GroupCode": "Quantum Physics",
   "GroupName": "quant-ph",
   "Discipline": "Physics"
 },
 {
   "Code": "math.AG",
   "Name": "Algebraic Geometry",
   "GroupCode": "Mathematics",
   "GroupName": "math",
   "Discipline": "Mathematics"
 },
 {
   "Code": "math.AT",
   "Name": "Algebraic Topology",
   "GroupCode": "Mathematics",
   "GroupName": "math",
   "Discipline": "Mathematics"
 },
 {
   "Code": "math.AP",
   "Name": "Analysis of PDEs",
   "GroupCode": "Mathematics",
   "GroupName": "math",
   "Discipline": "Mathematics"
 },
 {
   "Code": "math.CT",
   "Name": "Category Theory",
   "GroupCode": "Mathematics",
   "GroupName": "math",
   "Discipline": "Mathematics"
 },
 {
   "Code": "math.CA",
   "Name": "Classical Analysis and ODEs",
   "GroupCode": "Mathematics",
   "GroupName": "math",
   "Discipline": "Mathematics"
 },
 {
   "Code": "math.CO",
   "Name": "Combinatorics",
   "GroupCode": "Mathematics",
   "GroupName": "math",
   "Discipline": "Mathematics"
 },
 {
   "Code": "math.AC",
   "Name": "Commutative Algebra",
   "GroupCode": "Mathematics",
   "GroupName": "math",
   "Discipline": "Mathematics"
 },
 {
   "Code": "math.CV",
   "Name": "Complex Variables",
   "GroupCode": "Mathematics",
   "GroupName": "math",
   "Discipline": "Mathematics"
 },
 {
   "Code": "math.DG",
   "Name": "Differential Geometry",
   "GroupCode": "Mathematics",
   "GroupName": "math",
   "Discipline": "Mathematics"
 },
 {
   "Code": "math.DS",
   "Name": "Dynamical Systems",
   "GroupCode": "Mathematics",
   "GroupName": "math",
   "Discipline": "Mathematics"
 },
 {
   "Code": "math.FA",
   "Name": "Functional Analysis",
   "GroupCode": "Mathematics",
   "GroupName": "math",
   "Discipline": "Mathematics"
 },
 {
   "Code": "math.GM",
   "Name": "Generalmathematics",
   "GroupCode": "Mathematics",
   "GroupName": "math",
   "Discipline": "Mathematics"
 },
 {
   "Code": "math.GN",
   "Name": "General Topology",
   "GroupCode": "Mathematics",
   "GroupName": "math",
   "Discipline": "Mathematics"
 },
 {
   "Code": "math.GT",
   "Name": "Geometric Topology",
   "GroupCode": "Mathematics",
   "GroupName": "math",
   "Discipline": "Mathematics"
 },
 {
   "Code": "math.GR",
   "Name": "Group Theory",
   "GroupCode": "Mathematics",
   "GroupName": "math",
   "Discipline": "Mathematics"
 },
 {
   "Code": "math.HO",
   "Name": "History and Overview",
   "GroupCode": "Mathematics",
   "GroupName": "math",
   "Discipline": "Mathematics"
 },
 {
   "Code": "math.IT",
   "Name": "Information Theory",
   "GroupCode": "Mathematics",
   "GroupName": "math",
   "Discipline": "Mathematics"
 },
 {
   "Code": "math.KT",
   "Name": "K-Theory and Homology",
   "GroupCode": "Mathematics",
   "GroupName": "math",
   "Discipline": "Mathematics"
 },
 {
   "Code": "math.LO",
   "Name": "Logic",
   "GroupCode": "Mathematics",
   "GroupName": "math",
   "Discipline": "Mathematics"
 },
 {
   "Code": "math.MP",
   "Name": "Mathematical Physics",
   "GroupCode": "Mathematics",
   "GroupName": "math",
   "Discipline": "Mathematics"
 },
 {
   "Code": "math.MG",
   "Name": "Metric Geometry",
   "GroupCode": "Mathematics",
   "GroupName": "math",
   "Discipline": "Mathematics"
 },
 {
   "Code": "math.NT",
   "Name": "Number Theory",
   "GroupCode": "Mathematics",
   "GroupName": "math",
   "Discipline": "Mathematics"
 },
 {
   "Code": "math.NA",
   "Name": "Numerical Analysis",
   "GroupCode": "Mathematics",
   "GroupName": "math",
   "Discipline": "Mathematics"
 },
 {
   "Code": "math.OA",
   "Name": "Operator Algebras",
   "GroupCode": "Mathematics",
   "GroupName": "math",
   "Discipline": "Mathematics"
 },
 {
   "Code": "math.OC",
   "Name": "Optimization and Control",
   "GroupCode": "Mathematics",
   "GroupName": "math",
   "Discipline": "Mathematics"
 },
 {
   "Code": "math.PR",
   "Name": "Probability",
   "GroupCode": "Mathematics",
   "GroupName": "math",
   "Discipline": "Mathematics"
 },
 {
   "Code": "math.QA",
   "Name": "Quantum Algebra",
   "GroupCode": "Mathematics",
   "GroupName": "math",
   "Discipline": "Mathematics"
 },
 {
   "Code": "math.RT",
   "Name": "Representation Theory",
   "GroupCode": "Mathematics",
   "GroupName": "math",
   "Discipline": "Mathematics"
 },
 {
   "Code": "math.RA",
   "Name": "Rings and Algebras",
   "GroupCode": "Mathematics",
   "GroupName": "math",
   "Discipline": "Mathematics"
 },
 {
   "Code": "math.SP",
   "Name": "Spectral Theory",
   "GroupCode": "Mathematics",
   "GroupName": "math",
   "Discipline": "Mathematics"
 },
 {
   "Code": "math.ST",
   "Name": "Statistics Theory",
   "GroupCode": "Mathematics",
   "GroupName": "math",
   "Discipline": "Mathematics"
 },
 {
   "Code": "math.SG",
   "Name": "Symplectic Geometry",
   "GroupCode": "Mathematics",
   "GroupName": "math",
   "Discipline": "Mathematics"
 },
 {
   "Code": "cs.AI",
   "Name": "Artificial Intelligence",
   "GroupCode": "Computer Science",
   "GroupName": "cs",
   "Discipline": "Computer Science"
 },
 {
   "Code": "cs.CL",
   "Name": "Computation and Language",
   "GroupCode": "Computer Science",
   "GroupName": "cs",
   "Discipline": "Computer Science"
 },
 {
   "Code": "cs.CC",
   "Name": "Computational Complexity",
   "GroupCode": "Computer Science",
   "GroupName": "cs",
   "Discipline": "Computer Science"
 },
 {
   "Code": "cs.CE",
   "Name": "Computational Engineering Finance and Science",
   "GroupCode": "Computer Science",
   "GroupName": "cs",
   "Discipline": "Computer Science"
 },
 {
   "Code": "cs.CG",
   "Name": "Computational Geometry",
   "GroupCode": "Computer Science",
   "GroupName": "cs",
   "Discipline": "Computer Science"
 },
 {
   "Code": "cs.GT",
   "Name": "Computer Science and Game Theory",
   "GroupCode": "Computer Science",
   "GroupName": "cs",
   "Discipline": "Computer Science"
 },
 {
   "Code": "cs.CV",
   "Name": "Computer Vision and Pattern Recognition",
   "GroupCode": "Computer Science",
   "GroupName": "cs",
   "Discipline": "Computer Science"
 },
 {
   "Code": "cs.CY",
   "Name": "Computers and Society",
   "GroupCode": "Computer Science",
   "GroupName": "cs",
   "Discipline": "Computer Science"
 },
 {
   "Code": "cs.CR",
   "Name": "Cryptography and Security",
   "GroupCode": "Computer Science",
   "GroupName": "cs",
   "Discipline": "Computer Science"
 },
 {
   "Code": "cs.DS",
   "Name": "Data Structures and Algorithms",
   "GroupCode": "Computer Science",
   "GroupName": "cs",
   "Discipline": "Computer Science"
 },
 {
   "Code": "cs.DB",
   "Name": "Databases",
   "GroupCode": "Computer Science",
   "GroupName": "cs",
   "Discipline": "Computer Science"
 },
 {
   "Code": "cs.DL",
   "Name": "Digital Libraries",
   "GroupCode": "Computer Science",
   "GroupName": "cs",
   "Discipline": "Computer Science"
 },
 {
   "Code": "cs.DM",
   "Name": "Discrete Mathematics",
   "GroupCode": "Computer Science",
   "GroupName": "cs",
   "Discipline": "Computer Science"
 },
 {
   "Code": "cs.DC",
   "Name": "Distributed Parallel and Cluster Computing",
   "GroupCode": "Computer Science",
   "GroupName": "cs",
   "Discipline": "Computer Science"
 },
 {
   "Code": "cs.ET",
   "Name": "Emerging Technologies",
   "GroupCode": "Computer Science",
   "GroupName": "cs",
   "Discipline": "Computer Science"
 },
 {
   "Code": "cs.FL",
   "Name": "Formal Languages and Automata Theory",
   "GroupCode": "Computer Science",
   "GroupName": "cs",
   "Discipline": "Computer Science"
 },
 {
   "Code": "cs.GL",
   "Name": "General Literature",
   "GroupCode": "Computer Science",
   "GroupName": "cs",
   "Discipline": "Computer Science"
 },
 {
   "Code": "cs.GR",
   "Name": "Graphics",
   "GroupCode": "Computer Science",
   "GroupName": "cs",
   "Discipline": "Computer Science"
 },
 {
   "Code": "cs.AR",
   "Name": "Hardware Architecture",
   "GroupCode": "Computer Science",
   "GroupName": "cs",
   "Discipline": "Computer Science"
 },
 {
   "Code": "cs.HC",
   "Name": "Human-Computer Interaction",
   "GroupCode": "Computer Science",
   "GroupName": "cs",
   "Discipline": "Computer Science"
 },
 {
   "Code": "cs.IR",
   "Name": "Information Retrieval",
   "GroupCode": "Computer Science",
   "GroupName": "cs",
   "Discipline": "Computer Science"
 },
 {
   "Code": "cs.IT",
   "Name": "Information Theory",
   "GroupCode": "Computer Science",
   "GroupName": "cs",
   "Discipline": "Computer Science"
 },
 {
   "Code": "cs.LO",
   "Name": "Logic in Computer Science",
   "GroupCode": "Computer Science",
   "GroupName": "cs",
   "Discipline": "Computer Science"
 },
 {
   "Code": "cs.LG",
   "Name": "Machine Learning",
   "GroupCode": "Computer Science",
   "GroupName": "cs",
   "Discipline": "Computer Science"
 },
 {
   "Code": "cs.MS",
   "Name": "Mathematical Software",
   "GroupCode": "Computer Science",
   "GroupName": "cs",
   "Discipline": "Computer Science"
 },
 {
   "Code": "cs.MA",
   "Name": "Multiagent Systems",
   "GroupCode": "Computer Science",
   "GroupName": "cs",
   "Discipline": "Computer Science"
 },
 {
   "Code": "cs.MM",
   "Name": "Multimedia",
   "GroupCode": "Computer Science",
   "GroupName": "cs",
   "Discipline": "Computer Science"
 },
 {
   "Code": "cs.NI",
   "Name": "Networking and Internet Architecture",
   "GroupCode": "Computer Science",
   "GroupName": "cs",
   "Discipline": "Computer Science"
 },
 {
   "Code": "cs.NE",
   "Name": "Neural and Evolutionary Computing",
   "GroupCode": "Computer Science",
   "GroupName": "cs",
   "Discipline": "Computer Science"
 },
 {
   "Code": "cs.NA",
   "Name": "Numerical Analysis",
   "GroupCode": "Computer Science",
   "GroupName": "cs",
   "Discipline": "Computer Science"
 },
 {
   "Code": "cs.OS",
   "Name": "Operating Systems",
   "GroupCode": "Computer Science",
   "GroupName": "cs",
   "Discipline": "Computer Science"
 },
 {
   "Code": "cs.OH",
   "Name": "Other Computer Science",
   "GroupCode": "Computer Science",
   "GroupName": "cs",
   "Discipline": "Computer Science"
 },
 {
   "Code": "cs.PF",
   "Name": "Performance",
   "GroupCode": "Computer Science",
   "GroupName": "cs",
   "Discipline": "Computer Science"
 },
 {
   "Code": "cs.PL",
   "Name": "Programming Languages",
   "GroupCode": "Computer Science",
   "GroupName": "cs",
   "Discipline": "Computer Science"
 },
 {
   "Code": "cs.RO",
   "Name": "Robotics",
   "GroupCode": "Computer Science",
   "GroupName": "cs",
   "Discipline": "Computer Science"
 },
 {
   "Code": "cs.SI",
   "Name": "Social and Information Networks",
   "GroupCode": "Computer Science",
   "GroupName": "cs",
   "Discipline": "Computer Science"
 },
 {
   "Code": "cs.SE",
   "Name": "Software Engineering",
   "GroupCode": "Computer Science",
   "GroupName": "cs",
   "Discipline": "Computer Science"
 },
 {
   "Code": "cs.SD",
   "Name": "Sound",
   "GroupCode": "Computer Science",
   "GroupName": "cs",
   "Discipline": "Computer Science"
 },
 {
   "Code": "cs.SC",
   "Name": "Symbolic Computation",
   "GroupCode": "Computer Science",
   "GroupName": "cs",
   "Discipline": "Computer Science"
 },
 {
   "Code": "cs.SY",
   "Name": "Systems and Control",
   "GroupCode": "Computer Science",
   "GroupName": "cs",
   "Discipline": "Computer Science"
 },
 {
   "Code": "q-bio.BM",
   "Name": "Biomolecules",
   "GroupCode": "Quantitative Biology",
   "GroupName": "q-bio",
   "Discipline": "Quantitative Biology"
 },
 {
   "Code": "q-bio.CB",
   "Name": "Cell Behavior",
   "GroupCode": "Quantitative Biology",
   "GroupName": "q-bio",
   "Discipline": "Quantitative Biology"
 },
 {
   "Code": "q-bio.GN",
   "Name": "Genomics",
   "GroupCode": "Quantitative Biology",
   "GroupName": "q-bio",
   "Discipline": "Quantitative Biology"
 },
 {
   "Code": "q-bio.MN",
   "Name": "Molecular Networks",
   "GroupCode": "Quantitative Biology",
   "GroupName": "q-bio",
   "Discipline": "Quantitative Biology"
 },
 {
   "Code": "q-bio.NC",
   "Name": "Neurons and Cognition",
   "GroupCode": "Quantitative Biology",
   "GroupName": "q-bio",
   "Discipline": "Quantitative Biology"
 },
 {
   "Code": "q-bio.OT",
   "Name": "Other Quantitative Biology",
   "GroupCode": "Quantitative Biology",
   "GroupName": "q-bio",
   "Discipline": "Quantitative Biology"
 },
 {
   "Code": "q-bio.PE",
   "Name": "Populations and Evolution",
   "GroupCode": "Quantitative Biology",
   "GroupName": "q-bio",
   "Discipline": "Quantitative Biology"
 },
 {
   "Code": "q-bio.QM",
   "Name": "Quantitative Methods",
   "GroupCode": "Quantitative Biology",
   "GroupName": "q-bio",
   "Discipline": "Quantitative Biology"
 },
 {
   "Code": "q-bio.SC",
   "Name": "Subcellular Processes",
   "GroupCode": "Quantitative Biology",
   "GroupName": "q-bio",
   "Discipline": "Quantitative Biology"
 },
 {
   "Code": "q-bio.TO",
   "Name": "Tissues and Organs",
   "GroupCode": "Quantitative Biology",
   "GroupName": "q-bio",
   "Discipline": "Quantitative Biology"
 },
 {
   "Code": "q-fin.CP",
   "Name": "Computational Finance",
   "GroupCode": "Quantitative Finance",
   "GroupName": "q-fin",
   "Discipline": "Quantitative Finance"
 },
 {
   "Code": "q-fin.EC",
   "Name": "Economics",
   "GroupCode": "Quantitative Finance",
   "GroupName": "q-fin",
   "Discipline": "Quantitative Finance"
 },
 {
   "Code": "q-fin.GN",
   "Name": "General Finance",
   "GroupCode": "Quantitative Finance",
   "GroupName": "q-fin",
   "Discipline": "Quantitative Finance"
 },
 {
   "Code": "q-fin.MF",
   "Name": "Mathematical Finance",
   "GroupCode": "Quantitative Finance",
   "GroupName": "q-fin",
   "Discipline": "Quantitative Finance"
 },
 {
   "Code": "q-fin.PM",
   "Name": "Portfolio Management",
   "GroupCode": "Quantitative Finance",
   "GroupName": "q-fin",
   "Discipline": "Quantitative Finance"
 },
 {
   "Code": "q-fin.PR",
   "Name": "Pricing of Securities",
   "GroupCode": "Quantitative Finance",
   "GroupName": "q-fin",
   "Discipline": "Quantitative Finance"
 },
 {
   "Code": "q-fin.RM",
   "Name": "Risk Management",
   "GroupCode": "Quantitative Finance",
   "GroupName": "q-fin",
   "Discipline": "Quantitative Finance"
 },
 {
   "Code": "q-fin.ST",
   "Name": "Statistical Finance",
   "GroupCode": "Quantitative Finance",
   "GroupName": "q-fin",
   "Discipline": "Quantitative Finance"
 },
 {
   "Code": "q-fin.TR",
   "Name": "Trading and Market Microstructure",
   "GroupCode": "Quantitative Finance",
   "GroupName": "q-fin",
   "Discipline": "Quantitative Finance"
 },
 {
   "Code": "stat.AP",
   "Name": "Applications",
   "GroupCode": "Statistics",
   "GroupName": "stat",
   "Discipline": "Statistics"
 },
 {
   "Code": "stat.CO",
   "Name": "Computation",
   "GroupCode": "Statistics",
   "GroupName": "stat",
   "Discipline": "Statistics"
 },
 {
   "Code": "stat.ML",
   "Name": "Machine Learning",
   "GroupCode": "Statistics",
   "GroupName": "stat",
   "Discipline": "Statistics"
 },
 {
   "Code": "stat.ME",
   "Name": "Methodology",
   "GroupCode": "Statistics",
   "GroupName": "stat",
   "Discipline": "Statistics"
 },
 {
   "Code": "stat.OT",
   "Name": "Other Statistics",
   "GroupCode": "Statistics",
   "GroupName": "stat",
   "Discipline": "Statistics"
 },
 {
   "Code": "stat.TH",
   "Name": "Statistics Theory",
   "GroupCode": "Statistics",
   "GroupName": "stat",
   "Discipline": "Statistics"
 },
 {
   "Code": "eess.AS",
   "Name": "Audio and Speech Processing",
   "GroupCode": "Electrical Engineering and Systems Science",
   "GroupName": "eess",
   "Discipline": "Electrical Engineering and Systems Science"
 },
 {
   "Code": "eess.IV",
   "Name": "Image and Video Processing",
   "GroupCode": "Electrical Engineering and Systems Science",
   "GroupName": "eess",
   "Discipline": "Electrical Engineering and Systems Science"
 },
 {
   "Code": "eess.SP",
   "Name": "Signal Processing",
   "GroupCode": "Electrical Engineering and Systems Science",
   "GroupName": "eess",
   "Discipline": "Electrical Engineering and Systems Science"
 },
 {
   "Code": "eess.SY",
   "Name": "Systems and Control",
   "GroupCode": "Electrical Engineering and Systems Science",
   "GroupName": "eess",
   "Discipline": "Electrical Engineering and Systems Science"
 },
 {
   "Code": "econ.EM",
   "Name": "Econometrics",
   "GroupCode": "Economics",
   "GroupName": "econ",
   "Discipline": "Economics"
 },
 {
   "Code": "econ.GN",
   "Name": "General Economics",
   "GroupCode": "Economics",
   "GroupName": "econ",
   "Discipline": "Economics"
 },
 {
   "Code": "econ.TH",
   "Name": "Theoretical Economics",
   "GroupCode": "Economics",
   "GroupName": "econ",
   "Discipline": "Economics"
 }
]`

const arXivSubjectsTree = `[
    {
        "name": "Physics",
        "areas": [
            {
                "code": "astro-ph*",
                "name": "Astrophysics",
                "subareas": [
                    {
                        "code": "astro-ph.CO",
                        "name": "Cosmology and Extragalactic Astrophysics"
                    },
                    { "code": "astro-ph.EP", "name": "Earth and Planetary Astrophysics" },
                    { "code": "astro-ph.GA", "name": "Galaxy Astrophysics" },
                    {
                        "code": "astro-ph.HE",
                        "name": "High Energy Astrophysical Phenomena"
                    },
                    {
                        "code": "astro-ph.IM",
                        "name": "Instrumentation and Methods for Astrophysics"
                    },
                    { "code": "astro-ph.SR", "name": "Solar and Stella Astrophysics" }
                ]
            },
            {
                "code": "cond-mat*",
                "name": "Condensed Matter Physics",
                "subareas": [
                    {
                        "code": "cond-mat.dis-nn",
                        "name": "Disordered Systems and Neural Networks"
                    },
                    { "code": "cond-mat.mtrl-sci", "name": "Materials Science" },
                    {
                        "code": "cond-mat.mes-hall",
                        "name": "Mesoscale and Nanoscale Physics"
                    },
                    { "code": "cond-mat.other", "name": "Other Condensed Matter" },
                    { "code": "cond-mat.quant-gas", "name": "Quantum Gases" },
                    { "code": "soft", "name": "Soft Condensed Matter" },
                    { "code": "cond-mat.stat-mech", "name": "Statistical Mechanics" },
                    {
                        "code": "cond-mat.str-el",
                        "name": "Strongly Correlated Electrons"
                    },
                    { "code": "cond-mat.supr-con", "name": "Superconductivity" }
                ]
            },
            { "code": "gr-qc", "name": "General Relativity and Quantum Cosmology" },
            { "code": "hep-ex", "name": "High Energy Physics - Experiment" },
            { "code": "hep-lat", "name": "High Energy Physics - Lattice" },
            { "code": "hep-ph", "name": "High Energy Physics - Phenomenology" },
            { "code": "hep-th", "name": "High Energy Physics - Theory" },
            { "code": "math-ph", "name": "Mathematical Physics" },
            { "code": "nucl-ex", "name": "Nuclear Experiment" },
            { "code": "nucl-th", "name": "Nuclear Theory" },
            {
                "code": "physics*",
                "name": "Physics",
                "subareas": [
                    { "code": "physics.acc-ph", "name": "Accelerator Physics" },
                    {
                        "code": "physics.ao-ph",
                        "name": "Atmospheric and Oceanic Physics"
                    },
                    { "code": "physics.atom-ph", "name": "Atomic Physics" },
                    {
                        "code": "physics.atm-clus",
                        "name": "Atomic and Molecular Clusters"
                    },
                    { "code": "physics.bio-ph", "name": "Biological Physics" },
                    { "code": "physics.chem-ph", "name": "Chemical Physics" },
                    { "code": "physics.class-ph", "name": "Classical Physics" },
                    { "code": "physics.comp-ph", "name": "Computational Physics" },
                    {
                        "code": "physics.data-an",
                        "name": "Data Analysis, Statistics and Probability"
                    },
                    { "code": "physics.flu-dyn", "name": "Fluid Dynamics" },
                    { "code": "physics.gen-ph", "name": "General Physics" },
                    { "code": "physics.geo-ph", "name": "Geophysics" },
                    {
                        "code": "physics.hist-ph",
                        "name": "History and Philosophy of Physics"
                    },
                    {
                        "code": "physics.ins-det",
                        "name": "Instrumentation and Detectors"
                    },
                    { "code": "physics.med-ph", "name": "Medical Physics" },
                    { "code": "physics.optics", "name": "Optics" },
                    { "code": "physics.ed-ph", "name": "Physics Education" },
                    { "code": "physics.soc-ph", "name": "Physics and Society" },
                    { "code": "physics.plasma-ph", "name": "Plasma Physics" },
                    { "code": "physics.pop-ph", "name": "Popular Physics" },
                    { "code": "physics.space-ph", "name": "Space Physics" }
                ]
            },
            { "code": "quant-ph", "name": "Quantum Physics" }
        ]
    },
    {
        "name": "Mathematics",
        "areas": [
            {
                "code": "math*",
                "name": "Mathematics",
                "subareas": [
                    { "code": "math.AG", "name": "Algebraic Geometry" },
                    { "code": "math.AT", "name": "Algebraic Topology" },
                    { "code": "math.AP", "name": "Analysis of PDEs" },
                    { "code": "math.CT", "name": "Category Theory" },
                    { "code": "math.CA", "name": "Classical Analysis and ODEs" },
                    { "code": "math.CO", "name": "Combinatorics" },
                    { "code": "math.AC", "name": "Commutative Algebra" },
                    { "code": "math.CV", "name": "Complex Variables" },
                    { "code": "math.DG", "name": "Differential Geometry" },
                    { "code": "math.DS", "name": "Dynamical Systems" },
                    { "code": "math.FA", "name": "Functional Analysis" },
                    { "code": "math.GM", "name": "General Mathematics" },
                    { "code": "math.GN", "name": "General Topology" },
                    { "code": "math.GT", "name": "Geometric Topology" },
                    { "code": "math.GR", "name": "Group Theory" },
                    { "code": "math.HO", "name": "History and Overview" },
                    { "code": "math.IT", "name": "Information Theory" },
                    { "code": "math.KT", "name": "K-Theory and Homology" },
                    { "code": "math.LO", "name": "Logic" },
                    { "code": "math.MP", "name": "Mathematical Physics" },
                    { "code": "math.MG", "name": "Metric Geometry" },
                    { "code": "math.NT", "name": "Number Theory" },
                    { "code": "math.NA", "name": "Numerical Analysis" },
                    { "code": "math.OA", "name": "Operator Algebras" },
                    { "code": "math.OC", "name": "Optimization and Control" },
                    { "code": "math.PR", "name": "Probability" },
                    { "code": "math.QA", "name": "Quantum Algebra" },
                    { "code": "math.RT", "name": "Representation Theory" },
                    { "code": "math.RA", "name": "Rings and Algebras" },
                    { "code": "math.SP", "name": "Spectral Theory" },
                    { "code": "math.ST", "name": "Statistics Theory" },
                    { "code": "math.SG", "name": "Symplectic Geometry" }
                ]
            }
        ]
    },
    {
        "name": "Nonlinear Sciences",
        "areas": [
            {
                "code": "nlin*",
                "name": "Nonlinear Sciences",
                "subareas": [
                    {
                        "code": "nlin.AO",
                        "name": "Adaptation and Self-Organizing Systems"
                    },
                    { "code": "nlin.CG", "name": "Cellular Automata and Lattice Gases" },
                    { "code": "nlin.CD", "name": "Chaotic Dynamics" },
                    {
                        "code": "nlin.SI",
                        "name": "Exactly Solvable and Integrable Systems"
                    },
                    { "code": "nlin.PS", "name": "Pattern Formation and Solitons" }
                ]
            }
        ]
    },
    {
        "name": "Computer Science",
        "areas": [
            {
                "code": "cs*",
                "name": "Computer Research Repository",
                "subareas": [
                    { "code": "cs.AI", "name": "Artificial Intelligence" },
                    { "code": "cs.CL", "name": "Computation and Language" },
                    { "code": "cs.CC", "name": "Computation Complexity" },
                    {
                        "code": "cs.CE",
                        "name": "Computation Engineering, Finance and Science"
                    },
                    { "code": "cs.CG", "name": "Computational Geometry" },
                    { "code": "cs.GT", "name": "Computer Science and Game Theory" },
                    {
                        "code": "cs.CV",
                        "name": "Computer Vision and Pattern Recognition"
                    },
                    { "code": "cs.CY", "name": "Computers and Society" },
                    { "code": "cs.CR", "name": "Cryptography and Security" },
                    { "code": "cs.DS", "name": "Data Structures and Algorithms" },
                    { "code": "cs.DB", "name": "Databases" },
                    { "code": "cs.DL", "name": "Digital Libraries" },
                    { "code": "cs.DM", "name": "Discrete Mathematics" },
                    {
                        "code": "cs.DC",
                        "name": "Distributed, Parallel, and Cluster Computing"
                    },
                    { "code": "cs.ET", "name": "Emerging Technologies" },
                    { "code": "cs.FL", "name": "Formal Languages and Automata Theory" },
                    { "code": "cs.GL", "name": "General Literature" },
                    { "code": "cs.GR", "name": "Graphics" },
                    { "code": "cs.AR", "name": "Hardware Architecture" },
                    { "code": "cs.HC", "name": "Human-Computer Interaction" },
                    { "code": "cs.IR", "name": "Information Retrieval" },
                    { "code": "cs.LG", "name": "Learning" },
                    { "code": "cs.LO", "name": "Logic in Computer Science" },
                    { "code": "cs.MS", "name": "Mathematical Software" },
                    { "code": "cs.MA", "name": "Multiagent Systems" },
                    { "code": "cs.MM", "name": "Multimedia" },
                    { "code": "cs.NI", "name": "Networking and Internet Architecture" },
                    { "code": "cs.NE", "name": "Neural and Evolutionary Computing" },
                    { "code": "cs.NA", "name": "Numerical Analysis" },
                    { "code": "cs.OS", "name": "Operating System" },
                    { "code": "cs.OH", "name": "Other Computer Science" },
                    { "code": "cs.PF", "name": "Performance" },
                    { "code": "cs.PL", "name": "Programming Languages" },
                    { "code": "cs.RO", "name": "Robotics" },
                    { "code": "cs.SI", "name": "Social and Information Networks" },
                    { "code": "cs.SE", "name": "Software Engineering" },
                    { "code": "cs.SD", "name": "Sound" },
                    { "code": "cs.SC", "name": "Symbolic Computation" },
                    { "code": "cs.SY", "name": "Systems and Control" }
                ]
            }
        ]
    },
    {
        "name": "Quantitative Biology",
        "areas": [
            {
                "code": "q-bio*",
                "name": "Quantitative Biology",
                "subareas": [
                    { "code": "q-bio.BM", "name": "Biomolecules" },
                    { "code": "q-bio.CB", "name": "Cell Behavior" },
                    { "code": "q-bio.GN", "name": "Genomics" },
                    { "code": "q-bio.MN", "name": "Molecular Networks" },
                    { "code": "q-bio.NC", "name": "Neurons and Cognition" },
                    { "code": "q-bio.OT", "name": "Other Quatitative Biology" },
                    { "code": "q-bio.PE", "name": "Populations and Evolution" },
                    { "code": "q-bio.QM", "name": "Quantitative Methods" },
                    { "code": "q-bio.SC", "name": "Subcellular Processes" },
                    { "code": "q-bio.TO", "name": "Tissues and Organs" }
                ]
            }
        ]
    },
    {
        "name": "Quantitative Finance",
        "areas": [
            {
                "code": "q-fin*",
                "name": "Quantitative Finance",
                "subareas": [
                    { "code": "q-fin.CP", "name": "Computational Finance" },
                    { "code": "q-fin.GN", "name": "General Finance" },
                    { "code": "q-fin.PM", "name": "Portfolio Management" },
                    { "code": "q-fin.PR", "name": "Pricing of Securities" },
                    { "code": "q-fin.RM", "name": "Risk Management" },
                    { "code": "q-fin.ST", "name": "Statistical Finance" },
                    { "code": "q-fin.TR", "name": "Trading and Market Microstructure" }
                ]
            }
        ]
    },
    {
        "name": "Statistics",
        "areas": [
            {
                "code": "stat*",
                "name": "Statistics",
                "subareas": [
                    { "code": "stat.AP", "name": "Applications" },
                    { "code": "stat.CO", "name": "Computation" },
                    { "code": "stat.ML", "name": "Machine Learning" },
                    { "code": "stat.ME", "name": "Methodology" },
                    { "code": "stat.OT", "name": "Other Statistics" },
                    { "code": "stat.TH", "name": "Statistics Theory" }
                ]
            }
        ]
    }
]`
