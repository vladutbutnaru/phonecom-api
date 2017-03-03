'use strict';

exports.createAccountDevice = function(args, res, next) {
  /**
   * parameters expected in the args:
  * account_id (Integer)
  * data (CreateDeviceParams)
  **/
    var examples = {};
  examples['application/json'] = {
  "sip_authentication" : {
    "password" : "aeiou",
    "port" : "",
    "host" : "aeiou",
    "username" : "aeiou"
  },
  "name" : "aeiou",
  "id" : "",
  "lines" : [ {
    "extension" : {
      "extension" : "",
      "name" : "aeiou",
      "id" : ""
    },
    "line" : ""
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

exports.getAccountDevice = function(args, res, next) {
  /**
   * parameters expected in the args:
  * account_id (Integer)
  * device_id (Integer)
  **/
    var examples = {};
  examples['application/json'] = {
  "sip_authentication" : {
    "password" : "aeiou",
    "port" : "",
    "host" : "aeiou",
    "username" : "aeiou"
  },
  "name" : "aeiou",
  "id" : "",
  "lines" : [ {
    "extension" : {
      "extension" : "",
      "name" : "aeiou",
      "id" : ""
    },
    "line" : ""
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

exports.listAccountDevices = function(args, res, next) {
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

exports.replaceAccountDevice = function(args, res, next) {
  /**
   * parameters expected in the args:
  * account_id (Integer)
  * device_id (Integer)
  * data (CreateDeviceParams)
  **/
    var examples = {};
  examples['application/json'] = {
  "sip_authentication" : {
    "password" : "aeiou",
    "port" : "",
    "host" : "aeiou",
    "username" : "aeiou"
  },
  "name" : "aeiou",
  "id" : "",
  "lines" : [ {
    "extension" : {
      "extension" : "",
      "name" : "aeiou",
      "id" : ""
    },
    "line" : ""
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

