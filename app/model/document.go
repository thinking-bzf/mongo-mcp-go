package model

type FindDocumentRequest struct {
	Collection string                 `mapstructure:"collection" json:"collection" bson:"collection"`
	Filter     map[string]interface{} `mapstructure:"filter" json:"filter" bson:"filter"`
	Limit      int64                  `mapstructure:"limit" json:"limit" bson:"limit"`
	Projection map[string]interface{} `mapstructure:"projection" json:"projection" bson:"projection"`
}

type CountDocumentRequest struct {
	Collection string                 `mapstructure:"collection" json:"collection" bson:"collection"`
	Filter     map[string]interface{} `mapstructure:"filter" json:"filter" bson:"filter"`
	Limit      int64                  `mapstructure:"limit" json:"limit" bson:"limit"`
}

type InsertDocumentRequest struct {
	Collection string `mapstructure:"collection" json:"collection" bson:"collection"`
	Document   string `mapstructure:"document" json:"document" bson:"document"`
}

type DeleteDocumentRequest struct {
	Collection string                 `mapstructure:"collection" json:"collection" bson:"collection"`
	Filter     map[string]interface{} `mapstructure:"filter" json:"filter" bson:"filter"`
}

type UpdateDocumentRequest struct {
	Collection string                 `mapstructure:"collection" json:"collection"`
	Filter     map[string]interface{} `mapstructure:"filter" json:"filter" bson:"filter"`
	Update     map[string]interface{} `mapstructure:"update" json:"update" bson:"update"`
}
