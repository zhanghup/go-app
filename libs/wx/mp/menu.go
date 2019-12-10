package mp

import "fmt"

func (this *context) Menu() Imenu{
	return &menu{
		context: this,
	}
}

type Imenu interface {

}

type menu struct {
	context *context
}

func (this *menu) error(err interface{},fn string,i ... int) error{
	var s = ""
	switch err.(type) {
	case string:
		s = err.(string)
	case error:
		s = err.(error).Error()
	case Error:
		e := err.(Error)
		if e.Errcode == 0{
			return nil
		}
		s = fmt.Sprintf("%d: %s",e.Errcode,e.Errmsg)
	}

	if len(i) > 0{
		return fmt.Errorf("[自定义菜单_%d] - [%s] - %s",i[0],fn,s)
	}
	return fmt.Errorf("[自定义菜单] - [%s] - %s",fn,s)
}

