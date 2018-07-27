# ctxtimecop

ctxtimecop provides "time travel", "time freezing" capabilities with context for testing or anything else.
This is referenced from [go-timecop](https://github.com/bluele/go-timecop).

# Getting started

## Installation

```
$ go get -u github.com/malt03/ctxtimecop
```

## Example

```go
// examples/examples.go
package main

import (
	"context"
	"fmt"

	"github.com/malt03/ctxtimecop"
)

func main() {
	ctx := context.Background()
	ctxWithoutTravel := context.Background()
	t := timecop.Now(ctx)
	fmt.Printf("current: %v\n", t)
	fmt.Println("Dive into the future!")
	ctx = timecop.WithTravel(ctx, t.AddDate(1, 0, 0))

	for i := 0; i < 3; i++ {
		fmt.Printf("future!: %v\n", timecop.Now(ctx))
		fmt.Printf("now....: %v\n", timecop.Now(ctxWithoutTravel))
	}

	fmt.Println("Return to the current.")
	ctx = timecop.WithReturn(ctx)

	fmt.Printf("current: %v\n", timecop.Now(ctx))
}
```

```
$ go run examples/examples.go
current: 2018-07-27 04:28:40.401803 +0000 UTC
Dive into the future!
future!: 2019-07-27 04:28:40.401806 +0000 UTC
now....: 2018-07-27 04:28:40.402078 +0000 UTC
future!: 2019-07-27 04:28:40.401823 +0000 UTC
now....: 2018-07-27 04:28:40.402091 +0000 UTC
future!: 2019-07-27 04:28:40.401834 +0000 UTC
now....: 2018-07-27 04:28:40.402102 +0000 UTC
Return to the current.
current: 2018-07-27 04:28:40.402111 +0000 UTC
```

# Author

**Koji Murata**

* <http://github.com/malt03>
* <malt.koji@gmail.com>
