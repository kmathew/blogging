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
