package test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"records-app/internal/adapters/database"
	"records-app/internal/adapters/database/schemas"
	"records-app/internal/test/mock"
)

func TestGetAllRecordsDB_Success(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDB := mock.NewMockDB(mockCtrl)

	expectedRecords := []schemas.Records{{ID: 1, Title: "Record 1", Description: "Description 1"}}
	mockDB.EXPECT().GetAllRecordsDB(gomock.Any()).Return(expectedRecords, nil)

	testAdapter := &database.AdapterMethods{
		IAdapterRecordMethods: mockDB,
	}

	recordORM, _ := testAdapter.IAdapterRecordMethods.GetAllRecordsDB(context.Background())
	assert.Equal(t, expectedRecords, recordORM, "Returned records do not match expected records")
}

func TestGetAllRecordsDB_Failure(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDB := mock.NewMockDB(mockCtrl)

	expectedRecords := []schemas.Records{{ID: 1, Title: "Record 1", Description: "Description 1"}}
	unexpectedRecords := []schemas.Records{{ID: 2, Title: "Record 2", Description: "Description 2"}}
	mockDB.EXPECT().GetAllRecordsDB(gomock.Any()).Return(unexpectedRecords, nil)

	testAdapter := &database.AdapterMethods{
		IAdapterRecordMethods: mockDB,
	}

	recordORM, _ := testAdapter.IAdapterRecordMethods.GetAllRecordsDB(context.Background())
	assert.NotEqual(t, expectedRecords, recordORM, "Returned records match expected records")
}
