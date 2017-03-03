'use strict';

exports.createAccountSms = function(args, res, next) {
  /**
   * parameters expected in the args:
  * account_id (Integer)
  * data (CreateSmsParams)
  **/
    var examples = {};
  examples['application/json'] = {
  "created_at" : "2000-01-23T04:56:07.000+00:00",
  "from" : "aeiou",
  "id" : "aeiou",
  "to" : [ {
    "number" : "aeiou",
    "status" : "aeiou"
  } ],
  "text" : "aeiou",
  "direction" : "aeiou",
  "created_epoch" : ""
};
  if(Object.keys(examples).length > 0) {
    res.setHeader('Content-Type', 'application/json');
    res.end(JSON.stringify(examples[Object.keys(examples)[0]] || {}, null, 2));
  }
  else {
    res.end();
  }
  
}

exports.getAccountSms = function(args, res, next) {
  /**
   * parameters expected in the args:
  * account_id (Integer)
  * sms_id (Integer)
  **/
    var examples = {};
  examples['application/json'] = {
  "created_at" : "2000-01-23T04:56:07.000+00:00",
  "from" : "aeiou",
  "id" : "aeiou",
  "to" : [ {
    "number" : "aeiou",
    "status" : "aeiou"
  } ],
  "text" : "aeiou",
  "direction" : "aeiou",
  "created_epoch" : ""
};
  if(Object.keys(examples).length > 0) {
    res.setHeader('Content-Type', 'application/json');
    res.end(JSON.stringify(examples[Object.keys(examples)[0]] || {}, null, 2));
  }
  else {
    res.end();
  }
  
}

exports.listAccountSms = function(args, res, next) {
  /**
   * parameters expected in the args:
  * account_id (Integer)
  * filters[id] (List)
  * filters[direction] (String)
  * filters[from] (String)
  * sort[id] (String)
  * sort[created_at] (String)
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
    "from" : "aeiou",
    "id" : "aeiou",
    "direction" : "aeiou"
  },
  "sort" : {
    "created_at" : "aeiou",
    "id" : "aeiou"
  },
  "items" : [ {
    "created_at" : "2000-01-23T04:56:07.000+00:00",
    "from" : "aeiou",
    "id" : "aeiou",
    "to" : [ {
      "number" : "aeiou",
      "status" : "aeiou"
    } ],
    "text" : "aeiou",
    "direction" : "aeiou",
    "created_epoch" : ""
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

