package main

import (
	"context"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
	"time"
)

func main() {
	// 将 examplebucket-1250000000 和 COS_REGION 修改为用户真实的信息
	// 存储桶名称，由bucketname-appid 组成，appid必须填入，可以在COS控制台查看存储桶名称。https://console.cloud.tencent.com/cos5/bucket
	// COS_REGION 可以在控制台查看，https://console.cloud.tencent.com/cos5/bucket, 关于地域的详情见 https://cloud.tencent.com/document/product/436/6224
	u, err := url.Parse("https://happycar-1258086652.cos.ap-shanghai.myqcloud.com")
	if err != nil {
		panic(err)
	}
	// 用于Get Service 查询，默认全地域 service.cos.myqcloud.com
	su, _ := url.Parse("https://cos.COS_REGION.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u, ServiceURL: su}

	secID := "AKIDdIjwuc4cR0bpsgyFgMZ6kCqJkxf9MWYL"
	secretKey := "5zllVp30Tkl9OPANM0Rirrytdjo6RecM"
	// 1.永久密钥
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  secID,     // 替换为用户的 SecretId，请登录访问管理控制台进行查看和管理，https://console.cloud.tencent.com/cam/capi
			SecretKey: secretKey, // 替换为用户的 SecretKey，请登录访问管理控制台进行查看和管理，https://console.cloud.tencent.com/cam/capi
		},
	})

	// 获取预签名URL
	//presignedURL, err := client.Object.GetPresignedURL(
	//	context.Background(),
	//	http.MethodGet,
	//	"C7BF77A1-57AB-404E-9AC7-9B6F1400376A.jpeg",
	//	secID, secretKey, time.Second*20, nil,
	//)

	// 上传图片
	name := "abc.jpg"
	presignedURL, err := client.Object.GetPresignedURL(
		context.Background(),
		http.MethodPut,
		name,
		secID, secretKey, time.Hour*1, nil,
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(presignedURL)

	// 更安全的方式，我这边不用
	//client.Object.Get()
}
