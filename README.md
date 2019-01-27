# Blogging
### Objective 
 
Create a blog building platform using AWS. Build the backend apis that will eventually support a frontend
 
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

Please run create_tables.go to generate the tables once. No environment variables need to be set unless you plan to 
connect to a local docker container running dynamoDB.

For Deployment, create a function for each:
- authors
- blogs
- spaces
- approve

Upload each of their respective .zip's. If you wish to compile it on your own, you can use the makezip.sh.
After creating the functions, create the api gateway and select authentication. Create a 
functional test and run it after.

### Endpoints

Most methods for POST require Headers for Content-Type: application/json

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
Retrieves space by name
````
curl -X POST :/spaces  --data { "space_name": "FUNZONE", "owner_email": "yo@yo.com" }
````
##### Get Space by Name
````
curl :/spaces?space_name=<name-here>
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
Gets a list of approved blogs for a given space
````
curl :/blogs?space_name=<space-name-here>
````
##### Get Blogs by Author Email
Gets a list of blogs by author email regardless if approved or not
````
curl :/blogs?author_email=<email-address>
````
##### Approve Blog
Approves a blog to be published to a space
````
curl -X POST :/approve --data {"title": "fun1", "space_name": "FUNZONE", "yo@yo.com"}
````

## Notes
I have never worked with dynamoDB before. I had a rough understanding that it's meant to be used a Document store. 
I wanted to play around with the SecondaryIndexes and did just that in this project. I have never used AWS lambda
but I did learn how to use it to run my functions, creating api-gateways, and update IAM policies.

I would definitely do a few things differently in this project:

- Outside of glide I'm not familiar with dep or vgo. I would probably include one of those dependency managements. 

- I'm used to running stuff in Docker
on my local. While I did find a way to run my own dynamoDB on localhost, I couldn't really setup a docker to emulate
the lambda portion. I wish I had found AWS-SAM-CLI before I started this. It might have saved me some time
and helped with the testing/make stuff look cleaner

- A redesign for the Table structure for the data needs to be considered. I originally had designed the tables for
RDBMS instead of a denormalized one. A form of caching definitely needs to be used. Some of the functions have expensive
queries. Hopefully, the secondary indexes offset most.

- Unit tests.... I'm still not entirely sure how to properly unit test serverless applications without going out of my
way of changing the way my functions run in order to add the coverage.

- Blogging author authentication.. I just ignored this. As of right now, anyone can post blogs as other people etc

- Validations before doing a certain action.. Need more validations like when creating a blog post or registering a new 
author

- Requirements... I was confused with one of the requirements. Space owners can give others access, but authors have the
ability to create blog posts on anyone's page. Did spaces have to have a curated list of users before said authors could
publish a blog? Did the space owner approving blogs mean that it is giving access to other authors? I assumed the latter.
If it's the former, then I would have another field under Spaces that had a Hash Map of all the users with access. I would
not update dynamoDB schema etc. It would just be additional info in the document store. I would then pull space and get 
and use the map as a LUT to see if the person trying to publish is an author given access by a space owner. Wait, could
authors get more than one space?

- Logging... I wasn't sure if I should have using glog or logrus. I did see cloudfront having logs being published there.


## Credits
https://www.alexedwards.net/blog/serverless-api-with-go-and-aws-lambda
https://read.acloud.guru/serverless-golang-api-with-aws-lambda-34e442385a6a
https://hub.docker.com/r/amazon/dynamodb-local
https://docs.aws.amazon.com/sdk-for-go/api/service/dynamodb/

## Thanks
Justin and Madhu for giving me a shot at this. 
