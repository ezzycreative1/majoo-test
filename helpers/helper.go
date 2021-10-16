package helpers

import (
	"crypto/aes"
	"crypto/cipher"
	cRand "crypto/rand"
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"

	"github.com/gin-gonic/gin"
)

var client = &http.Client{}
var headers = map[string]string{
	"Content-Type":   "application/x-www-form-urlencoded",
	"TECHNICIAN_KEY": os.Getenv("KEY_API"),
}

func SDPCMDBRequest(param interface{}, URL, method, operation string) (*http.Response, []byte, error) {
	var client = &http.Client{}

	var req *http.Request
	data := url.Values{}
	if param != nil {
		payload, err := xml.Marshal(param)
		if err != nil {
			return nil, nil, err
		}

		fmt.Println(string(payload))
		data.Set("INPUT_DATA", string(payload))
	}
	data.Set("TECHNICIAN_KEY", os.Getenv("KEY_API"))
	data.Set("OPERATION_NAME", operation)

	body := strings.NewReader(data.Encode())

	reqs, err := http.NewRequest(method, URL, body)
	if err != nil {
		return nil, nil, err
	}
	req = reqs

	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}

	respBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, nil, err
	}

	fmt.Println("SDP CMDB Request Response:")
	fmt.Println(string(respBody))

	return res, respBody, nil
}

func SDPRequest(param interface{}, URL, method string, headers map[string]string) (*http.Response, []byte, error) {
	var client = &http.Client{}

	var req *http.Request

	if param != nil {
		payload, err := json.Marshal(param)
		if err != nil {
			return nil, nil, err
		}

		data := url.Values{}
		data.Set("input_data", string(payload))

		body := strings.NewReader(data.Encode())

		reqs, err := http.NewRequest(method, URL, body)
		if err != nil {
			return nil, nil, err
		}
		req = reqs
	} else {
		reqs, err := http.NewRequest(method, URL, nil)
		if err != nil {
			return nil, nil, err
		}
		req = reqs
	}

	if headers == nil {
		headers = map[string]string{
			"Content-Type":   "application/x-www-form-urlencoded",
			"TECHNICIAN_KEY": os.Getenv("KEY_API"),
			"format":         "json",
		}
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}

	respBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, nil, err
	}

	return res, respBody, nil
}

func SDPRequestWithoutToken(param interface{}, url, method string, headers map[string]string) (*http.Response, []byte, error) {
	var client = &http.Client{}

	var req *http.Request

	if param != nil {
		payload, err := json.Marshal(param)
		if err != nil {
			return nil, nil, err
		}

		body := strings.NewReader(string(payload))

		reqs, err := http.NewRequest(method, url, body)
		if err != nil {
			return nil, nil, err
		}
		req = reqs
	} else {
		reqs, err := http.NewRequest(method, url, nil)
		if err != nil {
			return nil, nil, err
		}
		req = reqs
	}

	if headers == nil {
		headers = map[string]string{
			"TECHNICIAN_KEY": os.Getenv("KEY_API"),
		}
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}

	respBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, nil, err
	}

	return res, respBody, nil
}

func SDPRequestResponseBody(param interface{}, token, url, method string, headers map[string]string) (*http.Response, []byte, error) {
	var client = &http.Client{}

	payload, err := json.Marshal(param)
	if err != nil {
		return nil, nil, err
	}

	body := strings.NewReader(string(payload))

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, nil, err
	}

	if headers == nil {
		headers = map[string]string{
			"Content-Type":  "application/json",
			"Accept":        "application/json",
			"Authorization": token,
		}
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}

	respBody, err := ioutil.ReadAll(res.Body)

	Models := make([]int, 0)
	json.Unmarshal(respBody, &Models)

	if len(Models) == 0 {
		respBody, err = json.Marshal(Models)

		if err != nil {
			return nil, nil, err
		}
	}

	return res, respBody, nil
}

// EncryptBase64 ..
func EncryptBase64(emailPlainText string) string {
	var encodedString = base64.StdEncoding.EncodeToString([]byte(emailPlainText))
	return encodedString
}

// DecryptBase64 ..
func DecryptBase64(emailChiperText string) string {
	var decodedByte, _ = base64.StdEncoding.DecodeString(emailChiperText)
	var decodedString = string(decodedByte)
	return decodedString
}

func ValidateImage(f multipart.File) bool {
	// maximize CPU for better performance
	runtime.GOMAXPROCS(runtime.NumCPU())

	buff := make([]byte, 512) // why 512 byte ? please read http://golang.org/pkg/net/http/#DetectContentType
	if _, err := f.Read(buff); err != nil {
		fmt.Println("========= level debug ==============")
		fmt.Println("user updaload image")
		fmt.Printf("message : %s\n", err.Error())
		return false
	}

	filetype := http.DetectContentType(buff)
	switch filetype {
	case "image/jpeg", "image/jpg":
		return true
	case "image/png":
		return true
	default:
		return false
	}
}

// FileUpload ..
func FileUpload(c *gin.Context) (error, string) {

	file, handler, err := c.Request.FormFile("filePhoto")
	if err != nil {
		return fmt.Errorf("cannot read request body"), ""
	}
	valid := ValidateImage(file)
	if !valid {
		return fmt.Errorf("format image not valid only jpeg, jpg, and png"), ""
	}
	defer file.Close()
	dir := fmt.Sprintf("%s/src/backend-sample-la/storage/photo", os.Getenv("GOPATH"))
	f, err := os.OpenFile(dir+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return err, ""
	}
	defer f.Close()
	io.Copy(f, file)
	newName := fmt.Sprintf("%s%s", GenerateRandString(40), filepath.Ext(handler.Filename))
	os.Rename(dir+handler.Filename, dir+newName)

	return nil, newName
}

// GenerateRandString ..
func GenerateRandString(lg int) string {

	var letterRunes = []rune("123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, lg)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)

}

// ValidEmail ..
func ValidEmail(email string) bool {

	reg := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	regMatch := reg.MatchString(email)

	return regMatch
}

func GenetateUrlImage(image string) string {

	payload := []byte(image)

	block, err := aes.NewCipher([]byte("inikey"))
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	ciphertext := make([]byte, aes.BlockSize+len(payload))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(cRand.Reader, iv); err != nil {
		fmt.Println(err.Error())
		return ""
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], payload)

	content := base64.URLEncoding.EncodeToString(ciphertext)
	url := os.Getenv("BASE_URL")

	return fmt.Sprintf("%s/user/image/%s", url, content)
}

func DecodeImageContent(cryptoText string) string {

	ciphertext, _ := base64.URLEncoding.DecodeString(cryptoText)

	block, err := aes.NewCipher([]byte(""))
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	if len(ciphertext) < aes.BlockSize {
		return ""
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	return fmt.Sprintf("%s", ciphertext)
}

func doRequest(method string, url string, body []byte) ([]byte, error) {

	payload := strings.NewReader(string(body))

	req, err := http.NewRequest(method, url, payload)
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	fmt.Println("LOG ERR", err)

	if err != nil {
		return nil, err
	}

	fmt.Println("LOG REQ", req)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	//fmt.Println(resp.Body)

	respBody, err := ioutil.ReadAll(resp.Body)
	if isStatusError(resp.StatusCode) {
		return nil, fmt.Errorf("Status error: %v %v", resp.StatusCode, string(respBody))
	}

	defer resp.Body.Close()
	return respBody, err
}

func doRequestWithHeader(method string, uri string, body []byte, header map[string]string) ([]byte, error) {

	fmt.Println("body ep => ", body)
	bodyConverted, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	//fmt.Println("bodyConverted", string(bodyConverted))

	data := url.Values{}
	data.Add("input_data", string(bodyConverted))

	payload := strings.NewReader(data.Encode())
	//fmt.Println("body", b)

	req, err := http.NewRequest(method, uri, payload)
	for k, v := range header {
		req.Header.Set(k, v)
	}

	fmt.Println("LOG ERR", err)

	if err != nil {
		return nil, err
	}

	fmt.Println("LOG REQ", req)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	fmt.Println(resp.Body)

	respBody, err := ioutil.ReadAll(resp.Body)
	if isStatusError(resp.StatusCode) {
		return nil, fmt.Errorf("Status error: %v %v", resp.StatusCode, string(respBody))
	}
	fmt.Println("RESPON BODY", respBody)
	defer resp.Body.Close()
	return respBody, err
}

func isStatusError(statusCode int) bool {
	return statusCode > http.StatusBadRequest
}

// Get do get http request
func Get(url string, body []byte) ([]byte, error) {
	return doRequest("GET", url, body)
}

// Post do post http request
func Post(url string, body []byte) ([]byte, error) {
	return doRequest("POST", url, body)
}
func PostWithHeader(url string, body []byte, header map[string]string) ([]byte, error) {
	return doRequestWithHeader("POST", url, body, header)
}
