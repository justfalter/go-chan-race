# Intro
I'm not sure if this describes a bug or not. I haven't found anything that 
says the behavior should be one way or another. I was just under the impression
that succesfully sending to a buffered channel meant that the next receive
from that channel (regardless of the goroutine or number of processors) would
find that item I sent.

# Description
Two goroutines. One buffered channel (chBuffered), one unbuffered channel (chUnbuffered). 
goroutine 1 sends to the chBuffered, and then sends to the chUnbuffered.
goroutine 2 receives chUnbuffered, and then tries to received from chBuffered (or times out immediately with <-time.After(0)).

# Expected behavior
goroutine 2 would always receive data in chBuffered after receiving data from chUnbuffered.

# Actual behavior
goroutine 2 sometimes doesn't receive data in chBuffered, instead going to the immediate timeout.

# Interesting
- With GOMAXPROCS=1, we always find data in chBuffered after chUnbuffered.
- The more recent version of go (1.9 and 1.10 seem to be more likely to run into
  this condition race than 1.8.x.

# Example runs
```
$ go build
$ ./go-chan-race
No data in chBuffered in iter 145
$ ./go-chan-race
No data in chBuffered in iter 617
$ GOMAXPROCS=1 ./go-chan-race
All done.
$ GOMAXPROCS=1 ./go-chan-race
All done.
```
