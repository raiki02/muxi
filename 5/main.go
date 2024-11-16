package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

type Response struct {
	MSG  string `json:"msg"`
	Data string `json:"data"`
}

type Data struct {
	Text      string `json:"text"`
	ExtraInfo string `json:"extra_info"`
}

func main() {

	base := "http://muxithief.muxixyz.com/api/v1/"
	client := http.Client{}
	req, err := http.NewRequest(
		http.MethodPost,
		base+"login",
		nil,
	)

	req.Header.Add("code", "114514")
	if err != nil {
		fmt.Println("req err: ", err)
		return
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("post err: ", err)
		return
	}
	defer func() { _ = res.Body.Close() }()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("read res.body err: ", err)
		return
	}
	fmt.Printf("body content: %s \n", body)
	//data:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzE3NTk0ODcsInN1YiI6IjExNDUxNCJ9.h-Iw99U3nK468NonOEvXYayRNhkCnxHelbxYMaenTqw
	/*msg:恭喜你:114514,
	你已经成功登陆了XXHBGS的内部系统,
	但是请小心不要被抓获,
	现在你需要把返回的token作为值加到Authorization请求头中去(以后的请求都要加上)，
	并尝试去破坏这个系统。这个系统我们曾经攻破过，在这个网站里似乎能找到什么线索：http://muxithief.muxixyz.com/api/v1/getStrategy*/
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzE3NjM5ODUsInN1YiI6IjExNDUxNCJ9.MNFgtYuh8SlC2GAf4hVsSF7sNna3tj64lD5AFNjLXiQ"
	//token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzE3NTk0ODcsInN1YiI6IjExNDUxNCJ9.h-Iw99U3nK468NonOEvXYayRNhkCnxHelbxYMaenTqw"
	req, _ = http.NewRequest(
		http.MethodGet,
		base+"getStrategy",
		nil,
	)
	req.Header.Add("Authorization", token)
	res, _ = client.Do(req)
	body, _ = io.ReadAll(res.Body)
	fmt.Printf("body: %s \n", body)
	/*msg:似乎是某人的十年前的日记:
	2014年11月1日:木犀成立了有几个月了,我们克服了很多困难,虽然日子还是很艰难,但是会好的。
	2014年11月5日:今天XXHBGS突然想要把我们干掉,还抢走了我们宝贵的文献,我们几个月的成果全毁了。
	2014年11月10日：经过全体后端组连续5天不眠不修的努力我们终于找到了夺回我们文献的办法，我们知道了怎么攻击XXHBGS的内部网站
	2014年11月11日：我们对：http://muxithief.muxixyz.com/api/v1/attack 网站发起了全面进攻，
					在进攻的过程中我们发现原来我们的文献藏在：http://muxithief.muxixyz.com/api/v1/paper 。
	2014年11月16日：经过不断的攻击和探索最终我们发现只要我们在向：http://muxithief.muxixyz.com/api/v1/attack 进行疯狂的攻击的同时，尝试去访问：http://muxithief.muxixyz.com/api/v1/paper 就能够成功找回我们被抢走的文献。
	具体攻击方法如下:访问：http://muxithief.muxixyz.com/api/v1/eyes 获取工具用的图片,通过短时间内并发请求发送攻击图片的方式来扰乱系统。PS:(3s内达到5次以上访问,别攻击太狠把服务器 打挂了)
	*/

	req, _ = http.NewRequest(
		http.MethodGet,
		base+"eyes",
		nil,
	)
	req.Header.Add("Authorization", token)
	res, _ = client.Do(req)
	body, _ = io.ReadAll(res.Body)
	fmt.Printf("body: %s \n", body)

	var response Response
	if err := json.Unmarshal(body, &response); err != nil {
		log.Fatalf("JSON 解析失败: %v", err)
	}
	imageData, err := base64.StdEncoding.DecodeString(response.Data)
	if err != nil {
		log.Fatalf("Base64 解码失败: %v", err)
	}
	if err := os.WriteFile("1145145.jpg", imageData, 0644); err != nil {
		log.Fatalf("写入文件失败: %v", err)
		return
	}
	file, err := os.Open("1145145.jpg")
	if err != nil {
		log.Fatal("打开图片文件失败:", err)
	}
	defer file.Close()

	var abody bytes.Buffer
	writer := multipart.NewWriter(&abody)

	// 创建文件字段并将文件内容写入
	fileWriter, err := writer.CreateFormFile("file", "1145145.jpg")
	if err != nil {
		log.Fatal("创建文件字段失败:", err)
	}

	_, err = io.Copy(fileWriter, file)
	if err != nil {
		log.Fatal("写入文件内容失败:", err)
	}

	// 关闭 multipart writer 以设置结束边界
	writer.Close()

	// 创建 POST 请求
	req3, err := http.NewRequest("POST", base+"attack", &abody)
	if err != nil {
		log.Fatal("创建请求失败:", err)
	}

	// 添加请求头，包括动态生成的 Content-Type
	req3.Header.Add("Content-Type", writer.FormDataContentType())
	req3.Header.Add("Authorization", token)

	// 发送请求
	for i := 0; i < 50; i++ {
		go func() {
			client.Do(req3)
		}()
	}

	req, _ = http.NewRequest(
		http.MethodGet,
		base+"paper",
		nil,
	)
	req.Header.Add("Authorization", token)
	res, _ = client.Do(req)
	body, _ = io.ReadAll(res.Body)
	fmt.Printf("%s", body)

}
