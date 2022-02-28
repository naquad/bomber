# Bomber

DDoS tool against Russian state propaganda.
Go version of this [page](https://omega1900.github.io/stop-russian-desinformation/).
The list of sites can be found in `main.go`.

# Installation

You can build statically linked binary optimized for size:

```
make
```

It can be distributed as needed.
(You can safely skip errors regarding UPX as it is used only for compression).

Or install it using `go get`:

```
$ go get github.com/naquad/bomber
```

Enjoy!

# Usage

Run the executable on the destination host:
```
$ ./bomber # or w/o path: bomber
```

If the resource consumption is more than you can spare then you can tweak some limits:
```
$ ./bomber -help
Usage of ./bomber:
  -connections int
    	number of simultaneous requests (default 1000)
  -cores int
    	number of CPU cores to use (default 8)
  -timeout int
    	HTTP connection timeout (in seconds) (default 1)
```

# Glory To Ukraine!

![Glory To Ukraine!](https://imgur.com/40DRyRQ.png)
