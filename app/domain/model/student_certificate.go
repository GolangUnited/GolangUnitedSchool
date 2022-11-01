package model

import "time"

type StudentCertificate struct {
	ID                    int64
	CreatedAt             time.Time
	CertificateTemplateID int64
	StudentID             int64
	CourseID              int64
}
