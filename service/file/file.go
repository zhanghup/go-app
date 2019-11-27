package file

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
	"github.com/nfnt/resize"
	"github.com/zhanghup/go-app"
	"github.com/zhanghup/go-tools"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"strconv"
	"strings"
)

func NewUploader(e *xorm.Engine) *Uploader {
	return &Uploader{e}
}

type Uploader struct {
	db *xorm.Engine
}

func (this *Uploader) Upload() func(c *gin.Context) {
	return func(c *gin.Context) {
		hd, err := c.FormFile("file")

		if err != nil {
			c.Fail400("读取文件失败【1】", err)
			return
		}

		fd, err := hd.Open()
		if err != nil {
			c.Fail400("读取文件失败【2】", err)
			return
		}

		data, err := ioutil.ReadAll(fd)
		if err != nil {
			c.Fail400("读取文件失败【3】", err)
			return
		}

		md5 := tools.MD5(data)
		old := app.Resource{}
		_, err = this.db.Table(old).Where("md5 = ?", md5).Get(&old)
		if err != nil {
			c.Fail400("读取文件失败【4】", err)
			return
		}

		if old.Id == nil {
			res := app.Resource{
				Bean: app.Bean{
					Id:     tools.ObjectString(),
					Status: tools.Ptr().Int(1),
				},
				ContentType: c.Request.Header.Get("Content-Type"),
				Type:        hd.Header.Get("Content-Type"),
				MD5:         md5,
				Name:        hd.Filename,
				Size:        int64(len(data)),
				Datas:       data,
			}
			_, err = this.db.Insert(res)
			if err != nil {
				c.Fail400("读取文件失败【5】", err)
				return
			}
			c.Success(*res.Id)
			return
		} else {
			c.Success(*old.Id)
			return
		}
	}

}
func (this *Uploader) Get() func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		if strings.LastIndex(id, ".") > 0 {
			id = id[:strings.LastIndex(id, ".")]
		}

		res := new(app.Resource)
		ok, err := this.db.Where("id = ?", id).Get(res)
		if err != nil || !ok {
			c.Fail400("查找文件失败", err)
			return
		}

		w := c.Writer
		w.Header().Set("Content-Type", res.Type)
		w.Header().Set("Content-Length", fmt.Sprintf("%d", res.Size))
		w.Header().Set("Content-Filename", res.Name)
		w.Write(res.Datas)
	}
}
func (this *Uploader) Resize() func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")

		widthStr := c.Param("width")
		width, err := strconv.ParseInt(widthStr, 10, 64)
		if err != nil {
			c.Fail400("输入的请求不合法【1】", err)
			return
		}

		heightStr := c.Param("width")
		if strings.LastIndex(heightStr, ".") > 0 {
			heightStr = heightStr[:strings.LastIndex(heightStr, ".")]
		}
		height, err := strconv.ParseInt(heightStr, 10, 64)
		if err != nil {
			c.Fail400("输入的请求不合法【2】", err)
			return
		}

		res := new(app.Resource)
		ok, err := this.db.Where("id = ?", id).Get(res)
		if err != nil || !ok {
			c.Fail400("查找文件失败【1】", err)
			return
		}

		var img image.Image
		switch res.Type {
		case "image/png":
			img, err = png.Decode(bytes.NewBuffer(res.Datas))
		case "image/gif":
			img, err = gif.Decode(bytes.NewBuffer(res.Datas))
		case "image/jpg":
			fallthrough
		case "image/jpeg":
			fallthrough
		default:
			img, err = jpeg.Decode(bytes.NewBuffer(res.Datas))
		}
		if err != nil {
			c.Fail400("图片转换失败【2】", err)
			return
		}

		m := resize.Resize(uint(width), uint(height), img, resize.Lanczos3)

		result := bytes.NewBuffer(nil)
		switch res.Type {
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
		w.Header().Set("Content-Type", res.Type)
		w.Header().Set("Content-Length", fmt.Sprintf("%d", result.Len()))
		w.Header().Set("Content-Filename", res.Name)
		w.Write(result.Bytes())
	}
}

func Gin(e *xorm.Engine, g *gin.Engine) {
	up := NewUploader(e)
	g.Group("/").POST("/upload", up.Upload())
	g.Group("/").GET("/upload/:id", up.Get())
	g.Group("/").GET("/image/:id/:width/:height", up.Resize())
}
