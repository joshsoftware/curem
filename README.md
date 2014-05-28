# curem

**CU**&#8203;stomer **RE**&#8203;lationship **M**&#8203;anagement made easy. Curem is a light weight Go CRM with jsonapi support backed with MongoDB. It can be consumed by any client which can understand and work with JSONapi. (jsonapi.org)

Basically a lead has 2 sides to it: external and internal.

* **External** side corresponds to the information about clients and customers. We can try to avoid duplication but it’s not necessary. For example, if we get multiple leads from same person / customer, even if there is duplication, I should be able to figure it out. Normalization is not necessary as it adds to complexity. 
* **Internal** side corresponds to the lead source, who’s handling it, what’s quoted, team size, start date and various comments. 

## Scope of Work

This is how we are building this application. 

### Phase 1
 Build a Go application that can manage leads. We write test cases for CRUD operations for lead management. Data should be persisted in MongoDB. NO APIs and NO WEB front.  

### Phase 2
 We write a jsonapi layer on top of this and expose the resources properly. If we can make this a Go package, it’s great, otherwise, it serves as an example. Lots of test cases that will fire Json api and get back results.

### Phase 3
 We write a the client side web front - responsive, cached using Angular, Ember etc. and showcase this. 
 
 
## Contributing

1. Fork it [https://github.com/joshsoftware/curem/fork](https://github.com/joshsoftware/curem/fork)
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request
