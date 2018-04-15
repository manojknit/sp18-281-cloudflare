## Research on how to deploy the MongoDB database to cloud.

#### MongoDB deployment research:

  - MongoDB can be deployed on the flexible AWS platform and can set up a fully customizable MongoDB cluster on demand. 	
  - The flexible AWS architecture allows us to choose the most appropriate network, compute, and storage infrastructure for our environment. 
  - Building a scalable, on-demand infrastructure on AWS provides a cost-effective solution for handling large-scale compute and storage requirements	
  - Need to follow the steps for the deployment of MondoDb cluster on AWS:
  - Launch the AWS CloudFormation template into the AWS account, specify parameter values, and create the stack.
  - Can choose any one of : 1)Deploy MongoDB into a new VPC on AWS 2)Deploy MongoDB into an existing VPC on AWS
  - If deploying MongoDB into an existing VPC, the VPC should be set up with two public subnets and three private subnets in different Availability Zones
  - The private subnets require NAT gateways or NAT instances in their route tables for outbound internet connectivity, and you must create bastion hosts and their associated security group for inbound SSH access.
  - On the Select Template page, keep the default setting for the template URL, and then choose Next.
  
#### Usage of MongoDB cluster in project:
  - The frontend will request the go API.
  - The go API will fetch or store data from the mongo cluster. 
  - As the Mongo cluster data is highly available, it will ensure availability of data. 
