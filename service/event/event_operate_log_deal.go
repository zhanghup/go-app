package event

//
//import (
//	"github.com/zhanghup/go-app/beans"
//	"github.com/zhanghup/go-tools"
//	"github.com/zhanghup/go-tools/tog"
//)
//
//const (
//	operate_log = "operate_log"
//)
//
//// ty: 数据库表名
//// opt: 操作 C/R/U/D/M
//// oid: 操作对象id
//// uid: 操作人id
//// msg: 操作结果
//// state: 操作结果
//// obj: 提交对象
//func OperateLogPush(ty, opt, oid, msg string, state int, obj interface{}, user beans.User) {
//	EventPublish(operate_log, ty, opt, oid, msg, state, obj, user)
//}
//func OperateLogSubscribe(fn func(ty, opt, oid, msg string, state int, obj interface{}, user beans.User)) {
//	EventSubscribe(operate_log, fn)
//}
//
//func operateLogSubscribeInit() {
//	OperateLogSubscribe(func(ty, opt, oid, msg string, state int, obj interface{}, user beans.User) {
//		model := beans.OperateLog{
//			Bean: beans.Bean{
//				Id:     tools.Ptr.Uid(),
//				Status: tools.Ptr.Int(1),
//			},
//			Type:   &ty,
//			Opt:    &opt,
//			Oid:    &oid,
//			Uid:    user.Id,
//			Uname:  user.Name,
//			State:  &state,
//			Msg:    &msg,
//			Object: tools.Ptr.String(tools.Str.JSONString(obj)),
//		}
//		_, err := db.Insert(model)
//		if err != nil {
//			tog.Error(err.Error())
//		}
//	})
//}
