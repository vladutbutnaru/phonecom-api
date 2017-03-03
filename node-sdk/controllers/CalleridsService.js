'use strict';

exports.getCallerIds = function(args, res, next) {
  /**
   * parameters expected in the args:
  * account_id (Integer)
  * extension_id (Integer)
  * filters[number] (List)
  * filters[name] (List)
  * sort[number] (String)
  * sort[name] (String)
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
    "number" : "aeiou",
    "name" : "aeiou"
  },
  "sort" : {
    "number" : "aeiou",
    "name" : "aeiou"
  },
  "items" : [ {
    "number" : "aeiou",
    "name" : "aeiou",
    "type" : "aeiou"
  } ]
};
  if(Object.keys(examples).length > 0) {
    res.setHeader('Content-Type', 'application/json');
    res.end(JSON.stringify(examples[Object.keys(examples)[0]] || {}, null, 2));
  }
  else {
    res.end();
  }
  
}

