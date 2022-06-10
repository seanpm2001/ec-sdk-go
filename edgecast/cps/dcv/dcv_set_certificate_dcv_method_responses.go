// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package dcv

// This file was generated by codegen-sdk-go.
// Any changes made to this file may be overwritten.

import (
	"github.com/EdgeCast/ec-sdk-go/edgecast/cps/models"
)

// NewDcvSetCertificateDcvMethodNoContent creates a DcvSetCertificateDcvMethodNoContent with default headers values
func NewDcvSetCertificateDcvMethodNoContent() *DcvSetCertificateDcvMethodNoContent {
	return &DcvSetCertificateDcvMethodNoContent{}
}

/* DcvSetCertificateDcvMethodNoContent describes a response with status code 204, with default header values.

Success
*/
type DcvSetCertificateDcvMethodNoContent struct {
}

// NewDcvSetCertificateDcvMethodBadRequest creates a DcvSetCertificateDcvMethodBadRequest with default headers values
func NewDcvSetCertificateDcvMethodBadRequest() *DcvSetCertificateDcvMethodBadRequest {
	return &DcvSetCertificateDcvMethodBadRequest{}
}

/* DcvSetCertificateDcvMethodBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DcvSetCertificateDcvMethodBadRequest struct {
	models.HyperionErrorReponse
}

// NewDcvSetCertificateDcvMethodNotFound creates a DcvSetCertificateDcvMethodNotFound with default headers values
func NewDcvSetCertificateDcvMethodNotFound() *DcvSetCertificateDcvMethodNotFound {
	return &DcvSetCertificateDcvMethodNotFound{}
}

/* DcvSetCertificateDcvMethodNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DcvSetCertificateDcvMethodNotFound struct {
	models.HyperionErrorReponse
}