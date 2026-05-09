package model

type SearchRequest struct {
	Prefix          string `form:"prefix" json:"prefix"`
	CountryCode     string `form:"country_code" json:"country_code,omitempty"`
	InstitutionName string `form:"institution_name" json:"institution_name,omitempty"`
	ActiveOnly      bool   `form:"active_only" json:"active_only"`
	Page            int    `form:"page" json:"page"`
	PageSize        int    `form:"page_size" json:"page_size"`
}

type SearchResponse struct {
	Data       []FinancialInstitution `json:"data"`
	Total      int64                  `json:"total"`
	Page       int                    `json:"page"`
	PageSize   int                    `json:"page_size"`
	TotalPages int                    `json:"total_pages"`
	CacheLevel string                 `json:"cache_level,omitempty"`
}

type GetResponse struct {
	Data       *FinancialInstitution `json:"data"`
	CacheLevel string                `json:"cache_level,omitempty"`
}
