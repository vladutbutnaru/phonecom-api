'use strict';

var url = require('url');


var Availablenumbers = require('./AvailablenumbersService');


module.exports.listAvailablePhoneNumbers = function listAvailablePhoneNumbers (req, res, next) {
  Availablenumbers.listAvailablePhoneNumbers(req.swagger.params, res, next);
};
