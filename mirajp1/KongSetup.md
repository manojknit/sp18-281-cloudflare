## Kong - Setup


## AWS installation of Kong Api gateway

+ Create an EC2 instance with default configuration (as done in cassandra installation)
+ SSH into the container using the key file selected and it's public ip address.
+ Download kong 0.10.3 version using following command.
```
    wget kong.bintray.com/kong-community-edition-rpm/centos/7/kong-community-edition-0.10.3.el7.noarch.rpm
``` 

+ Install kong from the downloaded rpm file using following command.
```
    sudo yum install kong-community-edition-0.10.3.el7.noarch.rpm --nogpgcheck
```

+ update the kong conf using following commands
```
    sudo mv /etc/kong/kong.default.conf /etc/kong/kong.conf
```
Edit the file with cassandra_contact_points pointing to your cluster nodes and database as cassandra.

```
    database = cassandra             # Determines which of PostgreSQL or Cassandra
    cassandra_contact_points =cassandra1,cassandra2,cassandra3,cassandra4,cassandra5   # A comma-separated list of contact
    cassandra_port = 9042           
    cassandra_keyspace = kong       # The keyspace to use in your cluster.
    cassandra_timeout = 5000        # Defines the timeout (in ms), for reading
    cassandra_ssl = off             # Toggles client-to-node TLS connections
    cassandra_ssl_verify = off      # Toggles server certificate verification if
    cassandra_consistency = ONE     # Consistency setting to use when reading/
    cassandra_lb_policy = RoundRobin  # Load balancing policy to use when
    cassandra_repl_strategy = SimpleStrategy  # When migrating for the first time,
    cassandra_repl_factor = 3       # When migrating for the first time, Kong
```

+ Run the kong migrations and start the kong using following commands.

```
    kong migrations up
    kong start
    kong health
```

[Kong Health](https://github.com/nguyensjsu/team281-cloudflare/blob/master/mirajp1/images/kong_health)
### References
1. [REST API](https://en.wikipedia.org/wiki/Representational_state_transfer)
2. [NoSQL Wiki](https://en.wikipedia.org/wiki/NoSQL)
3. [CassandraDB](https://hackernoon.com/using-apache-cassandra-a-few-things-before-you-start-ac599926e4b8)
4. [Kong doc](https://getkong.org/about/)
5. [API Gateways](https://docs.microsoft.com/en-us/azure/architecture/microservices/gateway)
