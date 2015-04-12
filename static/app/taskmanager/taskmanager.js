'use strict';

angular.module('tasksystem.taskmanager', ['ngRoute'])

    .config(['$routeProvider', function ($routeProvider) {
        $routeProvider.when('/taskmanager', {
            templateUrl: 'app/taskmanager/taskmanager.html',
            controller: 'TaskManagerCtrl'
        });
    }])

    .controller('TaskManagerCtrl', [function () {

    }]);