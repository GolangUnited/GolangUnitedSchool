package model

import "time"

type Student struct {
	ID int64
}

type StudentNote struct {
	ID        int64
	StudentID int64
	Note      string
	CreatedAt time.Time
}

type StudentCertificate struct {
	ID            int64
	StudentID     int64
	CourseID      int64
	CertificateID int64
	CreatedAt     time.Time
}
