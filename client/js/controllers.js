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

curemControllers.controller('contactDetailController', ['$scope','$routeParams','contactFactory', function ($scope, $routeParams, contactFactory) {
    $scope.slug = $routeParams.slug;

    contactFactory.get({slug:$routeParams.slug})
    .$promise.then(function(contact) {
	$scope.contact = contact;
    });

}]);

curemControllers.controller('newContactController',['$scope','$location','contactFactory', function($scope, $location, contactFactory) {
    $scope.createNewContact = function() {
	contactFactory.save($scope.contact);
	$location.path('/');
    };
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

    console.log($scope.leads);
}]);

curemControllers.controller('leadDetailController', ['$scope', '$routeParams', 'leadFactory', function($scope, $routeParams, leadFactory) {
    $scope.id = $routeParams.id;

    leadFactory.get({id: $routeParams.id})
    .$promise.then(function(lead) {
	$scope.lead = lead;
    });
    console.log($scope.lead);
}]);

curemControllers.controller('newLeadController', ['$scope', 'leadFactory', function($scope, leadFactory) {
    $scope.createNewLead = function() {
	leadFactory.save($scope.lead);
	$location.path('/leads');
    };
}]);
