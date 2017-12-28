package api

// UserInfo type
type UserInfo struct {
	StudentNo    string
	Name         string
	Faculty      string
	FieldOfStudy string
	Degree       string
}

// CourseInfo type
type CourseInfo struct {
	Code   string
	Title  string
	Grade  string
	Credit string
}
