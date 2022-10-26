The implementation of Median Filter scales as more workers are being used. 
 - This is due to how it uses multiple CPU cores to distribute the workload, compared to single threaded which forces all the work onto one core.
 - As worker threads increase, the amount of work allocated to threads decrease, causing the work to be done faster.

The plateau effect may be due to the machine only having 8 cores, so 8 -> 16 would still mean the same amount of work per core, but it may also be due to how the work has a minimum runtime, and multiple cores cannot speed it up. This includes unparallelisable tasks such as file reading or accessing mutex locks, which requires an unoptimisable amount of time.