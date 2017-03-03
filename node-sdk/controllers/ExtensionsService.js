'use strict';

exports.createAccountExtension = function(args, res, next) {
  /**
   * parameters expected in the args:
  * account_id (Integer)
  * data (CreateExtensionParams)
  **/
    var examples = {};
  examples['application/json'] = {
  "device_membership" : {
    "line" : "",
    "device" : {
      "name" : "aeiou",
      "id" : ""
    }
  },
  "local_area_code" : "aeiou",
  "call_notifications" : "",
  "extension" : "",
  "voicemail" : {
    "password" : "aeiou",
    "attachments" : "aeiou",
    "transcription" : "aeiou",
    "greeting" : {
      "standard" : "",
      "enable_leave_message_prompt" : true,
      "alternate" : "",
      "type" : "aeiou"
    },
    "enabled" : true,
    "notifications" : {
      "emails" : [ "aeiou" ],
      "sms" : "aeiou"
    }
  },
  "timezone" : "aeiou",
  "include_in_directory" : true,
  "full_name" : "aeiou",
  "enable_call_waiting" : true,
  "enable_outbound_calls" : true,
  "route" : {
    "name" : "aeiou",
    "id" : ""
  },
  "name_greeting" : {
    "name" : "aeiou",
    "id" : ""
  },
  "usage_type" : "aeiou",
  "caller_id" : "aeiou",
  "name" : "aeiou",
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

exports.getAccountExtension = function(args, res, next) {
  /**
   * parameters expected in the args:
  * account_id (Integer)
  * extension_id (Integer)
  **/
    var examples = {};
  examples['application/json'] = {
  "device_membership" : {
    "line" : "",
    "device" : {
      "name" : "aeiou",
      "id" : ""
    }
  },
  "local_area_code" : "aeiou",
  "call_notifications" : "",
  "extension" : "",
  "voicemail" : {
    "password" : "aeiou",
    "attachments" : "aeiou",
    "transcription" : "aeiou",
    "greeting" : {
      "standard" : "",
      "enable_leave_message_prompt" : true,
      "alternate" : "",
      "type" : "aeiou"
    },
    "enabled" : true,
    "notifications" : {
      "emails" : [ "aeiou" ],
      "sms" : "aeiou"
    }
  },
  "timezone" : "aeiou",
  "include_in_directory" : true,
  "full_name" : "aeiou",
  "enable_call_waiting" : true,
  "enable_outbound_calls" : true,
  "route" : {
    "name" : "aeiou",
    "id" : ""
  },
  "name_greeting" : {
    "name" : "aeiou",
    "id" : ""
  },
  "usage_type" : "aeiou",
  "caller_id" : "aeiou",
  "name" : "aeiou",
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

exports.listAccountExtensions = function(args, res, next) {
  /**
   * parameters expected in the args:
  * account_id (Integer)
  * filters[id] (List)
  * filters[extension] (List)
  * filters[name] (List)
  * sort[id] (String)
  * sort[extension] (String)
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
    "extension" : "aeiou",
    "name" : "aeiou",
    "id" : "aeiou"
  },
  "sort" : {
    "extension" : "aeiou",
    "name" : "aeiou",
    "id" : "aeiou"
  },
  "items" : [ {
    "device_membership" : {
      "line" : "",
      "device" : {
        "name" : "aeiou",
        "id" : ""
      }
    },
    "local_area_code" : "aeiou",
    "call_notifications" : "",
    "extension" : "",
    "voicemail" : {
      "password" : "aeiou",
      "attachments" : "aeiou",
      "transcription" : "aeiou",
      "greeting" : {
        "standard" : "",
        "enable_leave_message_prompt" : true,
        "alternate" : "",
        "type" : "aeiou"
      },
      "enabled" : true,
      "notifications" : {
        "emails" : [ "aeiou" ],
        "sms" : "aeiou"
      }
    },
    "timezone" : "aeiou",
    "include_in_directory" : true,
    "full_name" : "aeiou",
    "enable_call_waiting" : true,
    "enable_outbound_calls" : true,
    "route" : {
      "name" : "aeiou",
      "id" : ""
    },
    "name_greeting" : {
      "name" : "aeiou",
      "id" : ""
    },
    "usage_type" : "aeiou",
    "caller_id" : "aeiou",
    "name" : "aeiou",
    "id" : ""
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

exports.replaceAccountExtension = function(args, res, next) {
  /**
   * parameters expected in the args:
  * account_id (Integer)
  * extension_id (Integer)
  * data (ReplaceExtensionParams)
  **/
    var examples = {};
  examples['application/json'] = {
  "device_membership" : {
    "line" : "",
    "device" : {
      "name" : "aeiou",
      "id" : ""
    }
  },
  "local_area_code" : "aeiou",
  "call_notifications" : "",
  "extension" : "",
  "voicemail" : {
    "password" : "aeiou",
    "attachments" : "aeiou",
    "transcription" : "aeiou",
    "greeting" : {
      "standard" : "",
      "enable_leave_message_prompt" : true,
      "alternate" : "",
      "type" : "aeiou"
    },
    "enabled" : true,
    "notifications" : {
      "emails" : [ "aeiou" ],
      "sms" : "aeiou"
    }
  },
  "timezone" : "aeiou",
  "include_in_directory" : true,
  "full_name" : "aeiou",
  "enable_call_waiting" : true,
  "enable_outbound_calls" : true,
  "route" : {
    "name" : "aeiou",
    "id" : ""
  },
  "name_greeting" : {
    "name" : "aeiou",
    "id" : ""
  },
  "usage_type" : "aeiou",
  "caller_id" : "aeiou",
  "name" : "aeiou",
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

