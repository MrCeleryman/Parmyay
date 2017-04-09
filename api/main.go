package main
import (
	"parmyay"
)

func main() {
    parmyay.InitDb();
	router := parmyay.SetupRouter(false, true)
	router.Run(":8900")
}
