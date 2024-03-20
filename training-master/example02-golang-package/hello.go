/*
package main

import (
	"example/user/hello/morestrings"
	"fmt"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
}
*/

package main

import (
	"fmt"

	"example/user/hello/morestrings"
	"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
