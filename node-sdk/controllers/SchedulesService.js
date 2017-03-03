'use strict';

exports.getAccountSchedule = function(args, res, next) {
  /**
   * parameters expected in the args:
  * account_id (Integer)
  * schedule_id (Integer)
  **/
    var examples = {};
  examples['application/json'] = {
  "name" : "aeiou",
  "id" : ""
};
  if(Object.keys(examples).length > 0) {
    res.setHeader('Content-Type', 'application/json');
    res.end(JSON.stringify(examples[Object.keys(examples)[0]] || {}, null, 2));
  }
  else {
    res.end();
  }
  
}

exports.listAccountSchedules = function(args, res, next) {
  /**
   * parameters expected in the args:
  * account_id (Integer)
  * filters[id] (List)
  * filters[name] (List)
  * sort[id] (String)
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
    "name" : "aeiou",
    "id" : "aeiou"
  },
  "sort" : {
    "name" : "aeiou",
    "id" : "aeiou"
  },
  "items" : [ {
    "name" : "aeiou",
    "id" : ""
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

