package usecase

import (
	"context"
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"mail/internal/pkg/logger"
	"mail/internal/pkg/utils/constants"

	mock_repository "mail/internal/microservice/email/mock"
	domain "mail/internal/microservice/models/domain_models"
)

func GetCTX() context.Context {
	f, err := os.OpenFile("log_test.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Failed to create logfile" + "log.txt")
	}
	defer f.Close()

	ctx := context.WithValue(context.Background(), constants.LoggerKey, logger.InitializationBdLog(f))
	ctx2 := context.WithValue(ctx, constants.RequestIDKey, []string{"testID"})

	return ctx2
}

func TestNewEmailUseCase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockEmailRepository(ctrl)

	ExpextedEmailUseCase := EmailUseCase{
		repo: mockRepo,
	}

	EmailUseCase := NewEmailUseCase(mockRepo)

	assert.Equal(t, ExpextedEmailUseCase, *EmailUseCase)
}

func TestGetAllEmailsIncoming_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockEmailRepository(ctrl)
	useCase := NewEmailUseCase(mockRepo)

	login := "test@mailhub.su"
	expectedEmails := []*domain.Email{
		{Topic: "Topic 1", Text: "Text 1"},
		{Topic: "Topic 2", Text: "Text 2"},
	}
	ctx := GetCTX()

	sero := int64(0)

	mockRepo.EXPECT().GetAllIncoming(login, sero, sero, ctx).Return(expectedEmails, nil)

	emails, err := useCase.GetAllEmailsIncoming(login, sero, sero, ctx)

	assert.NoError(t, err)
	assert.Equal(t, expectedEmails, emails)
}

func TestAllEmailsIncoming_ErrorFromRepository(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockEmailRepository(ctrl)
	useCase := NewEmailUseCase(mockRepo)

	login := "test@mailhub.su"
	ctx := GetCTX()
	sero := int64(0)
	//mockRepo.EXPECT().GetAllIncoming(0, 0).Return(nil, errors.New("repository error"))
	mockRepo.EXPECT().GetAllIncoming(login, sero, sero, ctx).Return(nil, errors.New("repository error"))

	emails, err := useCase.GetAllEmailsIncoming(login, sero, sero, ctx)

	assert.Error(t, err)
	assert.Nil(t, emails)
}

func TestGetAllEmailsSent_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockEmailRepository(ctrl)
	useCase := NewEmailUseCase(mockRepo)

	login := "test@mailhub.su"
	expectedEmails := []*domain.Email{
		{Topic: "Topic 1", Text: "Text 1"},
		{Topic: "Topic 2", Text: "Text 2"},
	}
	ctx := GetCTX()
	sero := int64(0)
	mockRepo.EXPECT().GetAllSent(login, sero, sero, ctx).Return(expectedEmails, nil)

	emails, err := useCase.GetAllEmailsSent(login, sero, sero, ctx)

	assert.NoError(t, err)
	assert.Equal(t, expectedEmails, emails)
}

func TestGetAllEmailsSent_ErrorFromRepository(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockEmailRepository(ctrl)
	useCase := NewEmailUseCase(mockRepo)

	login := "test@mailhub.su"
	ctx := GetCTX()
	sero := int64(0)

	mockRepo.EXPECT().GetAllSent(login, sero, sero, ctx).Return(nil, errors.New("repository error"))

	emails, err := useCase.GetAllEmailsSent(login, sero, sero, ctx)

	assert.Error(t, err)
	assert.Nil(t, emails)
}

func TestGetAllEmailsDraft(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockEmailRepository(ctrl)
	useCase := NewEmailUseCase(mockRepo)

	login := "test@mailhub.su"
	expectedEmails := []*domain.Email{
		{Topic: "Topic 1", Text: "Text 1"},
		{Topic: "Topic 2", Text: "Text 2"},
	}
	ctx := GetCTX()
	sero := int64(0)

	mockRepo.EXPECT().GetAllDraft(login, sero, sero, ctx).Return(expectedEmails, nil)

	emails, err := useCase.GetAllDraftEmails(login, sero, sero, ctx)

	assert.NoError(t, err)
	assert.Equal(t, expectedEmails, emails)
}

func TestGetAllEmailsSpam(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockEmailRepository(ctrl)
	useCase := NewEmailUseCase(mockRepo)

	login := "test@mailhub.su"
	expectedEmails := []*domain.Email{
		{Topic: "Topic 1", Text: "Text 1"},
		{Topic: "Topic 2", Text: "Text 2"},
	}
	ctx := GetCTX()
	sero := int64(0)

	mockRepo.EXPECT().GetAllSpam(login, sero, sero, ctx).Return(expectedEmails, nil)

	emails, err := useCase.GetAllSpamEmails(login, sero, sero, ctx)

	assert.NoError(t, err)
	assert.Equal(t, expectedEmails, emails)
}

func TestGetEmailByID_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockEmailRepository(ctrl)
	useCase := NewEmailUseCase(mockRepo)

	login := "test@mailhub.su"
	ctx := GetCTX()

	expectedEmail := &domain.Email{ID: 1, Topic: "Topic 1", Text: "Text 1"}
	mockRepo.EXPECT().GetByID(uint64(1), login, ctx).Return(expectedEmail, nil)

	email, err := useCase.GetEmailByID(1, login, ctx)

	assert.NoError(t, err)
	assert.Equal(t, expectedEmail, email)
}

func TestGetEmailByID_ErrorFromRepository(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockEmailRepository(ctrl)
	useCase := NewEmailUseCase(mockRepo)

	login := "test@mailhub.su"
	ctx := GetCTX()

	mockRepo.EXPECT().GetByID(uint64(1), login, ctx).Return(nil, errors.New("repository error"))

	email, err := useCase.GetEmailByID(1, login, ctx)

	assert.Error(t, err)
	assert.Nil(t, email)
}

func TestCreateEmail_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockEmailRepository(ctrl)
	useCase := NewEmailUseCase(mockRepo)

	ctx := GetCTX()

	newEmail := &domain.Email{Topic: "Topic 1", Text: "Text 1"}
	mockRepo.EXPECT().Add(gomock.Any(), ctx).Return(uint64(1), newEmail, nil)

	id, emailRes, err := useCase.CreateEmail(newEmail, ctx)

	assert.Equal(t, uint64(1), id)
	assert.NoError(t, err)
	assert.Equal(t, newEmail, emailRes)
}

func TestCreateEmail_ErrorFromRepository(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockEmailRepository(ctrl)
	useCase := NewEmailUseCase(mockRepo)
	newEmail := &domain.Email{Topic: "Topic 1", Text: "Text 1"}

	ctx := GetCTX()

	mockRepo.EXPECT().Add(gomock.Any(), ctx).Return(uint64(1), newEmail, errors.New("repository error"))

	id, emailRes, err := useCase.CreateEmail(newEmail, ctx)

	assert.Equal(t, uint64(1), id)
	assert.Error(t, err)
	assert.Equal(t, newEmail, emailRes)
}

func TestCreateProfileEmail_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockEmailRepository(ctrl)
	useCase := NewEmailUseCase(mockRepo)

	email_id := uint64(1)
	sender := "test_sender@mailhub.su"
	recipient := "test_recipient@mailhub.su"
	ctx := GetCTX()

	mockRepo.EXPECT().AddProfileEmail(email_id, sender, recipient, ctx).Return(nil)

	err := useCase.CreateProfileEmail(email_id, sender, recipient, ctx)

	assert.NoError(t, err)
}

func TestCreateProfileEmail_ErrorFromRepository(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockEmailRepository(ctrl)
	useCase := NewEmailUseCase(mockRepo)

	email_id := uint64(1)
	sender := "test_sender@mailhub.su"
	recipient := "test_recipient@mailhub.su"
	ctx := GetCTX()

	mockRepo.EXPECT().AddProfileEmail(email_id, sender, recipient, ctx).Return(errors.New("repository error"))

	err := useCase.CreateProfileEmail(email_id, sender, recipient, ctx)

	assert.Error(t, err)
}

func TestCheckRecipientEmail_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockEmailRepository(ctrl)
	useCase := NewEmailUseCase(mockRepo)

	recipient := "test_recipient@mailhub.su"
	ctx := GetCTX()

	mockRepo.EXPECT().FindEmail(recipient, ctx).Return(nil)

	err := useCase.CheckRecipientEmail(recipient, ctx)

	assert.NoError(t, err)
}

func TestCheckRecipientEmail_ErrorFromRepository(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockEmailRepository(ctrl)
	useCase := NewEmailUseCase(mockRepo)

	recipient := "test_recipient@mailhub.su"
	ctx := GetCTX()

	mockRepo.EXPECT().FindEmail(recipient, ctx).Return(errors.New("repository error"))

	err := useCase.CheckRecipientEmail(recipient, ctx)

	assert.Error(t, err)
}

func TestUpdateEmail_ErrorFromRepository(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockEmailRepository(ctrl)
	useCase := NewEmailUseCase(mockRepo)

	newEmail := &domain.Email{ID: 1, Topic: "Topic 1", Text: "Text 1"}
	ctx := GetCTX()

	mockRepo.EXPECT().Update(gomock.Any(), ctx).Return(true, nil)

	emailRes, err := useCase.UpdateEmail(newEmail, ctx)

	assert.NoError(t, err)
	assert.Equal(t, true, emailRes)
}

func TestDeleteEmail_ErrorFromRepository(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockEmailRepository(ctrl)
	useCase := NewEmailUseCase(mockRepo)

	login := "test@mailhub.su"
	ctx := GetCTX()

	mockRepo.EXPECT().Delete(gomock.Any(), login, ctx).Return(true, nil)

	emailRes, err := useCase.DeleteEmail(uint64(1), login, ctx)

	assert.NoError(t, err)
	assert.Equal(t, true, emailRes)
}

func TestAddAttachment_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockEmailRepository(ctrl)
	useCase := NewEmailUseCase(mockRepo)

	fileID := "test_file_id"
	fileType := "pdf"
	emailID := uint64(123)
	ctx := context.Background()

	mockRepo.EXPECT().AddFile(fileID, fileType, ctx).Return(uint64(456), nil)
	mockRepo.EXPECT().AddAttachment(emailID, uint64(456), ctx).Return(nil)

	result, err := useCase.AddAttachment(fileID, fileType, emailID, ctx)

	assert.NoError(t, err)
	assert.Equal(t, uint64(456), result)
}

func TestAddAttachment_EmptyFileID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockEmailRepository(ctrl)
	useCase := NewEmailUseCase(mockRepo)

	fileID := ""
	fileType := "pdf"
	emailID := uint64(123)
	ctx := context.Background()

	result, err := useCase.AddAttachment(fileID, fileType, emailID, ctx)

	assert.Error(t, err)
	assert.Zero(t, result)
}

func TestAddAttachment_EmptyFileType(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockEmailRepository(ctrl)
	useCase := NewEmailUseCase(mockRepo)

	fileID := "test_file_id"
	fileType := ""
	emailID := uint64(123)
	ctx := context.Background()

	result, err := useCase.AddAttachment(fileID, fileType, emailID, ctx)

	assert.Error(t, err)
	assert.Zero(t, result)
}

func TestAddAttachment_InvalidEmailID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockEmailRepository(ctrl)
	useCase := NewEmailUseCase(mockRepo)

	fileID := "test_file_id"
	fileType := "pdf"
	emailID := uint64(0)
	ctx := context.Background()

	result, err := useCase.AddAttachment(fileID, fileType, emailID, ctx)

	assert.Error(t, err)
	assert.Zero(t, result)
}

func TestAddAttachment_ErrorAddingFile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockEmailRepository(ctrl)
	useCase := NewEmailUseCase(mockRepo)

	fileID := "test_file_id"
	fileType := "pdf"
	emailID := uint64(123)
	ctx := context.Background()

	mockRepo.EXPECT().AddFile(fileID, fileType, ctx).Return(uint64(0), errors.New("file adding error"))

	result, err := useCase.AddAttachment(fileID, fileType, emailID, ctx)

	assert.Error(t, err)
	assert.Zero(t, result)
}

func TestAddAttachment_ErrorAddingAttachment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockEmailRepository(ctrl)
	useCase := NewEmailUseCase(mockRepo)

	fileID := "test_file_id"
	fileType := "pdf"
	emailID := uint64(123)
	ctx := context.Background()

	mockRepo.EXPECT().AddFile(fileID, fileType, ctx).Return(uint64(456), nil)
	mockRepo.EXPECT().AddAttachment(emailID, uint64(456), ctx).Return(errors.New("attachment adding error"))

	result, err := useCase.AddAttachment(fileID, fileType, emailID, ctx)

	assert.Error(t, err)
	assert.Zero(t, result)
}

func TestGetFileByID_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockEmailRepository(ctrl)
	useCase := NewEmailUseCase(mockRepo)

	fileID := uint64(123)
	expectedFile := &domain.File{ID: fileID, FileId: "test_file"}

	ctx := context.Background()

	mockRepo.EXPECT().GetFileByID(fileID, ctx).Return(expectedFile, nil)

	file, err := useCase.GetFileByID(fileID, ctx)

	assert.NoError(t, err)
	assert.Equal(t, expectedFile, file)
}

func TestGetFileByID_InvalidFileID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockEmailRepository(ctrl)
	useCase := NewEmailUseCase(mockRepo)

	fileID := uint64(0)
	ctx := context.Background()

	file, err := useCase.GetFileByID(fileID, ctx)

	assert.Error(t, err)
	assert.Nil(t, file)
}

func TestGetFileByID_ErrorGettingFile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockEmailRepository(ctrl)
	useCase := NewEmailUseCase(mockRepo)

	fileID := uint64(123)
	ctx := context.Background()

	mockRepo.EXPECT().GetFileByID(fileID, ctx).Return(nil, errors.New("file retrieval error"))

	file, err := useCase.GetFileByID(fileID, ctx)

	assert.Error(t, err)
	assert.Nil(t, file)
}

func TestGetFilesByEmailID_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockEmailRepository(ctrl)
	useCase := NewEmailUseCase(mockRepo)

	emailID := uint64(123)
	expectedFiles := []*domain.File{
		{ID: 1, FileId: "file1.pdf"},
		{ID: 2, FileId: "file2.pdf"},
	}

	ctx := context.Background()

	mockRepo.EXPECT().GetFilesByEmailID(emailID, ctx).Return(expectedFiles, nil)

	files, err := useCase.GetFilesByEmailID(emailID, ctx)

	assert.NoError(t, err)
	assert.Equal(t, expectedFiles, files)
}

func TestGetFilesByEmailID_InvalidEmailID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockEmailRepository(ctrl)
	useCase := NewEmailUseCase(mockRepo)

	emailID := uint64(0)
	ctx := context.Background()

	files, err := useCase.GetFilesByEmailID(emailID, ctx)

	assert.Error(t, err)
	assert.Nil(t, files)
}

func TestGetFilesByEmailID_ErrorGettingFiles(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockEmailRepository(ctrl)
	useCase := NewEmailUseCase(mockRepo)

	emailID := uint64(123)
	ctx := context.Background()

	mockRepo.EXPECT().GetFilesByEmailID(emailID, ctx).Return(nil, errors.New("file retrieval error"))

	files, err := useCase.GetFilesByEmailID(emailID, ctx)

	assert.Error(t, err)
	assert.Nil(t, files)
}

func TestDeleteFileByID_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockEmailRepository(ctrl)
	useCase := NewEmailUseCase(mockRepo)

	fileID := uint64(123)
	ctx := context.Background()

	mockRepo.EXPECT().DeleteFileByID(fileID, ctx).Return(nil)

	deleted, err := useCase.DeleteFileByID(fileID, ctx)

	assert.NoError(t, err)
	assert.True(t, deleted)
}

func TestDeleteFileByID_InvalidFileID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockEmailRepository(ctrl)
	useCase := NewEmailUseCase(mockRepo)

	fileID := uint64(0)
	ctx := context.Background()

	deleted, err := useCase.DeleteFileByID(fileID, ctx)

	assert.Error(t, err)
	assert.False(t, deleted)
}

func TestDeleteFileByID_ErrorDeletingFile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockEmailRepository(ctrl)
	useCase := NewEmailUseCase(mockRepo)

	fileID := uint64(123)
	ctx := context.Background()

	mockRepo.EXPECT().DeleteFileByID(fileID, ctx).Return(errors.New("file deletion error"))

	deleted, err := useCase.DeleteFileByID(fileID, ctx)

	assert.Error(t, err)
	assert.False(t, deleted)
}

func TestUpdateFileByID_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockEmailRepository(ctrl)
	useCase := NewEmailUseCase(mockRepo)

	fileID := uint64(123)
	newFileID := "new_file_id"
	newFileType := "pdf"
	ctx := context.Background()

	oldFile := &domain.File{
		ID:       fileID,
		FileId:   "old_file_id",
		FileType: "old_type",
	}

	mockRepo.EXPECT().GetFileByID(fileID, ctx).Return(oldFile, nil)
	mockRepo.EXPECT().UpdateFileByID(fileID, newFileID, newFileType, ctx).Return(nil)

	updated, err := useCase.UpdateFileByID(fileID, newFileID, newFileType, ctx)

	assert.NoError(t, err)
	assert.True(t, updated)
	assert.Equal(t, newFileID, oldFile.FileId)
	assert.Equal(t, newFileType, oldFile.FileType)
}

func TestUpdateFileByID_InvalidFileID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockEmailRepository(ctrl)
	useCase := NewEmailUseCase(mockRepo)

	fileID := uint64(0)
	newFileID := "new_file_id"
	newFileType := "pdf"
	ctx := context.Background()

	updated, err := useCase.UpdateFileByID(fileID, newFileID, newFileType, ctx)

	assert.Error(t, err)
	assert.False(t, updated)
}

func TestUpdateFileByID_EmptyNewFileID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockEmailRepository(ctrl)
	useCase := NewEmailUseCase(mockRepo)

	fileID := uint64(123)
	newFileID := ""
	newFileType := "pdf"
	ctx := context.Background()

	updated, err := useCase.UpdateFileByID(fileID, newFileID, newFileType, ctx)

	assert.Error(t, err)
	assert.False(t, updated)
}

func TestUpdateFileByID_EmptyNewFileType(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockEmailRepository(ctrl)
	useCase := NewEmailUseCase(mockRepo)

	fileID := uint64(123)
	newFileID := "new_file_id"
	newFileType := ""
	ctx := context.Background()

	updated, err := useCase.UpdateFileByID(fileID, newFileID, newFileType, ctx)

	assert.Error(t, err)
	assert.False(t, updated)
}

func TestUpdateFileByID_ErrorGettingFile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockEmailRepository(ctrl)
	useCase := NewEmailUseCase(mockRepo)

	fileID := uint64(123)
	newFileID := "new_file_id"
	newFileType := "pdf"
	ctx := context.Background()

	mockRepo.EXPECT().GetFileByID(fileID, ctx).Return(nil, errors.New("file retrieval error"))

	updated, err := useCase.UpdateFileByID(fileID, newFileID, newFileType, ctx)

	assert.Error(t, err)
	assert.False(t, updated)
}

func TestUpdateFileByID_ErrorUpdatingFile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockEmailRepository(ctrl)
	useCase := NewEmailUseCase(mockRepo)

	fileID := uint64(123)
	newFileID := "new_file_id"
	newFileType := "pdf"
	ctx := context.Background()

	oldFile := &domain.File{
		ID:       fileID,
		FileId:   "old_file_id",
		FileType: "old_type",
	}

	mockRepo.EXPECT().GetFileByID(fileID, ctx).Return(oldFile, nil)
	mockRepo.EXPECT().UpdateFileByID(fileID, newFileID, newFileType, ctx).Return(errors.New("file update error"))

	updated, err := useCase.UpdateFileByID(fileID, newFileID, newFileType, ctx)

	assert.Error(t, err)
	assert.False(t, updated)
}
