package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	s3Bucket := os.Getenv("S3_BUCKET")
	secretKey := os.Getenv("SECRET_KEY")

	_ = godotenv.Load(".env")
	fmt.Println(s3Bucket)
	fmt.Println(secretKey)
}
