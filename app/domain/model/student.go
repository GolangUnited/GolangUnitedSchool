package model

import "time"

// no testing
type Student struct {
	StudentId int64 `json:"student_id" validate:"required,numeric,gte=0"`
}

type StudentsListDto struct {
	MetaData    PaginationResponse
	StudentList []Student
}

type StudentCertificate struct {
	ID                    int64
	CreatedAt             time.Time
	CertificateTemplateID int64
	StudentID             int64
	CourseID              int64
}

type StudentCertificatesListDto struct {
	Metadata                PaginationResponse
	StudentCertificatesList []StudentCertificate
}

type StudentHomework struct {
	ID         int64
	StudentID  int64
	HomeworkID int64
	StartedAt  time.Time
	FinishedAt time.Time
	Score      float32
}

type StudentHomeworksListDto struct {
	MetaData              PaginationResponse
	StudentsHomeworksList []StudentHomework
}
