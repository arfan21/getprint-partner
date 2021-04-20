package imgur

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

type Imgur interface {
	Upload(data string) (*ImgurResponse, error)
	Delete(deleteHash string) error
}

type imgur struct {
	clientId string
}

func NewImgur() Imgur {
	clientId := os.Getenv("IMGUR_CLIENT_ID")
	return &imgur{clientId: clientId}
}

func (img imgur) Upload(data string) (*ImgurResponse, error) {
	url := "https://api.imgur.com/3/upload"
	dataBase64 := strings.Split(data, ",")[1]

	payload := new(bytes.Buffer)
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("image", dataBase64)
	_ = writer.WriteField("type", "base64")
	err := writer.Close()
	if err != nil {
		return nil, err
	}

	client := new(http.Client)

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Client-ID %s", img.clientId))
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	imgurResponse := new(ImgurResponse)

	err = json.Unmarshal(resBody, &imgurResponse)
	if err != nil {
		return nil, err
	}

	if !imgurResponse.Success {
		return nil, fmt.Errorf("%s", imgurResponse.Data.Error)
	}

	return imgurResponse, nil
}

func (img imgur) Delete(deleteHash string) error {
	var url = "https://api.imgur.com/3/image/" + deleteHash

	client := new(http.Client)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Client-ID %s", img.clientId))

	res, err := client.Do(req)
	defer res.Body.Close()
	if err != nil {
		return err
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	imgurResponse := new(ImgurResponse)

	err = json.Unmarshal(resBody, &imgurResponse)
	if err != nil {
		return err
	}

	if !imgurResponse.Success {
		return fmt.Errorf("%s", imgurResponse.Data.Error)
	}

	return nil
}
