package services

type AwsFileStorage interface {
	UploadFile(filePath string, fileData []byte) (string, error)
	DeleteFile(filePath string) error
	DownloadFile(filePath string) ([]byte, error)
}
type awsFileStorage struct {
	// Add necessary fields here, e.g., AWS session, bucket name, etc.
}

func NewAwsFileStorage() AwsFileStorage {
	return &awsFileStorage{}
}

// Add any necessary fields, e.g., AWS session, bucket name, etc.
func (s *awsFileStorage) UploadFile(filePath string, fileData []byte) (string, error) {

	//TODO: Implement actual AWS S3 upload logic here
	return "https://s3.amazonaws.com/bucket/" + filePath, nil
}

func (s *awsFileStorage) DeleteFile(filePath string) error {

	return nil
}

func (s *awsFileStorage) DownloadFile(filePath string) ([]byte, error) {

	return []byte("file data"), nil
}
