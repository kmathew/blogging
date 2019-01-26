package models

type Blog struct {
	Title    string `json:"title" dynamodbav:"title"`
	Content  []byte `json:"content" dynamodbav:"content"`
	SpaceName  string `json:"space_name" dynamodbav:"space_name"`
	AuthorName string `json:"author_name" dynamodbav:"author_name"`
}

type Author struct {
	Name    string `json:"name" dynamodbav:"name"`
	DisplayName  string `json:"display_name" dynamodbav:"display_name"`
	SpaceName []string `json:"owned_space_id_list" dynamodbav:"owned_space_id_list"`
}

type Space struct {
	Name		 string `json:"name" dynamodbav:"name"`
	OwnerName      string `json:"owner_name" dynamodbav:"owner_name"`
	ApprovedList []string `json:"approved_blog_titles" dynamodbav:"approved_blog_titles"`
}

type Approval struct {
	//ApprovalID = SpaceName+BlogTitle
	ApprovalID string `json:"approval_id" dynamodbav:"name"`
	SpaceName    string  `json:"space_name" dynamodbav:"space_name"`
	BlogTitle     string `json:"blog_title" dynamodbav:"blog_title"`
	Status     string `json:"status" dynamodbav:"status"`
}

func RegisterAuthor(name string) {
	//check if author name is taken/already exists
	//if taken.. return error
	//if not ... add to author list
}

func CreateSpace(name string, authorID string) {
	//check if space exists for author already
	//if does error
	//if not create space/insert in table
}

func CreateBlog(title string, content []byte, spaceID string, authorID string) {
	//generate id
	//insert blog into table
	//if spaceid belongs to authorid.... then auto approve
	//if not then send blog for approval
}

func GetBlogsForSpaceID(spaceID string) {
	//return all blogs associated with spaceid that are approved
}

func GetUnapprovedBlogsForSpace(spaceID string) {
	//get unapproved blogs
}

func GetBlogsByAuthorID(authorID string) {
	//return all blogs by authorid
}

func GetBlogIDsByAuthorID(authorID string) {
	//return all blogs IDs by authorID
}

func GetBlogByBlogID(blogID string) {

}