package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)
type app struct {
	Id string `json:"id"`
	Secret string `json:"secret"`
	JsCode string `json:"js_code"`
}
type result struct {
	SessionKey string `json:"session_key"`
	Openid string `json:"openid"`
}
// 从前端获取小程序code、appid、小程序秘钥等组合后的url
//通过https://api.weixin.qq.com/sns/jscode2session 获取openid和session_key,给前端返回openid
func GetOpenId(c *gin.Context) {
	app := new(app)
	app.Id = c.Query("id")
	app.Secret = c.Query("secret")
	app.JsCode = c.Query("js_code")
	fmt.Println("\n",app)
	url := "https://api.weixin.qq.com/sns/jscode2session?appid=" + app.Id + "&secret=" + app.Secret + "&js_code=" + app.JsCode + "&grant_type=authorization_code"
	fmt.Println("url:"+url)
	response, err := http.Get(url)
	if err != nil  || response.StatusCode != http.StatusOK {
		ResponseError(c, CodeLogin)
		return
	} else {
		var res result
		body, _ := ioutil.ReadAll(response.Body)  //使用ioutil读取请求体重的所有内容,并使用body进行接受.
		fmt.Println(string(body))
		_ = json.Unmarshal(body,&res)
		fmt.Printf("%#v\n", res)
		defer response.Body.Close()
		ResponseSucess(c, gin.H{
			"openid":res.Openid,
		})
	}
}

