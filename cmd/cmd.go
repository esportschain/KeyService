package cmd

import (
    "flag"
    "net/http"
    "log"
    //"fmt"
    "pkey_svr/routers"
    "pkey_svr/config"
    "fmt"
)

type Cmd struct {}

var (
    Setting config.Conf
)

func (cmd *Cmd) Run()  {

    // 解析命令行参数
    cmd.parseCommandArgs()

    // 运行web server
    cmd.runWeb()
}

func (cmd *Cmd) parseCommandArgs()  {
    // 配置文件
    Setting.Version = "0.1"
    flag.IntVar(&Setting.Port,"port", 8080, "http listen port")
    flag.StringVar(&Setting.BindIp, "ip", "0.0.0.0", "http listen ip")

    flag.Parse()
}

func (cmd *Cmd) runWeb()  {
    http.HandleFunc("/get_pkey", routers.GetPkey)

    BindAdd := fmt.Sprintf("%s:%d", Setting.BindIp, Setting.Port)
    log.Printf("listen %s\n", BindAdd)
    err := http.ListenAndServe(BindAdd, nil)
    if err != nil {
        log.Fatalln(err)
    }
}