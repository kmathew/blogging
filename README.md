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

##### Register Author
curl <uri>:/authors
POST --data {
             	"name": "kev",
             	"display_name": "sdfsf",
             	"email": "yo@yo.com"
             }
##### Get Author Obj
curl :/authors?email=<email-address>

##### Create Space
curl -X POST :/spaces  --data { "space_name": "FUNZONE", "owner_email": "yo@yo.com" }

##### Get Space by Author Email
curl :/spaces?owner_email=<email-address>

##### Create Blog
curl -X POST :/blogs --date {"title": "fun1", "content": "bytes of data here", "space_name": "FUNZONE", "yo@yo.com" }

##### Get Blog by Title
curl :/blogs?title=<title-name-here>
