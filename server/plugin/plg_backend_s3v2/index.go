package plg_backend_s3v2

import (
	"io"
	"os"

	. "github.com/mickael-kerjean/filestash/server/common"
)

type S3V2Backend struct {
}

func init() {
	Backend.Register("s3v2", S3V2Backend{})
}

func (this S3V2Backend) Init(params map[string]string, app *App) (IBackend, error) {
	return nil, nil
}

func (this S3V2Backend) LoginForm() Form {
	return Form{
		Elmnts: []FormElement{
			FormElement{
				Name:  "type",
				Type:  "hidden",
				Value: "s3",
			},
			FormElement{
				Name:        "access_key_id",
				Type:        "text",
				Placeholder: "Access Key ID*",
			},
			FormElement{
				Name:        "secret_access_key",
				Type:        "password",
				Placeholder: "Secret Access Key*",
			},
			FormElement{
				Name:        "advanced",
				Type:        "enable",
				Placeholder: "Advanced",
				Target: []string{
					"s3_region", "s3_endpoint", "s3_role_arn", "s3_session_token",
					"s3_path", "s3_encryption_key", "s3_number_thread", "s3_timeout",
				},
			},
			FormElement{
				Id:          "s3_region",
				Name:        "region",
				Type:        "text",
				Placeholder: "Region",
			},
			FormElement{
				Id:          "s3_endpoint",
				Name:        "endpoint",
				Type:        "text",
				Placeholder: "Endpoint",
			},
			FormElement{
				Id:          "s3_role_arn",
				Name:        "role_arn",
				Type:        "text",
				Placeholder: "Role ARN",
			},
			FormElement{
				Id:          "s3_session_token",
				Name:        "session_token",
				Type:        "text",
				Placeholder: "Session Token",
			},
			FormElement{
				Id:          "s3_path",
				Name:        "path",
				Type:        "text",
				Placeholder: "Path",
			},
			FormElement{
				Id:          "s3_encryption_key",
				Name:        "encryption_key",
				Type:        "text",
				Placeholder: "Encryption Key",
			},
			FormElement{
				Id:          "s3_number_thread",
				Name:        "number_thread",
				Type:        "text",
				Placeholder: "Num. Thread",
			},
			FormElement{
				Id:          "s3_timeout",
				Name:        "timeout",
				Type:        "number",
				Placeholder: "List Object Timeout",
			},
		},
	}
}

func (this S3V2Backend) Meta(path string) Metadata {
	if path == "/" {
		return Metadata{
			CanCreateFile: NewBool(false),
			CanRename:     NewBool(false),
			CanMove:       NewBool(false),
			CanUpload:     NewBool(false),
		}
	}
	return Metadata{}
}

func (this S3V2Backend) Ls(path string) (file []os.FileInfo, err error) {
	return nil, nil
}

func (this S3V2Backend) Stat(path string) (os.FileInfo, error) {
	return nil, nil
}

func (this S3V2Backend) Cat(path string) (io.ReadCloser, error) {
	return nil, nil
}

func (this S3V2Backend) Mkdir(path string) error {
	return nil
}

func (this S3V2Backend) Rm(path string) error {
	return nil
}

func (this S3V2Backend) Mv(from string, to string) error {
	return nil
}

func (this S3V2Backend) Touch(path string) error {
	return nil
}

func (this S3V2Backend) Save(path string, file io.Reader) error {
	return nil
}
