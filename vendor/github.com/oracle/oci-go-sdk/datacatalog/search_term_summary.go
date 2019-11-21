// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// DataCatalog API
//
// A description of the DataCatalog API
//

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/common"
)

// SearchTermSummary Summary of a Term associated with an object. This is a brief summary returned as part of the search result
type SearchTermSummary struct {

	// The display name of a user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Unique Term key that is immutable.
	Key *string `mandatory:"false" json:"key"`

	// Unique id of the parent Glossary.
	GlossaryKey *string `mandatory:"false" json:"glossaryKey"`

	// Name of the parent Glossary.
	GlossaryName *string `mandatory:"false" json:"glossaryName"`

	// This terms parent term key. Will be null if the term has no parent term.
	ParentTermKey *string `mandatory:"false" json:"parentTermKey"`

	// Name of the parent term key. Will be null if the term has no parent term.
	ParentTermName *string `mandatory:"false" json:"parentTermName"`
}

func (m SearchTermSummary) String() string {
	return common.PointerString(m)
}