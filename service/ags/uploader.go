package ags

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

type IUploader interface {
	UploadIO(read io.Reader, name, contentType string) (string, error)
	GetFile(id string) (*beans.Resource, *os.File, error)

	GinUpload() func(c *gin.Context)
	GinGet() func(c *gin.Context)
	GinResize() func(c *gin.Context)
	GinRouter(auth gin.IRouter, any gin.IRouter)
}

type uploader struct {
	db *xorm.Engine
}

func (this *uploader) UploadIO(read io.Reader, name, contentType string) (string, error) {
	data, err := ioutil.ReadAll(read)
	if err != nil {
		return "[UploadIO] 读取文件失败【1】", err
	}

	md5 := tools.MD5(data)
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
				Id:     tools.PtrOfUUID(),
				Status: tools.PtrOfString("1"),
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
		_ = os.MkdirAll(fmt.Sprintf("upload/%s/%s", tools.Time.YM(), contentType), 0777)
		f, err := os.Create(fmt.Sprintf("upload/%s/%s/%s.%s", tools.Time.YM(), contentType, *res.Id, endStr))
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
func (this *uploader) GetFile(id string) (*beans.Resource, *os.File, error) {
	res := new(beans.Resource)
	ok, err := this.db.Where("id = ?", id).Get(res)
	if err != nil {
		return nil, nil, err
	}
	if !ok {
		return nil, nil, errors.New("文件不存在")
	}
	f, err := os.Open(fmt.Sprintf("upload/%s/%s/%s.%s", tools.Time.UnixYM(*res.Created), res.ContentType, *res.Id, res.FileEnd))
	return res, f, err
}

func (this *uploader) GinUpload() func(c *gin.Context) {
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
func (this *uploader) GinGet() func(c *gin.Context) {
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
func (this *uploader) GinResize() func(c *gin.Context) {
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
func (this *uploader) GinRouter(auth gin.IRouter, any gin.IRouter) {
	auth.POST("/upload", this.GinUpload())
	any.GET("/upload/:id", this.GinGet())
	any.GET("/upload/:id/:width/:height", this.GinResize())
	any.GET("/upload/:id/:width", this.GinResize())
}

var defaultUploader IUploader

/*
	初始化上传工具
	@db: db为空，初始化默认上传器
		 db不为空，返回一个新的上传器，但是默认的不会被替换
*/
func UploaderNew(db ...*xorm.Engine) IUploader {
	if defaultUploader != nil && len(db) == 0 {
		return defaultUploader
	}

	var newUp *xorm.Engine
	if len(db) == 0 {
		newUp = defaultDB
	} else {
		newUp = db[0]
	}

	up := &uploader{
		db: newUp,
	}

	if len(db) == 0 {
		defaultUploader = up
		return defaultUploader
	} else {
		return up
	}
}
