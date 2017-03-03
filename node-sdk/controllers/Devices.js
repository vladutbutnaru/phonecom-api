'use strict';

var url = require('url');


var Devices = require('./DevicesService');


module.exports.createAccountDevice = function createAccountDevice (req, res, next) {
  Devices.createAccountDevice(req.swagger.params, res, next);
};

module.exports.getAccountDevice = function getAccountDevice (req, res, next) {
  Devices.getAccountDevice(req.swagger.params, res, next);
};

module.exports.listAccountDevices = function listAccountDevices (req, res, next) {
  Devices.listAccountDevices(req.swagger.params, res, next);
};

module.exports.replaceAccountDevice = function replaceAccountDevice (req, res, next) {
  Devices.replaceAccountDevice(req.swagger.params, res, next);
};
