var curemApp = angular.module('curemApp', ['ngResource']);

curemApp.factory('contactFactory', ['$resource', function($resource) {
      return $resource(
	  'http://localhost:3000/contacts/:slug', 
	  {slug: '@slug'},
	  {
	      'update': { method:'PUT' }
	  }
      );
}]);


curemApp.controller('contactsController', ['$scope', 'contactFactory', function ($scope, contactFactory) {

    $scope.contacts = contactFactory.query();

    console.log($scope.contacts)
}]);
