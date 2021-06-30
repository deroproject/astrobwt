DERO AstroBWT CPU Mining Proof-of-Work
======================================  

#### [AstroBWT Testnet](https://github.com/deroproject/derosuite_AstroBWT_testnet)
### [AstroBWT Mainnet HardFork on block 4550555,  March 7,2020. ~0200-GMT.](DERO HardFork on block 4550555,  March 7,2020. ~0200-GMT.)  

#### Table of Contents
1. [AstroBWT BUILDING](#astrobwt-building) 
1. [Sample Output](#sample-output) 
1. [AstroBWT Pseudo CODE](#astrobwt-pseudo-code) 
1. [Explaining AstroBWT](#explaining-astrobwt) 
1. [AstroBWT Hash Rates](#astrobwt-hash-rates) 


### AstroBWT BUILDING  
```go get -u github.com/deroproject/astrobwt/miner  ```

### SAMPLE OUTPUT  
``` ./miner  
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
  ``` 
### AstroBWT Pseudo CODE  
```
1. Calulate SHA3-256 of input data
2. Expand data using Salsa20  cipher  69371  bytes
3. Calculate BWT of step 2
4. Calculate SHA3-256 of BWT data
5. Expand data using Salsa20  cipher  69371  + random number based on step 4
6. Calculate BWT of data from step 5
7. Calculate SHA3-256 of BWT data from step 6  
```
[More about BWT here](https://en.wikipedia.org/wiki/Burrows%E2%80%93Wheeler_transform).

### Explaining AstroBWT  

AstroBWT is not a product of the current cryptosphere approach. It has roots in Information Theory and the Compression Domains.  
1. AstroBWT is based on mathematical proofs and research, unlike many other CPU mining ALGOs.  
1. All current mining algorithms are static. To explain more simply, all current cryptocurrency mining algorithms follow data dependent branches, loops or conditions. Dero also built out a RandomX implementation in Golang that has been discarded for use but open sourced to the public [RandomX in Golang](https://git.dero.io/DERO_Foundation/RandomX).    
1. AstroBWT, as the name implies, has BWT at it's core. BWT has been in research for the last 3 decades and numerous optimization attempts of GPU/FPGA have taken place. For more information on that topic, please refer to literature in the Information Theory and Compression Domains. However, all known implementations to this date, could not deliver an improvement even twice that of CPUs.  
1. All major providers (such as INTEL, NVIDIA etc) have already provided optimized implementations of BWT. Since, BWT has been used quite often in the general information theory and compression domains, it has been a subject of intensive studies. https://software.intel.com/en-us/ipp-dev-reference-burrows-wheeler-transform  
1. In the coming months or years that AstroBWT is optimized or shown to have a significant performance boost on FPGAs, ASICs, or GPUs there would be a benefit for everyone. Such an advancement could even trigger the next revolution in Bioinformatics, Signal Processing, DNA Sequencing or other numerous domains where BWT is used. So, AstroBWT will not only serve as an ASIC/FPGA and GPU resisitant algorithm but it will aslo succeed in helping scientific research if it is optimized for these things.  

Read more about research of BWT FPGA implementation [here](https://www.sfu.ca/~zhenman/files/SC3-FCCM2019-BWT.pdf) and [local copy of this research pdf](https://github.com/deroproject/astrobwt/blob/master/SC3-FCCM2019-BWT.pdf).

### AstroBWT Hash Rates 
[ 1. Submit and compare hashes with other devices(On Github).](https://github.com/deroproject/astrobwt/issues/2):  
[ 2. Submit and compare hashes with other devices((https://benchmark.dero.network/).](https://benchmark.dero.network/)

|SNo.| 	CPU  |Miner| Hash Rates |Threads |	OS |  
|----|-------|-----|------------|--------|-----|  
|1| AMD Ryzen Threadripper 1950X | Official |560 H/s | -| Linux|  
|2| Ryzen 7 2700x DDR4 2666 "Stock settings" | XMRig-5.9.0 |610 H/s | -| Windows 10|  
|3| Qualcomm SDM845 Snapdragon 845 (10 nm) CPU Octa-core (4x2.8 GHz Kryo 385 Gold & 4x1.7 GHz Kryo 385 Silver) | XMRigCC |220 H/s | -| Ubuntu|  
|4| Ryzen 7 3800X 4.4Ghz  | XMRig |970 H/s | -| Linux|  


