package http

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/denisbrodbeck/striphtmltags"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"io/ioutil"
	"log"
	apiModels "mail/internal/models/delivery_models"
	"mail/internal/models/response"
	gmailAuth "mail/internal/pkg/gmail/gmail_auth/delivery/http"
	domainSession "mail/internal/pkg/session/interface"
	"mail/internal/pkg/utils/validators"
	"net/http"
	"os"
	"time"
)

// GMailEmailHandler handles user-related HTTP requests.
type GMailEmailHandler struct {
	Sessions     domainSession.SessionsManager
	GMailService *gmail.Service
}

// GetIncoming displays the list of email messages.
// @Summary Display the list of email messages
// @Description Get a list of all email messages
// @Tags emails-gmail
// @Produce json
// @Param X-Csrf-Token header string true "CSRF Token"
// @Success 200 {object} response.Response "List of all email messages"
// @Failure 400 {object} response.Response "Bad request"
// @Failure 401 {object} response.Response "Not Authorized"
// @Failure 500 {object} response.Response "JSON encoding error"
// @Router /api/v1/gmail/emails/incoming [get]
func (g *GMailEmailHandler) GetIncoming(w http.ResponseWriter, r *http.Request) {
	login, err := g.Sessions.GetLoginBySession(r, r.Context())
	if err != nil {
		response.HandleError(w, http.StatusBadRequest, "Bad user session")
		return
	}
	if !validators.IsValidEmailFormatGmail(login) {
		response.HandleError(w, http.StatusBadRequest, "Login must end with @gmail.com")
		return
	}

	srv, err := GetSRV(login)
	if err != nil {
		response.HandleError(w, http.StatusInternalServerError, "Unable to retrieve Gmail client")
		return
	}

	req, err := srv.Users.Messages.List("me").Q("label:inbox").Do()
	if err != nil {
		response.HandleError(w, http.StatusInternalServerError, "Error receiving list messages")
		return
	}

	emailsApi := make([]*apiModels.OtherEmail, len(req.Messages))
	for i, m := range req.Messages {
		msg, err := srv.Users.Messages.Get("me", m.Id).Format("full").Do()
		if err != nil {
			response.HandleError(w, http.StatusInternalServerError, "Error receiving messages")
			return
		}
		email := CreateEmailStruct(msg)
		emailsApi[i] = email
	}
	response.HandleSuccess(w, http.StatusOK, map[string]interface{}{"emails": emailsApi})
}

// GetSent displays the list of email messages.
// @Summary Display the list of email messages
// @Description Get a list of all email messages
// @Tags emails-gmail
// @Produce json
// @Param X-Csrf-Token header string true "CSRF Token"
// @Success 200 {object} response.Response "List of all email messages"
// @Failure 400 {object} response.Response "Bad request"
// @Failure 401 {object} response.Response "Not Authorized"
// @Failure 500 {object} response.Response "JSON encoding error"
// @Router /api/v1/gmail/emails/sent [get]
func (g *GMailEmailHandler) GetSent(w http.ResponseWriter, r *http.Request) {
	login, err := g.Sessions.GetLoginBySession(r, r.Context())
	if err != nil {
		response.HandleError(w, http.StatusBadRequest, "Bad user session")
		return
	}
	if !validators.IsValidEmailFormatGmail(login) {
		response.HandleError(w, http.StatusBadRequest, "Login must end with @gmail.com")
		return
	}

	srv, err := GetSRV(login)
	if err != nil {
		response.HandleError(w, http.StatusInternalServerError, "Unable to retrieve Gmail client")
		return
	}

	req, err := srv.Users.Messages.List("me").Q("label:sent").Do()
	if err != nil {
		response.HandleError(w, http.StatusInternalServerError, "Error receiving list messages")
		return
	}

	emailsApi := make([]*apiModels.OtherEmail, len(req.Messages))
	for i, m := range req.Messages {
		msg, err := srv.Users.Messages.Get("me", m.Id).Format("full").Do()
		if err != nil {
			response.HandleError(w, http.StatusInternalServerError, "Error receiving messages")
			return
		}
		email := CreateEmailStruct(msg)
		emailsApi[i] = email
	}
	response.HandleSuccess(w, http.StatusOK, map[string]interface{}{"emails": emailsApi})
}

// GetSpam displays the list of email messages.
// @Summary Display the list of email messages
// @Description Get a list of all email messages
// @Tags emails-gmail
// @Produce json
// @Param X-Csrf-Token header string true "CSRF Token"
// @Success 200 {object} response.Response "List of all email messages"
// @Failure 400 {object} response.Response "Bad request"
// @Failure 401 {object} response.Response "Not Authorized"
// @Failure 500 {object} response.Response "JSON encoding error"
// @Router /api/v1/gmail/emails/spam [get]
func (g *GMailEmailHandler) GetSpam(w http.ResponseWriter, r *http.Request) {
	login, err := g.Sessions.GetLoginBySession(r, r.Context())
	if err != nil {
		response.HandleError(w, http.StatusBadRequest, "Bad user session")
		return
	}
	if !validators.IsValidEmailFormatGmail(login) {
		response.HandleError(w, http.StatusBadRequest, "Login must end with @gmail.com")
		return
	}

	srv, err := GetSRV(login)
	if err != nil {
		response.HandleError(w, http.StatusInternalServerError, "Unable to retrieve Gmail client")
		return
	}

	req, err := srv.Users.Messages.List("me").Q("label:spam").Do()
	if err != nil {
		response.HandleError(w, http.StatusInternalServerError, "Error receiving list messages")
		return
	}

	emailsApi := make([]*apiModels.OtherEmail, len(req.Messages))
	for i, m := range req.Messages {
		msg, err := srv.Users.Messages.Get("me", m.Id).Format("full").Do()
		if err != nil {
			response.HandleError(w, http.StatusInternalServerError, "Error receiving messages")
			return
		}
		email := CreateEmailStruct(msg)
		emailsApi[i] = email
	}

	for i, _ := range emailsApi {
		emailsApi[i].SpamStatus = true
	}

	response.HandleSuccess(w, http.StatusOK, map[string]interface{}{"emails": emailsApi})
}

// GetDrafts displays the list of email messages.
// @Summary Display the list of email messages
// @Description Get a list of all email messages
// @Tags emails-gmail
// @Produce json
// @Param X-Csrf-Token header string true "CSRF Token"
// @Success 200 {object} response.Response "List of all email messages"
// @Failure 400 {object} response.Response "Bad request"
// @Failure 401 {object} response.Response "Not Authorized"
// @Failure 500 {object} response.Response "JSON encoding error"
// @Router /api/v1/gmail/emails/draft [get]
func (g *GMailEmailHandler) GetDrafts(w http.ResponseWriter, r *http.Request) {
	login, err := g.Sessions.GetLoginBySession(r, r.Context())
	if err != nil {
		response.HandleError(w, http.StatusBadRequest, "Bad user session")
		return
	}
	if !validators.IsValidEmailFormatGmail(login) {
		response.HandleError(w, http.StatusBadRequest, "Login must end with @gmail.com")
		return
	}

	srv, err := GetSRV(login)
	if err != nil {
		response.HandleError(w, http.StatusInternalServerError, "Unable to retrieve Gmail client")
		return
	}

	req, err := srv.Users.Messages.List("me").Q("label:drafts").Do()
	if err != nil {
		response.HandleError(w, http.StatusInternalServerError, "Error receiving list messages")
		return
	}

	emailsApi := make([]*apiModels.OtherEmail, len(req.Messages))
	for i, m := range req.Messages {
		msg, err := srv.Users.Messages.Get("me", m.Id).Format("full").Do()
		if err != nil {
			response.HandleError(w, http.StatusInternalServerError, "Error receiving messages")
			return
		}
		email := CreateEmailStruct(msg)
		emailsApi[i] = email
	}

	for i, _ := range emailsApi {
		emailsApi[i].DraftStatus = true
	}

	response.HandleSuccess(w, http.StatusOK, map[string]interface{}{"emails": emailsApi})
}

// GetById returns an email message by its ID.
// @Summary Get an email message by ID
// @Description Get an email message by its unique identifier
// @Tags emails-gmail
// @Produce json
// @Param id path integer true "ID of the email message"
// @Param X-Csrf-Token header string true "CSRF Token"
// @Success 200 {object} response.Response "Email message data"
// @Failure 400 {object} response.Response "Bad id in request"
// @Failure 401 {object} response.Response "Not Authorized"
// @Failure 404 {object} response.Response "Email not found"
// @Router /api/v1/gmail/email/{id} [get]
func (g *GMailEmailHandler) GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	messageID, ok := vars["id"]
	if !ok {
		response.HandleError(w, http.StatusBadRequest, "Bad id in request")
		return
	}

	login, err := g.Sessions.GetLoginBySession(r, r.Context())
	if err != nil {
		response.HandleError(w, http.StatusBadRequest, "Bad user session")
		return
	}
	if !validators.IsValidEmailFormatGmail(login) {
		response.HandleError(w, http.StatusBadRequest, "Login must end with @gmail.com")
		return
	}

	srv, err := GetSRV(login)
	if err != nil {
		response.HandleError(w, http.StatusInternalServerError, "Unable to retrieve Gmail client")
		return
	}

	msg, err := srv.Users.Messages.Get("me", messageID).Format("full").Do()
	if err != nil {
		response.HandleError(w, http.StatusInternalServerError, "Error receiving messages")
		return
	}

	email := CreateEmailStruct(msg)

	response.HandleSuccess(w, http.StatusOK, map[string]interface{}{"email": email})
}

// Send adds a new email message.
// @Summary Send a new email message
// @Description Send a new email message to the system
// @Tags emails-gmail
// @Accept json
// @Produce json
// @Param X-Csrf-Token header string true "CSRF Token"
// @Param email body response.EmailSwag true "Email message in JSON format"
// @Success 200 {object} response.Response "ID of the send email message"
// @Failure 400 {object} response.Response "Bad JSON in request"
// @Failure 401 {object} response.Response "Not Authorized"
// @Failure 500 {object} response.Response "Failed to add email message"
// @Router /api/v1/gmail/email/send [post]
func (g *GMailEmailHandler) Send(w http.ResponseWriter, r *http.Request) {
	var newEmail apiModels.OtherEmail
	decoder := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)
	err := json.NewDecoder(r.Body).Decode(&newEmail)
	if err != nil {
		response.HandleError(w, http.StatusBadRequest, "Bad JSON in request")
		return
	}

	login, err := g.Sessions.GetLoginBySession(r, r.Context())
	if err != nil {
		response.HandleError(w, http.StatusBadRequest, "Bad sender session")
		return
	}

	err = g.Sessions.CheckLogin(newEmail.SenderEmail, r, r.Context())
	if err != nil {
		response.HandleError(w, http.StatusBadRequest, "Bad sender login")
		return
	}

	srv, err := GetSRV(login)
	if err != nil {
		response.HandleError(w, http.StatusInternalServerError, "Unable to retrieve Gmail client")
		return
	}

	input := base64.RawStdEncoding.EncodeToString([]byte(fmt.Sprintf("From: %v\r\nTo: %v\r\nSubject: %v\r\n\r\n%v", newEmail.SenderEmail, newEmail.RecipientEmail, newEmail.Topic, newEmail.Text)))
	/*
		input := base64.RawStdEncoding.EncodeToString([]byte(
			"From: fedasovsergey00@gmail.com\r\n" +
				"To: fedasov03@inbox.ru\r\n" +
				fmt.Sprintf("Subject: Hello\n\n", sub) +
				"Это тестовое сообщение, отправленное через Gmail API."))
	*/

	message := &gmail.Message{
		Raw: input,
	}

	_, err = srv.Users.Messages.Send("me", message).Do()
	if err != nil {
		response.HandleError(w, http.StatusInternalServerError, "Error sending the message")
		return
	}

	response.HandleSuccess(w, http.StatusOK, map[string]interface{}{"email": newEmail})
}

// Delete deletes an email message.
// @Summary Delete an email message
// @Description Delete an email message based on its identifier
// @Tags emails-gmail
// @Produce json
// @Param id path integer true "ID of the email message"
// @Param X-Csrf-Token header string true "CSRF Token"
// @Success 200 {object} response.Response "Deletion success status"
// @Failure 400 {object} response.Response "Bad id"
// @Failure 401 {object} response.Response "Not Authorized"
// @Failure 500 {object} response.Response "Failed to delete email message"
// @Router /api/v1/gmail/email/delete/{id} [delete]
func (g *GMailEmailHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	messageID, ok := vars["id"]
	if !ok {
		response.HandleError(w, http.StatusBadRequest, "Bad id in request")
		return
	}

	login, err := g.Sessions.GetLoginBySession(r, r.Context())
	if err != nil {
		response.HandleError(w, http.StatusBadRequest, "Bad user session")
		return
	}
	if !validators.IsValidEmailFormatGmail(login) {
		response.HandleError(w, http.StatusBadRequest, "Login must end with @gmail.com")
		return
	}

	srv, err := GetSRV(login)
	if err != nil {
		response.HandleError(w, http.StatusInternalServerError, "Unable to retrieve Gmail client")
		return
	}

	err = srv.Users.Messages.Delete("me", messageID).Do()
	if err != nil {
		response.HandleError(w, http.StatusInternalServerError, "Error deleting a message")
		return
	}

	response.HandleSuccess(w, http.StatusOK, map[string]interface{}{"Success": true})
}

func CreateEmailStruct(msg *gmail.Message) *apiModels.OtherEmail {
	email := &apiModels.OtherEmail{}
	email.ID = msg.Id
	email.DateOfDispatch = time.Unix(msg.InternalDate/1000, 0)

	fmt.Println(msg.Payload.MimeType)
	if msg.Payload.MimeType == "text/plain" {
		email = ParserMessageHeadres(email, msg)
		data, err := base64.URLEncoding.DecodeString(msg.Payload.Body.Data)
		if err != nil {
			fmt.Println("Error: ", err)
		}
		email.Text = striphtmltags.StripTags(string(data))
	} else if msg.Payload.MimeType == "text/html" {
		data, err := base64.URLEncoding.DecodeString(msg.Payload.Body.Data)
		if err != nil {
			fmt.Println("Error: ", err)
		}
		email.Text = string(data)
		email = ParserMessageHeadres(email, msg)
	} else if len(msg.Payload.Parts) != 0 {
		for _, part := range msg.Payload.Parts {
			if part.MimeType == "text/html" {
				data, err := base64.URLEncoding.DecodeString(part.Body.Data)
				if err != nil {
					fmt.Println("Error: ", err)
				}
				email.Text = string(data)
				email = ParserMessageHeadres(email, msg)
			}
		}
	}

	return email
}

func ParserMessageHeadres(email *apiModels.OtherEmail, msg *gmail.Message) *apiModels.OtherEmail {
	for _, mes := range msg.Payload.Headers {
		if mes.Name == "To" {
			email.RecipientEmail = mes.Value
		}
		if mes.Name == "From" {
			email.SenderEmail = mes.Value
		}
		if mes.Name == "Subject" {
			email.Topic = striphtmltags.StripTags(mes.Value)
		}
	}
	return email
}

func getClient() (*oauth2.Config, *oauth2.Token) {
	b, err := os.ReadFile("cmd/configs/credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	config, err := google.ConfigFromJSON(b, gmail.GmailReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	tok := &oauth2.Token{}
	data, _ := ioutil.ReadFile("token.json")
	err = json.Unmarshal(data, tok)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return config, tok
}

func GetSRV(login string) (*gmail.Service, error) {
	srv, ok := gmailAuth.MapOAuthCongig[login]
	if !ok {
		return nil, errors.New("Unable to retrieve Gmail client")
	}
	return srv, nil
}
