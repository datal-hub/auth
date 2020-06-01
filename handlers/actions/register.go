package actions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/lib/pq"

	"github.com/datal-hub/auth/handlers/middleware"
	"github.com/datal-hub/auth/models"
	log "github.com/datal-hub/auth/pkg/logger"
)

func Register(w http.ResponseWriter, r *http.Request) {
	logDetails := middleware.HttpLogDetails(r)

	if r.Body == nil {
		log.ErrorF("Register: nil body", logDetails)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message":"nil body"}`))
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logDetails["message"] = err.Error()
		log.ErrorF("Register: error reading body.", logDetails)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message":"invalid body"}`))
		return
	}
	cred := &models.Credentials{}
	if err := json.Unmarshal(body, cred); err != nil {
		logDetails["message"] = err.Error()
		log.ErrorF("Register: Unmarshal auth data error", logDetails)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message":"invalid body"}`))
		return
	}

	if err := cred.IsValid(); err != nil {
		logDetails["message"] = err.Error()
		log.ErrorF("Register: invalid credentials", logDetails)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf(`{"message":"%s"}`, err.Error())))
		return
	}
	if err := cred.SetHash(); err != nil {
		logDetails["message"] = err.Error()
		log.ErrorF("Register: set hash error", logDetails)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"internal error"}`))
		return
	}

	db := middleware.DBFromContext(r.Context())
	cred.CreateDttm = time.Now().UTC()
	err = db.Save(cred)
	if err != nil {
		if err, ok := err.(*pq.Error); ok && err.Code == "23505" {
			log.ErrorF("Register: user already exists", logDetails)
			w.WriteHeader(http.StatusConflict)
			w.Write([]byte(`{"message":"user already exists"}`))
			return
		} else {
			logDetails["message"] = err.Error()
			log.ErrorF("Register: save credentials error", logDetails)
			w.WriteHeader(http.StatusServiceUnavailable)
			w.Write([]byte(`{"message":"service unavailable"}`))
			return
		}
	}
	w.WriteHeader(http.StatusOK)
}
