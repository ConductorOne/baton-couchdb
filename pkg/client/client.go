package client

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/conductorone/baton-sdk/pkg/uhttp"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
)

type CouchDBClient struct {
	client      *uhttp.BaseHttpClient
	InstanceURL string
	Token       string
}

type Option func(c *CouchDBClient)

func WithBasicAuth(username, password string) Option {
	return func(c *CouchDBClient) {
		c.Token = basicAuth(username, password)
	}
}

func WithInstanceURL(url string) Option {
	return func(c *CouchDBClient) {
		c.InstanceURL = url
	}
}

func New(ctx context.Context, opts ...Option) (*CouchDBClient, error) {
	httpClient, err := uhttp.NewClient(ctx, uhttp.WithLogger(true, ctxzap.Extract(ctx)))
	if err != nil {
		return nil, err
	}

	cli, err := uhttp.NewBaseHttpClientWithContext(ctx, httpClient)
	if err != nil {
		return nil, err
	}

	cdbClient := CouchDBClient{
		client: cli,
	}

	for _, o := range opts {
		o(&cdbClient)
	}

	return &cdbClient, nil
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func (c *CouchDBClient) ListAllDataBases(ctx context.Context) ([]string, error) {
	var dbList []string
	endPoint := "/_all_dbs"

	url, err := url.JoinPath(c.InstanceURL, endPoint)
	if err != nil {
		return nil, err
	}

	_, err = c.doRequest(ctx, http.MethodGet, url, &dbList, nil)
	if err != nil {
		return nil, err
	}

	return dbList, nil
}

func (c *CouchDBClient) GetSecurityObject(ctx context.Context, dbName string) (*SecurityObject, error) {
	var dbSecObject *SecurityObject
	endPoint := "/_security"

	url, err := url.JoinPath(c.InstanceURL, dbName, endPoint)
	if err != nil {
		return nil, err
	}

	_, err = c.doRequest(ctx, http.MethodGet, url, &dbSecObject, nil)
	if err != nil {
		return nil, err
	}

	return dbSecObject, nil
}

func (c *CouchDBClient) doRequest(
	ctx context.Context,
	method string,
	endpointUrl string,
	res interface{},
	body interface{},
) (http.Header, error) {
	var (
		resp *http.Response
		err  error
	)

	urlAddress, err := url.Parse(endpointUrl)
	if err != nil {
		return nil, err
	}

	req, err := c.client.NewRequest(
		ctx,
		method,
		urlAddress,
		uhttp.WithAcceptJSONHeader(),
		uhttp.WithContentTypeJSONHeader(),
		uhttp.WithHeader("Authorization", "Basic "+c.Token),
		uhttp.WithJSONBody(body),
	)
	if err != nil {
		return nil, err
	}

	resp, err = c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if res != nil {
		bodyContent, err := io.ReadAll(resp.Body)

		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(bodyContent, &res)
		if err != nil {
			return nil, err
		}
	}

	return resp.Header, nil
}
