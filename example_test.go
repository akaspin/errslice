package errslice_test

import (
	"fmt"
	"github.com/akaspin/errslice"
)

func ExampleAppend() {
	var err error
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			err = errslice.Append(err, fmt.Errorf("bad:%d", i))
		}
	}
	fmt.Println(err)
	// Output:
	// bad:0,bad:2,bad:4
}
