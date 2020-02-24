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
1. All current mining algorithms are static. To explain more simply, all current cryptocurrency mining algorithms follow data dependent branches, loops or conditions. We evaluated RandomX too for true CPU only mining also but it's design too was not great. DeroProject has also implemented [RandomX in Golang](https://git.dero.io/DERO_Foundation/RandomX).    
1. AstroBWT, as the name implies, has BWT at it's core. BWT has been in research for the last 3 decades and numerous optimization attempts of GPU/FPGA have taken place. For more information on that topic, please refer to literature in the Information Theory and Compression Domains. However, all known implementations to this date, could not deliver an improvement even twice that of CPUs.  
1. All major providers (such as INTEL, NVIDIA etc) have already provided optimized implementations of BWT. Since, BWT has been used quite often in the general information theory and compression domains, it has been a subject of intensive studies. https://software.intel.com/en-us/ipp-dev-reference-burrows-wheeler-transform  
1. In the coming months or years that AstroBWT is optimized or shown to have a significant performance boost on FPGAs, ASICs, or GPUs there would be a benefit for everyone. Such an advancement could even trigger the next revolution in Bioinformatics, Signal Processing, DNA Sequencing or other numerous domains where BWT is used. So, AstroBWT will not only serve as an ASIC/FPGA and GPU resisitant algorithm but it will aslo succeed in helping scientific research if it is optimized for these things.  

Read more about research of BWT FPGA implementation [here](http://www.sfu.ca/~zhenman/files/C16-FCCM2019-BWT.pdf).

### AstroBWT Hash Rates 
[Submitted by public](https://github.com/deroproject/astrobwt/issues/2):  
|SNo.| 	CPU  |Hash Rates |Threads |	OS |  
|------|-------|--------|-------|---|  
| 1| AMD Ryzen Threadripper 1950X |  ~240 H/s|15 Threads|Linux|  
| 2| Intel Core i7-6700 |   ~108 H/s|-|-|  
| 3| Intel® Xeon® E5-4657L v2 | ~127-160 H/s|-|-|  
| 4| Intel(R) Core(TM) i7-87s50H @ 2.20GHz |     ~75 H/s | 6 cores/12 threads|-|  
| 5| AMD RYZEN 7 3700 X | ~180 H/S|-|-|  
| 6| AMD Ryzen 5 3600 (3.6GHz) | ~70-95 H/s|-|-|  
| 7|  Intel Core i5-2500K (3.3GHz)   |    ~50 H/s|-|-|  
| 8|  Intel Xeon E31240 (3.3GHz)  |    ~60 H/s|-|-|  
| 9|   AMD A4-7300 (3.8GHz)  |    ~12 H/s |-|-|  
| 10|   Ryzen 5 3600 @ stock 26 watt  |   ~180-200 H/s  |-|-|  
| 11|    Ryzen @ 3950 34 watt  |    ~220-240 H/s |-|-|  
| 12|    Ryzen @ 4150 48 watt  |    ~230-250 H/s |-|-|  
| 13|   Ryzen 5 2600 @ 3900   |   ~140-160 H/s |-|-|  
| 14|    Ryzen 3 1200 @ 2000  |   ~45-55 H/s  |-|-|  
| 15|    Ryzen 3 120 @ 3900   |   ~80-90 H/s |-|-|  
| 16|   I7-7700k   |    ~100-120 H/s |-|-|  
| 17|   Intel Core i7-6700K 4.20GHz|     ~120-138 H/s   |-|-|  
| 18|   Intel Atom x5-Z8330  1.44Ghz   |    ~2-6 H/s |-|-|  
| 19|    Intel Core M 5Y10  0.80Ghz  |    ~12-33 H/s  |-|-|  
| 20|   i7-7700HQ @ 2.80GHz  |    ~90 H/s |-|-|  
| 21| i3 8100(4 core)  |    ~75 H/s |3 core|proxmox virtualization|  
| 22|  Intel i5-4570 @ 3.2GhZ  |    ~56.5 H/s |-|iMac late 2013 8Gb|  
| 23|   Intel i7-3770K @3.5GHz |    ~72-92 H/s |-|-|  
| 24|   Ryzen 2700 3.9GHz |     ~172 H/s |-|-|  



