'use strict';

var url = require('url');


var Queues = require('./QueuesService');


module.exports.createAccountQueue = function createAccountQueue (req, res, next) {
  Queues.createAccountQueue(req.swagger.params, res, next);
};

module.exports.deleteAccountQueue = function deleteAccountQueue (req, res, next) {
  Queues.deleteAccountQueue(req.swagger.params, res, next);
};

module.exports.getAccountQueue = function getAccountQueue (req, res, next) {
  Queues.getAccountQueue(req.swagger.params, res, next);
};

module.exports.listAccountQueues = function listAccountQueues (req, res, next) {
  Queues.listAccountQueues(req.swagger.params, res, next);
};

module.exports.replaceAccountQueue = function replaceAccountQueue (req, res, next) {
  Queues.replaceAccountQueue(req.swagger.params, res, next);
};
