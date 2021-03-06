/*
 * S3 asset uploader REST API
 *
 * an S3 asset uploader API
 *
 * API version: 1.0.0
 * Contact: bill.hongwu.chen@gmail.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"time"
)

type Asset struct {

	// UUID of asset
	Id string `json:"id,omitempty"`

	// A s3 signed url for upload
	UploadUrl string `json:"upload_url,omitempty"`

	// A s3 signed url for download
	DownloadUrl string `json:"download_url,omitempty"`

	// timeout in seconds when the download url is still working. default is 60 seconds if not psecified.
	Timeout int32 `json:"timeout,omitempty"`

	// The status of file upload operation to s3
	Status string `json:"status,omitempty"`

	// created timestamp of the asset
	Created time.Time `json:"created,omitempty"`
}
