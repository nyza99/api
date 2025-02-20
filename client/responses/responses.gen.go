// Package responses provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.0 DO NOT EDIT.
package responses

import (
	externalRef0 "github.com/observatorium/api/client/models"
)

// LabelValuesResponse defines model for LabelValuesResponse.
type LabelValuesResponse struct {
	Data      *[]string `json:"data,omitempty"`
	Error     *string   `json:"error,omitempty"`
	ErrorType *string   `json:"errorType,omitempty"`
	Status    string    `json:"status"`
	Warnings  *[]string `json:"warnings,omitempty"`
}

// LabelsResponse defines model for LabelsResponse.
type LabelsResponse struct {
	Data      *[]string `json:"data,omitempty"`
	Error     *string   `json:"error,omitempty"`
	ErrorType *string   `json:"errorType,omitempty"`
	Status    string    `json:"status"`
	Warnings  *[]string `json:"warnings,omitempty"`
}

// QueryRangeResponse defines model for QueryRangeResponse.
type QueryRangeResponse struct {
	Data      *externalRef0.RangeQueryResponse `json:"data,omitempty"`
	Error     *string                          `json:"error,omitempty"`
	ErrorType *string                          `json:"errorType,omitempty"`
	Status    string                           `json:"status"`
	Warnings  *[]string                        `json:"warnings,omitempty"`
}

// QueryResponse defines model for QueryResponse.
type QueryResponse struct {
	Data      *externalRef0.InstantQueryResponse `json:"data,omitempty"`
	Error     *string                            `json:"error,omitempty"`
	ErrorType *string                            `json:"errorType,omitempty"`
	Status    string                             `json:"status"`
	Warnings  *[]string                          `json:"warnings,omitempty"`
}

// RulesResponse defines model for RulesResponse.
type RulesResponse struct {
	Data      *externalRef0.Rules `json:"data,omitempty"`
	Error     *string             `json:"error,omitempty"`
	ErrorType *string             `json:"errorType,omitempty"`
	Status    string              `json:"status"`
	Warnings  *[]string           `json:"warnings,omitempty"`
}

// SeriesResponse defines model for SeriesResponse.
type SeriesResponse struct {
	Data *[]struct {
		AdditionalProperties map[string]string `json:"-"`
	} `json:"data,omitempty"`
	Error     *string   `json:"error,omitempty"`
	ErrorType *string   `json:"errorType,omitempty"`
	Status    string    `json:"status"`
	Warnings  *[]string `json:"warnings,omitempty"`
}
