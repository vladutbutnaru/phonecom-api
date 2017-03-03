'use strict';

exports.createAccountMenu = function(args, res, next) {
  /**
   * parameters expected in the args:
  * account_id (Integer)
  * data (CreateMenuParams)
  **/
    var examples = {};
  examples['application/json'] = {
  "allow_extension_dial" : true,
  "greeting" : {
    "name" : "aeiou",
    "id" : ""
  },
  "timeout_handler" : {
    "name" : "aeiou",
    "id" : ""
  },
  "name" : "aeiou",
  "options" : "",
  "id" : "",
  "keypress_wait_time" : "",
  "keypress_error" : ""
};
  if(Object.keys(examples).length > 0) {
    res.setHeader('Content-Type', 'application/json');
    res.end(JSON.stringify(examples[Object.keys(examples)[0]] || {}, null, 2));
  }
  else {
    res.end();
  }
  
}

exports.deleteAccountMenu = function(args, res, next) {
  /**
   * parameters expected in the args:
  * account_id (Integer)
  * menu_id (Integer)
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

exports.getAccountMenu = function(args, res, next) {
  /**
   * parameters expected in the args:
  * account_id (Integer)
  * menu_id (Integer)
  **/
    var examples = {};
  examples['application/json'] = {
  "allow_extension_dial" : true,
  "greeting" : {
    "name" : "aeiou",
    "id" : ""
  },
  "timeout_handler" : {
    "name" : "aeiou",
    "id" : ""
  },
  "name" : "aeiou",
  "options" : "",
  "id" : "",
  "keypress_wait_time" : "",
  "keypress_error" : ""
};
  if(Object.keys(examples).length > 0) {
    res.setHeader('Content-Type', 'application/json');
    res.end(JSON.stringify(examples[Object.keys(examples)[0]] || {}, null, 2));
  }
  else {
    res.end();
  }
  
}

exports.listAccountMenus = function(args, res, next) {
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

exports.replaceAccountMenu = function(args, res, next) {
  /**
   * parameters expected in the args:
  * account_id (Integer)
  * menu_id (Integer)
  * data (ReplaceMenuParams)
  **/
    var examples = {};
  examples['application/json'] = {
  "allow_extension_dial" : true,
  "greeting" : {
    "name" : "aeiou",
    "id" : ""
  },
  "timeout_handler" : {
    "name" : "aeiou",
    "id" : ""
  },
  "name" : "aeiou",
  "options" : "",
  "id" : "",
  "keypress_wait_time" : "",
  "keypress_error" : ""
};
  if(Object.keys(examples).length > 0) {
    res.setHeader('Content-Type', 'application/json');
    res.end(JSON.stringify(examples[Object.keys(examples)[0]] || {}, null, 2));
  }
  else {
    res.end();
  }
  
}

