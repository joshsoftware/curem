var curemApp = angular.module('curemApp', ['ngRoute','curemControllers']);

curemApp.config(['$routeProvider',
  function($routeProvider) {
      $routeProvider.
	  when('/contacts', {
	      templateUrl: 'partials/contacts-list.html',
	      controller: 'contactsController'
	  }).
	  when('/contacts/:slug', {
	      templateUrl: 'partials/contact-detail.html',
	      controller: 'contactsDetailController'
	  }).
	  otherwise ({
	      redirectTo: '/contacts'
	  });
}]);
