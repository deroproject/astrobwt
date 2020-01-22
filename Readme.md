DERO AstroBWT Proof-of-Work
============================

The proof of work is based on Information Theory Domain.
The algorithm is
1. Calulate SHA3-256 of input data
2. Expand data using Salsa0  cipher  69371  bytes
3. Calculate BWT of step 2
4. Calculate SHA3-256 of BWT data
5. Expand data using Salsa20  cipher  69371  + random number based on step 4
6. Calculate BWT of data from step 5
7. Calculate SHA3-256 of BWT data from step 6

Reference: https://en.wikipedia.org/wiki/Burrows%E2%80%93Wheeler_transform64efe2cd45eb900614c89e8b7ef7dba9c83
