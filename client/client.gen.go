// Package client provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.0 DO NOT EDIT.
package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/ghodss/yaml"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	externalRef0 "github.com/observatorium/api/client/models"
	externalRef1 "github.com/observatorium/api/client/parameters"
	externalRef2 "github.com/observatorium/api/client/responses"
)

// GetLabelValuesParams defines parameters for GetLabelValues.
type GetLabelValuesParams struct {
	// Repeated series selector argument
	Match *externalRef1.OptionalSeriesMatcher `json:"match[],omitempty"`

	// Start timestamp
	Start *externalRef1.StartTS `json:"start,omitempty"`

	// End timestamp
	End *externalRef1.EndTS `json:"end,omitempty"`
}

// GetLabelsParams defines parameters for GetLabels.
type GetLabelsParams struct {
	// Repeated series selector argument
	Match *externalRef1.OptionalSeriesMatcher `json:"match[],omitempty"`

	// Start timestamp
	Start *externalRef1.StartTS `json:"start,omitempty"`

	// End timestamp
	End *externalRef1.EndTS `json:"end,omitempty"`
}

// GetInstantQueryParams defines parameters for GetInstantQuery.
type GetInstantQueryParams struct {
	// query to fetch result for
	Query *externalRef1.Query `json:"query,omitempty"`

	// Evaluation timeout
	Timeout *externalRef1.QueryTimeout `json:"timeout,omitempty"`

	// Evaluation timestamp
	Time *string `json:"time,omitempty"`
}

// GetRangeQueryParams defines parameters for GetRangeQuery.
type GetRangeQueryParams struct {
	// query to fetch result for
	Query *externalRef1.Query `json:"query,omitempty"`

	// Start timestamp
	Start *externalRef1.StartTS `json:"start,omitempty"`

	// End timestamp
	End *externalRef1.EndTS `json:"end,omitempty"`

	// Evaluation timeout
	Timeout *externalRef1.QueryTimeout `json:"timeout,omitempty"`

	// Query resolution step width
	Step *string `json:"step,omitempty"`
}

// GetRulesParams defines parameters for GetRules.
type GetRulesParams struct {
	// type of Rules
	Type *string `json:"type,omitempty"`

	// Repeated label selector argument
	Match *[]string `json:"match[],omitempty"`
}

// GetSeriesParams defines parameters for GetSeries.
type GetSeriesParams struct {
	// Repeated series selector argument
	Match externalRef1.SeriesMatcher `json:"match[]"`

	// Start timestamp
	Start *externalRef1.StartTS `json:"start,omitempty"`

	// End timestamp
	End *externalRef1.EndTS `json:"end,omitempty"`
}

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// GetLabelValues request
	GetLabelValues(ctx context.Context, tenant externalRef1.Tenant, labelName string, params *GetLabelValuesParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetLabels request
	GetLabels(ctx context.Context, tenant externalRef1.Tenant, params *GetLabelsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetInstantQuery request
	GetInstantQuery(ctx context.Context, tenant externalRef1.Tenant, params *GetInstantQueryParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetRangeQuery request
	GetRangeQuery(ctx context.Context, tenant externalRef1.Tenant, params *GetRangeQueryParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetRules request
	GetRules(ctx context.Context, tenant externalRef1.Tenant, params *GetRulesParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetRawRules request
	GetRawRules(ctx context.Context, tenant externalRef1.Tenant, reqEditors ...RequestEditorFn) (*http.Response, error)

	// SetRawRules request with any body
	SetRawRulesWithBody(ctx context.Context, tenant externalRef1.Tenant, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetSeries request
	GetSeries(ctx context.Context, tenant externalRef1.Tenant, params *GetSeriesParams, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetLabelValues(ctx context.Context, tenant externalRef1.Tenant, labelName string, params *GetLabelValuesParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetLabelValuesRequest(c.Server, tenant, labelName, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetLabels(ctx context.Context, tenant externalRef1.Tenant, params *GetLabelsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetLabelsRequest(c.Server, tenant, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetInstantQuery(ctx context.Context, tenant externalRef1.Tenant, params *GetInstantQueryParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetInstantQueryRequest(c.Server, tenant, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetRangeQuery(ctx context.Context, tenant externalRef1.Tenant, params *GetRangeQueryParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetRangeQueryRequest(c.Server, tenant, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetRules(ctx context.Context, tenant externalRef1.Tenant, params *GetRulesParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetRulesRequest(c.Server, tenant, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetRawRules(ctx context.Context, tenant externalRef1.Tenant, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetRawRulesRequest(c.Server, tenant)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) SetRawRulesWithBody(ctx context.Context, tenant externalRef1.Tenant, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewSetRawRulesRequestWithBody(c.Server, tenant, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetSeries(ctx context.Context, tenant externalRef1.Tenant, params *GetSeriesParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetSeriesRequest(c.Server, tenant, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetLabelValuesRequest generates requests for GetLabelValues
func NewGetLabelValuesRequest(server string, tenant externalRef1.Tenant, labelName string, params *GetLabelValuesParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "tenant", runtime.ParamLocationPath, tenant)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "label_name", runtime.ParamLocationPath, labelName)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/metrics/v1/%s/api/v1/label/%s/values", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if params.Match != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "match[]", runtime.ParamLocationQuery, *params.Match); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.Start != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "start", runtime.ParamLocationQuery, *params.Start); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.End != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "end", runtime.ParamLocationQuery, *params.End); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetLabelsRequest generates requests for GetLabels
func NewGetLabelsRequest(server string, tenant externalRef1.Tenant, params *GetLabelsParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "tenant", runtime.ParamLocationPath, tenant)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/metrics/v1/%s/api/v1/labels", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if params.Match != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "match[]", runtime.ParamLocationQuery, *params.Match); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.Start != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "start", runtime.ParamLocationQuery, *params.Start); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.End != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "end", runtime.ParamLocationQuery, *params.End); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetInstantQueryRequest generates requests for GetInstantQuery
func NewGetInstantQueryRequest(server string, tenant externalRef1.Tenant, params *GetInstantQueryParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "tenant", runtime.ParamLocationPath, tenant)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/metrics/v1/%s/api/v1/query", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if params.Query != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "query", runtime.ParamLocationQuery, *params.Query); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.Timeout != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "timeout", runtime.ParamLocationQuery, *params.Timeout); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.Time != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "time", runtime.ParamLocationQuery, *params.Time); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetRangeQueryRequest generates requests for GetRangeQuery
func NewGetRangeQueryRequest(server string, tenant externalRef1.Tenant, params *GetRangeQueryParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "tenant", runtime.ParamLocationPath, tenant)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/metrics/v1/%s/api/v1/query_range", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if params.Query != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "query", runtime.ParamLocationQuery, *params.Query); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.Start != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "start", runtime.ParamLocationQuery, *params.Start); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.End != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "end", runtime.ParamLocationQuery, *params.End); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.Timeout != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "timeout", runtime.ParamLocationQuery, *params.Timeout); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.Step != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "step", runtime.ParamLocationQuery, *params.Step); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetRulesRequest generates requests for GetRules
func NewGetRulesRequest(server string, tenant externalRef1.Tenant, params *GetRulesParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "tenant", runtime.ParamLocationPath, tenant)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/metrics/v1/%s/api/v1/rules", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if params.Type != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "type", runtime.ParamLocationQuery, *params.Type); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.Match != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "match[]", runtime.ParamLocationQuery, *params.Match); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetRawRulesRequest generates requests for GetRawRules
func NewGetRawRulesRequest(server string, tenant externalRef1.Tenant) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "tenant", runtime.ParamLocationPath, tenant)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/metrics/v1/%s/api/v1/rules/raw", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewSetRawRulesRequestWithBody generates requests for SetRawRules with any type of body
func NewSetRawRulesRequestWithBody(server string, tenant externalRef1.Tenant, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "tenant", runtime.ParamLocationPath, tenant)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/metrics/v1/%s/api/v1/rules/raw", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewGetSeriesRequest generates requests for GetSeries
func NewGetSeriesRequest(server string, tenant externalRef1.Tenant, params *GetSeriesParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "tenant", runtime.ParamLocationPath, tenant)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/metrics/v1/%s/api/v1/series", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "match[]", runtime.ParamLocationQuery, params.Match); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	if params.Start != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "start", runtime.ParamLocationQuery, *params.Start); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.End != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "end", runtime.ParamLocationQuery, *params.End); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// GetLabelValues request
	GetLabelValuesWithResponse(ctx context.Context, tenant externalRef1.Tenant, labelName string, params *GetLabelValuesParams, reqEditors ...RequestEditorFn) (*GetLabelValuesResponse, error)

	// GetLabels request
	GetLabelsWithResponse(ctx context.Context, tenant externalRef1.Tenant, params *GetLabelsParams, reqEditors ...RequestEditorFn) (*GetLabelsResponse, error)

	// GetInstantQuery request
	GetInstantQueryWithResponse(ctx context.Context, tenant externalRef1.Tenant, params *GetInstantQueryParams, reqEditors ...RequestEditorFn) (*GetInstantQueryResponse, error)

	// GetRangeQuery request
	GetRangeQueryWithResponse(ctx context.Context, tenant externalRef1.Tenant, params *GetRangeQueryParams, reqEditors ...RequestEditorFn) (*GetRangeQueryResponse, error)

	// GetRules request
	GetRulesWithResponse(ctx context.Context, tenant externalRef1.Tenant, params *GetRulesParams, reqEditors ...RequestEditorFn) (*GetRulesResponse, error)

	// GetRawRules request
	GetRawRulesWithResponse(ctx context.Context, tenant externalRef1.Tenant, reqEditors ...RequestEditorFn) (*GetRawRulesResponse, error)

	// SetRawRules request with any body
	SetRawRulesWithBodyWithResponse(ctx context.Context, tenant externalRef1.Tenant, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*SetRawRulesResponse, error)

	// GetSeries request
	GetSeriesWithResponse(ctx context.Context, tenant externalRef1.Tenant, params *GetSeriesParams, reqEditors ...RequestEditorFn) (*GetSeriesResponse, error)
}

type GetLabelValuesResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON2XX      *externalRef2.LabelValuesResponse
}

// Status returns HTTPResponse.Status
func (r GetLabelValuesResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetLabelValuesResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetLabelsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON2XX      *externalRef2.LabelsResponse
}

// Status returns HTTPResponse.Status
func (r GetLabelsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetLabelsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetInstantQueryResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON2XX      *externalRef2.QueryResponse
}

// Status returns HTTPResponse.Status
func (r GetInstantQueryResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetInstantQueryResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetRangeQueryResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON2XX      *externalRef2.QueryRangeResponse
}

// Status returns HTTPResponse.Status
func (r GetRangeQueryResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetRangeQueryResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetRulesResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON2XX      *externalRef2.RulesResponse
}

// Status returns HTTPResponse.Status
func (r GetRulesResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetRulesResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetRawRulesResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	YAML200      *externalRef0.RulesRaw
}

// Status returns HTTPResponse.Status
func (r GetRawRulesResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetRawRulesResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type SetRawRulesResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r SetRawRulesResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r SetRawRulesResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetSeriesResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON2XX      *externalRef2.SeriesResponse
}

// Status returns HTTPResponse.Status
func (r GetSeriesResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetSeriesResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetLabelValuesWithResponse request returning *GetLabelValuesResponse
func (c *ClientWithResponses) GetLabelValuesWithResponse(ctx context.Context, tenant externalRef1.Tenant, labelName string, params *GetLabelValuesParams, reqEditors ...RequestEditorFn) (*GetLabelValuesResponse, error) {
	rsp, err := c.GetLabelValues(ctx, tenant, labelName, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetLabelValuesResponse(rsp)
}

// GetLabelsWithResponse request returning *GetLabelsResponse
func (c *ClientWithResponses) GetLabelsWithResponse(ctx context.Context, tenant externalRef1.Tenant, params *GetLabelsParams, reqEditors ...RequestEditorFn) (*GetLabelsResponse, error) {
	rsp, err := c.GetLabels(ctx, tenant, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetLabelsResponse(rsp)
}

// GetInstantQueryWithResponse request returning *GetInstantQueryResponse
func (c *ClientWithResponses) GetInstantQueryWithResponse(ctx context.Context, tenant externalRef1.Tenant, params *GetInstantQueryParams, reqEditors ...RequestEditorFn) (*GetInstantQueryResponse, error) {
	rsp, err := c.GetInstantQuery(ctx, tenant, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetInstantQueryResponse(rsp)
}

// GetRangeQueryWithResponse request returning *GetRangeQueryResponse
func (c *ClientWithResponses) GetRangeQueryWithResponse(ctx context.Context, tenant externalRef1.Tenant, params *GetRangeQueryParams, reqEditors ...RequestEditorFn) (*GetRangeQueryResponse, error) {
	rsp, err := c.GetRangeQuery(ctx, tenant, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetRangeQueryResponse(rsp)
}

// GetRulesWithResponse request returning *GetRulesResponse
func (c *ClientWithResponses) GetRulesWithResponse(ctx context.Context, tenant externalRef1.Tenant, params *GetRulesParams, reqEditors ...RequestEditorFn) (*GetRulesResponse, error) {
	rsp, err := c.GetRules(ctx, tenant, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetRulesResponse(rsp)
}

// GetRawRulesWithResponse request returning *GetRawRulesResponse
func (c *ClientWithResponses) GetRawRulesWithResponse(ctx context.Context, tenant externalRef1.Tenant, reqEditors ...RequestEditorFn) (*GetRawRulesResponse, error) {
	rsp, err := c.GetRawRules(ctx, tenant, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetRawRulesResponse(rsp)
}

// SetRawRulesWithBodyWithResponse request with arbitrary body returning *SetRawRulesResponse
func (c *ClientWithResponses) SetRawRulesWithBodyWithResponse(ctx context.Context, tenant externalRef1.Tenant, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*SetRawRulesResponse, error) {
	rsp, err := c.SetRawRulesWithBody(ctx, tenant, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseSetRawRulesResponse(rsp)
}

// GetSeriesWithResponse request returning *GetSeriesResponse
func (c *ClientWithResponses) GetSeriesWithResponse(ctx context.Context, tenant externalRef1.Tenant, params *GetSeriesParams, reqEditors ...RequestEditorFn) (*GetSeriesResponse, error) {
	rsp, err := c.GetSeries(ctx, tenant, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetSeriesResponse(rsp)
}

// ParseGetLabelValuesResponse parses an HTTP response from a GetLabelValuesWithResponse call
func ParseGetLabelValuesResponse(rsp *http.Response) (*GetLabelValuesResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetLabelValuesResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode/100 == 2:
		var dest externalRef2.LabelValuesResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON2XX = &dest

	}

	return response, nil
}

// ParseGetLabelsResponse parses an HTTP response from a GetLabelsWithResponse call
func ParseGetLabelsResponse(rsp *http.Response) (*GetLabelsResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetLabelsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode/100 == 2:
		var dest externalRef2.LabelsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON2XX = &dest

	}

	return response, nil
}

// ParseGetInstantQueryResponse parses an HTTP response from a GetInstantQueryWithResponse call
func ParseGetInstantQueryResponse(rsp *http.Response) (*GetInstantQueryResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetInstantQueryResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode/100 == 2:
		var dest externalRef2.QueryResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON2XX = &dest

	}

	return response, nil
}

// ParseGetRangeQueryResponse parses an HTTP response from a GetRangeQueryWithResponse call
func ParseGetRangeQueryResponse(rsp *http.Response) (*GetRangeQueryResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetRangeQueryResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode/100 == 2:
		var dest externalRef2.QueryRangeResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON2XX = &dest

	}

	return response, nil
}

// ParseGetRulesResponse parses an HTTP response from a GetRulesWithResponse call
func ParseGetRulesResponse(rsp *http.Response) (*GetRulesResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetRulesResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode/100 == 2:
		var dest externalRef2.RulesResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON2XX = &dest

	}

	return response, nil
}

// ParseGetRawRulesResponse parses an HTTP response from a GetRawRulesWithResponse call
func ParseGetRawRulesResponse(rsp *http.Response) (*GetRawRulesResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetRawRulesResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "yaml") && rsp.StatusCode == 200:
		var dest externalRef0.RulesRaw
		if err := yaml.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.YAML200 = &dest

	}

	return response, nil
}

// ParseSetRawRulesResponse parses an HTTP response from a SetRawRulesWithResponse call
func ParseSetRawRulesResponse(rsp *http.Response) (*SetRawRulesResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &SetRawRulesResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// ParseGetSeriesResponse parses an HTTP response from a GetSeriesWithResponse call
func ParseGetSeriesResponse(rsp *http.Response) (*GetSeriesResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetSeriesResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode/100 == 2:
		var dest externalRef2.SeriesResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON2XX = &dest

	}

	return response, nil
}
