// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Object Storage Service API
//
// APIs for managing buckets and objects.
//

package objectstorage

import (
	"bitbucket.aka.lgl.grungy.us/golang-sdk2/common"
)

// CommitMultipartUploadPartDetails. To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// [Getting Started with Policies]({{DOC_SERVER_URL}}/Content/Identity/Concepts/policygetstarted.htm).
type CommitMultipartUploadPartDetails struct {

	// The part number for this part.
	PartNum *int `mandatory:"true" json:"partNum,omitempty"`

	// The ETag returned when this part was uploaded.
	Etag *string `mandatory:"true" json:"etag,omitempty"`
}

func (model CommitMultipartUploadPartDetails) String() string {
	return common.PointerString(model)
}