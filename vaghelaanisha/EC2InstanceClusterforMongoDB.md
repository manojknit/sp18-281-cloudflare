## Creation of EC2 nodes with mongoDB and forming a cluster of nodes.

#### Creation of EC2 AWS instances:
  - Login and go to the AWS console. Then select EC2 and go to the EC2 dashboard.
  - Create new EC2 instance with the required configuration:
  - Launch the instance.
  - Give the name as "master-node" for this instance.
  - Create the AMI image for the EC2 instance created.
  
#### Installation of Mongo in EC2 instance:
  - Install mongo on EC2 using following command:
  - sudo yum update -y  
  
#### Creation of 4 other EC2 AWS instances for the cluster:
  - Select the AMI image created in the previous step and create 4 other EC2 instances with the same configuration from this image.
  - Name the 4 other EC2 instances as "Slave-node-1", "slave-node-2", "slave-node-3" and "arbiter".

#### Creation of the cluster(Setting up Replica set for mongo db in the EC2 instances):
  - Connect to the master node using ssh command
  - Connect to mongodb in the master node using:"mongo"
  - Create a mongodb user
  
  - Terminate any process running on port 27017 using following command on all nodes:
```
sudo service mongod stop
```
  - Start mongodb on all nodes by binding the ip
  - Commands to run on master node to initiate the replica set and add nodes:
```
rs.initiate()
rs.add('slave-IP:27017');
rs.addArb('arbiter-IP:27017');
```
- Commands to run on slave node to connect to master node:
```
use <dbname>
rs.slaveOk()
```
