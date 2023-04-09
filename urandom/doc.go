// Package urandom provides a function returning a random uint64 from the OS.
// Since urn only supports RHEL, this is implemented by reading bytes from /dev/urandom. Information on /dev/urandom can be found [here].
//
// [here]: https://stochastics.ai/Edge_2022.html
package urandom
