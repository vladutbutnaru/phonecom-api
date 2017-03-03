'use strict';

exports.createRoute = function(args, res, next) {
  /**
   * parameters expected in the args:
  * account_id (Integer)
  * data (CreateRouteParams)
  **/
    var examples = {};
  examples['application/json'] = {
  "extension" : {
    "extension" : "",
    "name" : "aeiou",
    "id" : ""
  },
  "name" : "aeiou",
  "rules" : [ {
    "filter" : {
      "schedule" : {
        "name" : "aeiou",
        "id" : ""
      },
      "contact" : {
        "prefix" : "aeiou",
        "nickname" : "aeiou",
        "last_name" : "aeiou",
        "company" : "aeiou",
        "id" : "",
        "middle_name" : "aeiou",
        "suffix" : "aeiou",
        "first_name" : "aeiou"
      },
      "type" : "aeiou",
      "group" : {
        "name" : "aeiou",
        "id" : ""
      }
    },
    "actions" : [ {
      "duration" : "",
      "extension" : "",
      "hold_music" : {
        "name" : "aeiou",
        "id" : ""
      },
      "greeting" : "",
      "action" : "aeiou",
      "menu" : {
        "name" : "aeiou",
        "id" : ""
      },
      "trunk" : {
        "name" : "aeiou",
        "id" : ""
      },
      "items" : [ {
        "number" : "aeiou",
        "extension" : "",
        "screening" : true,
        "caller_id" : "aeiou",
        "voice_tag" : "aeiou",
        "distinctive_ring" : "aeiou",
        "type" : "aeiou"
      } ],
      "timeout" : "",
      "queue" : {
        "name" : "aeiou",
        "id" : ""
      }
    } ]
  } ],
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

exports.deleteAccountRoute = function(args, res, next) {
  /**
   * parameters expected in the args:
  * account_id (Integer)
  * route_id (Integer)
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

exports.getAccountRoute = function(args, res, next) {
  /**
   * parameters expected in the args:
  * account_id (Integer)
  * route_id (Integer)
  **/
    var examples = {};
  examples['application/json'] = {
  "extension" : {
    "extension" : "",
    "name" : "aeiou",
    "id" : ""
  },
  "name" : "aeiou",
  "rules" : [ {
    "filter" : {
      "schedule" : {
        "name" : "aeiou",
        "id" : ""
      },
      "contact" : {
        "prefix" : "aeiou",
        "nickname" : "aeiou",
        "last_name" : "aeiou",
        "company" : "aeiou",
        "id" : "",
        "middle_name" : "aeiou",
        "suffix" : "aeiou",
        "first_name" : "aeiou"
      },
      "type" : "aeiou",
      "group" : {
        "name" : "aeiou",
        "id" : ""
      }
    },
    "actions" : [ {
      "duration" : "",
      "extension" : "",
      "hold_music" : {
        "name" : "aeiou",
        "id" : ""
      },
      "greeting" : "",
      "action" : "aeiou",
      "menu" : {
        "name" : "aeiou",
        "id" : ""
      },
      "trunk" : {
        "name" : "aeiou",
        "id" : ""
      },
      "items" : [ {
        "number" : "aeiou",
        "extension" : "",
        "screening" : true,
        "caller_id" : "aeiou",
        "voice_tag" : "aeiou",
        "distinctive_ring" : "aeiou",
        "type" : "aeiou"
      } ],
      "timeout" : "",
      "queue" : {
        "name" : "aeiou",
        "id" : ""
      }
    } ]
  } ],
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

exports.listAccountRoutes = function(args, res, next) {
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

exports.replaceAccountRoute = function(args, res, next) {
  /**
   * parameters expected in the args:
  * account_id (Integer)
  * route_id (Integer)
  * data (CreateRouteParams)
  **/
    var examples = {};
  examples['application/json'] = {
  "extension" : {
    "extension" : "",
    "name" : "aeiou",
    "id" : ""
  },
  "name" : "aeiou",
  "rules" : [ {
    "filter" : {
      "schedule" : {
        "name" : "aeiou",
        "id" : ""
      },
      "contact" : {
        "prefix" : "aeiou",
        "nickname" : "aeiou",
        "last_name" : "aeiou",
        "company" : "aeiou",
        "id" : "",
        "middle_name" : "aeiou",
        "suffix" : "aeiou",
        "first_name" : "aeiou"
      },
      "type" : "aeiou",
      "group" : {
        "name" : "aeiou",
        "id" : ""
      }
    },
    "actions" : [ {
      "duration" : "",
      "extension" : "",
      "hold_music" : {
        "name" : "aeiou",
        "id" : ""
      },
      "greeting" : "",
      "action" : "aeiou",
      "menu" : {
        "name" : "aeiou",
        "id" : ""
      },
      "trunk" : {
        "name" : "aeiou",
        "id" : ""
      },
      "items" : [ {
        "number" : "aeiou",
        "extension" : "",
        "screening" : true,
        "caller_id" : "aeiou",
        "voice_tag" : "aeiou",
        "distinctive_ring" : "aeiou",
        "type" : "aeiou"
      } ],
      "timeout" : "",
      "queue" : {
        "name" : "aeiou",
        "id" : ""
      }
    } ]
  } ],
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

