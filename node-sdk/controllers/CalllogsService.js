'use strict';

exports.getAccountCallLog = function(args, res, next) {
  /**
   * parameters expected in the args:
  * account_id (Integer)
  * callLog_id (Integer)
  **/
    var examples = {};
  examples['application/json'] = {
  "call_number" : "aeiou",
  "call_recording" : "aeiou",
  "extension" : {
    "extension" : "",
    "name" : "aeiou",
    "id" : ""
  },
  "caller_cnam" : "aeiou",
  "created_at" : "aeiou",
  "type" : "aeiou",
  "uuid" : "aeiou",
  "call_duration" : "",
  "final_action" : "aeiou",
  "start_time" : "aeiou",
  "is_monitored" : "aeiou",
  "caller_id" : "aeiou",
  "details" : [ {
    "start_time" : "",
    "voip_phone_id" : "",
    "type" : "aeiou",
    "id_value" : "",
    "voip_id" : ""
  } ],
  "id" : "aeiou",
  "called_number" : "aeiou",
  "direction" : "aeiou"
};
  if(Object.keys(examples).length > 0) {
    res.setHeader('Content-Type', 'application/json');
    res.end(JSON.stringify(examples[Object.keys(examples)[0]] || {}, null, 2));
  }
  else {
    res.end();
  }
  
}

exports.listAccountCallLogs = function(args, res, next) {
  /**
   * parameters expected in the args:
  * account_id (Integer)
  * filters[id] (List)
  * filters[start_time] (List)
  * filters[created_at] (String)
  * filters[direction] (String)
  * filters[called_number] (String)
  * filters[type] (String)
  * filters[extension] (List)
  * sort[id] (String)
  * sort[start_time] (String)
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
    "start_time" : "aeiou",
    "created_at" : "aeiou",
    "id" : "aeiou",
    "called_number" : "aeiou",
    "type" : "aeiou",
    "direction" : "aeiou"
  },
  "sort" : {
    "start_time" : "aeiou",
    "created_at" : "aeiou",
    "id" : "aeiou"
  },
  "items" : [ {
    "call_number" : "aeiou",
    "call_recording" : "aeiou",
    "extension" : {
      "extension" : "",
      "name" : "aeiou",
      "id" : ""
    },
    "caller_cnam" : "aeiou",
    "created_at" : "aeiou",
    "type" : "aeiou",
    "uuid" : "aeiou",
    "call_duration" : "",
    "final_action" : "aeiou",
    "start_time" : "aeiou",
    "is_monitored" : "aeiou",
    "caller_id" : "aeiou",
    "details" : [ {
      "start_time" : "",
      "voip_phone_id" : "",
      "type" : "aeiou",
      "id_value" : "",
      "voip_id" : ""
    } ],
    "id" : "aeiou",
    "called_number" : "aeiou",
    "direction" : "aeiou"
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

