package main

import (
  "fmt"
  "os"

  "github.com/urfave/cli"
  "phonecom-go-sdk"
)

func main() {
  app := cli.NewApp()

  app.Flags = []cli.Flag {
    cli.StringFlag{
      Name: "command",
      Value: "list-accounts",
        Usage: "API command that you want to execute",
    },
  }

  app.Action = func(c *cli.Context) error {
   
  
    if c.String("command") == "list-media" {
        var mediaApi swagger.MediaApi
        mySlice1 := make([]string, 0)
        x,response,error := mediaApi.ListAccountMedia(1145, mySlice1, mySlice1, "", "", 3, 1, "")
        fmt.Println(x, response, error)
        
      
    } else {
      fmt.Println("Command not valid")
    }
    return nil
  }

  app.Run(os.Args)
}