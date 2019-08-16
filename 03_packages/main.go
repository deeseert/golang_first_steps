package main

// To import more than one package you need to put them in parenteses
// no comas
import (
	"fmt"
	"math"
	// This is the package you have created and imported here
	"github.com/deeseert/go_crash_course/03_packages/strutil"
)

func main() {
	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Ceil(2.7))
	fmt.Println(math.Sqrt(16))
	// and packages used here
	fmt.Println(strutil.Reverse("Hello World!"))
}
