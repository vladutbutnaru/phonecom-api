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
   
  
    if c.String("command") == "list-extensions" {
        var extensionsApi swagger.ExtensionsApi = *swagger.NewExtensionsApi()
        mySlice1 := make([]string, 0)
        x,response,error := extensionsApi.ListAccountExtensions(1315091, mySlice1, mySlice1, mySlice1, "", "", "", 3, 1, "")
        fmt.Printf("%+v\n",x)
        fmt.Println("")
        fmt.Println("Response:")
        fmt.Printf("%+v\n", response)
        fmt.Println(error)
        
      
    } else {
      fmt.Println("Command not valid")
    }
    return nil
  }

  app.Run(os.Args)
}
