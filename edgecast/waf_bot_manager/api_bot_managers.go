// Code generated by the Code Generation Kit (CGK) using OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
//
// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

/*
WAF API

The WAF API is a RESTful server application for managing customer configuration settings.

API version: 1.0
*/

package waf_bot_manager

import (
	"fmt"

	"github.com/EdgeCast/ec-sdk-go/edgecast/internal/ecclient"
	"github.com/go-openapi/errors"
)

// BotManagersClient is the concrete client implementation for BotManagers
type BotManagersClient struct {
	apiClient  ecclient.APIClient
	baseAPIURL string
}

// NewBotManagersClient creates a new instance of BotManagersClient
func NewBotManagersClient(
	c ecclient.APIClient,
	baseAPIURL string,
) BotManagersClient {
	return BotManagersClient{c, baseAPIURL}
}

// BotManagersClientService defines the operations for BotManagers
type BotManagersClientService interface {
	CustIdBotManagerBotManagerIdDelete(
		params CustIdBotManagerBotManagerIdDeleteParams,
	) error

	CustIdBotManagerBotManagerIdGet(
		params CustIdBotManagerBotManagerIdGetParams,
	) (*BotManager, error)

	CustIdBotManagerBotManagerIdPut(
		params CustIdBotManagerBotManagerIdPutParams,
	) error

	CustIdBotManagerGet(
		params CustIdBotManagerGetParams,
	) ([]ObjShort, error)

	CustIdBotManagerPost(
		params CustIdBotManagerPostParams,
	) (*BotManager, error)
}

// CustIdBotManagerBotManagerIdDeleteParams contains the parameters for CustIdBotManagerBotManagerIdDelete
type CustIdBotManagerBotManagerIdDeleteParams struct {
	// The customer id
	CustId string

	// The Bot Manager id
	BotManagerId string
}

// NewCustIdBotManagerBotManagerIdDeleteParams creates a new instance of CustIdBotManagerBotManagerIdDeleteParams
func NewCustIdBotManagerBotManagerIdDeleteParams() CustIdBotManagerBotManagerIdDeleteParams {
	return CustIdBotManagerBotManagerIdDeleteParams{}
}

// CustIdBotManagerBotManagerIdDelete - DELETE Bot Manager Object
//
//	Delete a Bot Manager object identified by id.
func (c BotManagersClient) CustIdBotManagerBotManagerIdDelete(
	params CustIdBotManagerBotManagerIdDeleteParams,
) error {
	req, err := buildCustIdBotManagerBotManagerIdDeleteRequest(params, c.baseAPIURL)
	if err != nil {
		return err
	}

	_, err = c.apiClient.SubmitRequest(*req)

	if err != nil {
		return fmt.Errorf("CustIdBotManagerBotManagerIdDelete: %w", err)
	}

	return nil
}

func buildCustIdBotManagerBotManagerIdDeleteRequest(
	p CustIdBotManagerBotManagerIdDeleteParams,
	baseAPIURL string,
) (*ecclient.SubmitRequestParams, error) {
	req := ecclient.NewSubmitRequestParams()
	req.Path = baseAPIURL + "/{cust_id}/bot_manager/{bot_manager_id}"
	errs := make([]error, 0)

	method, err := ecclient.ToHTTPMethod("Delete")
	if err != nil {
		errs = append(errs, fmt.Errorf("CustIdBotManagerBotManagerIdDelete: %w", err))
	}

	req.Method = method

	req.PathParams["custId"] = p.CustId

	req.PathParams["botManagerId"] = p.BotManagerId

	if len(errs) > 0 {
		return nil, errors.CompositeValidationError(errs...)
	}

	return &req, nil
}

// CustIdBotManagerBotManagerIdGetParams contains the parameters for CustIdBotManagerBotManagerIdGet
type CustIdBotManagerBotManagerIdGetParams struct {
	// The customer id
	CustId string

	// The Bot Manager id
	BotManagerId string
}

// NewCustIdBotManagerBotManagerIdGetParams creates a new instance of CustIdBotManagerBotManagerIdGetParams
func NewCustIdBotManagerBotManagerIdGetParams() CustIdBotManagerBotManagerIdGetParams {
	return CustIdBotManagerBotManagerIdGetParams{}
}

// CustIdBotManagerBotManagerIdGet - GET Bot Manager Object
//
//	Get a Bot Manager object from a Bot Manager id.
func (c BotManagersClient) CustIdBotManagerBotManagerIdGet(
	params CustIdBotManagerBotManagerIdGetParams,
) (*BotManager, error) {
	req, err := buildCustIdBotManagerBotManagerIdGetRequest(params, c.baseAPIURL)
	if err != nil {
		return nil, err
	}

	parsedResponse := BotManager{}
	req.ParsedResponse = &parsedResponse

	_, err = c.apiClient.SubmitRequest(*req)

	if err != nil {
		return nil, fmt.Errorf("CustIdBotManagerBotManagerIdGet: %w", err)
	}

	return &parsedResponse, nil
}

func buildCustIdBotManagerBotManagerIdGetRequest(
	p CustIdBotManagerBotManagerIdGetParams,
	baseAPIURL string,
) (*ecclient.SubmitRequestParams, error) {
	req := ecclient.NewSubmitRequestParams()
	req.Path = baseAPIURL + "/{cust_id}/bot_manager/{bot_manager_id}"
	errs := make([]error, 0)

	method, err := ecclient.ToHTTPMethod("Get")
	if err != nil {
		errs = append(errs, fmt.Errorf("CustIdBotManagerBotManagerIdGet: %w", err))
	}

	req.Method = method

	req.PathParams["custId"] = p.CustId

	req.PathParams["botManagerId"] = p.BotManagerId

	if len(errs) > 0 {
		return nil, errors.CompositeValidationError(errs...)
	}

	return &req, nil
}

// CustIdBotManagerBotManagerIdPutParams contains the parameters for CustIdBotManagerBotManagerIdPut
type CustIdBotManagerBotManagerIdPutParams struct {
	// The customer id
	CustId string

	// The Bot Manager id
	BotManagerId string

	BotManagerInfo BotManager
}

// NewCustIdBotManagerBotManagerIdPutParams creates a new instance of CustIdBotManagerBotManagerIdPutParams
func NewCustIdBotManagerBotManagerIdPutParams() CustIdBotManagerBotManagerIdPutParams {
	return CustIdBotManagerBotManagerIdPutParams{}
}

// CustIdBotManagerBotManagerIdPut - PUT Bot Manager Object
//
//	Modify a Bot Manager object identified by id.
func (c BotManagersClient) CustIdBotManagerBotManagerIdPut(
	params CustIdBotManagerBotManagerIdPutParams,
) error {
	req, err := buildCustIdBotManagerBotManagerIdPutRequest(params, c.baseAPIURL)
	if err != nil {
		return err
	}

	_, err = c.apiClient.SubmitRequest(*req)

	if err != nil {
		return fmt.Errorf("CustIdBotManagerBotManagerIdPut: %w", err)
	}

	return nil
}

func buildCustIdBotManagerBotManagerIdPutRequest(
	p CustIdBotManagerBotManagerIdPutParams,
	baseAPIURL string,
) (*ecclient.SubmitRequestParams, error) {
	req := ecclient.NewSubmitRequestParams()
	req.Path = baseAPIURL + "/{cust_id}/bot_manager/{bot_manager_id}"
	errs := make([]error, 0)

	method, err := ecclient.ToHTTPMethod("Put")
	if err != nil {
		errs = append(errs, fmt.Errorf("CustIdBotManagerBotManagerIdPut: %w", err))
	}

	req.Method = method

	req.PathParams["custId"] = p.CustId

	req.PathParams["botManagerId"] = p.BotManagerId

	req.RawBody = p.BotManagerInfo

	if len(errs) > 0 {
		return nil, errors.CompositeValidationError(errs...)
	}

	return &req, nil
}

// CustIdBotManagerGetParams contains the parameters for CustIdBotManagerGet
type CustIdBotManagerGetParams struct {
	// The customer id
	CustId string
}

// NewCustIdBotManagerGetParams creates a new instance of CustIdBotManagerGetParams
func NewCustIdBotManagerGetParams() CustIdBotManagerGetParams {
	return CustIdBotManagerGetParams{}
}

// CustIdBotManagerGet - GET Bot Managers
//
//	List all Bot Managers for a given customer.
func (c BotManagersClient) CustIdBotManagerGet(
	params CustIdBotManagerGetParams,
) ([]ObjShort, error) {
	req, err := buildCustIdBotManagerGetRequest(params, c.baseAPIURL)
	if err != nil {
		return nil, err
	}

	parsedResponse := make([]ObjShort, 0)
	req.ParsedResponse = &parsedResponse

	_, err = c.apiClient.SubmitRequest(*req)

	if err != nil {
		return nil, fmt.Errorf("CustIdBotManagerGet: %w", err)
	}

	return parsedResponse, nil
}

func buildCustIdBotManagerGetRequest(
	p CustIdBotManagerGetParams,
	baseAPIURL string,
) (*ecclient.SubmitRequestParams, error) {
	req := ecclient.NewSubmitRequestParams()
	req.Path = baseAPIURL + "/{cust_id}/bot_manager"
	errs := make([]error, 0)

	method, err := ecclient.ToHTTPMethod("Get")
	if err != nil {
		errs = append(errs, fmt.Errorf("CustIdBotManagerGet: %w", err))
	}

	req.Method = method

	req.PathParams["custId"] = p.CustId

	if len(errs) > 0 {
		return nil, errors.CompositeValidationError(errs...)
	}

	return &req, nil
}

// CustIdBotManagerPostParams contains the parameters for CustIdBotManagerPost
type CustIdBotManagerPostParams struct {
	// The customer id
	CustId string

	BotManagerInfo BotManager
}

// NewCustIdBotManagerPostParams creates a new instance of CustIdBotManagerPostParams
func NewCustIdBotManagerPostParams() CustIdBotManagerPostParams {
	return CustIdBotManagerPostParams{}
}

// CustIdBotManagerPost - POST Bot Manager
//
//	Create a new Bot Manager for a given customer.
func (c BotManagersClient) CustIdBotManagerPost(
	params CustIdBotManagerPostParams,
) (*BotManager, error) {
	req, err := buildCustIdBotManagerPostRequest(params, c.baseAPIURL)
	if err != nil {
		return nil, err
	}

	parsedResponse := BotManager{}
	req.ParsedResponse = &parsedResponse

	_, err = c.apiClient.SubmitRequest(*req)

	if err != nil {
		return nil, fmt.Errorf("CustIdBotManagerPost: %w", err)
	}

	return &parsedResponse, nil
}

func buildCustIdBotManagerPostRequest(
	p CustIdBotManagerPostParams,
	baseAPIURL string,
) (*ecclient.SubmitRequestParams, error) {
	req := ecclient.NewSubmitRequestParams()
	req.Path = baseAPIURL + "/{cust_id}/bot_manager"
	errs := make([]error, 0)

	method, err := ecclient.ToHTTPMethod("Post")
	if err != nil {
		errs = append(errs, fmt.Errorf("CustIdBotManagerPost: %w", err))
	}

	req.Method = method

	req.PathParams["custId"] = p.CustId

	req.RawBody = p.BotManagerInfo

	if len(errs) > 0 {
		return nil, errors.CompositeValidationError(errs...)
	}

	return &req, nil
}
