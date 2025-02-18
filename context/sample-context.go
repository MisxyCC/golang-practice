package context

import (
	"context"
	"fmt"
)

func Demo() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "myKey", "sample-value")
	doSomething(ctx)
}

func doSomething(ctx context.Context) {
	fmt.Printf("Do something mykey's value is: %s\n", ctx.Value("myKey"))
	anotherContext := context.WithValue(ctx, "myKey", "anotherValue")
	doAnother(anotherContext)
	fmt.Printf("Do something mykey's value is: %s\n", ctx.Value("myKey"))
}

func doAnother(ctx context.Context) {
	fmt.Printf("doAnother: myKey's value is %s\n", ctx.Value("myKey"))
}