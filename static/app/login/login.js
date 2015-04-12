'use strict';

angular.module('tasksystem.login', ['ngRoute'])

    .config(['$routeProvider', function ($routeProvider) {
        $routeProvider.when('/login', {
            templateUrl: 'app/login/login.html',
            controller: 'LoginCtrl'
        });
    }])

    .controller('LoginCtrl', function ($scope) {
        $scope.getValideCode = function () {
            var codeArray = ['A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J',
                'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V',
                'W', 'X', 'Y', 'Z'];
            var code = '';
            for (var i = 0; i < 4; i++) {
                code = code + codeArray[Math.round(Math.random() * 25)];
            }
            $scope.valideCode = code;
        };
        $scope.getValideCode();

    });