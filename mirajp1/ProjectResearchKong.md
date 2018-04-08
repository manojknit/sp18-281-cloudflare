## Project Research - Kong API Gateway

## REST API?
API is Application Programming Interface which acts as an endpoint for resources in the web
Representational State Transfer (REST) defines a set of constraints and properties based on HTTP for API end points.
REST-compliant web services allow the requesting systems to access and manipulate textual representations of web resources by using a uniform and predefined set of stateless operations.
By using a stateless protocol and standard operations, REST systems aim for fast performance, reliability, and the ability to grow, by re-using components that can be managed and updated without affecting the system as a whole, even while it is running.

## API Gateway
In a microservices architecture, a client might interact with more than one front-end service. Given this fact, how does a client know what endpoints to call? What happens when new services are introduced, or existing services are refactored? An API gateway can help to address these challenges.

An API gateway sits between clients and services. It acts as a reverse proxy, routing requests from clients to services. It may also perform various cross-cutting tasks such as authentication, SSL termination, and rate limiting. 

The API gateway handles requests in one of two ways. Some requests are simply proxied/routed to the appropriate service. It handles other requests by fanning out to multiple services.

![API Gateways](https://docs.microsoft.com/en-us/azure/architecture/microservices/images/gateway.png)

## Advantages of API Gateway
* Use it as a reverse proxy to route requests to one or more backend services. This provides a single endpoint for clients, and decouples clients from services.
* Use the gateway to aggregate multiple individual requests into a single request.
* Use the gateway to offload functionality from individual services to the gateway like Authentication, logging, caching, monitoring, firewall etc.

## Kong API Gateway
Kong is a scalable API Gateway. It runs in front of RESTful APIs and can be extended through Plugins.
* Scalable: Kong easily scales horizontally by simply adding more machines, meaning your platform can handle virtually any load while keeping latency low.

* Modular: Kong can be extended by adding new plugins, which are easily configured through a RESTful Admin API like file logs,authentication etc.

* Runs on any infrastructure: Kong runs anywhere. You can deploy Kong in the cloud or on-premise environments, including single or multi-datacenter setups and for public, private or invite-only APIs.

![KONG](https://getkong.org/assets/images/docs/kong-architecture.jpg)

Client requests to the kong servers(clusters or single node) and is routed to appropriate backends as configured by the admin. It will execute installed plugins. Kong effectively becomes the entry point for every API request.
![workflow](https://getkong.org/assets/images/docs/kong-simple.png)

### Kong Deployment
* Kong API Gateway is backed by the CassandraDB or Postgres(we will use cassandra for this project)
* Kong itself can be formed into clusters and the cassandra which backs it can be formed into clusters.
* kong can be deployed to docker cloud as in class Labs as there is readily available containers for cassandra and kong
    - However, in my personal experience, there were difficulties in setting up kong in the docker cloud due to IP mismatch errors.
* Other way to go for deployment is through AWS EC2 and form a cassandra cluster there (as in personal project)
    - open up ports used by cassandra and kong in security group(8000,8001,8443,8444 etc).
    - make the ec2 publicly availble which the front end can access(from heroku deployed front end)


### References
1. [REST API](https://en.wikipedia.org/wiki/Representational_state_transfer)
2. [NoSQL Wiki](https://en.wikipedia.org/wiki/NoSQL)
3. [CassandraDB](https://hackernoon.com/using-apache-cassandra-a-few-things-before-you-start-ac599926e4b8)
4. [Kong doc](https://getkong.org/about/)
5. [API Gateways](https://docs.microsoft.com/en-us/azure/architecture/microservices/gateway)
