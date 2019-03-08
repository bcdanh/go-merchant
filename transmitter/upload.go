package transmitter

import (
	"bytes"
	"log"
	"os"

	drive "google.golang.org/api/drive/v3"
)

const folderID = "1p4J2b8cDhAA9NjKGd0xNfx91BA_nFpvc"

var dbFileID = "1AqR6URNNnWPxdFxkd-tBrP2Uv1CAJAjy"
var logFileID = "1UhSUwq0uz6dUsXeHfA7GTRDvzmcMhtZ7"

func UploadDBFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		log.Printf("ERROR: opening %q: %v\n", filename, err)
		return err
	}
	defer file.Close()

	dbFile := drive.File{Name: "RealtimeData.db"}
	_, err = srv.Files.Update(dbFileID, &dbFile).AddParents(folderID).Media(file).Do()

	if err != nil {
		log.Println("Unable to create file: ", err)
	}

	return err
}

func UpdateFile(filepath string, fileID string, filename string) error {
	f, err := os.Open(filepath)
	if err != nil {
		log.Printf("ERROR: opening %q: %v\n", filepath, err)
		return err
	}
	defer f.Close()

	dFile := drive.File{Name: filename}
	_, err = srv.Files.Update(fileID, &dFile).AddParents(folderID).Media(f).Do()

	if err != nil {
		log.Println("ERROR: drive: Unable to create file: ", err)

	}
	return err
}

func CreateFile(filename string) string {
	file := bytes.NewReader([]byte{})

	driveFile := drive.File{Name: filename, Parents: []string{folderID}}
	onDrive, err := srv.Files.Create(&driveFile).Media(file).Do()
	if err != nil {
		log.Println("ERROR: create file on drive: ", err)
	}
	return onDrive.Id
}
