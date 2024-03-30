package dotenv_test

import (
	"flag"
	"fmt"
	"os"
	"testing"

	"github.com/yenole/dotenv.git"
)

func TestOk(t *testing.T) {
	dotenv.Ok()
	fmt.Printf("os.Getenv(\"LIST\"): %v\n", os.Getenv("LIST"))
}

var (
	name = flag.String("name", "", "name is who")
)

func TestFlagParse(t *testing.T) {
	dotenv.FlagParse()
	fmt.Printf("name: %v\n", *name)
}

func TestFlagParsePrefix(t *testing.T) {
	dotenv.FlagParse(dotenv.Option{Envs: []string{".env"}, Prefix: "DOTENV_"})
	fmt.Printf("name: %v\n", *name)
}
