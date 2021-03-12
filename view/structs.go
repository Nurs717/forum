package view

// User user
type User struct {
	ID       int
	Email    string
	Username string
	Password string
}

// SessionID sesID
type SessionID struct {
	Email     string
	SessionID string
}

// Post post
type Post struct { //Y
	PostID       int
	UserID       int
	PostBody     string
	PostDate     string
	PostTime     string
	UserName     string
	Category     string
	LikeCounts   int
	UserNameLogo string

	Comments []Comment
}

// Comment comment
type Comment struct {
	CommentID   int
	CommentBody string
	UserID      int
	PostID      int
	UserName    string
	LikeCounts  int
}

// Like likes
type Like struct {
	UserID int
	PostID int
}

// CommentLike comment likes
type CommentLike struct {
	CommentID int
	UserID    int
}
