package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	t := time.Now().Add(500 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), t)
	// 尽管ctx可能会过期, 调用他的cancel可以防止父ctx和ctx生命周期过长
	defer cancel()
	select {
	case <-time.After(time.Second):
		fmt.Println("king")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}

}
