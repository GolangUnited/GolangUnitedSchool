package model

type PaginationParams struct {
	Page         int64 `json:"page"`
	PageSize     int64 `json:"page_size"`
	WithMetadata bool  `json:"with_total_count"`
}

type PaginationResponse struct {
	Page       int64 `json:"page"`
	PageSize   int64 `json:"page_size"`
	PageCount  int64 `json:"page_count"`
	TotalCount int64 `json:"total_count"`
}

// properly pagination example
/*
{
  "_metadata":
  {
      "page": 5,
      "per_page": 20,
      "page_count": 20,
      "total_count": 521,
      "Links": [
        {"self": "/products?page=5&per_page=20"},
        {"first": "/products?page=0&per_page=20"},
        {"previous": "/products?page=4&per_page=20"},
        {"next": "/products?page=6&per_page=20"},
        {"last": "/products?page=26&per_page=20"},
      ]
  },
  "records": [
    {
      "id": 1,
      "name": "Widget #1",
      "uri": "/products/1"
    },
    {
      "id": 2,
      "name": "Widget #2",
      "uri": "/products/2"
    },
    {
      "id": 3,
      "name": "Widget #3",
      "uri": "/products/3"
    }
  ]
}
https://stackoverflow.com/questions/12168624/pagination-response-payload-from-a-restful-api
*/
