# urn
![image](urn_small.png)

## Goal
The goal of this project is to provide reference-grade, pure Golang implementations of the highest quality (not necessarily fastest) random number generation algorithms known for all commonly used probability distributions. This will be accomplished by
* writing clear, idiomatic code in the standard Golang format, avoiding the use of type inference and difficult-to-understand constructs
* providing thorough test suites of all implemented algorithms
* clearly and thoroughly documenting all algorithm implementations, APIs, and tests, including references to (publicly available) papers or books describing the underlying algorithms

As I update this frequently, projects using it may need to run
```
GOPROXY=direct go get -u
```
periodically to ensure an old, cached version is not being used.

This project targets RHEL 9 on x86_64. At this time there are no plans to support other platforms, though most likely it will work correctly on any POSIX-compliant x86_64 Linux distribution.

For C implementations of some of the same algorithms, see the sister project ![here](https://github.com/kmazza2/crng).

## TODO
* ~~Implement function which returns random number provided by the OS~~
* Implement LXM algorithm as an alternative to SplitMix64
* ~~Implement xoshiro256**~~
* Implement general gamma distribution
* Implement jump functions (for xoshiro256**)
* Implement discrete uniform distribution
* Implement general finite discrete distribution with given masses
* Implement hypergeometric distribution
* Consider unexporting interfaces (Float64rng, Uint64rng) if possible
* Write tests for dunif, normal, ...
* Check documentation
* Prevent Uniform(0,1) generators from returning 0 or 1
