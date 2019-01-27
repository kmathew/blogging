# Blogging
### Objective 
 
Create a blog building platform using AWS. Build the backend apis that will eventually support a frontend (building a frontend is not part of this exercise).  
 
The platform provides the following features: A blog author is able to register to the platform 
- A blog author may create a space to publish their blogs 
- A blog author can grant access to other authors to publish to their space 
- A blog author must approve any blog entries published in their space 
- A registered author can create a new blog entry in their own, or another author’s space 
- A registered author can fetch a list of their blog entries, select one and view. 
- Any user can view and read blog entries in any space. 

### Instructions
Make sure that your aws client is configured. In order to run the code locally, a config context file is required with a
configuration set for an IAM policy user with dynamoDB read and write and lambda function execute/add access.
https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-welcome.html

Setup IAM Policy
https://docs.aws.amazon.com/IAM/latest/UserGuide/id_users_create.html

In the root directory, there's a file called trust.json. This is a policy that needs to be attached to the role for the
lambda functions to work properly without resulting in Internal Server Errors. In order to get the resource id,
after deploying, test the api gateway. If you get a 500 Internal Server error, go and check the logs
for to see what permission is missing for a specific resource. It might be one of the tables for this
service to work.

Please run create_tables.go to generate the tables once. No environment variables need to be set unless you plan to connect to a local docker
container running dynamoDB.

For Deployment, create a function for each:
- authors
- blogs
- spaces
- approve

Upload each of their respective .zip's. If you wish to compile it on your own, you can use the makezip.sh script.
After creating the functions, create the api gateway and select authentication. Create a 
functional test and run it after.

### Endpoints

##### Register Author
Registers user
```
curl <uri>:/authors
POST --data {
             	"name": "kev",
             	"display_name": "sdfsf",
             	"email": "yo@yo.com"
             }
```
             
##### Get Author Obj
Get Author Object
```
curl :/authors?email=<email-address>
```
##### Create Space
Creates a Space under the user.
````
curl -X POST :/spaces  --data { "space_name": "FUNZONE", "owner_email": "yo@yo.com" }
````
##### Get Space by Author Email
Retrieves space by given owner email
````
curl :/spaces?owner_email=<email-address>
````
##### Create Blog
Creates a blog that is not yet approved.
````
curl -X POST :/blogs --date {"title": "fun1", "content": "bytes of data here", "space_name": "FUNZONE", "yo@yo.com" }
````
##### Get Blog by Title
Gets a blog for a given title
````
curl :/blogs?title=<title-name-here>
````
##### Get Blogs by Space Name
Gets a list of blogs for a given space
````
curl :/blogs?space_name=<space-name-here>
````
##### Get Blogs by Author Email
Gets a list of blogs by author email
````
curl :/blogs?author_email=<email-address>
````
##### Approve Blog
Approves a blog to be published to a space
````
curl -X POST :/approve --data {"title": "fun1", "space_name": "FUNZONE", "yo@yo.com"}
````