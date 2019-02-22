# eBird Data

This folder contains two sample eBird datasets.  One is the downloadable sample
from the download page at https://ebird.org/science/download-ebird-data-products.

The other is a subset of the relAug-2018 data for just Idaho.

## Citation

eBird provides free access to the basic dataset for non-commercial use.  The
citation for the bulk download is:

	eBird Basic Dataset
	Version: EBD_relAug-2018.
	Cornell Lab of Ornithology, Ithaca, New York. May 2017.

For other citations, see: https://ebird.org/science/citation

### Escaping quotes

Unix: `sed 's/\"/\\\"/g`
Windows(?): `get-content somefile.txt | %{$_ -replace "expression","replace"}`
