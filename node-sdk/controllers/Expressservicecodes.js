'use strict';

var url = require('url');


var Expressservicecodes = require('./ExpressservicecodesService');


module.exports.getAccountExpressSrvCode = function getAccountExpressSrvCode (req, res, next) {
  Expressservicecodes.getAccountExpressSrvCode(req.swagger.params, res, next);
};

module.exports.listAccountExpressSrvCodes = function listAccountExpressSrvCodes (req, res, next) {
  Expressservicecodes.listAccountExpressSrvCodes(req.swagger.params, res, next);
};
