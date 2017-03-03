'use strict';

exports.listAvailablePhoneNumbers = function(args, res, next) {
  /**
   * parameters expected in the args:
  * filters[phone_number] (List)
  * filters[country_code] (List)
  * filters[npa] (List)
  * filters[nxx] (List)
  * filters[xxxx] (List)
  * filters[city] (List)
  * filters[province] (List)
  * filters[country] (List)
  * filters[price] (List)
  * filters[category] (List)
  * filters[is_toll_free] (List)
  * sort[internal] (String)
  * sort[price] (String)
  * sort[phone_number] (String)
  * limit (Integer)
  * offset (Integer)
  * fields (String)
  **/
    var examples = {};
  examples['application/json'] = {
  "total" : "",
  "offset" : "",
  "limit" : "",
  "filters" : {
    "country_code" : "aeiou",
    "country" : "aeiou",
    "province" : "aeiou",
    "city" : "aeiou",
    "price" : "aeiou",
    "xxxx" : "aeiou",
    "phone_number" : "aeiou",
    "category" : "aeiou",
    "npa" : [ "" ],
    "nxx" : "aeiou"
  },
  "sort" : {
    "internal" : "aeiou",
    "price" : "aeiou",
    "phone_number" : "aeiou"
  },
  "items" : ""
};
  if(Object.keys(examples).length > 0) {
    res.setHeader('Content-Type', 'application/json');
    res.end(JSON.stringify(examples[Object.keys(examples)[0]] || {}, null, 2));
  }
  else {
    res.end();
  }
  
}

