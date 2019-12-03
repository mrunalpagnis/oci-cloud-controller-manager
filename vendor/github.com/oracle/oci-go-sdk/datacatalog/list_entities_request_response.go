// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListEntitiesRequest wrapper for the ListEntities operation
type ListEntitiesRequest struct {

	// unique Catalog identifier
	CatalogId *string `mandatory:"true" contributesTo:"path" name:"catalogId"`

	// Unique Data Asset key.
	DataAssetKey *string `mandatory:"true" contributesTo:"path" name:"dataAssetKey"`

	// A filter to return only resources that match the entire display name given. The match is not case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only resources that match the specified lifecycle state. The value is case insensitive.
	LifecycleState LifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Time that the Resource was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeCreated"`

	// Time that the Resource was updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeUpdated"`

	// Id (OCID) of the user who created the resource.
	CreatedById *string `mandatory:"false" contributesTo:"query" name:"createdById"`

	// Id of the user who updated the resource.
	UpdatedById *string `mandatory:"false" contributesTo:"query" name:"updatedById"`

	// Unique external identifier of this resource in the external source system.
	ExternalKey *string `mandatory:"false" contributesTo:"query" name:"externalKey"`

	// Last modified timestamp of this object in the external system.
	TimeExternal *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeExternal"`

	// Time that the resource's status was last updated. An RFC3339 formatted datetime string.
	TimeStatusUpdated *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeStatusUpdated"`

	// Identifies if the object is a physical object (materialized) or virtual/logical object defined on other objects.
	IsLogical *bool `mandatory:"false" contributesTo:"query" name:"isLogical"`

	// Identifies if an object is a sub object (partition) of a physical or materialized parent object.
	IsPartition *bool `mandatory:"false" contributesTo:"query" name:"isPartition"`

	// Key of the associated folder.
	FolderKey *string `mandatory:"false" contributesTo:"query" name:"folderKey"`

	// Full path of the resource for resources that support paths.
	Path *string `mandatory:"false" contributesTo:"query" name:"path"`

	// Harvest Status of the harvestable resource as updated by the harvest process.
	HarvestStatus HarvestStatusEnum `mandatory:"false" contributesTo:"query" name:"harvestStatus" omitEmpty:"true"`

	// Key of the last harvest process to update this resource.
	LastJobKey *string `mandatory:"false" contributesTo:"query" name:"lastJobKey"`

	// Used to control which fields are returned in an Entity summary response.
	Fields []ListEntitiesFieldsEnum `contributesTo:"query" name:"fields" omitEmpty:"true" collectionFormat:"multi"`

	// The field to sort by. Only one sort order may be provided. Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. If no value is specified TIMECREATED is default.
	SortBy ListEntitiesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListEntitiesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListEntitiesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListEntitiesRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListEntitiesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListEntitiesResponse wrapper for the ListEntities operation
type ListEntitiesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []EntitySummary instances
	Items []EntitySummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListEntitiesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListEntitiesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListEntitiesFieldsEnum Enum with underlying type: string
type ListEntitiesFieldsEnum string

// Set of constants representing the allowable values for ListEntitiesFieldsEnum
const (
	ListEntitiesFieldsKey            ListEntitiesFieldsEnum = "key"
	ListEntitiesFieldsDisplayname    ListEntitiesFieldsEnum = "displayName"
	ListEntitiesFieldsDescription    ListEntitiesFieldsEnum = "description"
	ListEntitiesFieldsDataassetkey   ListEntitiesFieldsEnum = "dataAssetKey"
	ListEntitiesFieldsTimecreated    ListEntitiesFieldsEnum = "timeCreated"
	ListEntitiesFieldsLifecyclestate ListEntitiesFieldsEnum = "lifecycleState"
	ListEntitiesFieldsFolderkey      ListEntitiesFieldsEnum = "folderKey"
	ListEntitiesFieldsExternalkey    ListEntitiesFieldsEnum = "externalKey"
	ListEntitiesFieldsPath           ListEntitiesFieldsEnum = "path"
	ListEntitiesFieldsUri            ListEntitiesFieldsEnum = "uri"
)

var mappingListEntitiesFields = map[string]ListEntitiesFieldsEnum{
	"key":            ListEntitiesFieldsKey,
	"displayName":    ListEntitiesFieldsDisplayname,
	"description":    ListEntitiesFieldsDescription,
	"dataAssetKey":   ListEntitiesFieldsDataassetkey,
	"timeCreated":    ListEntitiesFieldsTimecreated,
	"lifecycleState": ListEntitiesFieldsLifecyclestate,
	"folderKey":      ListEntitiesFieldsFolderkey,
	"externalKey":    ListEntitiesFieldsExternalkey,
	"path":           ListEntitiesFieldsPath,
	"uri":            ListEntitiesFieldsUri,
}

// GetListEntitiesFieldsEnumValues Enumerates the set of values for ListEntitiesFieldsEnum
func GetListEntitiesFieldsEnumValues() []ListEntitiesFieldsEnum {
	values := make([]ListEntitiesFieldsEnum, 0)
	for _, v := range mappingListEntitiesFields {
		values = append(values, v)
	}
	return values
}

// ListEntitiesSortByEnum Enum with underlying type: string
type ListEntitiesSortByEnum string

// Set of constants representing the allowable values for ListEntitiesSortByEnum
const (
	ListEntitiesSortByTimecreated ListEntitiesSortByEnum = "TIMECREATED"
	ListEntitiesSortByDisplayname ListEntitiesSortByEnum = "DISPLAYNAME"
)

var mappingListEntitiesSortBy = map[string]ListEntitiesSortByEnum{
	"TIMECREATED": ListEntitiesSortByTimecreated,
	"DISPLAYNAME": ListEntitiesSortByDisplayname,
}

// GetListEntitiesSortByEnumValues Enumerates the set of values for ListEntitiesSortByEnum
func GetListEntitiesSortByEnumValues() []ListEntitiesSortByEnum {
	values := make([]ListEntitiesSortByEnum, 0)
	for _, v := range mappingListEntitiesSortBy {
		values = append(values, v)
	}
	return values
}

// ListEntitiesSortOrderEnum Enum with underlying type: string
type ListEntitiesSortOrderEnum string

// Set of constants representing the allowable values for ListEntitiesSortOrderEnum
const (
	ListEntitiesSortOrderAsc  ListEntitiesSortOrderEnum = "ASC"
	ListEntitiesSortOrderDesc ListEntitiesSortOrderEnum = "DESC"
)

var mappingListEntitiesSortOrder = map[string]ListEntitiesSortOrderEnum{
	"ASC":  ListEntitiesSortOrderAsc,
	"DESC": ListEntitiesSortOrderDesc,
}

// GetListEntitiesSortOrderEnumValues Enumerates the set of values for ListEntitiesSortOrderEnum
func GetListEntitiesSortOrderEnumValues() []ListEntitiesSortOrderEnum {
	values := make([]ListEntitiesSortOrderEnum, 0)
	for _, v := range mappingListEntitiesSortOrder {
		values = append(values, v)
	}
	return values
}