# urn
![image](urn_small.png)

## Goal
The goal of this project is to provide reference-grade, pure Golang implementations of the highest quality (not necessarily fastest) random number generation algorithms known for all commonly encountered probability distributions. This will be accomplished by
* writing clear, idiomatic code in the standard Golang format, avoiding the use of type inference and difficult-to-understand constructs
* providing thorough test suites of all implementated algorithms
* clearly and thoroughly documenting all algorithm implementations and tests, including references to (publicly available) papers or books describing the underlying algorithms

As I update this frequently, projects using it may need to run
```
GOPROXY=direct go get -u
```
periodically to ensure an old, cached version is not being used.

This project targets RHEL 9 on x86_64. At this time there are no plans to support other platforms, though most likely it will work correctly on any POSIX-compliant x86_64 Linux distribution.

## TODO
* Implement LXM algorithm as an alternative to SplitMix64
