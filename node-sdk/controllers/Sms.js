'use strict';

var url = require('url');


var Sms = require('./SmsService');


module.exports.createAccountSms = function createAccountSms (req, res, next) {
  Sms.createAccountSms(req.swagger.params, res, next);
};

module.exports.getAccountSms = function getAccountSms (req, res, next) {
  Sms.getAccountSms(req.swagger.params, res, next);
};

module.exports.listAccountSms = function listAccountSms (req, res, next) {
  Sms.listAccountSms(req.swagger.params, res, next);
};
