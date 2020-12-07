package initia

import (
	"fmt"
	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/database/txorm"
	"math/rand"
	"xorm.io/xorm"
)

func InitUser(db *xorm.Engine) {
	ok, err := db.Table(beans.User{}).Where("id = ?", "root").Exist()
	if err != nil {
		panic(err)
	}
	if ok {
		return
	}

	err = txorm.NewEngine(db).TS(func(sess txorm.ISession) error {

		user := beans.User{
			Bean: beans.Bean{
				Id:     tools.Ptr.String("root"),
				Status: tools.Ptr.String("1"),
				Weight: tools.Ptr.Int(0),
			},
		}
		err := sess.Insert(user)
		if err != nil {
			return err
		}

		salt := tools.Str.Uid()
		password := tools.Crypto.Password("Aa123456.", salt)
		err = sess.Insert(beans.Account{
			Bean: beans.Bean{
				Id:     tools.Ptr.String("root"),
				Status: tools.Ptr.String("1"),
				Weight: tools.Ptr.Int(0),
			},
			Type:     tools.Ptr.String("password"),
			Uid:      user.Id,
			Username: tools.Ptr.String("root"),
			Password: &password,
			Salt:     &salt,
		})
		return err
	})

	if err != nil {
		panic(err)
	}
}

func InitTest(db *xorm.Engine) {
	InitTestDept(db)
	InitTestUser(db)
	InitTestRole(db)
}

func InitTestDept(db *xorm.Engine) {
	db.Table(beans.Dept{}).Insert(map[string]interface{}{"id": "1", "name": "浙江", "type": "1", "code": "zj","status":"1"})
	db.Table(beans.Dept{}).Insert(map[string]interface{}{"id": "11", "name": "嘉兴", "pid": "1", "type": "1", "code": "jx","status":"1"})
	db.Table(beans.Dept{}).Insert(map[string]interface{}{"id": "12", "name": "宁波", "pid": "1", "type": "1", "code": "nb","status":"1"})
	db.Table(beans.Dept{}).Insert(map[string]interface{}{"id": "13", "name": "杭州", "pid": "1", "type": "1", "code": "hz","status":"1"})
	db.Table(beans.Dept{}).Insert(map[string]interface{}{"id": "14", "name": "绍兴", "pid": "1", "type": "1", "code": "sx","status":"1"})
	db.Table(beans.Dept{}).Insert(map[string]interface{}{"id": "15", "name": "舟山", "pid": "1", "type": "1", "code": "zs","status":"1"})
	db.Table(beans.Dept{}).Insert(map[string]interface{}{"id": "16", "name": "衢州", "pid": "1", "type": "1", "code": "qz","status":"1"})
	db.Table(beans.Dept{}).Insert(map[string]interface{}{"id": "2", "name": "江苏", "type": "1", "code": "js","status":"1"})
	db.Table(beans.Dept{}).Insert(map[string]interface{}{"id": "21", "name": "无锡", "pid": "2", "type": "1", "code": "wx","status":"1"})
	db.Table(beans.Dept{}).Insert(map[string]interface{}{"id": "22", "name": "苏州", "pid": "2", "type": "1", "code": "sz","status":"1"})
	db.Table(beans.Dept{}).Insert(map[string]interface{}{"id": "23", "name": "南京", "pid": "2", "type": "1", "code": "nj","status":"1"})
	db.Table(beans.Dept{}).Insert(map[string]interface{}{"id": "24", "name": "徐州", "pid": "2", "type": "1", "code": "xz","status":"1"})
	db.Table(beans.Dept{}).Insert(map[string]interface{}{"id": "25", "name": "常州", "pid": "2", "type": "1", "code": "cz","status":"1"})
	db.Table(beans.Dept{}).Insert(map[string]interface{}{"id": "26", "name": "扬州", "pid": "2", "type": "1", "code": "yz","status":"1"})
	db.Table(beans.Dept{}).Insert(map[string]interface{}{"id": "3", "name": "上海", "type": "1", "code": "sh","status":"1"})
	db.Table(beans.Dept{}).Insert(map[string]interface{}{"id": "31", "name": "嘉定", "pid": "3", "type": "1", "code": "jd","status":"1"})
	db.Table(beans.Dept{}).Insert(map[string]interface{}{"id": "32", "name": "普陀", "pid": "3", "type": "1", "code": "pt","status":"1"})
	db.Table(beans.Dept{}).Insert(map[string]interface{}{"id": "33", "name": "浦东", "pid": "3", "type": "1", "code": "pd","status":"1"})
	db.Table(beans.Dept{}).Insert(map[string]interface{}{"id": "34", "name": "徐汇", "pid": "3", "type": "1", "code": "xh","status":"1"})
	db.Table(beans.Dept{}).Insert(map[string]interface{}{"id": "35", "name": "宝山", "pid": "3", "type": "1", "code": "bs","status":"1"})
	db.Table(beans.Dept{}).Insert(map[string]interface{}{"id": "36", "name": "松江", "pid": "3", "type": "1", "code": "sj","status":"1"})
	db.Table(beans.Dept{}).Insert(map[string]interface{}{"id": "4", "name": "安徽", "type": "1", "code": "ah","status":"1"})
	db.Table(beans.Dept{}).Insert(map[string]interface{}{"id": "41", "name": "合肥", "pid": "4", "type": "1", "code": "hf","status":"1"})
	db.Table(beans.Dept{}).Insert(map[string]interface{}{"id": "42", "name": "芜湖", "pid": "4", "type": "1", "code": "wh","status":"1"})
	db.Table(beans.Dept{}).Insert(map[string]interface{}{"id": "43", "name": "蚌埠", "pid": "4", "type": "1", "code": "ph","status":"1"})
	db.Table(beans.Dept{}).Insert(map[string]interface{}{"id": "44", "name": "淮南", "pid": "4", "type": "1", "code": "hn","status":"1"})
	db.Table(beans.Dept{}).Insert(map[string]interface{}{"id": "45", "name": "淮北", "pid": "4", "type": "1", "code": "hb","status":"1"})
	db.Table(beans.Dept{}).Insert(map[string]interface{}{"id": "46", "name": "六安", "pid": "4", "type": "1", "code": "la","status":"1"})
}

func InitTestUser(db *xorm.Engine) {
	names := []string{"澄邈", "德泽", "海超", "海阳", "海荣", "海逸", "海昌", "瀚钰", "瀚文", "涵亮", "涵煦", "明宇", "涵衍", "浩皛", "浩波", "浩博", "浩初", "浩宕", "浩歌", "浩广", "浩邈", "浩气", "浩思", "浩言", "鸿宝", "鸿波", "鸿博", "鸿才", "鸿畅", "鸿畴", "鸿达", "鸿德", "鸿飞", "鸿风", "鸿福", "鸿光", "鸿晖", "鸿朗", "鸿文", "鸿轩", "鸿煊", "鸿骞", "鸿远", "鸿云", "鸿哲", "鸿祯", "鸿志", "鸿卓", "嘉澍", "光济", "澎湃", "彭泽", "鹏池", "鹏海", "浦和", "浦泽", "瑞渊", "越泽", "博耘", "德运", "辰宇", "辰皓", "辰钊", "辰铭", "辰锟", "辰阳", "辰韦", "辰良", "辰沛", "晨轩", "晨涛", "晨濡", "晨潍", "鸿振", "吉星", "铭晨", "起运", "运凡", "运凯", "运鹏", "运浩", "运诚", "运良", "运鸿", "运锋", "运盛", "运升", "运杰", "运珧", "运骏", "运凯", "运乾", "维运", "运晟", "运莱", "运华", "耘豪", "星爵", "星腾", "星睿", "星泽", "星鹏", "星然", "震轩", "震博", "康震", "震博", "振强", "振博", "振华", "振锐", "振凯", "振海", "振国", "振平", "昂然", "昂雄", "昂杰", "昂熙", "昌勋", "昌盛", "昌淼", "昌茂", "昌黎", "昌燎", "昌翰", "晨朗", "德明", "德昌", "德曜", "范明", "飞昂", "高旻", "晗日", "昊然", "昊天", "昊苍", "昊英", "昊宇", "昊嘉", "昊明", "昊伟", "昊硕", "昊磊", "昊东", "鸿晖", "鸿朗", "华晖", "金鹏", "晋鹏", "敬曦", "景明", "景天", "景浩", "俊晖", "君昊", "昆琦", "昆鹏", "昆纬", "昆宇", "昆锐", "昆卉", "昆峰", "昆颉", "昆谊", "昆皓", "昆鹏", "昆明", "昆杰", "昆雄", "昆纶", "鹏涛", "鹏煊", "曦晨", "曦之", "新曦", "旭彬", "旭尧", "旭鹏", "旭东", "旭炎", "炫明", "宣朗", "学智", "轩昂", "彦昌", "曜坤", "曜栋", "曜文", "曜曦", "曜灿", "曜瑞", "智伟", "智杰", "智刚", "智阳", "昌勋", "昌盛", "昌茂", "昌黎", "昌燎", "昌翰", "晨朗", "昂然", "昂雄", "昂杰", "昂熙", "范明", "飞昂", "高朗", "高旻", "德明", "德昌", "德曜", "智伟", "智杰", "智刚", "智阳", "瀚彭", "旭炎", "宣朗", "学智", "昊然", "昊天", "昊苍", "昊英", "昊宇", "昊嘉", "昊明", "昊伟", "鸿朗", "华晖", "金鹏", "晋鹏", "敬曦", "景明", "景天", "景浩", "景行", "景中", "景逸", "景彰", "昆鹏", "昆明", "昆杰", "昆雄", "昆纶", "鹏涛", "鹏煊", "景平", "俊晖", "君昊", "昆琦", "昆鹏", "昆纬", "昆宇", "昆锐", "昆卉", "昆峰", "昆颉", "昆谊", "轩昂", "彦昌", "曜坤", "曜文", "曜曦", "曜灿", "曜瑞", "曦晨", "曦之", "新曦", "鑫鹏", "旭彬", "旭尧", "旭鹏", "旭东", "浩轩", "浩瀚", "浩慨", "浩阔", "鸿熙", "鸿羲", "鸿禧", "鸿信", "泽洋", "泽雨", "哲瀚", "胤运", "佑运", "允晨", "运恒", "运发", "云天", "耘志", "耘涛", "振荣", "振翱", "中震", "子辰", "晗昱", "瀚玥", "瀚昂", "瀚彭", "景行", "景中", "景逸", "景彰", "绍晖", "文景", "曦哲", "永昌", "子昂", "智宇", "智晖", "晗日", "晗昱", "瀚昂", "昊硕", "昊磊", "昊东", "鸿晖", "绍晖", "文昂", "文景", "曦哲", "永昌", "子昂", "智宇", "智晖", "浩然", "鸿运", "辰龙", "运珹", "振宇", "高朗", "景平", "鑫鹏", "昌淼", "炫明", "昆皓", "曜栋", "文昂", "治汇", "恨桃", "依秋", "依波", "香巧", "紫萱", "涵易", "忆之", "幻巧", "美倩", "安寒", "白亦", "惜玉", "碧春", "怜雪", "听南", "念蕾", "紫夏", "凌旋", "芷梦", "凌寒", "梦竹", "千凡", "丹蓉", "慧贞", "思菱", "平卉", "笑柳", "雪卉", "南蓉", "谷梦", "巧兰", "绿蝶", "飞荷", "佳蕊", "芷荷", "怀瑶", "慕易", "若芹", "紫安", "曼冬", "寻巧", "雅昕", "尔槐", "以旋", "初夏", "依丝", "怜南", "傲菡", "谷蕊", "笑槐", "飞兰", "笑卉", "迎荷", "佳音", "梦君", "妙绿", "觅雪", "寒安", "沛凝", "白容", "乐蓉", "映安", "依云", "映冬", "凡雁", "梦秋", "梦凡", "秋巧", "若云", "元容", "怀蕾", "灵寒", "天薇", "翠安", "乐琴", "宛南", "怀蕊", "白风", "访波", "亦凝", "易绿", "夜南", "曼凡", "亦巧", "青易", "冰真", "白萱", "友安", "海之", "小蕊", "又琴", "天风", "若松", "盼菡", "秋荷", "香彤", "语梦", "惜蕊", "迎彤", "沛白", "雁彬", "易蓉", "雪晴", "诗珊", "春冬", "晴钰", "冰绿", "半梅", "笑容", "沛凝", "映秋", "盼烟", "晓凡", "涵雁", "问凝", "冬萱", "晓山", "雁蓉", "梦蕊", "山菡", "南莲", "飞双", "凝丝", "思萱", "怀梦", "雨梅", "冷霜", "向松", "迎丝", "迎梅", "雅彤", "香薇", "以山", "碧萱", "寒云", "向南", "书雁", "怀薇", "思菱", "忆文", "翠巧", "书文", "若山", "向秋", "凡白", "绮烟", "从蕾", "天曼", "又亦", "从语", "绮彤", "之玉", "凡梅", "依琴", "沛槐", "又槐", "元绿", "安珊", "夏之", "易槐", "宛亦", "白翠", "丹云", "问寒", "易文", "傲易", "青旋", "思真", "雨珍", "幻丝", "代梅", "盼曼", "妙之", "半双", "若翠", "初兰", "惜萍", "初之", "宛丝", "寄南"}

	for _, name := range names {
		nn := fmt.Sprintf("%s", name)
		dept := fmt.Sprintf("%d%d", rand.Intn(4)+1, rand.Intn(6)+1)
		py := tools.Pin.Py(nn)
		pinyin := tools.Pin.Pinyin(nn)
		db.Insert(beans.User{
			Bean: beans.Bean{
				Id:     tools.Ptr.Uid(),
				Status: tools.Ptr.String("1"),
			},
			Type:   tools.Ptr.String("1"),
			Mobile: tools.Ptr.String(fmt.Sprintf("15%d%d%d%d%d%d%d%d%d", rand.Intn(10), rand.Intn(10), rand.Intn(10), rand.Intn(10), rand.Intn(10), rand.Intn(10), rand.Intn(10), rand.Intn(10), rand.Intn(10))),
			Name:   &nn,
			Dept:   &dept,
			Py:     &py,
			Pinyin: &pinyin,
		})
	}
}

func InitTestRole(db *xorm.Engine) {
	db.Table(beans.Role{}).Insert(map[string]interface{}{"id": "1", "name": "管理员", "desc": "管理员拥有极大的权限","status":"1"})
	db.Table(beans.Role{}).Insert(map[string]interface{}{"id": "2", "name": "小职员", "desc": "","status":"1"})
	db.Table(beans.Role{}).Insert(map[string]interface{}{"id": "3", "name": "经理", "desc": "","status":"1"})
	db.Table(beans.Role{}).Insert(map[string]interface{}{"id": "4", "name": "总经理", "desc": "","status":"1"})
	db.Table(beans.Role{}).Insert(map[string]interface{}{"id": "5", "name": "董事长", "desc": "","status":"1"})
	db.Table(beans.Role{}).Insert(map[string]interface{}{"id": "6", "name": "特派员", "desc": "","status":"1"})
	db.Table(beans.Role{}).Insert(map[string]interface{}{"id": "7", "name": "打字员", "desc": "","status":"1"})
	db.Table(beans.Role{}).Insert(map[string]interface{}{"id": "8", "name": "程序员", "desc": "","status":"1"})
	db.Table(beans.Role{}).Insert(map[string]interface{}{"id": "9", "name": "研发部", "desc": "","status":"1"})
	db.Table(beans.Role{}).Insert(map[string]interface{}{"id": "10", "name": "销售部", "desc": "","status":"1"})
	db.Table(beans.Role{}).Insert(map[string]interface{}{"id": "11", "name": "实施部", "desc": "","status":"1"})
	db.Table(beans.Role{}).Insert(map[string]interface{}{"id": "12", "name": "人事部", "desc": "","status":"1"})
	db.Table(beans.Role{}).Insert(map[string]interface{}{"id": "13", "name": "测试部", "desc": "","status":"1"})
}
