// Copyright 2016 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may not
// use this file except in compliance with the License. A copy of the
// License is located at
//
// http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
// either express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Package service wraps SSM service
package service

import (
	"testing"

	"github.com/aws/amazon-ssm-agent/agent/contracts"
	"github.com/aws/amazon-ssm-agent/agent/log"
	"github.com/aws/amazon-ssm-agent/agent/sdkutil"
	ssmSvc "github.com/aws/amazon-ssm-agent/agent/ssm"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	instanceID = "i-test"
)

var ssmMock = ssmSvc.NewMockDefault()
var logMock = log.NewMockLog()

func TestListAssociations(t *testing.T) {
	service := AssociationService{
		ssmSvc:     ssmMock,
		stopPolicy: &sdkutil.StopPolicy{},
	}

	associationName := "test"
	association := ssm.InstanceAssociationSummary{
		Name: &associationName,
	}

	output := ssm.ListInstanceAssociationsOutput{
		Associations: []*ssm.InstanceAssociationSummary{&association},
	}
	getDocumentOutput := ssm.GetDocumentOutput{
		Name: &associationName,
	}

	ssmMock.On("ListInstanceAssociations", mock.AnythingOfType("*log.Mock"), mock.AnythingOfType("string")).Return(&output, nil)
	ssmMock.On("GetDocument", mock.AnythingOfType("*log.Mock"), mock.AnythingOfType("string")).Return(&getDocumentOutput, nil)

	_, err := service.ListInstanceAssociations(logMock, instanceID)

	assert.NoError(t, err)
}

//func TestLoadAssociationDetails(t *testing.T) {
//	service := AssociationService{
//		ssmSvc:     ssmMock,
//		stopPolicy: &sdkutil.StopPolicy{},
//	}
//
//	associationName := "test"
//	documentContent := "document content"
//	assocRawData := model.AssociationRawData{}
//	assocRawData.Association = &ssm.Association{}
//	assocRawData.Association.Name = &associationName
//	assocRawData.Association.InstanceId = &instanceID
//
//	getDocumentOutput := ssm.GetDocumentOutput{
//		Content: &documentContent,
//	}
//
//	associationOutput := ssm.DescribeAssociationOutput{
//		AssociationDescription: &ssm.AssociationDescription{},
//	}
//
//	ssmMock.On("GetDocument", mock.AnythingOfType("*log.Mock"), mock.AnythingOfType("string")).Return(&getDocumentOutput, nil)
//	ssmMock.On("DescribeAssociation", mock.AnythingOfType("*log.Mock"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(&associationOutput, nil)
//
//	err := service.LoadAssociationDetail(logMock, &assocRawData)
//
//	assert.NoError(t, err)
//	assert.NotNil(t, assocRawData.Parameter)
//}

func TestUpdateAssociationStatus(t *testing.T) {
	service := AssociationService{
		ssmSvc:     ssmMock,
		stopPolicy: &sdkutil.StopPolicy{},
	}

	associationName := "test"
	status := ssm.AssociationStatusNamePending
	output := ssm.UpdateAssociationStatusOutput{
		AssociationDescription: &ssm.AssociationDescription{
			Status: &ssm.AssociationStatus{
				Name: &status,
			},
		},
	}
	info := contracts.AgentInfo{}

	ssmMock.On("UpdateAssociationStatus",
		mock.AnythingOfType("*log.Mock"),
		mock.AnythingOfType("string"),
		mock.AnythingOfType("string"),
		mock.AnythingOfType("*ssm.AssociationStatus")).Return(&output, nil)

	result, err := service.UpdateAssociationStatus(
		logMock,
		instanceID,
		associationName,
		status,
		"TestMessage",
		&info)

	assert.NotNil(t, result)
	assert.NoError(t, err)
	assert.Equal(t, *result.AssociationDescription.Status.Name, status)
}
