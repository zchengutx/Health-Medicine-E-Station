package comment

import (
	"context"
	"errors"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"net/http"

	"mime"
	"path/filepath"
)

const (
	MinioEndpoint  = "14.103.134.228:9000"
	MinioAccessKey = "EsC0ClOoyswxNguwC6JP"
	MinioSecretKey = "6OOrbmDCIh1lZ0urTdLtkz4iWR8EPauw1Hyxtfzj"
	BucketName     = "chen"
)

func Upload(ctx context.Context, c *http.Request) (err error, url string) {
	// 1. 从请求中获取文件
	_, m, err := c.FormFile("file")
	if err != nil {
		return err, ""
	}
	// 2. 校验文件大小
	if m.Size >= 10*1024*1024 {
		return errors.New("file too big"), ""
	}

	// 3. 初始化minio客户端
	client, err := minio.New(MinioEndpoint, &minio.Options{
		Creds: credentials.NewStaticV4(MinioAccessKey, MinioSecretKey, ""),
	})
	if err != nil {
		return err, ""
	}
	// 4. 打开文件流
	open, err := m.Open()
	defer open.Close()
	// 5. 生成文件名称
	//location, _ := time.LoadLocation("Asia/Shanghai")
	// 6. 生成文件后缀
	//Now := time.Now().In(location).Truncate(time.Second)
	// 7. 生成文件名称
	//format := Now.Format("20060102150405")
	ext := filepath.Ext(m.Filename)
	fmt.Println(ext)
	extension := mime.TypeByExtension(ext)
	//sprintf := fmt.Sprintf("%s/%s", format, m.Filename)
	// 8. 上传文件
	client.PutObject(ctx, BucketName, m.Filename, open, m.Size, minio.PutObjectOptions{ContentType: mime.TypeByExtension(extension)})

	Url := fmt.Sprintf("%s/%s/%s", MinioEndpoint, BucketName, m.Filename)

	return nil, Url
}
