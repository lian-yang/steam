# Steam

Steam is a library for interactions with [Steam](https://steamcommunity.com), it's written in Go.  
Steam tries to keep-it-simple and does not add extra non-sense.  There are absolutely no internal-polling or such,
      everything is up to you, all it does is wrap around Steam API.

## Why?

- You don't want a library to be "re-trying" automatically
- You don't want a library to be doing your homework
- You are an on-point person and just want stuff that works as-needed

## Install

```
go get github.com/lian-yang/steam
```

## Example

```go
package main

import (
	"log"
	"os"

	"github.com/lian-yang/steam"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	timeTip, err := steam.GetTimeTip()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Time tip: %#v\n", timeTip)
	timeDiff := time.Duration(timeTip.Time - time.Now().Unix())
	
	session := steam.NewSession(&http.Client{}, "")
	if err := session.Login(os.Getenv("steamAccount"), os.Getenv("steamPassword"), os.Getenv("steamSharedSecret"), timeDiff); err != nil {
		log.Fatal(err)
	}
	log.Print("Login successful")
}
```

Find more examples in the examples/ directory.  Even better is to read through the source code, it's simple and
straight-forward to understand.

## License

LGPL 2.1
