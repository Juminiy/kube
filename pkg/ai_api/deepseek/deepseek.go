package deepseek

import (
	"github.com/Juminiy/kube/pkg/internal_api"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/boltdb/bolt"
	"github.com/go-resty/resty/v2"
	"golang.org/x/net/context"
)

type Client struct {
	cli *resty.Client
	ctx context.Context

	baseUrl string
	apiKey  string

	store *bolt.DB
}

func New(baseUrl, apiKey string) (*Client, error) {
	store, err := bolt.Open(_DefaultLocalStorage, internal_api.FilePerm, nil)
	if err != nil {
		return nil, err
	}
	err = store.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(_BucketCompletions)
		if err != nil {
			return err
		}
		_, err = tx.CreateBucketIfNotExists(_BucketBetaCompletions)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &Client{
		cli: resty.New().
			SetScheme("https").
			SetBaseURL(baseUrl).
			SetAuthScheme("Bearer").
			SetAuthToken(apiKey).
			SetTimeout(_DefaultHTTPTimeout).
			SetAllowGetMethodPayload(true).
			SetHeader("Content-Type", "application/json").
			SetHeader("Accept", "application/json"),
		ctx:     context.Background(),
		baseUrl: baseUrl,
		apiKey:  apiKey,
		store:   store,
	}, nil
}

var _DefaultHTTPTimeout = util.TimeSecond(60)
var _DefaultLocalStorage = "deepseek.db"
