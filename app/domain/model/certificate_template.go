package model

type CertificateTemplate struct {
	ID       int64  `json:"id"`
	Template string `json:"template"`
}

type CertificateTemplateList struct {
	Metadata             PaginationResponse    `json:"_metadata"`
	CertificateTemplates []CertificateTemplate `json:"certificate_templates"`
}
