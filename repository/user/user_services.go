package user

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/arfan21/getprint-partner/models"
)

type UserServices interface {
	VerificationUser(id string) error
}

type userServices struct {
}

func NewUserServices() UserServices {
	return &userServices{}
}

func (srv *userServices) VerificationUser(id string) error {
	var url = os.Getenv("SERVICE_USER")
	res, err := http.Get(url + "/user/" + id)

	if err != nil {
		if strings.Contains(err.Error(), "connection refused") {
			return models.ErrInternalServerError
		}
		return err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return err
	}

	decodedJSON := make(map[string]interface{})

	err = json.Unmarshal(body, &decodedJSON)

	if err != nil {
		return err
	}

	if !(res.StatusCode >= 200 && res.StatusCode < 300) {
		return errors.New(decodedJSON["message"].(string))
	}

	return nil
}
