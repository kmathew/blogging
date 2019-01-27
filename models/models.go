package models

type Blog struct {
	Title    string `json:"title" dynamodbav:"title"`
	Content  []byte `json:"content" dynamodbav:"content"`
	SpaceName  string `json:"space_name" dynamodbav:"space_name"`
	AuthorName string `json:"author_name" dynamodbav:"author_name"` //local second
	Approved bool `json:"author_name" dynamodbav:"author_name"`
	//APPROVED = local secondary index
	//global secondary index author
	//global secondary index space
	//local secondary index for status of approval
	//global secondary index for space
}

type Author struct {
	Name    string `json:"name" dynamodbav:"name"`
	DisplayName  string `json:"display_name" dynamodbav:"display_name"`
	Email string `json:"email" dynamodbav:"email"`
	SpaceName string `json:"owned_space_name" dynamodbav:"owned_space_name"`
	//spaces is global secondary index
}

type Space struct {
	Name		 string `json:"name" dynamodbav:"name"`
	OwnerName      string `json:"owner_name" dynamodbav:"owner_name"`
	//blogs is a
}

type Approval struct {
	//ApprovalID = SpaceName+BlogTitle
	ApprovalID string `json:"approval_id" dynamodbav:"name"`
	SpaceName    string  `json:"space_name" dynamodbav:"space_name"`
	RequestorName string `json:"requestor_name" dynamodbav:"requestor_name"`
	ApproverName string `json:"approver_name" dynamodbav:"approver_name"`
	BlogTitle     string `json:"blog_title" dynamodbav:"blog_title"`
	Status     string `json:"status" dynamodbav:"status"`
}