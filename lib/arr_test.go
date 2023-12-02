package lib

import (
	"fmt"
	"strings"
)

func main() {
	websites := []string{"http://foo.com", "https://bar.com", "https://gosamples.dev"}
	httpsWebsites := filter(websites, func(v string) bool {
		return strings.HasPrefix(v, "https://")
	})
	fmt.Println(httpsWebsites)

	numbers := []int{1, 2, 3, 4, 5, 6}
	divisibleBy2 := filter(numbers, func(v int) bool {
		return v%2 == 0
	})
	fmt.Println(divisibleBy2)
}
