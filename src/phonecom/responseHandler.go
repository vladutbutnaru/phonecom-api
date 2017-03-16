package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/waiyuen/Phone.com-API-SDK-go"
	"strings"
     "github.com/imdario/mergo"
     
)

type ResponseHandler struct {
	param CliParams
}



func (h *ResponseHandler) handle(
	x interface{},
	response *swagger.APIResponse,
	error error) (error, map[string]interface{}) {

	if error != nil {

		if h.param.verbose {
			fmt.Println("Error while getting response")
		}

		return error, nil
	}

	payload := response.Payload

	if payload == nil {

		if h.param.verbose {
			fmt.Println("Null response payload")
		}

		return errors.New(msgCouldNotGetResponse), nil
	}

	validatedJson := validateJson(string(payload))

	if validatedJson == nil {

		if h.param.verbose {
			fmt.Println("Could not unmarshal API json response")
		}

		return errors.New(msgInvalidJson), nil
	}

	message := validateResponse(validatedJson)

	if message != "" {

		if h.param.verbose {
			fmt.Printf("%+v\n%s\n", x, response)
		}

		return errors.New(message), nil
	} else {

		var jsonObject interface{}
		if !h.param.fullList && validatedJson["items"] != nil {
			jsonObject = validatedJson["items"]
		} else {
			jsonObject = validatedJson
		}

		if h.param.verbose {
			fmt.Printf("\nAPI Response [Verbose]:\n%+v\n", x)
			fmt.Println()
		}

		fmt.Println("API Response:")
		fmt.Println()

		var outputType = "json"
		if strings.EqualFold(h.param.outputFormat, "csv") {
			outputType = "csv"
		}
        
       
        
        if validatedJson["total"] != nil {
            var currentNum = 0
            var items = validatedJson["items"].([]interface{})
            currentNum = len(items)
           
               firstId :=validatedJson["total"].(json.Number)
            totalGood, _ := json.Number.Int64(firstId)
            var total = int(totalGood)
       
           
            
            for  total  >   currentNum {
               
                
                 err,respPage := h.handle(x,response,error)
                if err != nil {
                    return err, nil
                    
                }
                 if err := mergo.Merge(&validatedJson, respPage); err != nil {
                     
                    }   
  
                
                //join validatedJson with respPage and that is it
                 currentNum = currentNum + currentNum
            }
            
           
        }
        
        
		if outputType == "csv" {
			exportToCsv(jsonObject)
		} else if outputType == "json" {
			jsonOutput, err := json.MarshalIndent(jsonObject, "", "  ")
			if err != nil {
				return err, nil
			}
			fmt.Printf("%s\n", jsonOutput)
		} else {
			fmt.Println("Invalid output type")
		}

	}

	return nil, validatedJson
}
