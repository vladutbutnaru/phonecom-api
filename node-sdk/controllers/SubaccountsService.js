'use strict';

exports.createAccountSubaccount = function(args, res, next) {
  /**
   * parameters expected in the args:
  * account_id (Integer)
  * data (CreateSubaccountParams)
  **/
    var examples = {};
  examples['application/json'] = {
  "password" : "aeiou",
  "billing_contact" : "",
  "contact" : {
    "address" : {
      "line_1" : "aeiou",
      "line_2" : "aeiou",
      "country" : "aeiou",
      "province" : "aeiou",
      "city" : "aeiou",
      "postal_code" : "aeiou"
    },
    "phone" : "aeiou",
    "name" : "aeiou",
    "company" : "aeiou",
    "primary_email" : "aeiou",
    "fax" : "aeiou",
    "alternate_email" : "aeiou"
  },
  "name" : "aeiou",
  "master_account" : {
    "name" : "aeiou",
    "id" : ""
  },
  "id" : "",
  "username" : "aeiou"
};
  if(Object.keys(examples).length > 0) {
    res.setHeader('Content-Type', 'application/json');
    res.end(JSON.stringify(examples[Object.keys(examples)[0]] || {}, null, 2));
  }
  else {
    res.end();
  }
  
}

exports.listAccountSubaccounts = function(args, res, next) {
  /**
   * parameters expected in the args:
  * account_id (Integer)
  * filters[id] (List)
  * sort[id] (String)
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
    "id" : "aeiou"
  },
  "sort" : {
    "id" : "aeiou"
  },
  "items" : [ {
    "password" : "aeiou",
    "billing_contact" : "",
    "contact" : {
      "address" : {
        "line_1" : "aeiou",
        "line_2" : "aeiou",
        "country" : "aeiou",
        "province" : "aeiou",
        "city" : "aeiou",
        "postal_code" : "aeiou"
      },
      "phone" : "aeiou",
      "name" : "aeiou",
      "company" : "aeiou",
      "primary_email" : "aeiou",
      "fax" : "aeiou",
      "alternate_email" : "aeiou"
    },
    "name" : "aeiou",
    "master_account" : {
      "name" : "aeiou",
      "id" : ""
    },
    "id" : "",
    "username" : "aeiou"
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

