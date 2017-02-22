package main

import (
    "fmt"
    "os"
    "encoding/json"
    "io/ioutil"
    "phonecom-go-sdk"
)

//func createCallParams(inputFile string) swagger.CreateCallParams {
//
///*
//	var params swagger.CreateCallParams
//
//	params.CallerPhoneNumber = "+17324568888"
//	params.CallerExtension = 1749841
//	params.CallerCallerId = "+17329980138"
//	params.CallerPrivate = true
//	params.CalleePhoneNumber = "+17329980138"
//	params.CalleeExtension = 1749842
//	params.CalleeCallerId = "+17324568888"
//	params.CalleePrivate = true
//*/
//
//
//     file, e := ioutil.ReadFile(inputFile)
//    if e != nil {
//        fmt.Printf("File error: %v\n", e)
//        os.Exit(1)
//    }
//    fmt.Printf("%s\n", string(file))
//
//    var params  swagger.CreateCallParams
//    json.Unmarshal(file, &params)
//
//	return params
//}

func createDeviceParams(inputFile string) swagger.CreateDeviceParams {

	

	
     file, e := ioutil.ReadFile(inputFile)
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }
    fmt.Printf("%s\n", string(file))
    
    var params  swagger.CreateDeviceParams
    json.Unmarshal(file, &params)

	return params
}
func createMenuParams(inputFile string) swagger.CreateMenuParams {
    
    file, e := ioutil.ReadFile(inputFile)
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }
    fmt.Printf("%s\n", string(file))
    
    var params  swagger.CreateMenuParams
    json.Unmarshal(file, &params)

	return params


}

func replaceMenuParams(inputFile string) swagger.ReplaceMenuParams{
    
	 file, e := ioutil.ReadFile(inputFile)
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }
    fmt.Printf("%s\n", string(file))
    
    var params  swagger.ReplaceMenuParams
    json.Unmarshal(file, &params)

	return params





}
func createExtensionsParams(inputFile string) swagger.CreateExtensionParams {

	 file, e := ioutil.ReadFile(inputFile)
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }
    fmt.Printf("%s\n", string(file))
    
    var params  swagger.CreateExtensionParams
    json.Unmarshal(file, &params)

	return params
}
func createQueueParams(inputFile string) swagger.CreateQueueParams {
	 file, e := ioutil.ReadFile(inputFile)
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }
    fmt.Printf("%s\n", string(file))
    
    var params  swagger.CreateQueueParams
    json.Unmarshal(file, &params)

	return params

}

func createRouteParams(inputFile string) swagger.CreateRouteParams {
	 file, e := ioutil.ReadFile(inputFile)
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }
    fmt.Printf("%s\n", string(file))
    
    var params  swagger.CreateRouteParams
    json.Unmarshal(file, &params)

	return params


}

func createSmsParams(inputFile string) swagger.CreateSmsParams {
		 file, e := ioutil.ReadFile(inputFile)
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }
    fmt.Printf("%s\n", string(file))
    
    var params  swagger.CreateSmsParams
    json.Unmarshal(file, &params)

	return params
}

func createSubaccountParams(inputFile string) swagger.CreateSubaccountParams{
		 file, e := ioutil.ReadFile(inputFile)
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }
    fmt.Printf("%s\n", string(file))
    
    var params  swagger.CreateSubaccountParams
    json.Unmarshal(file, &params)

	return params


}

func createTrunkParams(inputFile string) swagger.CreateTrunkParams{
		 file, e := ioutil.ReadFile(inputFile)
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }
    fmt.Printf("%s\n", string(file))
    
    var params  swagger.CreateTrunkParams
    json.Unmarshal(file, &params)

	return params


}


func replaceExtensionParams(inputFile string) swagger.ReplaceExtensionParams {
	 file, e := ioutil.ReadFile(inputFile)
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }
    fmt.Printf("%s\n", string(file))
    
    var params  swagger.ReplaceExtensionParams
    json.Unmarshal(file, &params)

	return params
}

func createContactParams(inputFile string) swagger.CreateContactParams {

 file, e := ioutil.ReadFile(inputFile)
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }
    fmt.Printf("%s\n", string(file))
    
    var params  swagger.CreateContactParams
    json.Unmarshal(file, &params)

	return params
}

func createGroupParams(inputFile string) swagger.CreateGroupParams {

	file, e := ioutil.ReadFile(inputFile)
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }
    fmt.Printf("%s\n", string(file))
    
    var params  swagger.CreateGroupParams
    json.Unmarshal(file, &params)

	return params
}

func createPhoneNumberParams(inputFile string) swagger.CreatePhoneNumberParams {

	file, e := ioutil.ReadFile(inputFile)
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }
    fmt.Printf("%s\n", string(file))
    
    var params  swagger.CreatePhoneNumberParams
    json.Unmarshal(file, &params)

	return params
}

func replacePhoneNumberParams(inputFile string) swagger.ReplacePhoneNumberParams {

	file, e := ioutil.ReadFile(inputFile)
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }
    fmt.Printf("%s\n", string(file))
    
    var params  swagger.ReplacePhoneNumberParams
    json.Unmarshal(file, &params)

	return params
}
