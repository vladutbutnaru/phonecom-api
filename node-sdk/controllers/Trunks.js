'use strict';

var url = require('url');


var Trunks = require('./TrunksService');


module.exports.createAccountTrunk = function createAccountTrunk (req, res, next) {
  Trunks.createAccountTrunk(req.swagger.params, res, next);
};

module.exports.deleteAccountTrunk = function deleteAccountTrunk (req, res, next) {
  Trunks.deleteAccountTrunk(req.swagger.params, res, next);
};

module.exports.getAccountTrunk = function getAccountTrunk (req, res, next) {
  Trunks.getAccountTrunk(req.swagger.params, res, next);
};

module.exports.listAccountTrunks = function listAccountTrunks (req, res, next) {
  Trunks.listAccountTrunks(req.swagger.params, res, next);
};

module.exports.replaceAccountTrunk = function replaceAccountTrunk (req, res, next) {
  Trunks.replaceAccountTrunk(req.swagger.params, res, next);
};
