package main

import (
  "fmt"
  "os"

  "github.com/urfave/cli"
  "phonecom-go-sdk"
)

func main() {
  app := cli.NewApp()

  app.Action = func(c *cli.Context) error {
    AccountFull GetAccount(1145)
      
    fmt.Printf("Hello %q", c.Args().Get(0))
   
 
      
      
      
      
    return nil
  }

  app.Run(os.Args)
}