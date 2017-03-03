'use strict';

var url = require('url');


var Callerids = require('./CalleridsService');


module.exports.getCallerIds = function getCallerIds (req, res, next) {
  Callerids.getCallerIds(req.swagger.params, res, next);
};
