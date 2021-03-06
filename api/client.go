package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/sirupsen/logrus"
	"github.com/zxldev/fenbeitong/util"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
	"time"
)

/**
@see http://docs.open.fenbeitong.com/open-api/
*/

type FenBeiTong struct {
	AppId         string    `json:"app_id"`
	AppKey        string    `json:"app_key"`
	SignKey       string    `json:"sign_key"`
	AdminEmployId string    `json:"admin_phone"`
	AccessToken   string    `json:"access_token"`
	TokenExpire   time.Time `json:"token_expire"`
	ApiUrl        string    `json:"api_url"`
	TokenLock     sync.WaitGroup
	DepartmentMap map[string]string `json:"department_map"`
}

var FenBeiTongClient FenBeiTong

func (d *FenBeiTong) Init(AppId, AppKey, SignKey, AdminEmployId, ApiUrl string) {
	if ApiUrl == "" {
		d.ApiUrl = ServerApi
	} else {
		d.ApiUrl = ApiUrl
	}
	d.AppId = AppId
	d.AppKey = AppKey
	d.SignKey = SignKey
	d.AdminEmployId = AdminEmployId
	//d.GenDepartmentMap()
}

/**
判断是否同一个部门
*/
func (d *FenBeiTong) IsSameDepartment(OutBudgetId, DepartmentId string) bool {
	if v, ok := d.DepartmentMap[OutBudgetId]; ok {
		return v == DepartmentId
	} else {
		return false
	}
}

func (d *FenBeiTong) GetToken() (string, error) {
	if d.AccessToken == "" || time.Now().Unix() > d.TokenExpire.Unix() {
		//获取新Token，并且重置过期时间
		r, err := d.PostAuth("/open/api/auth/v1/dispense",
			url.Values{
				"app_id":  []string{d.AppId},
				"app_key": []string{d.AppKey},
			})
		if err != nil {
			return "", ErrorNetWork
		}
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return "", ErrorDecode
		}
		logrus.Info(string(b))
		tokenResp := AuthorizeResp{}

		err = json.Unmarshal(b, &tokenResp)
		if err != nil {
			return "", ErrorDecode
		}

		token := Token{}
		err = json.Unmarshal(bytes.NewBufferString(tokenResp.Data).Bytes(), &token)
		if err != nil {
			return "", ErrorDecode
		}
		if token.AccessToken == "" {
			return "", ErrorGetAccessToken
		}
		d.AccessToken = token.AccessToken
		d.TokenExpire = time.Now().Add(2*time.Hour - 60)
	}

	return d.AccessToken, nil

}

const (
	ServerApi = "https://open.fenbeitong.com"
)

func (d *FenBeiTong) PostAuth(url string, data url.Values) (resp *http.Response, err error) {
	return http.PostForm(d.ApiUrl+url, data)
}

func (d *FenBeiTong) Post(url string, data interface{}) (ret interface{}, err error) {
	token, err := d.GetToken()
	if err != nil {
		return "", err
	}
	resp, err := http.PostForm(d.ApiUrl+url, util.SignRequest(data, d.SignKey, token, d.AdminEmployId))

	if err != nil {
		return "", ErrorNetWork
	}

	baseresp := BaseResponse{}
	json.NewDecoder(resp.Body).Decode(&baseresp)

	if baseresp.Code == 0 {
		return baseresp.Data, nil
	} else {
		return nil, errors.New(baseresp.Msg)
	}
}
