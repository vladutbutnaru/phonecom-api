'use strict';

var url = require('url');


var Extensions = require('./ExtensionsService');


module.exports.createAccountExtension = function createAccountExtension (req, res, next) {
  Extensions.createAccountExtension(req.swagger.params, res, next);
};

module.exports.getAccountExtension = function getAccountExtension (req, res, next) {
  Extensions.getAccountExtension(req.swagger.params, res, next);
};

module.exports.listAccountExtensions = function listAccountExtensions (req, res, next) {
  Extensions.listAccountExtensions(req.swagger.params, res, next);
};

module.exports.replaceAccountExtension = function replaceAccountExtension (req, res, next) {
  Extensions.replaceAccountExtension(req.swagger.params, res, next);
};
