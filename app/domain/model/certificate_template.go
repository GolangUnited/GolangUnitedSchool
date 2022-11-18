package model

type CertificateTemplate struct {
	ID       int64
	Template string
}

type CertificateTemplatesListDto struct {
	Metadata                 PaginationResponse
	CertificateTemplatesList []CertificateTemplate
}
