DERO AstroBWT Proof-of-Work
============================  

### AstroBWT BUILDING  
go get -u github.com/deroproject/astrobwt/miner  

### SAMPLE OUTPUT  
./miner
DERO AstroBWT Miner v0.01 alpha
CPU: Intel(R) Xeon(R) CPU E3-1270 v6 @ 3.80GHz    PhysicalThreads:1
             Threads           Total Time     Total Iterations            Time/PoW         Hash Rate/Sec 
                   1         3.272996982s                  100          32.729969ms                 30.6 
                   2         3.572288466s                  200          17.861442ms                 56.0 
                   3         4.013980986s                  300          13.379936ms                 74.7 
                   4         4.704899609s                  400          11.762249ms                 85.0 
                   5         5.784798143s                  500          11.569596ms                 86.4 
                   6         6.629462384s                  600          11.049103ms                 90.5 
                   7         8.351780961s                  700          11.931115ms                 83.8 
                   8         10.49473002s                  800          13.118412ms                 76.2


The proof of work is based on Information Theory Domain.  
## AstroBWT Pseudo CODE:  
1. Calulate SHA3-256 of input data
2. Expand data using Salsa0  cipher  69371  bytes
3. Calculate BWT of step 2
4. Calculate SHA3-256 of BWT data
5. Expand data using Salsa20  cipher  69371  + random number based on step 4
6. Calculate BWT of data from step 5
7. Calculate SHA3-256 of BWT data from step 6



Reference: https://en.wikipedia.org/wiki/Burrows%E2%80%93Wheeler_transform
