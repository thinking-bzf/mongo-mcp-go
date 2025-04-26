package model

type FindDocumentRequest struct {
	Collection string                 `mapstructure:"collection" json:"collection"`
	Filter     map[string]interface{} `mapstructure:"filter" json:"filter"`
	Limit      int64                  `mapstructure:"limit" json:"limit"`
	Projection map[string]interface{} `mapstructure:"projection" json:"projection"`
}

type CountDocumentRequest struct {
	Collection string                 `mapstructure:"collection" json:"collection"`
	Filter     map[string]interface{} `mapstructure:"filter" json:"filter"`
	Limit      int64                  `mapstructure:"limit" json:"limit"`
}

type InsertDocumentRequest struct {
	Collection string `mapstructure:"collection" json:"collection"`
	Document   string `mapstructure:"document" json:"document"`
}

type DeleteDocumentRequest struct {
	Collection string                 `mapstructure:"collection" json:"collection"`
	Filter     map[string]interface{} `mapstructure:"filter" json:"filter"`
}

type UpdateDocumentRequest struct {
	Collection string                 `mapstructure:"collection" json:"collection"`
	Filter     map[string]interface{} `mapstructure:"filter" json:"filter"`
	Update     map[string]interface{} `mapstructure:"update" json:"update"`
}
