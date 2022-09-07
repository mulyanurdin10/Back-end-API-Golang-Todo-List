package business

import (
	"errors"
	"testcode/features/activitys"
	"testcode/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAllData(t *testing.T) {
	repo := new(mocks.ActivityData)
	GetAllData := []activitys.Core{{ID: 1, Email: "mulyanurdin10@gmail.com", Title: "test-1"}}

	t.Run("Success Get All Data", func(t *testing.T) {
		repo.On("GetAllData").Return(GetAllData, nil).Once()

		srv := NewActivityBusiness(repo)

		res, err := srv.GetAllData()
		assert.NoError(t, err)
		assert.Equal(t, GetAllData[0].ID, res[0].ID)
		repo.AssertExpectations(t)
	})

	t.Run("Error Get All Data", func(t *testing.T) {
		repo.On("GetAllData").Return(nil, errors.New("Failed to get all data")).Once()

		srv := NewActivityBusiness(repo)

		res, err := srv.GetAllData()
		assert.Error(t, err)
		assert.Nil(t, res)
		repo.AssertExpectations(t)
	})
}

func TestGetData(t *testing.T) {
	repo := new(mocks.ActivityData)
	GetData := activitys.Core{ID: 1, Email: "mulyanurdin10@gmail.com", Title: "test-1"}

	t.Run("Success Get Data", func(t *testing.T) {
		repo.On("GetData", mock.Anything).Return(GetData, 1, nil).Once()

		srv := NewActivityBusiness(repo)

		data, res, err := srv.GetData(1)
		assert.NoError(t, err)
		assert.Equal(t, GetData.ID, res, data)
		repo.AssertExpectations(t)
	})

	t.Run("Error Get Data", func(t *testing.T) {
		repo.On("GetData", mock.Anything).Return(GetData, 0, errors.New("Failed to get data")).Once()

		srv := NewActivityBusiness(repo)

		data, res, err := srv.GetData(2)
		assert.Error(t, err)
		assert.Equal(t, 0, res, data)
		repo.AssertExpectations(t)
	})
}

func TestInsertData(t *testing.T) {
	repo := new(mocks.ActivityData)
	insertData := activitys.Core{ID: 1, Email: "mulyanurdin10@gmail.com", Title: "test-1"}

	t.Run("Success Insert Data", func(t *testing.T) {
		repo.On("InsertData", mock.Anything).Return(insertData, 1, nil).Once()
		repo.On("UniqueData", mock.Anything).Return(0, nil).Once()
		srv := NewActivityBusiness(repo)

		data, res, err := srv.InsertData(insertData)
		assert.NoError(t, err)
		assert.Equal(t, 1, res, data)
		repo.AssertExpectations(t)
	})

	t.Run("Validasi Insert Data Email", func(t *testing.T) {
		srv := NewActivityBusiness(repo)
		validasiInsertData := activitys.Core{ID: 1, Email: "mulyanurdin10@gmail", Title: "test-1"}
		data, res, err := srv.InsertData(validasiInsertData)
		assert.NotNil(t, err)
		assert.Equal(t, -1, res, data)
		repo.AssertExpectations(t)
	})

	t.Run("Validasi Insert Data Title", func(t *testing.T) {
		srv := NewActivityBusiness(repo)
		validasiInsertData := activitys.Core{ID: 1, Email: "mulyanurdin10@gmail.com", Title: ""}
		data, res, err := srv.InsertData(validasiInsertData)
		assert.NotNil(t, err)
		assert.Equal(t, -1, res, data)
		repo.AssertExpectations(t)
	})

	t.Run("Validasi Insert Data Email Unique", func(t *testing.T) {
		repo.On("UniqueData", mock.Anything).Return(1, nil, errors.New("Email already exists")).Once()
		srv := NewActivityBusiness(repo)

		data, res, err := srv.InsertData(insertData)
		assert.Error(t, err)
		assert.Equal(t, -1, res, data)
		repo.AssertExpectations(t)
	})
}

func TestUpdateData(t *testing.T) {
	repo := new(mocks.ActivityData)
	updateData := activitys.Core{ID: 1, Email: "mulyanurdin10@gmail.com", Title: "test-1"}

	t.Run("Success Update Data", func(t *testing.T) {
		repo.On("UpdateData", mock.Anything, mock.Anything).Return(updateData, 1, nil).Once()
		repo.On("GetData", mock.Anything).Return(updateData, 1, nil).Once()
		repo.On("UniqueData", mock.Anything).Return(0, nil).Once()

		srv := NewActivityBusiness(repo)

		data, res, err := srv.UpdateData(1, updateData)
		assert.NoError(t, err)
		assert.Equal(t, 1, res, data)
		repo.AssertExpectations(t)
	})

	t.Run("Validasi Insert Data Email", func(t *testing.T) {
		srv := NewActivityBusiness(repo)
		validasiInsertData := activitys.Core{ID: 1, Email: "mulyanurdin10@gmail", Title: "test-1"}
		repo.On("GetData", mock.Anything).Return(updateData, 1, nil).Once()
		data, res, err := srv.UpdateData(1, validasiInsertData)
		assert.NotNil(t, err)
		assert.Equal(t, -1, res, data)
		repo.AssertExpectations(t)
	})

	t.Run("Validasi Insert Data Email Unique", func(t *testing.T) {
		repo.On("GetData", mock.Anything).Return(updateData, 1, nil).Once()
		repo.On("UniqueData", mock.Anything).Return(1, nil, errors.New("Email already exists")).Once()
		srv := NewActivityBusiness(repo)

		data, res, err := srv.UpdateData(1, updateData)
		assert.Error(t, err)
		assert.Equal(t, -1, res, data)
		repo.AssertExpectations(t)
	})

	t.Run("Validasi Insert Data Email Empty", func(t *testing.T) {
		srv := NewActivityBusiness(repo)
		validasiInsertData := activitys.Core{ID: 1, Email: "", Title: "test-1"}
		repo.On("UpdateData", mock.Anything, mock.Anything).Return(validasiInsertData, 1, nil).Once()
		repo.On("GetData", mock.Anything).Return(updateData, 1, nil).Once()
		data, res, err := srv.UpdateData(1, validasiInsertData)
		assert.NoError(t, err)
		assert.Equal(t, 1, res, data)
		repo.AssertExpectations(t)
	})

	t.Run("Validasi Insert Data Title Empty", func(t *testing.T) {
		srv := NewActivityBusiness(repo)
		validasiInsertData := activitys.Core{ID: 1, Email: "mulyanurdin10@gmail.com", Title: ""}
		repo.On("UpdateData", mock.Anything, mock.Anything).Return(validasiInsertData, 1, nil).Once()
		repo.On("GetData", mock.Anything).Return(updateData, 1, nil).Once()
		repo.On("UniqueData", mock.Anything).Return(0, nil).Once()
		data, res, err := srv.UpdateData(1, validasiInsertData)
		assert.NoError(t, err)
		assert.Equal(t, 1, res, data)
		repo.AssertExpectations(t)
	})
}

func TestDeleteData(t *testing.T) {
	repo := new(mocks.ActivityData)

	t.Run("Success Delete Data", func(t *testing.T) {
		repo.On("DeleteData", mock.Anything).Return(1, nil).Once()

		srv := NewActivityBusiness(repo)

		res, err := srv.DeleteData(1)
		assert.NoError(t, err)
		assert.Equal(t, 1, res)
		repo.AssertExpectations(t)
	})

	t.Run("Error Delete Data", func(t *testing.T) {
		repo.On("DeleteData", mock.Anything).Return(0, errors.New("Failed to delete data")).Once()

		srv := NewActivityBusiness(repo)

		res, err := srv.DeleteData(2)
		assert.Error(t, err)
		assert.Equal(t, 0, res)
		repo.AssertExpectations(t)
	})
}
