'use strict';

var url = require('url');


var Media = require('./MediaService');


module.exports.getAccountMedia = function getAccountMedia (req, res, next) {
  Media.getAccountMedia(req.swagger.params, res, next);
};

module.exports.listAccountMedia = function listAccountMedia (req, res, next) {
  Media.listAccountMedia(req.swagger.params, res, next);
};
