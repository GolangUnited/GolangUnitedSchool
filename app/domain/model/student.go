package model

import "time"

type Student struct {
	StudentId int64
}

type StudentCertificate struct {
	ID            int64
	StudentID     int64
	CourseID      int64
	CertificateID int64
	CreatedAt     time.Time
}
