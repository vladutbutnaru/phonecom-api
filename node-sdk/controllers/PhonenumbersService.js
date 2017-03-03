'use strict';

exports.createAccountPhoneNumber = function(args, res, next) {
  /**
   * parameters expected in the args:
  * account_id (Integer)
  * data (CreatePhoneNumberParams)
  **/
    var examples = {};
  examples['application/json'] = {
  "call_notifications" : {
    "emails" : [ "aeiou" ],
    "sms" : "aeiou"
  },
  "route" : {
    "name" : "aeiou",
    "id" : ""
  },
  "block_anonymous" : true,
  "caller_id" : {
    "name" : "aeiou",
    "type" : "aeiou"
  },
  "name" : "aeiou",
  "phone_number" : "aeiou",
  "id" : "",
  "block_incoming" : true,
  "sms_forwarding" : {
    "extension" : {
      "extension" : "",
      "name" : "aeiou",
      "id" : ""
    },
    "application" : {
      "name" : "aeiou",
      "id" : ""
    },
    "type" : "aeiou"
  }
};
  if(Object.keys(examples).length > 0) {
    res.setHeader('Content-Type', 'application/json');
    res.end(JSON.stringify(examples[Object.keys(examples)[0]] || {}, null, 2));
  }
  else {
    res.end();
  }
  
}

exports.getAccountPhoneNumber = function(args, res, next) {
  /**
   * parameters expected in the args:
  * account_id (Integer)
  * number_id (Integer)
  **/
    var examples = {};
  examples['application/json'] = {
  "call_notifications" : {
    "emails" : [ "aeiou" ],
    "sms" : "aeiou"
  },
  "route" : {
    "name" : "aeiou",
    "id" : ""
  },
  "block_anonymous" : true,
  "caller_id" : {
    "name" : "aeiou",
    "type" : "aeiou"
  },
  "name" : "aeiou",
  "phone_number" : "aeiou",
  "id" : "",
  "block_incoming" : true,
  "sms_forwarding" : {
    "extension" : {
      "extension" : "",
      "name" : "aeiou",
      "id" : ""
    },
    "application" : {
      "name" : "aeiou",
      "id" : ""
    },
    "type" : "aeiou"
  }
};
  if(Object.keys(examples).length > 0) {
    res.setHeader('Content-Type', 'application/json');
    res.end(JSON.stringify(examples[Object.keys(examples)[0]] || {}, null, 2));
  }
  else {
    res.end();
  }
  
}

exports.listAccountPhoneNumbers = function(args, res, next) {
  /**
   * parameters expected in the args:
  * account_id (Integer)
  * filters[id] (List)
  * filters[name] (List)
  * filters[phone_number] (List)
  * sort[id] (String)
  * sort[name] (String)
  * sort[phone_number] (String)
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
    "phone_number" : "aeiou",
    "id" : "aeiou"
  },
  "sort" : {
    "name" : "aeiou",
    "phone_number" : "aeiou",
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

exports.replaceAccountPhoneNumber = function(args, res, next) {
  /**
   * parameters expected in the args:
  * account_id (Integer)
  * number_id (Integer)
  * data (ReplacePhoneNumberParams)
  **/
    var examples = {};
  examples['application/json'] = {
  "call_notifications" : {
    "emails" : [ "aeiou" ],
    "sms" : "aeiou"
  },
  "route" : {
    "name" : "aeiou",
    "id" : ""
  },
  "block_anonymous" : true,
  "caller_id" : {
    "name" : "aeiou",
    "type" : "aeiou"
  },
  "name" : "aeiou",
  "phone_number" : "aeiou",
  "id" : "",
  "block_incoming" : true,
  "sms_forwarding" : {
    "extension" : {
      "extension" : "",
      "name" : "aeiou",
      "id" : ""
    },
    "application" : {
      "name" : "aeiou",
      "id" : ""
    },
    "type" : "aeiou"
  }
};
  if(Object.keys(examples).length > 0) {
    res.setHeader('Content-Type', 'application/json');
    res.end(JSON.stringify(examples[Object.keys(examples)[0]] || {}, null, 2));
  }
  else {
    res.end();
  }
  
}

