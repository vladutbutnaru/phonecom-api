'use strict';

var url = require('url');


var Schedules = require('./SchedulesService');


module.exports.getAccountSchedule = function getAccountSchedule (req, res, next) {
  Schedules.getAccountSchedule(req.swagger.params, res, next);
};

module.exports.listAccountSchedules = function listAccountSchedules (req, res, next) {
  Schedules.listAccountSchedules(req.swagger.params, res, next);
};
