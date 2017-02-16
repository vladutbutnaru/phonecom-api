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
       cli.IntFlag{
      Name: "id",
      Value: 0,
        Usage: "ID of entity you want to operate",
    },
        cli.IntFlag{
      Name: "limit",
      Value: 5,
        Usage: "Upper limit of results you want to fetch",
    },
        cli.IntFlag{
      Name: "offset",
      Value: 0,
        Usage: "Offset of results you want to fetch",
    },
  }

  app.Action = func(c *cli.Context) error {
   
      switch c.String("command") {
          case "list-media":
     
     
       
        var mediaApi swagger.MediaApi = *swagger.NewMediaApi()
        mySlice1 := make([]string, 0)
        x,response,error := mediaApi.ListAccountMedia(1315091, mySlice1, mySlice1, "", "", int32( c.Int("limit")), int32( c.Int("offset")), "")
        if error!=nil {
            panic(error)
            return nil
            
        }
           _ = response
      fmt.Printf("%+v\n",x)
       
        
          case "get-recording":
           var mediaApi swagger.MediaApi = *swagger.NewMediaApi()
          x,response,error := mediaApi.GetAccountMedia(1315091, int32(c.Int("id")))
          if error!=nil {
            panic(error)
            return nil
            
        }
           _ = response
      fmt.Printf("%+v\n",x)
          
          case "list-menus":
           var menusApi swagger.MenusApi = *swagger.NewMenusApi()
           mySlice1 := make([]string, 0)
        x,response,error := menusApi.ListAccountMenus(1315091, mySlice1, mySlice1, "", "", int32( c.Int("limit")), int32( c.Int("offset")), "")
        if error!=nil {
            panic(error)
            return nil
            
        }
           _ = response
      fmt.Printf("%+v\n",x)
          
          case "get-menu":
          var menusApi swagger.MenusApi = *swagger.NewMenusApi()
          
          x,response,error := menusApi.GetAccountMenu(1315091, int32(c.Int("id")))
        if error!=nil {
            panic(error)
            return nil
            
        }
           _ = response
      fmt.Printf("%+v\n",x)
          
          case "list-queues":
           var queuesApi swagger.QueuesApi = *swagger.NewQueuesApi()
           mySlice1 := make([]string, 0)
        x,response,error := queuesApi.ListAccountQueues(1315091, mySlice1, mySlice1, "", "", int32( c.Int("limit")), int32( c.Int("offset")), "")
        if error!=nil {
            panic(error)
            return nil
            
        }
           _ = response
      fmt.Printf("%+v\n",x)
          
           case "get-queue":
         var queuesApi swagger.QueuesApi = *swagger.NewQueuesApi()
          
          x,response,error := queuesApi.GetAccountQueue(1315091, int32(c.Int("id")))
        if error!=nil {
            panic(error)
            return nil
            
        }
           _ = response
      fmt.Printf("%+v\n",x)
          
                    case "list-routes":
           var routesApi swagger.RoutesApi = *swagger.NewRoutesApi()
           mySlice1 := make([]string, 0)
        x,response,error := routesApi.ListAccountRoutes(1315091, mySlice1, mySlice1, "", "", int32( c.Int("limit")), int32( c.Int("offset")), "")
        if error!=nil {
            panic(error)
            return nil
            
        }
           _ = response
      fmt.Printf("%+v\n",x)
          
           case "get-route":
         var routesApi swagger.RoutesApi = *swagger.NewRoutesApi()
          
          x,response,error := routesApi.GetAccountRoute(1315091, int32(c.Int("id")))
        if error!=nil {
            panic(error)
            return nil
            
        }
           _ = response
      fmt.Printf("%+v\n",x)
          
           case "list-schedules":
           var schedulesApi swagger.SchedulesApi = *swagger.NewSchedulesApi()
           mySlice1 := make([]string, 0)
        x,response,error := schedulesApi.ListAccountSchedules(1315091, mySlice1, mySlice1, "", "", int32( c.Int("limit")), int32( c.Int("offset")), "")
        if error!=nil {
            panic(error)
            return nil
            
        }
           _ = response
      fmt.Printf("%+v\n",x)
          
              case "get-schedule":
        var schedulesApi swagger.SchedulesApi = *swagger.NewSchedulesApi()
          
          x,response,error := schedulesApi.GetAccountSchedule(1315091, int32(c.Int("id")))
        if error!=nil {
            panic(error)
            return nil
            
        }
           _ = response
      fmt.Printf("%+v\n",x)
          
           case "list-sms":
           var smsApi swagger.SmsApi = *swagger.NewSmsApi()
           mySlice1 := make([]string, 0)
        x,response,error := smsApi.ListAccountSms(1315091, mySlice1, "","", "", "", int32( c.Int("limit")), int32( c.Int("offset")), "")
        if error!=nil {
            panic(error)
            return nil
            
        }
           _ = response
      fmt.Printf("%+v\n",x)
          
                case "get-sms":
      var smsApi swagger.SmsApi = *swagger.NewSmsApi()
          
          x,response,error := smsApi.GetAccountSms(1315091, int32(c.Int("id")))
        if error!=nil {
            panic(error)
            return nil
            
        }
           _ = response
      fmt.Printf("%+v\n",x)
          
           case "list-available-phone-numbers":
           var availableNumbersApi swagger.AvailablenumbersApi = *swagger.NewAvailablenumbersApi()
           mySlice1 := make([]string, 0)
        x,response,error := availableNumbersApi.ListAvailablePhoneNumbers(mySlice1, mySlice1,mySlice1,mySlice1,mySlice1,mySlice1,mySlice1,mySlice1,mySlice1,mySlice1,"", "", "", int32( c.Int("limit")), int32( c.Int("offset")), "")
        if error!=nil {
            panic(error)
            return nil
            
        }
           _ = response
      fmt.Printf("%+v\n",x)
          
              case "list-subaccounts":
           var subaccountsApi swagger.SubaccountsApi = *swagger.NewSubaccountsApi()
           mySlice1 := make([]string, 0)
        x,response,error := subaccountsApi.ListAccountSubaccounts(1315091, mySlice1, mySlice1, "", "", int32( c.Int("limit")), int32( c.Int("offset")), "")
        if error!=nil {
            panic(error)
            return nil
            
        }
           _ = response
      fmt.Printf("%+v\n",x)
          
          
      default:
          fmt.Println("Command not valid")
      }
    return nil
  }

  app.Run(os.Args)
}