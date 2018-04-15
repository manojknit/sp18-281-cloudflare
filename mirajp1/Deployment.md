## Kong - Deployment

## On Docker
* On first view, deployment of kong on docker cloud seems to be the best way to go forward as it is very easy to configure it on containers.
* However, after trying to deploy I faced lot of configuration problems with kong and cassandra because of IP address issues and the deployment failed as shown below.
* After trying various fixes, it didn't seem to work on docker-cloud.


## AWS Deployment on EC2 instances
* Spin of one ec2 t2.micro instance.
* SSH into the ec2 instance and install java 8 and cassandra into it.
* Stop the instance and create an AMI image out of it.
* start 4 another instances using the AMI image
* Better to allocate Elastic IPs to the instances so that they do not change.
* Change the cassandra.yaml file on each ec2 to setup a cluster. change the seeds IP addresses as required


### References
1. [REST API](https://en.wikipedia.org/wiki/Representational_state_transfer)
2. [NoSQL Wiki](https://en.wikipedia.org/wiki/NoSQL)
3. [CassandraDB](https://hackernoon.com/using-apache-cassandra-a-few-things-before-you-start-ac599926e4b8)
4. [Kong doc](https://getkong.org/about/)
5. [API Gateways](https://docs.microsoft.com/en-us/azure/architecture/microservices/gateway)
