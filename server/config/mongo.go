package config

type Mongo struct {
	MongoDB  string `mapstructure:"db" json:"db" yaml:"db"`    // 数据库名
	MongoURL string `mapstructure:"url" json:"url" yaml:"url"` // Mongo_URL
}
