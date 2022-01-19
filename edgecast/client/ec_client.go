package client

/*
	This file contains the concrete client implementation for the EC SDK
*/

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/EdgeCast/ec-sdk-go/edgecast/auth"
	"github.com/EdgeCast/ec-sdk-go/edgecast/eclog"
	"github.com/EdgeCast/ec-sdk-go/edgecast/internal/collectionhelper"
	"github.com/EdgeCast/ec-sdk-go/edgecast/internal/jsonhelper"
)

const (
	defaultHeaderAccept      string = "application/json"
	defaultHeaderContentType string = "application/json"
)

// ECClient -
type ECClient struct {
	reqBuilder requestBuilder
	reqSender  requestSender
	config     ClientConfig
}

// NewECClient creates a default instance of ECClient using the provided
// configuration
func NewECClient(config ClientConfig) ECClient {
	return ECClient{
		reqBuilder: newECRequestBuilder(config),
		reqSender:  newECRequestSender(config),
		config:     config,
	}
}

// Do invokes an HTTP request with the given parameters
func (c ECClient) SubmitRequest(params SubmitRequestParams) (*Response, error) {
	req, err := c.reqBuilder.buildRequest(buildRequestParams{
		method:      params.Method.String(),
		path:        params.Path,
		rawBody:     params.Body,
		queryParams: params.QueryParams,
		pathParams:  params.PathParams,
		userAgent:   c.config.UserAgent,
	})

	// Provides an object to be filled in when unmarshaling the API response
	req.parsedResponse = params.ParsedResponse

	if err != nil {
		return nil, fmt.Errorf("ECClient.Do: %v", err)
	}

	c.config.Logger.Debug(
		"[REQUEST-URI]:[%s] %s\n",
		req.method,
		req.url.String())
	c.config.Logger.Debug("[REQUEST-BODY]:%v\n", req.rawBody)
	c.config.Logger.Debug("[REQUEST-HEADERS]:%+v\n", req.headers)

	resp, err := c.reqSender.sendRequest(*req)
	if err != nil {
		return nil, fmt.Errorf("ECClient.Do: %v", err)
	}
	bodyAsString, _ := jsonhelper.ConvertToJSONString(resp.Data, true)
	c.config.Logger.Debug("[RESPONSE-BODY]:%s\n", bodyAsString)
	return resp, nil
}

// ecRequestBuilder builds requests to be sent to the Edgecast API
type ecRequestBuilder struct {
	baseAPIURL   url.URL
	authProvider *auth.AuthorizationProvider
	userAgent    string
	logger       eclog.Logger
}

// newECRequestBuilder creates a default instance of ecRequestBuilder using the
// provided configuration
func newECRequestBuilder(config ClientConfig) ecRequestBuilder {
	return ecRequestBuilder{
		baseAPIURL:   config.BaseAPIURL,
		authProvider: &config.AuthProvider,
		userAgent:    config.UserAgent,
		logger:       config.Logger,
	}
}

// buildRequest creates a new Request for the Edgecast API with query params,
// adding appropriate headers
func (eb ecRequestBuilder) buildRequest(
	params buildRequestParams,
) (*Request, error) {
	relativeURL, err := url.Parse(params.path)
	if err != nil {
		return nil,
			fmt.Errorf("ecRequestBuilder.buildRequest: url.Parse: %v", err)
	}

	req := Request{
		method:  params.method,
		url:     eb.baseAPIURL.ResolveReference(relativeURL),
		headers: make(map[string]string),
	}

	err = req.setPathParams(params.pathParams)
	if err != nil {
		return nil,
			fmt.Errorf("ecRequestBuilder.buildRequest: %v", err)
	}
	req.setQueryParams(params.queryParams)

	req.headers["User-Agent"] = params.userAgent
	req.headers["Accept"] = defaultHeaderAccept

	if params.rawBody != nil {
		err := req.setBody(params.rawBody)
		if err != nil {
			return nil, fmt.Errorf(
				"ecRequestBuilder.buildRequest: %v", err)
		}
	}

	if eb.authProvider != nil {
		err := req.setAuthorization(*eb.authProvider)
		if err != nil {
			return nil, fmt.Errorf(
				"ecRequestBuilder.buildRequest: %v", err)
		}
	}

	return &req, nil
}

func (req *Request) setPathParams(params map[string]string) error {
	// Apply path parameters
	// e.g.
	// path = "/customers/{customer_id}/policies/{policy_id}""
	// params = { "customer_id": "1", "policy_id": "99" }
	// result = "customers/1/policies/99"
	for k, v := range params {
		searchKey := fmt.Sprintf("{%s}", k)

		if !strings.Contains(req.url.Path, searchKey) {
			return fmt.Errorf(
				"Request.setPathParams: param not found in path: %s",
				k)
		}

		req.url.Path = strings.Replace(
			req.url.Path,
			searchKey,
			fmt.Sprintf("%v", v),
			-1)
	}
	return nil
}

func (req *Request) setQueryParams(queryParams map[string]string) {
	// Adding Query Params
	query := req.url.Query()
	for k, v := range queryParams {
		query.Add(k, v)
	}
	// Encode the parameters and set the URL
	req.url.RawQuery = query.Encode()
}

func (req *Request) setBody(rawBody interface{}) error {
	if req.headers == nil {
		req.headers = make(map[string]string)
	}
	switch b := rawBody.(type) {
	case string:
		req.rawBody = []byte(b)
		req.headers["Content-Type"] = "text/plain; charset=utf-8"
		req.headers["Accept"] = "application/json, text/html"
	default:
		buf := new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(rawBody)
		if err != nil {
			return err
		}
		req.rawBody = buf
		req.headers["Accept"] = defaultHeaderAccept
		req.headers["Content-Type"] = defaultHeaderContentType
	}
	return nil
}

func (req *Request) setAuthorization(
	auth auth.AuthorizationProvider,
) error {
	authHeader, err := auth.GetAuthorizationHeader()
	if err != nil {
		return fmt.Errorf(
			"request.setAuthorization: failed to get authorization: %v",
			err)
	}
	if req.headers == nil {
		req.headers = make(map[string]string)
	}
	req.headers["Authorization"] = authHeader
	return nil
}

// ecRequestSender sends requests to the Edgecast API
type ecRequestSender struct {
	clientAdapter clientAdapter
	logger        eclog.Logger
}

// newECRequestSender creates a default instance of ecRequestSender
func newECRequestSender(config ClientConfig) ecRequestSender {
	return ecRequestSender{
		clientAdapter: newRetryableHTTPClientAdapter(config),
		logger:        config.Logger,
	}
}

// literalResponse is used for unmarshaling response data
// that is in an unrecognized format
type literalResponse struct {
	value interface{}
}

// sendRequest sends a Request and returns the Response.
// If Request.ParsedResponse is non-nil, then the response body will be
// unmarshaled to it.
// Response.Data will always have the unmarshaled response body. Note that if
// Request.ParsedResponse is nil, then Response.Data will be a map[string]string
// as a result of unmarshaling JSON.
func (es ecRequestSender) sendRequest(req Request) (*Response, error) {
	httpResp, err := es.clientAdapter.do(req)
	if err != nil {
		return nil, fmt.Errorf("SendRequest:%v", err)
	}

	defer httpResp.Body.Close()
	body, err := ioutil.ReadAll(httpResp.Body)
	bodyAsString := string(body)
	if httpResp.StatusCode >= 400 && httpResp.StatusCode <= 599 {
		if err != nil {
			return nil,
				fmt.Errorf(
					"ecRequestSender.sendRequest: ioutil.ReadAll: %v",
					err)
		}
		return nil,
			fmt.Errorf(
				"ecRequestSender.sendRequest failed: %s", bodyAsString)
	}

	var temp interface{}
	if jsonUnmarshalErr := json.Unmarshal(body, &temp); err != nil {
		return nil,
			fmt.Errorf(
				"ecRequestSender.sendRequest: malformed JSON response: %v",
				jsonUnmarshalErr)
	}

	if collectionhelper.IsInterfaceArray(temp) {
		if err := json.Unmarshal([]byte(body), &req.parsedResponse); err != nil {
			return nil,
				fmt.Errorf(
					"ecRequestSender.sendRequest: malformed Json Array response:%v",
					err)
		}
	} else {
		if jsonhelper.IsJSONString(bodyAsString) {
			err = json.Unmarshal([]byte(bodyAsString), &req.parsedResponse)
			if err != nil {
				return nil, fmt.Errorf(
					"ecRequestSender.sendRequest: Decode error: %v",
					err)
			}
		} else {
			// if response is not json string
			switch v := req.parsedResponse.(type) {
			case literalResponse:
				rs, ok := req.parsedResponse.(literalResponse)
				if ok {
					rs.value = bodyAsString
					req.parsedResponse = rs
				}
			case float64:
				fmt.Println("float64:", v)
			default:
				fmt.Println("unknown")
			}
		}
	}
	return &Response{
		Data:         req.parsedResponse,
		HTTPResponse: httpResp,
	}, nil
}

// sendRequestWithStringResponse sends an HTTP request and returns the response
// body as a string
func (es ecRequestSender) sendRequestWithStringResponse(
	req Request,
) (*string, error) {
	httpResp, err := es.clientAdapter.do(req)
	if err != nil {
		return nil, fmt.Errorf(
			"ecRequestSender.sendRequestWithStringResponse: %v",
			err)
	}
	defer httpResp.Body.Close()
	body, err := ioutil.ReadAll(httpResp.Body)
	bodyAsString := string(body)
	if httpResp.StatusCode >= 400 && httpResp.StatusCode <= 599 {
		if err != nil {
			return nil, fmt.Errorf(
				"ecRequestSender.sendRequestWithStringResponse: ioutil.ReadAll: %v",
				err)
		}

		return nil, fmt.Errorf(
			"ecRequestSender.sendRequestWithStringResponse failed: %s",
			bodyAsString)
	}
	return &bodyAsString, nil
}
