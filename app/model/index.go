package model

type CreateIndexRequest struct {
	Collection string                 `mapstructure:"collection" json:"collection"`
	IndexSpec  map[string]interface{} `mapstructure:"index_spec" json:"index_spec"`
}
type DropIndexRequest struct {
	Collection string `mapstructure:"collection" json:"collection"`
	IndexName  string `mapstructure:"index_name" json:"index_name"`
}
