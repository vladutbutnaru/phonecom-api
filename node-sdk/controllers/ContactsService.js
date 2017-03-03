'use strict';

exports.createAccountExtensionContact = function(args, res, next) {
  /**
   * parameters expected in the args:
  * account_id (Integer)
  * extension_id (Integer)
  * data (CreateContactParams)
  **/
    var examples = {};
  examples['application/json'] = {
  "addresses" : "",
  "prefix" : "aeiou",
  "last_name" : "aeiou",
  "created_at" : "",
  "middle_name" : "aeiou",
  "suffix" : "aeiou",
  "emails" : "",
  "phone_numbers" : [ {
    "number" : "aeiou",
    "normalized" : "aeiou",
    "type" : "aeiou"
  } ],
  "phonetic_first_name" : "aeiou",
  "phonetic_middle_name" : "aeiou",
  "updated_at" : "",
  "nickname" : "aeiou",
  "company" : "aeiou",
  "id" : "",
  "department" : "aeiou",
  "first_name" : "aeiou",
  "phonetic_last_name" : "aeiou",
  "job_title" : "aeiou",
  "group" : {
    "name" : "aeiou",
    "id" : ""
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

exports.deleteAccountExtensionContact = function(args, res, next) {
  /**
   * parameters expected in the args:
  * account_id (Integer)
  * extension_id (Integer)
  * contact_id (Integer)
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

exports.getAccountExtensionContact = function(args, res, next) {
  /**
   * parameters expected in the args:
  * account_id (Integer)
  * extension_id (Integer)
  * contact_id (Integer)
  **/
    var examples = {};
  examples['application/json'] = {
  "addresses" : "",
  "prefix" : "aeiou",
  "last_name" : "aeiou",
  "created_at" : "",
  "middle_name" : "aeiou",
  "suffix" : "aeiou",
  "emails" : "",
  "phone_numbers" : [ {
    "number" : "aeiou",
    "normalized" : "aeiou",
    "type" : "aeiou"
  } ],
  "phonetic_first_name" : "aeiou",
  "phonetic_middle_name" : "aeiou",
  "updated_at" : "",
  "nickname" : "aeiou",
  "company" : "aeiou",
  "id" : "",
  "department" : "aeiou",
  "first_name" : "aeiou",
  "phonetic_last_name" : "aeiou",
  "job_title" : "aeiou",
  "group" : {
    "name" : "aeiou",
    "id" : ""
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

exports.listAccountExtensionContacts = function(args, res, next) {
  /**
   * parameters expected in the args:
  * account_id (Integer)
  * extension_id (Integer)
  * filters[id] (List)
  * filters[group_id] (List)
  * filters[updated_at] (List)
  * sort[id] (String)
  * sort[updated_at] (String)
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
    "updated_at" : "aeiou",
    "group_id" : "aeiou",
    "id" : "aeiou"
  },
  "sort" : {
    "updated_at" : "aeiou",
    "id" : "aeiou"
  },
  "items" : [ {
    "addresses" : "",
    "prefix" : "aeiou",
    "last_name" : "aeiou",
    "created_at" : "",
    "middle_name" : "aeiou",
    "suffix" : "aeiou",
    "emails" : "",
    "phone_numbers" : [ {
      "number" : "aeiou",
      "normalized" : "aeiou",
      "type" : "aeiou"
    } ],
    "phonetic_first_name" : "aeiou",
    "phonetic_middle_name" : "aeiou",
    "updated_at" : "",
    "nickname" : "aeiou",
    "company" : "aeiou",
    "id" : "",
    "department" : "aeiou",
    "first_name" : "aeiou",
    "phonetic_last_name" : "aeiou",
    "job_title" : "aeiou",
    "group" : {
      "name" : "aeiou",
      "id" : ""
    }
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

exports.replaceAccountExtensionContact = function(args, res, next) {
  /**
   * parameters expected in the args:
  * account_id (Integer)
  * extension_id (Integer)
  * data (CreateContactParams)
  **/
    var examples = {};
  examples['application/json'] = {
  "addresses" : "",
  "prefix" : "aeiou",
  "last_name" : "aeiou",
  "created_at" : "",
  "middle_name" : "aeiou",
  "suffix" : "aeiou",
  "emails" : "",
  "phone_numbers" : [ {
    "number" : "aeiou",
    "normalized" : "aeiou",
    "type" : "aeiou"
  } ],
  "phonetic_first_name" : "aeiou",
  "phonetic_middle_name" : "aeiou",
  "updated_at" : "",
  "nickname" : "aeiou",
  "company" : "aeiou",
  "id" : "",
  "department" : "aeiou",
  "first_name" : "aeiou",
  "phonetic_last_name" : "aeiou",
  "job_title" : "aeiou",
  "group" : {
    "name" : "aeiou",
    "id" : ""
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

