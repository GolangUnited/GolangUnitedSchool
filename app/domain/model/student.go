package model

import "time"

type Student struct {
	StudentId int64
}

type StudentCertificate struct {
	ID                    int64
	CreatedAt             time.Time
	CertificateTemplateID int64
	StudentID             int64
	CourseID              int64
}

type StudentHomework struct {
	ID         int64
	StudentID  int64
	HomeworkID int64
	StartedAt  time.Time
	FinishedAt time.Time
	Score      float32
}
