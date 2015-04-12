'use strict';

// Declare app level module which depends on views, and components
angular.module('tasksystem', [
    'ngRoute',
    'jsTree.directive',
    'tasksystem.login',
    'tasksystem.main',
    'tasksystem.taskmanager'
]).
    config(['$routeProvider', function($routeProvider) {
        $routeProvider.otherwise({redirectTo: '/login'});
    }]);
