# Starbucks Payment API [Dhanashree Gaonkar]

## Week 1
### Project Research
- [x] 1. Understanding the requirements, architecture of the project through the article  https://www.infoq.com/articles/jepsen (Links to an external site.) 
- [x] 2. Research on the nosql db(allowing AP) to be used. --Update( Using Mongodb for the project) 
- [x] 3. Understanding partition tolerance and its pitfalls. 

## Week 2
### Mongodb cluster setup with replicaset configuration
- [x] 1. Research on the nosql db(allowing AP) to be used. 
- [x] 2. Research on how to use redis cache along with Mongodb 
- [x] 3. Initial setup to run goapi. 
- [x] 4. Function to connect to redis cache.  
- [x] 5. Method for connecting mongodb using mgo. 
- [x] 6. Setup an AMI for Ec2 having mongodb installed. Boot up 5 different Ec2's with the AMI.
- [x] 7. Set a common VPC for all the Ec2 instances.
- [x] 8. mongodb configuration for replicationset. 

## Week 3
### CRUD operations and error handling
- [x] 1. Implementing basic CRUD operations for Mongodb in GO. 
- [x] 2. Testing replication amongst the replicasets. 
- [x] 3. Creating a partition and testing AP property for the cluster. 
- [x] 4. Testing leader election after bringing down primary. 
- [x] 5. Testing high availability of mongodb by partitioning secondary from only primary. 

## Week 4
### Payments API specific features created
- [x] 1. Implementing Payment API specific features. 
- [x] 2. Update the order status for "pending" orders.
- [x] 3. Save Payment details to the database.
- [x] 4. Validate user session with redis cache. 

## Week 5
### Module integration and testing
- [x] 1. Wrote unit testcases for unit testing.
- [x] 2. Did module integration testing for our microservices.
- [x] 3. Created a Postman collection for Go Api testing.
