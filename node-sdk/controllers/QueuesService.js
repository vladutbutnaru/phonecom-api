'use strict';

exports.createAccountQueue = function(args, res, next) {
  /**
   * parameters expected in the args:
  * account_id (Integer)
  * data (CreateQueueParams)
  **/
    var examples = {};
  examples['application/json'] = {
  "caller_id_type" : "aeiou",
  "hold_music" : {
    "name" : "aeiou",
    "id" : ""
  },
  "max_hold_time" : "",
  "greeting" : {
    "name" : "aeiou",
    "id" : ""
  },
  "members" : "",
  "name" : "aeiou",
  "ring_time" : "",
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

exports.deleteAccountQueue = function(args, res, next) {
  /**
   * parameters expected in the args:
  * account_id (Integer)
  * queue_id (Integer)
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

exports.getAccountQueue = function(args, res, next) {
  /**
   * parameters expected in the args:
  * account_id (Integer)
  * queue_id (Integer)
  **/
    var examples = {};
  examples['application/json'] = {
  "caller_id_type" : "aeiou",
  "hold_music" : {
    "name" : "aeiou",
    "id" : ""
  },
  "max_hold_time" : "",
  "greeting" : {
    "name" : "aeiou",
    "id" : ""
  },
  "members" : "",
  "name" : "aeiou",
  "ring_time" : "",
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

exports.listAccountQueues = function(args, res, next) {
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

exports.replaceAccountQueue = function(args, res, next) {
  /**
   * parameters expected in the args:
  * account_id (Integer)
  * queue_id (Integer)
  * data (CreateQueueParams)
  **/
    var examples = {};
  examples['application/json'] = {
  "caller_id_type" : "aeiou",
  "hold_music" : {
    "name" : "aeiou",
    "id" : ""
  },
  "max_hold_time" : "",
  "greeting" : {
    "name" : "aeiou",
    "id" : ""
  },
  "members" : "",
  "name" : "aeiou",
  "ring_time" : "",
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

