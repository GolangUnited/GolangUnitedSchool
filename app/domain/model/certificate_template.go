package model

type CertificateTemplate struct {
	ID       int64  `json:"id" example:"1"`
	Template string `json:"template" example:"some template"`
}

type CertificateTemplateList struct {
	Metadata             PaginationResponse    `json:"_metadata"`
	CertificateTemplates []CertificateTemplate `json:"certificate_templates"`
}

type CertificateTemplatesListDto struct {
	Metadata                 PaginationResponse
	CertificateTemplatesList []CertificateTemplate
}
