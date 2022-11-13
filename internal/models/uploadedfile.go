package models

//	 UploadedFile is a struct used to save information about a file

type UploadedFile struct {
	NewFileName      string
	OriginalFilename string
	FileSize         int64
}
