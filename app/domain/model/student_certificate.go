package model

import "time"

type StudentCertificate struct {
	ID                    int64     `json:"id"`
	CreatedAt             time.Time `json:"created_at"`
	CertificateTemplateId int64     `json:"certificate_template_id"`
	StudentId             int64     `json:"student_id"`
	CourseId              int64     `json:"course_id"`
}
