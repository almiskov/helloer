package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/almiskov/helloer/internal/app"
	"github.com/almiskov/helloer/internal/sayer"
)

const version = "0.0.0"

func main() {
	var say string

	flagset := flag.NewFlagSet(fmt.Sprintf("flagset %s", version), flag.ExitOnError)
	flagset.StringVar(&say, "say", "nothing to say", "what to say")
	flagset.Parse(os.Args[1:])

	s := sayer.NewWriterSayer(os.Stdout)

	a := app.New(s)

	a.Run(say)
}
