package util

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"time"
)

func Md5(content string) string {
	w := md5.New()
	io.WriteString(w, content)
	return fmt.Sprintf("%x", w.Sum(nil))
}

/**
参数签名算法
*/
func GenSign(t time.Time,data ,signKey string ) string {
	return Md5(fmt.Sprintf("timestamp=%d&data=%s&sign_key=%s",t.Unix()*1000,data,signKey))
}

func SignRequest(data interface{}, signKey ,access_token,employee_id string) url.Values {
	b, _ := json.Marshal(data)
	t := time.Now()

	v := url.Values{
		"timestamp":[]string{fmt.Sprint("%d",t.Unix()*1000)},
		"access_token":[]string{access_token},
		"sign":[]string{GenSign(t,string(b), signKey)},
		"data":[]string{string(b)},
		"employee_id":[]string{employee_id},
		"employee_type":[]string{"0"},
	}

	return v
}
