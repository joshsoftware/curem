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
	      controller: 'contactDetailController'
	  }).
	  when('/new/contact', {
	      templateUrl: 'partials/new-contact.html',
	      controller: 'newContactController'
	  }).
	  otherwise ({
	      redirectTo: '/contacts'
	  });
}]);
