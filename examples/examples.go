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
