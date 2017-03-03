'use strict';

exports.listAvailablePhoneNumberRegions = function(args, res, next) {
  /**
   * parameters expected in the args:
  * filters[country_code] (List)
  * filters[npa] (List)
  * filters[nxx] (List)
  * filters[is_toll_free] (List)
  * filters[city] (List)
  * filters[province_postal_code] (List)
  * filters[country_postal_code] (List)
  * sort[country_code] (String)
  * sort[npa] (String)
  * sort[nxx] (String)
  * sort[is_toll_free] (String)
  * sort[city] (String)
  * sort[province_postal_code] (String)
  * sort[country_postal_code] (String)
  * limit (Integer)
  * offset (Integer)
  * fields (String)
  * group_by (List)
  **/
    var examples = {};
  examples['application/json'] = {
  "total" : "",
  "offset" : "",
  "limit" : "",
  "group_by" : "",
  "filters" : {
    "country_code" : "aeiou",
    "province_postal_code" : "aeiou",
    "country_postal_code" : "aeiou",
    "city" : "aeiou",
    "is_toll_free" : "aeiou",
    "npa" : [ "" ],
    "nxx" : "aeiou"
  },
  "sort" : {
    "country_code" : "aeiou",
    "province_postal_code" : "aeiou",
    "country_postal_code" : "aeiou",
    "city" : "aeiou",
    "is_toll_free" : "aeiou",
    "npa" : "aeiou",
    "nxx" : "aeiou"
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

