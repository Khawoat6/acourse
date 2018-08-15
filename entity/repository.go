package entity

import "github.com/lib/pq"

// RegisterCourse type
type RegisterCourse struct {
	UserID    string
	Title     string
	ShortDesc string
	LongDesc  string
	Image     string
	Start     pq.NullTime
}

// UpdateCourse type
type UpdateCourse struct {
	ID        string
	Title     string
	ShortDesc string
	LongDesc  string
	Start     pq.NullTime
}

// RegisterCourseContent type
type RegisterCourseContent struct {
	CourseID  string
	Title     string
	LongDesc  string
	VideoID   string
	VideoType int
}