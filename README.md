# curem

[![Build Status](https://drone.io/github.com/joshsoftware/curem/status.png)](https://drone.io/github.com/joshsoftware/curem/latest)

**CU**&#8203;stomer **RE**&#8203;lationship **M**&#8203;anagement made easy. Curem is a light weight CRM written in Go, with JSON support and backed with MongoDB. It can be consumed by any client which can understand and work with JSON.

Basically a lead has 2 sides to it: external and internal.

* **External** side corresponds to the information about clients and customers. We can try to avoid duplication but it’s not necessary. For example, if we get multiple leads from same person / customer, even if there is duplication, I should be able to figure it out. Normalization is not necessary as it adds to complexity. 
* **Internal** side corresponds to the lead source, who’s handling it, what’s quoted, team size, start date and various comments. 

## TODO

- [ ] Configure using an external file
- [ ] Add Project Title field in lead type
- [ ] Add text indexes, full text search for multiple fields (contact person, contact company, lead title)

## Contributing

1. Fork it [https://github.com/joshsoftware/curem/fork](https://github.com/joshsoftware/curem/fork)
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request