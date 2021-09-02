// Copyright Edgecast, Licensed under the terms of the Apache 2.0 license. See LICENSE file in project root for terms.

package auth

// AuthorizationProvider defines structs that can provide Authorization headers
type AuthorizationProvider interface {
	GetAuthorizationHeader() (string, error)
}
