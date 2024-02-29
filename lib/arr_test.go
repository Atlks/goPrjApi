package lib

import (
	"fmt"
	"strings"
	"testing"
)

func TestHandler10016(t *testing.T) {
	websites := []string{"http://foo.com", "https://bar.com", "https://gosamples.dev"}
	httpsWebsites := Filter(websites, func(v string) bool {
		return strings.HasPrefix(v, "https://")
	})
	fmt.Println(httpsWebsites)

	numbers := []int{1, 2, 3, 4, 5, 6}
	divisibleBy2 := Filter(numbers, func(v int) bool {
		return v%2 == 0
	})
	fmt.Println(divisibleBy2)
}
