package routers

import (
    "net/http"
    "io/ioutil"
    "log"
    "encoding/json"
    "strings"
    //"github.com/vmihailenco/msgpack"
    "fmt"
    "time"
    "math/rand"
)

type Token struct  {
    Sig string `json:"sig"`
    Tk string `json:"tk"`       // 请求的token
}

// 处理秘钥请求
func GetPkey(resp http.ResponseWriter, req *http.Request)  {
    var tk Token
    err := readBody(resp, req, &tk)
    if err != nil {
        return
    }

    tk.Sig = strings.TrimSpace(tk.Sig)
    tk.Tk = strings.TrimSpace(tk.Tk)
    if tk.Tk == "" {
        resp.Write(generateFailureBody("tk 不能为空"))
        return
    }

    if tk.Sig == "" {
        resp.Write(generateFailureBody("sig 不能为空"))
        return
    }

    log.Printf("add req#%+v\n", req)
    log.Printf("add tk#%+v\n", tk)

    if err != nil {
        resp.Write(generateFailureBody("failure"))
    } else {
        privateKey := []byte("*************************************************")
        str := authKey(privateKey)
        resp.Write(generateSuccessBody("success", []byte(str)))
    }
}


type ResponseBody struct {
    Code int       `msg:"code"`
    Message string `msg:"message"`
    Data []byte `msg:"data"`
}

func readBody(resp http.ResponseWriter, req *http.Request, v interface{}) error {
    body, err := ioutil.ReadAll(req.Body)
    if err != nil {
        log.Printf("read body fail#%s", err.Error())
        resp.Write(generateFailureBody("read request body fail"))
        return err
    }
    err = json.Unmarshal(body, v)
    if err != nil {
        log.Printf("json fail#%s", err.Error())
        resp.Write(generateFailureBody("json fail"))
        return err
    }

    return nil
}

func generateSuccessBody(msg string, data []byte) ([]byte)  {
     return generateResponseBody(0, msg, data)
}

func generateFailureBody(msg string) ([]byte) {
    return generateResponseBody(1, msg, nil)
}

func generateResponseBody(code int, msg string, data []byte) ([]byte)  {
    body := &ResponseBody{}
    body.Code = code
    body.Message = msg
    body.Data = data

    bytes, err := json.Marshal(body)
    if err != nil {
        log.Printf("response body to json fail#%s", err.Error())
        return []byte(`{"code":"1", "message": "body fail", "data":[]}`)
    }

    return bytes
}

/**
混淆私钥
 */
func authKey(privateKey []byte) string {
    var str string
    for i := range privateKey {
        str = fmt.Sprintf("%s%c", str, privateKey[i])
        if i % 7 == 0 {
            str = fmt.Sprintf("%s%s", str, randStr(1))
        }
    }

    return str
}


/**
生成随机字符串
 */
func randStr(strLen int) string {
    data := make([]byte, strLen)
    rand.Seed(time.Now().UnixNano())
    var num int
    for i := 0; i < strLen; i++ {
        num = rand.Intn(57) + 65
        for {
            if num > 90 && num < 97 {
                num = rand.Intn(57) + 65
            } else {
                break
            }
        }
        data[i] = byte(num)
    }
    return strings.ToLower(string(data))
}
