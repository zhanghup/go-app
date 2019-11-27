package file

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
	"github.com/nfnt/resize"
	"github.com/pkg/errors"
	"github.com/zhanghup/go-app"
	"github.com/zhanghup/go-tools"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"io/ioutil"
	"os"
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
		idx := strings.LastIndex(hd.Filename, ".")
		endStr := ""
		if idx > -1 {
			endStr = hd.Filename[idx+1:]
		}

		old := app.Resource{}
		_, err = this.db.Table(old).Where("md5 = ?", md5).Get(&old)
		if err != nil {
			c.Fail400("读取文件失败【4】", err)
			return
		}

		ct := hd.Header.Get("Content-Type")

		if old.Id == nil {
			res := app.Resource{
				Bean: app.Bean{
					Id:     tools.ObjectString(),
					Status: tools.Ptr().Int(1),
				},
				ContentType: ct,
				MD5:         md5,
				Name:        hd.Filename,
				Size:        int64(len(data)),
				FileEnd:     endStr,
			}
			_, err = this.db.Insert(res)
			if err != nil {
				c.Fail400("读取文件失败【5】", err)
				return
			}
			_ = os.MkdirAll(fmt.Sprintf("upload/%s", ct), 0777)
			f, err := os.Create(fmt.Sprintf("upload/%s/%s.%s", ct, *res.Id, endStr))
			if err != nil {
				c.Fail400("读取写入失败【1】", err)
				return
			}
			defer f.Close()
			_, err = io.Copy(f, bytes.NewBuffer(data))
			if err != nil {
				c.Fail400("读取写入失败【2】", err)
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
func (this *Uploader) GetFile(id string) (*app.Resource, *os.File, error) {
	res := new(app.Resource)
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
		id := c.Param("id")
		if strings.LastIndex(id, ".") > 0 {
			id = id[:strings.LastIndex(id, ".")]
		}

		res, f, err := this.GetFile(id)
		if err != nil {
			c.Fail400("文件不存在", err)
			return
		}

		w := c.Writer
		w.Header().Set("Content-Type", res.ContentType)
		w.Header().Set("Content-Length", fmt.Sprintf("%d", res.Size))
		w.Header().Set("Content-Filename", res.Name)
		io.Copy(w, f)
	}
}

func (this *Uploader) Resize() func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")

		widthStr := c.Param("width")
		if strings.LastIndex(widthStr, ".") > 0 {
			widthStr = widthStr[:strings.LastIndex(widthStr, ".")]
		}
		width, err := strconv.ParseInt(widthStr, 10, 64)
		if err != nil {
			c.Fail400("输入的请求不合法【1】", err)
			return
		}

		var height int64
		heightStr := c.Param("height")
		if len(heightStr) != 0 {
			if strings.LastIndex(heightStr, ".") > 0 {
				heightStr = heightStr[:strings.LastIndex(heightStr, ".")]
			}
			height, err = strconv.ParseInt(heightStr, 10, 64)
			if err != nil {
				c.Fail400("输入的请求不合法【2】", err)
				return
			}
		}

		res, f, err := this.GetFile(id)
		if err != nil {
			c.Fail400("文件不存在", err)
			return
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
			c.Fail400("图片转换失败【2】", err)
			return
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
	}
}

func Gin(e *xorm.Engine, g *gin.Engine) {
	up := NewUploader(e)
	g.Group("/").POST("/upload", up.Upload())
	g.Group("/").GET("/upload/:id", up.Get())
	g.Group("/").GET("/upload/:id/:width/:height", up.Resize())
	g.Group("/").GET("/upload/:id/:width", up.Resize())
}
