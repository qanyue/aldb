package config

import "os"

func GetEnv() *Config {
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		uri = "mongodb://localhost:27017"
	}
	db := os.Getenv("MONGO_DB")
	if db == "" {
		db = "aldb"
	}
	bucket := os.Getenv("COS_BUCKET")
	region := os.Getenv("COS_REGION")
	secretID := os.Getenv("COS_SECRET_ID")
	secretKey := os.Getenv("COS_SECRET_KEY")
	baseUrl := os.Getenv("COS_BASE_URL")
	pathPrefix := os.Getenv("COS_PATH_PREFIX")
	return &Config{
		Mongo:      Mongo{uri, db},
		TencentCOS: TencentCOS{bucket, region, secretID, secretKey, baseUrl, pathPrefix},
	}
}
