package api

import (
	"flag"
	"testing"
)

func init() {
	FenBeiTongClient.Init(*flag.String("app_id", "", ""),
		*flag.String("app_key", "", ""),
		*flag.String("sign_key", "", ""),
		*flag.String("admin", "", ""),
		*flag.String("url", "", ""))
}

func TestFenBeiTong_ThridInfo(t *testing.T) {
	info, err := FenBeiTongClient.ThridInfo()
	t.Log(info, err)
}
