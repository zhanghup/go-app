package file

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nfnt/resize"
	"github.com/pkg/errors"
	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/tgin"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"xorm.io/xorm"
)

func NewUploader(db *xorm.Engine) *Uploader {
	return &Uploader{db: db}
}

type Uploader struct {
	db *xorm.Engine
}

func (this *Uploader) UploadIO(read io.Reader, name, contentType string) (string, error) {
	data, err := ioutil.ReadAll(read)
	if err != nil {
		return "[UploadIO] 读取文件失败【1】", err
	}

	md5 := tools.Crypto.MD5(data)
	idx := strings.LastIndex(name, ".")
	endStr := ""
	if idx > -1 {
		endStr = name[idx+1:]
	}

	old := beans.Resource{}
	_, err = this.db.Table(old).Where("md5 = ?", md5).Get(&old)
	if err != nil {
		return "[UploadIO] 读取文件失败【2】", err
	}

	if old.Id == nil {
		res := beans.Resource{
			Bean: beans.Bean{
				Id:     tools.Ptr.Uid(),
				Status: tools.Ptr.Int(1),
			},
			ContentType: contentType,
			MD5:         md5,
			Name:        name,
			Size:        int64(len(data)),
			FileEnd:     endStr,
		}
		_, err = this.db.Insert(res)
		if err != nil {
			return "[UploadIO] 写入文件失败【1】", err
		}
		_ = os.MkdirAll(fmt.Sprintf("upload/%s", contentType), 0777)
		f, err := os.Create(fmt.Sprintf("upload/%s/%s.%s", contentType, *res.Id, endStr))
		if err != nil {
			return "[UploadIO] 写入文件失败【2】", err
		}
		defer f.Close()
		_, err = io.Copy(f, bytes.NewBuffer(data))
		if err != nil {
			return "[UploadIO] 写入文件失败【3】", err
		}

		return *res.Id, nil
	} else {
		return *old.Id, nil
	}
}

func (this *Uploader) Upload() func(c *gin.Context) {
	return func(c *gin.Context) {
		tgin.Do(c, func(c *gin.Context) (interface{}, string) {
			hd, err := c.FormFile("file")
			if err != nil {
				return err.Error(), "读取文件失败【1】"
			}

			fd, err := hd.Open()
			if err != nil {
				return err.Error(), "读取文件失败【2】"
			}

			id, err := this.UploadIO(fd, hd.Filename, hd.Header.Get("Content-Type"))
			if err != nil {
				return "", err.Error()
			}
			return id, ""

		})

	}

}
func (this *Uploader) GetFile(id string) (*beans.Resource, *os.File, error) {
	res := new(beans.Resource)
	ok, err := this.db.Where("id = ?", id).Get(res)
	if err != nil {
		return nil, nil, err
	}
	if !ok {
		return nil, nil, errors.New("文件不存在")
	}
	f, err := os.Open(fmt.Sprintf("upload/%s/%s.%s", res.ContentType, *res.Id, res.FileEnd))
	return res, f, err
}

func (this *Uploader) Get() func(c *gin.Context) {
	return func(c *gin.Context) {
		tgin.DoCustom(c, func(c *gin.Context) (interface{}, string) {
			id := c.Param("id")
			if strings.LastIndex(id, ".") > 0 {
				id = id[:strings.LastIndex(id, ".")]
			}

			res, f, err := this.GetFile(id)
			if err != nil {
				return err.Error(), "文件不存在"
			}

			w := c.Writer
			w.Header().Set("Content-Type", res.ContentType)
			w.Header().Set("Content-Length", fmt.Sprintf("%d", res.Size))
			w.Header().Set("Content-Filename", res.Name)
			io.Copy(w, f)
			return nil, ""
		})

	}
}

func (this *Uploader) Resize() func(c *gin.Context) {
	return func(c *gin.Context) {
		tgin.DoCustom(c, func(c *gin.Context) (interface{}, string) {
			id := c.Param("id")

			widthStr := c.Param("width")
			if strings.LastIndex(widthStr, ".") > 0 {
				widthStr = widthStr[:strings.LastIndex(widthStr, ".")]
			}
			width, err := strconv.ParseInt(widthStr, 10, 64)
			if err != nil {
				return err.Error(), "请求不合法【1】"
			}

			var height int64
			heightStr := c.Param("height")
			if len(heightStr) != 0 {
				if strings.LastIndex(heightStr, ".") > 0 {
					heightStr = heightStr[:strings.LastIndex(heightStr, ".")]
				}
				height, err = strconv.ParseInt(heightStr, 10, 64)
				if err != nil {
					return err.Error(), "请求不合法【2】"
				}
			}

			res, f, err := this.GetFile(id)
			if err != nil {
				return err.Error(), "文件不存在"
			}

			var img image.Image
			switch res.ContentType {
			case "image/png":
				img, err = png.Decode(f)
			case "image/gif":
				img, err = gif.Decode(f)
			case "image/jpg":
				fallthrough
			case "image/jpeg":
				fallthrough
			default:
				img, err = jpeg.Decode(f)
			}
			if err != nil {
				return err.Error(), "图片转换失败【1】"
			}

			if len(heightStr) == 0 {
				height = int64(int(width) / img.Bounds().Dx() * img.Bounds().Dy())
			}

			m := resize.Resize(uint(width), uint(height), img, resize.Lanczos3)

			result := bytes.NewBuffer(nil)
			switch res.ContentType {
			case "image/png":
				err = png.Encode(result, m)
			case "image/gif":
				err = gif.Encode(result, m, nil)
			case "image/jpg":
				fallthrough
			case "image/jpeg":
				fallthrough
			default:
				err = jpeg.Encode(result, m, nil)
			}
			w := c.Writer
			w.Header().Set("Content-Type", res.ContentType)
			w.Header().Set("Content-Length", fmt.Sprintf("%d", result.Len()))
			w.Header().Set("Content-Filename", res.Name)
			w.Write(result.Bytes())
			return nil, ""
		})
	}
}

func Gin(auth, any gin.IRouter, db *xorm.Engine) {
	up := NewUploader(db)
	auth.POST("/upload", up.Upload())
	any.GET("/upload/:id", up.Get())
	any.GET("/upload/:id/:width/:height", up.Resize())
	any.GET("/upload/:id/:width", up.Resize())
}
