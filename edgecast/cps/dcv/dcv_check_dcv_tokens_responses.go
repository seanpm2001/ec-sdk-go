// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package dcv

// This file was generated by codegen-sdk-go.
// Any changes made to this file may be overwritten.

import (
	"github.com/EdgeCast/ec-sdk-go/edgecast/cps/models"
)

// NewDcvCheckDcvTokensBadRequest creates a DcvCheckDcvTokensBadRequest with default headers values
func NewDcvCheckDcvTokensBadRequest() *DcvCheckDcvTokensBadRequest {
	return &DcvCheckDcvTokensBadRequest{}
}

/* DcvCheckDcvTokensBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DcvCheckDcvTokensBadRequest struct {
	models.HyperionErrorReponse
}

// NewDcvCheckDcvTokensNotFound creates a DcvCheckDcvTokensNotFound with default headers values
func NewDcvCheckDcvTokensNotFound() *DcvCheckDcvTokensNotFound {
	return &DcvCheckDcvTokensNotFound{}
}

/* DcvCheckDcvTokensNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DcvCheckDcvTokensNotFound struct {
	models.HyperionErrorReponse
}
