package azureblobremote

import (
	"fmt"
	"github.com/VictoriaMetrics/VictoriaMetrics/lib/backup/common"
	"github.com/VictoriaMetrics/VictoriaMetrics/lib/logger"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"io"
	"github.com/Azure/azure-storage-blob-go/azblob"
	"context"
	"log"
	"net/url"
	"strings"
	"os"
)

// FS represents filesystem for backups in GCS.
//
// Init must be called before calling other FS methods.
type FS struct {
	// Azure Blob Storage Account
	storageAccount string

	// Azure Blob Access Key
	storageAccessKey string

	// Azure Blob container to use.
	Container string

	// Directory in the container to write to.
	Dir string

	containerURL *azblob.ContainerURL

}

func handleErrors(err error) {
	if err != nil {
		if serr, ok := err.(azblob.StorageError); ok { // This error is a Service-specific
			switch serr.ServiceCode() { // Compare serviceCode to ServiceCodeXxx constants
			case azblob.ServiceCodeContainerAlreadyExists:
				fmt.Println("Received 409. Container already exists")
				return
			}
		}
		log.Fatal(err)
	}
}

func (fs *FS) Init() error {
	if fs.containerURL != nil {
		logger.Panicf("BUG: Init is already called")
	}
	for strings.HasPrefix(fs.Dir, "/") {
		fs.Dir = fs.Dir[1:]
	}
	if !strings.HasSuffix(fs.Dir, "/") {
		fs.Dir += "/"
	}

	// Check if credentials flags are set else try from environment variables.
	if len(fs.storageAccount) == 0 || len(fs.storageAccessKey) == 0 {
		accountName, accountKey := os.Getenv("AZURE_STORAGE_ACCOUNT"), os.Getenv("AZURE_STORAGE_ACCESS_KEY")
		if len(accountName) == 0 || len(accountKey) == 0 {
			fs.storageAccount = accountName
			fs.storageAccessKey = accountKey
		}
	}

	// Create a default request pipeline using your storage account name and account key.
	credential, err := azblob.NewSharedKeyCredential(fs.storageAccount, fs.storageAccessKey)
	if err != nil {
		log.Fatal("Invalid credentials with error: " + err.Error())
	}
	p := azblob.NewPipeline(credential, azblob.PipelineOptions{})

	// Construct blob url
	URL, _ := url.Parse(
		fmt.Sprintf("https://%s.blob.core.windows.net/%s", fs.storageAccount, fs.Container))

	// Create a ContainerURL object that wraps the container URL and a request
	// pipeline to make requests.
	containerURL := azblob.NewContainerURL(*URL, p)

	// Create the container
	// fmt.Printf("Creating a container named %s\n", fs.Container)
	// _, err = containerURL.Create(fs.ctx, azblob.Metadata{}, azblob.PublicAccessNone)
	// handleErrors(err)

	fs.containerURL = &containerURL
	return nil
}

// MustStop stops fs.
func (fs *FS) MustStop() {
	fs.containerURL = nil
}

// String returns human-readable description for fs.
func (fs *FS) String() string {
	return fmt.Sprintf("AZUREBLOB{bucket: %q, dir: %q}", fs.Container, fs.Dir)
}

func (fs *FS) ListParts() ([]common.Part, error) {
	dir := fs.Dir
	ctx := context.Background()
	for marker := (azblob.Marker{}); marker.NotDone(); {
		// Get a result segment starting with the blob indicated by the current Marker.
		listBlob, err := fs.containerURL.ListBlobsFlatSegment(ctx, marker, azblob.ListBlobsSegmentOptions{})
		handleErrors(err)

		// ListBlobs returns the start of the next segment; you MUST use this to get
		// the next segment (after processing the current result segment).
		marker = listBlob.NextMarker

		// Process the blobs returned in this result segment (if the segment is empty, the loop body won't execute)
		for _, blobInfo := range listBlob.Segment.BlobItems {
			fmt.Print("	Blob name: " + blobInfo.Name + "\n")
		}
	}

	panic("implement me")
}

func (fs *FS) DeletePart(p common.Part) error {
	panic("implement me")
}

func (fs *FS) RemoveEmptyDirs() error {
	panic("implement me")
}

func (fs *FS) CopyPart(dstFS common.OriginFS, p common.Part) error {
	panic("implement me")
}

func (fs *FS) DownloadPart(p common.Part, w io.Writer) error {
	panic("implement me")
}

func (fs *FS) UploadPart(p common.Part, r io.Reader) error {
	panic("implement me")
}

func (fs *FS) DeleteFile(filePath string) error {
	panic("implement me")
}

func (fs *FS) CreateFile(filePath string, data []byte) error {
	panic("implement me")
}

func (fs *FS) HasFile(filePath string) (bool, error) {
	panic("implement me")
}

