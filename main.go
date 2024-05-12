package main

import (
  "io"
  "io/ioutil"
  "log"
  "net/http"
)

func main() {
  http.HandleFunc("/get", handler) // 设置路由
  log.Println("Server is running at http://localhost:8000")
  log.Fatal(http.ListenAndServe(":8000", nil)) // 启动服务器
}

func handler(w http.ResponseWriter, r *http.Request) {
  // 检查请求方法是否为POST
  if r.Method != http.MethodPost {
    http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
    return
  }

  // 读取请求体中的URL
  body, err := ioutil.ReadAll(r.Body)
  if err != nil {
    http.Error(w, "Error reading request body", http.StatusInternalServerError)
    return
  }
  defer r.Body.Close()

  url := string(body) // 将请求体内容作为URL

  // 使用从请求体中获取的URL发送GET请求
  resp, err := http.Get(url)
  if err != nil {
    http.Error(w, "Unable to request data from the URL provided", http.StatusInternalServerError)
    return
  }
  defer resp.Body.Close()

  // 将响应内容直接复制到客户端响应体中
  _, err = io.Copy(w, resp.Body)
  if err != nil {
    http.Error(w, "Failed to read response from the URL provided", http.StatusInternalServerError)
  }
}
