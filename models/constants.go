package models

//indexes
const LocalIndexAuthor string = "local_index_author"
const LocalIndexApproved string = "local_index_approved"

const GlobalIndexSpaceName string = "global_index_space_name"
const GlobalIndexAuthorEmail string = "global_index_author_email"
const GlobalIndexApproved string = "global_index_approved"

const GlobalIndexOwnerEmail string = "global_index_owner_email"

//Bool
const True string = "true"
const False string = "false"

//Tables
const SpacesTable string = "Spaces"
const AuthorsTable string = "Authors"
const BlogsTable string = "Blogs"

//Keys
//Authors
const Email string = "email"
const DispName string = "display_name"

//Blogs
const Title string = "title"
const SpaceName string = "space_name"
const AuthorEmail string = "author_email"
const Approved string = "approved"

//Space
const OwnerEmail string = "owner_email"