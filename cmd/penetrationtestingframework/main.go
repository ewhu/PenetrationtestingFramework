// cmd/penetrationtestingframework/main.go
package main

import (
"flag"
"log"
"os"

"penetrationtestingframework/internal/penetrationtestingframework"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := penetrationtestingframework.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
