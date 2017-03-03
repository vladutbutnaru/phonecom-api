'use strict';

exports.getAccountExpressSrvCode = function(args, res, next) {
  /**
   * parameters expected in the args:
  * account_id (Integer)
  * code_id (Integer)
  **/
    var examples = {};
  examples['application/json'] = {
  "expire_date" : "",
  "id" : "",
  "express_service_code" : "aeiou"
};
  if(Object.keys(examples).length > 0) {
    res.setHeader('Content-Type', 'application/json');
    res.end(JSON.stringify(examples[Object.keys(examples)[0]] || {}, null, 2));
  }
  else {
    res.end();
  }
  
}

exports.listAccountExpressSrvCodes = function(args, res, next) {
  /**
   * parameters expected in the args:
  * account_id (Integer)
  * filters[id] (List)
  **/
    var examples = {};
  examples['application/json'] = {
  "total" : "",
  "offset" : "",
  "limit" : "",
  "filters" : {
    "id" : "aeiou"
  },
  "sort" : "",
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

