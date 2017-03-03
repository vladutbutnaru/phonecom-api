'use strict';

var url = require('url');


var Calllogs = require('./CalllogsService');


module.exports.getAccountCallLog = function getAccountCallLog (req, res, next) {
  Calllogs.getAccountCallLog(req.swagger.params, res, next);
};

module.exports.listAccountCallLogs = function listAccountCallLogs (req, res, next) {
  Calllogs.listAccountCallLogs(req.swagger.params, res, next);
};
