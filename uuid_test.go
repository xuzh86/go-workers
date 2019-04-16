package workers

import (
	"fmt"
)

func main() {
	u1 := Must(NewV4())
	id := fmt.Sprintf("%s", u1)
	fmt.Println(id)
}