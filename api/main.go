package main
import (
	"parmyay"
	"os"
)

func main() {
	os.Setenv("DB_NAME", "data.db");
    parmyay.InitDb();
	router := parmyay.SetupRouter(false, true)
	router.Run(":8900")
}
