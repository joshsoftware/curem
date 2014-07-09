var curemControllers = angular.module('curemControllers', ['ngResource', 'ngRoute']);

curemControllers.factory('contactFactory', ['$resource', function($resource) {
      return $resource(
	  'http://localhost:3000/contacts/:slug', 
	  {slug: '@slug'},
	  {
	      'update': { method:'PATCH' }
	  }
      );
}]);


curemControllers.controller('contactsController', ['$scope', 'contactFactory', function ($scope, contactFactory) {

    $scope.contacts = contactFactory.query();
    $scope.orderProp = '-updatedAt';
    console.log($scope.contacts)
}]);

curemControllers.controller('contactsDetailController', ['$scope', '$routeParams', function ($scope, $routeParams) {
    $scope.slug = $routeParams.slug;
}]);

curemControllers.factory('leadFactory', ['$resource', function($resource) {
    return $resource(
	'http://localhost:3000/leads/:id',
	{id: '@id'},
	{
	    'update': {method: 'PATCH'}
	}
    );
}]);

curemControllers.controller('leadsController', ['$scope', 'leadFactory', function ($scope, leadFactory) {

    $scope.leads = leadFactory.query();

    console.log($scope.leads)
}]);
