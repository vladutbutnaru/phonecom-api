'use strict';

var url = require('url');


var Numberregions = require('./NumberregionsService');


module.exports.listAvailablePhoneNumberRegions = function listAvailablePhoneNumberRegions (req, res, next) {
  Numberregions.listAvailablePhoneNumberRegions(req.swagger.params, res, next);
};
