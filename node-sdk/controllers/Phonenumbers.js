'use strict';

var url = require('url');


var Phonenumbers = require('./PhonenumbersService');


module.exports.createAccountPhoneNumber = function createAccountPhoneNumber (req, res, next) {
  Phonenumbers.createAccountPhoneNumber(req.swagger.params, res, next);
};

module.exports.getAccountPhoneNumber = function getAccountPhoneNumber (req, res, next) {
  Phonenumbers.getAccountPhoneNumber(req.swagger.params, res, next);
};

module.exports.listAccountPhoneNumbers = function listAccountPhoneNumbers (req, res, next) {
  Phonenumbers.listAccountPhoneNumbers(req.swagger.params, res, next);
};

module.exports.replaceAccountPhoneNumber = function replaceAccountPhoneNumber (req, res, next) {
  Phonenumbers.replaceAccountPhoneNumber(req.swagger.params, res, next);
};
