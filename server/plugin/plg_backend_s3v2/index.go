package plg_backend_s3v2

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	. "github.com/mickael-kerjean/filestash/server/common"
)

var S3Cache AppCache

type S3V2Backend struct {
	config     *aws.Config
	params     map[string]string
	app        *App
	threadSize int
	timeout    time.Duration
}

func init() {
	Backend.Register("s3v2", S3V2Backend{})
	S3Cache = NewAppCache(2, 1)
}

func (this S3V2Backend) Init(params map[string]string, app *App) (IBackend, error) {
	// TODO: Might be better to use optFns here and override all the options in there
	config, err := config.LoadDefaultConfig(this.app.Context)
	if err != nil {
		return nil, NewError(fmt.Sprintf("Error loading default AWS config"), 400)
	}

	if params["encryption_key"] != "" && len(params["encryption_key"]) != 32 {
		return nil, NewError(fmt.Sprintf("Encryption key needs to be 32 characters (current: %d)", len(params["encryption_key"])), 400)
	}
	region := params["region"]
	if region == "" {
		region = "us-east-1"
		if strings.HasSuffix(params["endpoint"], ".cloudflarestorage.com") {
			region = "auto"
		}
		config.Region = region
	}
	// TODO: Not sure this CredentialsProvider thing is correct in v2
	if params["access_key_id"] != "" || params["secret_access_key"] != "" {
		creds := credentials.StaticCredentialsProvider{}

		creds.Value.AccessKeyID = params["access_key_id"]
		creds.Value.SecretAccessKey = params["secret_access_key"]
		creds.Value.SessionToken = params["session_token"]

		config.Credentials = creds
	}

	if params["endpoint"] != "" {
		config.BaseEndpoint = aws.String(params["endpoint"])
	}

	var timeout time.Duration
	if secs, err := strconv.Atoi(params["timeout"]); err == nil {
		timeout = time.Duration(secs) * time.Second
	}
	threadSize, err := strconv.Atoi(params["number_thread"])
	if err != nil {
		threadSize = 50
	} else if threadSize > 5000 || threadSize < 1 {
		threadSize = 2
	}
	
	backend := &S3V2Backend{
		config: &config,
		params: params,
		app: app,
		threadSize: threadSize,
		timeout: timeout,
	}

	return backend, nil
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
					"s3_region", "s3_endpoint", "s3_session_token",
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
