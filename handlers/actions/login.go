package actions

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/datal-hub/auth/handlers/middleware"
	"github.com/datal-hub/auth/models"
	log "github.com/datal-hub/auth/pkg/logger"
)

func Login(w http.ResponseWriter, r *http.Request) {
	logDetails := middleware.HttpLogDetails(r)

	if r.Body == nil {
		log.ErrorF("Login: nil body", logDetails)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message":"nil body"}`))
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logDetails["message"] = err.Error()
		log.ErrorF("Login: error reading body.", logDetails)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message":"invalid body"}`))
		return
	}
	cred := &models.Credentials{}
	if err := json.Unmarshal(body, cred); err != nil {
		logDetails["message"] = err.Error()
		log.ErrorF("Login: Unmarshal auth data error", logDetails)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message":"invalid body"}`))
		return
	}

	if cred.Login == "" || cred.Password == "" {
		log.ErrorF("Login: invalid credentials", logDetails)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message":"empty fields"}`))
		return
	}

	db := middleware.DBFromContext(r.Context())
	user, err := db.GetCredentials(cred.Login)
	if user == nil && err == nil {
		log.ErrorF("Login: credentials not found", logDetails)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message":"credentials not found"}`))
		return
	}
	if err != nil {
		logDetails["message"] = err.Error()
		log.ErrorF("Login: get credentials error", logDetails)
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte(`{"message":"service unavailable"}`))
		return
	}
	user.Password = cred.Password
	if err := user.CheckPassword(); err != nil {
		logDetails["message"] = err.Error()
		log.ErrorF("Login: check password error", logDetails)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message":"invalid password"}`))
		return
	}
	user.Password = ""
	response, err := json.Marshal(user)
	if err != nil {
		logDetails["message"] = err.Error()
		log.ErrorF("Login: error marshaling info.", logDetails)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "internal error"}`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
