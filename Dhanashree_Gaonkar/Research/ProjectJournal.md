# **Objectives:**
In this project, you will be testing the Partition Tolerance of a NoSQL DB using the procedures described in the following article:  https://www.infoq.com/articles/jepsen (Links to an external site.)Links to an external site..  In addition, you will be developing an Data Service API in Go and deploying the API on AWS.

# **Requirements:**
* Select one of:  Redis or Riak  (other NoSQL DBs that can be configured in AP mode are also allowed).
* Set up the Five Nodes per the article's direction on AWS EC2 Instances.
* Set up the Test cases, run the Experiments and Record results.
* Develop the Go API and integrate the API with the Team SaaS App from the Hackathon Project.
* Project must be use in a GitHub Repo (assigned by TA) to maintain source code, documentation and design diagrams.
* Repo will be maintain in:  https://github.com/nguyensjsu (Links to an external site.)Links to an external site.
* Keep a Project Journal (as a markdown document) recording weekly progress, challenges, tests and test results.
* All documentation must be written in GitHub Markdown (including diagrams) and committed to GitHub
* https://help.github.com/articles/about-writing-and-formatting-on-github/Links to an external site.

#  Week1: 
* **Objective** ---> Initial Research on nosql db, CAP theorem, mongodb Vs Riak.
# Task1:

1. Understanding CAP Theorem.
1. Understanding the requirements, architecture of the project through the article  https://www.infoq.com/articles/jepsen (Links to an external site.) **Done**
2. Research on the nosql db(allowing AP) to be used. --Update( Using Mongodb for the project) **Done**
3. Understanding partition tolerance and its pitfalls. **Done**

# Week2: 
* **Objective** ---> Mongodb cluster setup with replicaset configuration.

# Task:
1. Setup an AMI for Ec2 having mongodb installed. Boot up 5 different Ec2's with the AMI. **Done**
2. Set a common VPC for all the Ec2 instances. **Done**
3. mongodb configuration for replicationset. **Done**

## Create Mongodb cluster using Ec2 instances:
1. Create Ec2 instance using the "Amazon Linux AMI 2017.09.1 (HVM), SSD Volume Type" image.
2. Create a security group with rule to allow all ip addresses with port 22, 27017.
3. Use common VPC for all instances which has public and private subnets
  * Enable auto-assign public ip
4. Use default storage config.
5. Add name tags for every instance like - Primary, Secondary_1, Secondary_2, Secondary_3,Secondary_4 and Arbiter.
 #### In case of primary node going down, there is a chance that 2 nodes get equal number of votes. To remove this imbalance of votes arbiter is added.
 
 ## Installation mongodb in the primary node:
 1. sudo yum update -y
 2. Create a /etc/yum.repos.d/mongodb-org-3.6.repo file 
    sudo touch /etc/yum.repos.d/mongodb-org-3.6.repo
    sudo vi /etc/yum.repos.d/mongodb-org-3.6.repo
 3.  copy below content in this file :
    [mongodb-org-3.6]
    name=MongoDB Repository
    baseurl=https://repo.mongodb.org/yum/amazon/2013.03/mongodb-org/3.6/x86_64/
    gpgcheck=1
    enabled=1
    gpgkey=https://www.mongodb.org/static/pgp/server-3.6.asc
4. Install mongodb: sudo yum install -y mongodb-org
5. Create default path for mongodb : 
   sudo mkdir -p /data/db
   sudo chown ec2-user /data/db
   ls -ld /data/db
6. Run mongodb in background:
   mongod &

* After mongodb installation is done in Primary.
    Create an image of the primary node and then create secondary nodes using this image. 
 
##  Authentication using a secretkey file.
 1. Create a shared secret-key file in primary node:
    sudo chmod 777 /etc/ssl/
    sudo openssl rand -base64 756 > /etc/ssl/mongodb-internal.key
    sudo chown mongod:mongod /etc/ssl/mongodb-internal.key
    sudo chmod 777 /etc/ssl/mongodb-internal.key

2. Copy this file in all the remaining nodes
Run this command on the AWS node where the mongodb-internal.key has been generated (key generated in previous step)
   scp -i <private_key> /etc/ssl/mongodb-internal.key ec2-user@<ip_address_of_AWS_instance_where_key_needs_to_be_copied>:/tmp/
Run this command on the AWS node where the mongodb-internal.key has been copied
    mv /tmp/mongodb-internal.key /etc/ssl/

## Replicaset configuration in mongod instances:

1. Edit /etc/mongod.conf file in all nodes:
 Refer to the external file.
2. Comment bindip to listen on all 
3. Add key filename and enable authorization
4. Add replica set name 
5. Edit /etc/hosts file in all nodes:
   127.0.0.1   localhost localhost.localdomain localhost4 localhost4.localdomain4
   10.0.2.99 p
   10.0.2.156 s1
   10.0.2.58 s2
   10.0.2.201 s3
   10.0.2.120 a
   ::1         localhost6 localhost6.localdomain6
6. kill all processes running on port 27017
   kill $(lsof -t -i:27017)
7. Run following command to start mongod process on all nodes:
   mongod --bind_ip_all --replSet rs0 &
8. Run following commands on primary node to initiate replicate set and add nodes: 
   rs.initiate()
   rs.add('primary ip:27017');
   rs.addArb('arbiter ip:27017');
9. open mongo shell using command mongo
   change to admin database - use admin
10. To start replication run command on slave: 
   rs.slaveOk()

# Week3: 
* **Objective** ---> CRUD operations and error handling.

# Task:
1. Research on how to use redis cache along with Mongodb **Done**
2. Initial setup to run goapi. **Done**
3. Function to connect to redis cache.  **Done**
4. Method for connecting mongodb using mgo. **Done**
5. Implementing basic CRUD operations for Mongodb in GO. **Done**

# Week4: 
* **Objective** ---> Testing AP property and Youtube video production.

# Task: 
1. Testing replication amongst the replicasets. **Done**
2. Creating a partition and testing AP property for the cluster. **Done**
3. Youtube video creation **In Progress** 

## AP property testing -usecases
Write data in primary nodes and read replicated data from secondary:
    1. Write data in primary node using following command

Test results on generated replica set:
  Test 1. Step down the primary for 12 seconds using following command:
* rs.stepDown(12);

              Expected Output                |              Actual Output                
---------------------------------------------------------------------------------------------
      Other nodes from all secondary         |        Other nodes from all secondary         
      should become primary                  |        becomes Primary

    Test passed.


  Test 2. Set Firewall using following command in one of the secondary which blocks inbound message from primary:
iptables -A INPUT -s <ip_address_of_AWS_instance_where_primary_node_of_repica_set_is_running> -j DROP

                Expected Output                |              Actual Output                
  ---------------------------------------------------------------------------------------------
        Should able to read stale data         |        Stale data is being read               

    Test passed.


To drop Firewall created in previous step setting use following command:
iptables -D INPUT -s <ip_address> -j DROP


189 - 54.153.65.106 - p
250 - 54.193.4.116 - s1
149 - 54.183.210.248 - a









   

   

