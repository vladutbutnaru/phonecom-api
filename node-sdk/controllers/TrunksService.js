'use strict';

exports.createAccountTrunk = function(args, res, next) {
  /**
   * parameters expected in the args:
  * account_id (Integer)
  * data (CreateTrunkParams)
  **/
    var examples = {};
  examples['application/json'] = {
  "error_message" : "",
  "max_minutes_per_month" : "",
  "greeting" : {
    "name" : "aeiou",
    "id" : ""
  },
  "codecs" : [ "aeiou" ],
  "name" : "aeiou",
  "id" : "",
  "max_concurrent_calls" : "",
  "uri" : "aeiou"
};
  if(Object.keys(examples).length > 0) {
    res.setHeader('Content-Type', 'application/json');
    res.end(JSON.stringify(examples[Object.keys(examples)[0]] || {}, null, 2));
  }
  else {
    res.end();
  }
  
}

exports.deleteAccountTrunk = function(args, res, next) {
  /**
   * parameters expected in the args:
  * account_id (Integer)
  * trunk_id (Integer)
  **/
    var examples = {};
  examples['application/json'] = {
  "success" : true
};
  if(Object.keys(examples).length > 0) {
    res.setHeader('Content-Type', 'application/json');
    res.end(JSON.stringify(examples[Object.keys(examples)[0]] || {}, null, 2));
  }
  else {
    res.end();
  }
  
}

exports.getAccountTrunk = function(args, res, next) {
  /**
   * parameters expected in the args:
  * account_id (Integer)
  * trunk_id (Integer)
  **/
    var examples = {};
  examples['application/json'] = {
  "error_message" : "",
  "max_minutes_per_month" : "",
  "greeting" : {
    "name" : "aeiou",
    "id" : ""
  },
  "codecs" : [ "aeiou" ],
  "name" : "aeiou",
  "id" : "",
  "max_concurrent_calls" : "",
  "uri" : "aeiou"
};
  if(Object.keys(examples).length > 0) {
    res.setHeader('Content-Type', 'application/json');
    res.end(JSON.stringify(examples[Object.keys(examples)[0]] || {}, null, 2));
  }
  else {
    res.end();
  }
  
}

exports.listAccountTrunks = function(args, res, next) {
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
    "error_message" : "",
    "max_minutes_per_month" : "",
    "greeting" : {
      "name" : "aeiou",
      "id" : ""
    },
    "codecs" : [ "aeiou" ],
    "name" : "aeiou",
    "id" : "",
    "max_concurrent_calls" : "",
    "uri" : "aeiou"
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

exports.replaceAccountTrunk = function(args, res, next) {
  /**
   * parameters expected in the args:
  * account_id (Integer)
  * trunk_id (Integer)
  * data (CreateTrunkParams)
  **/
    var examples = {};
  examples['application/json'] = {
  "error_message" : "",
  "max_minutes_per_month" : "",
  "greeting" : {
    "name" : "aeiou",
    "id" : ""
  },
  "codecs" : [ "aeiou" ],
  "name" : "aeiou",
  "id" : "",
  "max_concurrent_calls" : "",
  "uri" : "aeiou"
};
  if(Object.keys(examples).length > 0) {
    res.setHeader('Content-Type', 'application/json');
    res.end(JSON.stringify(examples[Object.keys(examples)[0]] || {}, null, 2));
  }
  else {
    res.end();
  }
  
}

