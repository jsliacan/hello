package hello

import (
	"fmt"

	"github.com/jsliacan/hello/utils"
)

func Greet() {
	fmt.Println("Hello stranger!")
	fmt.Println(utils.noBigDeal())
}
