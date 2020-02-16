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

### Explaining AstroBWT:
AstroBWT is not from crypto world and has roots in Information Theory and Compression Domain and this domain has much bigger mcap.  

1. AstroBWT is based on proofs and research unlike other CPU mining ALGOs.  
1. Second, All algorithms are static (At best, they follow data dependent branches, loops or conditions). This includes even RandomX also.(In case you are missing, @DeroProject has already implemented [RandomX in golang](https://git.dero.io/DERO_Foundation/RandomX), there are just obfuscated data dependent branches, nothing more in RandomX. So, We found RandomX is not for DERO.)  
1. BWT is the core/crux of AstroBWT. It has been in research for the last 3 decades and numerous optimization attempts of GPU/FPGA  have taken place. Please refer to literature in the Information Theory and Compression Domain. However, all known implementations till date, could never deliver a speedup of even 2 times over processors.  
1. All major providers such as INTEL, NVIDIA etc. already provide optimized implementations of BWT. Since, BWT has so much use in the general information theory. Eg. https://software.intel.com/en-us/ipp-dev-reference-burrows-wheeler-transform  
1. Let's assume if in coming months/years AstroBWT is optimized/speedup to give several times performance on FPGA/ASIC/GPU/etc. then you would be doing great service & might even trigger the next revolution in Bioinformatics/SignalProcessing/DNA Sequencing etc. or in other numerous domains where BWT is used. So, AstroBWT will have subserved the society more practically in other ways also.

Read more about research of BTW FPGA implementation [here](http://www.sfu.ca/~zhenman/files/C16-FCCM2019-BWT.pdf).
