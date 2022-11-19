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
	ID                    int64     `json:"id"`
	CreatedAt             time.Time `json:"created_at"`
	CertificateTemplateID int64     `json:"certificate_template_id"`
	StudentID             int64     `json:"student_id"`
	CourseID              int64     `json:"course_id"`
}

type StudentCertificateList struct {
	Metadata            PaginationResponse   `json:"_metadata"`
	StudentCertificates []StudentCertificate `json:"student_certificates"`
}

type StudentCertificatesListDto struct {
	Metadata                PaginationResponse
	StudentCertificatesList []StudentCertificate
}

type StudentHomework struct {
	ID         int64     `json:"id"`
	StudentID  int64     `json:"student_id"`
	HomeworkID int64     `json:"homework_id"`
	StartedAt  time.Time `json:"started_at"`
	FinishedAt time.Time `json:"finished_at"`
	Score      float32   `json:"score"`
}

type StudentHomeworkList struct {
	Metadata         PaginationResponse `json:"_metadata"`
	StudentHomeworks []StudentHomework  `json:"student_homeworks"`
}

type StudentHomeworksListDto struct {
	MetaData              PaginationResponse
	StudentsHomeworksList []StudentHomework
}
