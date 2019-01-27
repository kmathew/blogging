package models

type Blog struct {
	Title    string `json:"title" dynamodbav:"title"`
	Content  []byte `json:"content" dynamodbav:"content"`
	SpaceName  string `json:"space_name" dynamodbav:"space_name"`
	AuthorEmail string `json:"author_email" dynamodbav:"author_email"` //local second
	Approved string `json:"approved" dynamodbav:"approved"`
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
	//spaces is global secondary index
}

type Space struct {
	Name		 string `json:"space_name" dynamodbav:"space_name"`
	OwnerEmail      string `json:"owner_email" dynamodbav:"owner_email"`
	//blogs is a
}

type Approval struct {
	//ApprovalID = SpaceName+BlogTitle
	SpaceName    string  `json:"space_name" dynamodbav:"space_name"`
	ApproverEmail string `json:"approver_email" dynamodbav:"approver_email"`
	Title     string `json:"title" dynamodbav:"title"`
}