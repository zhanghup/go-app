package initia

import (
	"fmt"
	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-tools"
	"math/rand"
	"strings"
	"xorm.io/xorm"
)

var test_depts = []map[string]interface{}{
	{"id": "1", "name": "销售部", "type": "1", "code": "1", "status": "1"},
	{"id": "11", "name": "业务拓展部", "pid": "1", "type": "1", "code": "11", "status": "1"},
	{"id": "12", "name": "市场部", "pid": "1", "type": "1", "code": "12", "status": "1"},
	{"id": "13", "name": "客服部", "pid": "1", "type": "1", "code": "13", "status": "1"},
	{"id": "14", "name": "成品仓", "pid": "1", "type": "1", "code": "14", "status": "1"},
	{"id": "15", "name": "策划部", "pid": "1", "type": "1", "code": "15", "status": "1"},
	{"id": "2", "name": "人力资源部", "type": "1", "code": "2", "status": "1"},
	{"id": "21", "name": "人事部", "pid": "2", "type": "1", "code": "21", "status": "1"},
	{"id": "22", "name": "行政部", "pid": "2", "type": "1", "code": "22", "status": "1"},
	{"id": "221", "name": "保安部", "pid": "22", "type": "1", "code": "221", "status": "1"},
	{"id": "222", "name": "清洁工", "pid": "22", "type": "1", "code": "222", "status": "1"},
	{"id": "223", "name": "食堂", "pid": "22", "type": "1", "code": "223", "status": "1"},
	{"id": "3", "name": "生产部", "type": "1", "code": "3", "status": "1"},
	{"id": "31", "name": "裁床", "pid": "3", "type": "1", "code": "31", "status": "1"},
	{"id": "32", "name": "车间", "pid": "3", "type": "1", "code": "32", "status": "1"},
	{"id": "321", "name": "成件组", "pid": "32", "type": "1", "code": "321", "status": "1"},
	{"id": "322", "name": "流水组", "pid": "32", "type": "1", "code": "322", "status": "1"},
	{"id": "33", "name": "后道", "pid": "3", "type": "1", "code": "33", "status": "1"},
	{"id": "331", "name": "后道1", "pid": "33", "type": "1", "code": "331", "status": "1"},
	{"id": "332", "name": "后道2", "pid": "33", "type": "1", "code": "332", "status": "1"},
	{"id": "34", "name": "辅料仓", "pid": "3", "type": "1", "code": "34", "status": "1"},
	{"id": "35", "name": "采购部", "pid": "3", "type": "1", "code": "35", "status": "1"},
	{"id": "36", "name": "品控部", "pid": "3", "type": "1", "code": "36", "status": "1"},
	{"id": "361", "name": "跟单组", "pid": "36", "type": "1", "code": "361", "status": "1"},
	{"id": "4", "name": "财务部", "type": "1", "code": "4", "status": "1"},
	{"id": "41", "name": "总账会计", "pid": "4", "type": "1", "code": "41", "status": "1"},
	{"id": "42", "name": "应收应付会计", "pid": "4", "type": "1", "code": "42", "status": "1"},
	{"id": "43", "name": "出纳", "pid": "4", "type": "1", "code": "43", "status": "1"},
	{"id": "44", "name": "会计文员", "pid": "4", "type": "1", "code": "44", "status": "1"},
	{"id": "5", "name": "设计部", "type": "1", "code": "5", "status": "1"},
	{"id": "51", "name": "设计部", "pid": "5", "type": "1", "code": "51", "status": "1"},
	{"id": "52", "name": "技术部", "pid": "5", "type": "1", "code": "52", "status": "1"},
}

func InitTest(db *xorm.Engine) {
	InitTestDept(db)
	InitTestUser(db)
	InitTestRole(db)
}

func InitTestDept(db *xorm.Engine) {
	datas := make([]interface{}, 0)
	for _, s := range test_depts {
		datas = append(datas, s)
	}
	db.Table(beans.Dept{}).Insert(datas...)
}

func InitTestUser(db *xorm.Engine) {
	deptIds := make([]string, 0)
	for _, d := range test_depts {
		deptIds = append(deptIds, d["id"].(string))
	}

	names := []string{"张千渔", "张小驭", "张开娥", "张思宏", "张启然", "张子鑫", "张今", "张中文", "张高炎", "张永匀", "张晨荣", "张辉君", "张宏泽", "张明莲", "张昊钟", "张文鸿", "张骏意", "张泓超", "张漾东", "张仕松", "张丽菱", "张观桓", "张悠君", "张子兵", "张医水", "张之石", "张丰隽", "张亚林", "张炅学", "张梅宇", "张同结", "张华明", "张汶洲", "张悠梅", "张昕腾", "张才涛", "张同丽", "张依飞", "张雨勇", "张泓瑶", "张彦贞", "张红强", "张皙萱", "张建英", "张福轩", "张业丽", "张佳芮", "张梓怡", "张学美", "张盛平", "张剑开", "张芮仪", "张泞惠", "张璞", "张点丽", "张思崴", "张姿宇", "张恩才", "张青霏", "张愉", "张维叶", "张秀诚", "张一宾", "张相时", "张昱哲", "张采铭", "张翌天", "张海岳", "张亮杰", "张存原", "张雁舒", "张乐锋", "张韶惠", "张亚博", "张旎婷", "张梓远", "张十锦", "张丹祥", "张季霖", "张义涵", "张雄月", "张俊竣", "张琼铭", "张秋茵", "张亦熠", "张晨华", "张俞麾", "张琦娟", "张瀚梅", "张凡方", "张皓仪", "张春晶", "张辰冉", "蒋梓润", "蒋畅", "蒋贺", "蒋君浩", "蒋俊博", "蒋小萱", "蒋意辉", "蒋箐", "蒋锦", "蒋迩谷", "蒋坤", "蒋俊琪", "蒋新红", "蒋全领", "蒋增佳", "蒋庸", "蒋斐", "蒋开力", "蒋文兴", "蒋天彬", "蒋士倩", "蒋克", "蒋婉", "蒋锦明", "蒋晏", "蒋泉良", "蒋臻设", "蒋程林", "蒋咏爱", "蒋鹏明", "蒋耀儒", "蒋韩博", "蒋艺欣", "蒋巨麟", "蒋良珍", "蒋虹胜", "蒋潇宏", "蒋安彤", "蒋亦良", "蒋润全", "蒋立棠", "蒋自芮", "蒋乃盈", "蒋霓", "蒋桂骏", "蒋雅凤", "蒋斌", "蒋亦鹏", "蒋昀熙", "蒋政玉", "蒋颜", "蒋晓希", "蒋宗丽", "蒋浩钰", "蒋小微", "蒋奕", "蒋桢宸", "蒋中今", "蒋朝仪", "蒋晓新", "蒋子铄", "蒋星舒", "蒋子语", "蒋怡晴", "蒋春德", "蒋姗", "蒋金煜", "蒋帅雯", "蒋贞元", "蒋俏瞳", "蒋丽铭", "蒋军", "蒋春荷", "蒋岩家", "蒋蕴", "蒋梓睿", "蒋晓彤", "蒋雨萍", "蒋涵涛", "蒋佳雯", "蒋雍宇", "蒋佳嘉", "蒋宝萌", "蒋迪", "蒋锌", "蒋坡", "蒋金", "蒋广琪", "蒋双明", "蒋嘉葭", "蒋泓杰", "蒋宁", "马雨坤", "马芯", "马钰林", "马玉景", "马春淇", "马友成", "马梅", "马麟", "马程辉", "马荣平", "马尉文", "马丽杨", "马晟玉", "马怡菡", "马凯诚", "马誉红", "马彦潼", "马逸霞", "马梓杭", "马夕苗", "马煦", "马思海", "马米宵", "马竞阳", "马大羽", "马爱冰", "马亚茜", "马克鑫", "马一涵", "马志玉", "马瀛鑫", "马树芝", "马妸轩", "马雨东", "马凤帆", "马颖亮", "马上枫", "马宏杰", "马俊伟", "马蕾", "马煜杰", "马秦莹", "马剑轩", "马知亮", "马宇婷", "马煜元", "马婧莲", "马冲卫", "马之", "马晚鹏", "马鑫希", "马静岚", "马泊标", "马中玲", "马鹏艳", "马恩华", "马靖文", "马芳", "马岽楷", "马龙慧", "马心紫", "马思佳", "马丹杰", "马佳冰", "马玳娜", "马了", "马家涵", "马应洪", "马小欲", "马秀治", "马增琪", "马渔", "马嘉实", "马嗣华", "马思颖", "马永婷", "马兴心", "马磬", "马媛", "马修辰", "马墨杰", "马志雨", "马洪思", "马涪辉", "马恩玲", "马鸿泰", "马梦玟", "马漩", "马晓娟", "马淄", "马向煌", "马木文", "马明杰", "房景军", "房紫君", "房善晶", "房真", "房昊松", "房微", "房丁佳", "房浩荃", "房紫童", "房杰舜", "房良翔", "房少芮", "房建涵", "房仁鸣", "房傧", "房晓隽", "房彤媛", "房祁", "房福畅", "房亮", "房翊鸽", "房书宇", "房淑泽", "房俊", "房泰六", "房思榕", "房秀媚", "房佳然", "房誉霏", "房伟海", "房具军", "房河涵", "房小文", "房竹芳", "房雨聍", "曾怀荷", "曾彧明", "曾承彤", "曾睿淮", "曾冠义", "曾敏震", "曾宇涛", "曾兰坚", "曾筵明", "曾烁", "曾寰棉", "曾缜华", "曾世旎", "曾映林", "曾敬绩", "曾凝", "曾卓涵", "曾畅", "曾俊锋", "曾腊婷", "曾尚", "曾湛洋", "曾梅颖", "曾雅亮", "曾鸾凯", "曾雨梅", "曾城", "曾香波", "曾鸣佳", "曾香文", "曾悦薇", "曾文来", "曾智桥", "曾珍朵", "曾霆英", "曾天涵", "曾饴芹", "曾成红", "曾梓轩", "曾佳宇", "曾芙", "曾明帆", "曾子惠", "曾文芸", "曾青澄", "曾胜砚", "曾嘉旭", "曾超", "曾一迪", "曾莉", "曾梓渊", "曾昱博", "曾广", "曾锦鸣", "曾庆坤", "曾茸", "曾殷琨", "曾玉峰", "曾亮", "曾昶邑", "曾城涵", "曾梓含", "曾明莉", "曾秀一", "曾梦如", "曾彬晨", "曾云淼", "曾洪心", "曾哲钰", "曾鹏轩", "曾佳涛", "曾梦然", "曾冠樱", "曾驿轩", "曾雅媚", "曾瑜", "曾婉沁", "曾红源", "曾佳泽", "曾文馨", "曾嘉琳", "曾兆蕾", "曾晓健", "曾波", "曾睿燕", "曾原", "曾泺", "曾明浩", "曾彤", "曾钥", "曾雅岚", "曾凯勋", "曾乔", "韩爽", "韩瑞宇", "韩缘", "韩勐爱", "韩昌彬", "韩文鑫", "韩文尧", "韩雨舒", "韩祯智", "韩焯杰", "韩石泉", "韩则凡", "韩媛硕", "韩君升", "韩铭莹", "韩庄伟", "韩雨华", "韩婕", "韩尚珏", "韩钦", "韩彦荣", "韩小冠", "韩鑫齐", "韩惠聪", "韩于", "韩毓尔", "韩建平", "韩彬锛", "韩雨", "韩展恒", "韩烨", "韩晓明", "韩琼", "韩龙", "韩潼", "韩栩轩", "韩千宾", "韩怡龙", "韩七", "韩华弘", "韩童", "韩瑞", "韩丽豪", "韩旗", "韩晨", "韩维", "韩林谆", "韩宝宁", "韩新娇", "韩艺青", "韩梅", "韩丽丛", "韩昭亮", "韩桃恒", "韩永子", "韩妤", "韩学仁", "韩军霞", "韩静", "韩进锐", "韩嘉平", "韩秀清", "韩得林", "韩薇鸣", "韩梅渊", "韩庆怡", "韩佳升", "韩宏明", "韩鸣迪", "韩馨伦", "韩子瑶", "韩贵睿", "韩琼晨", "韩大民", "韩丽真", "韩希英", "韩彦武", "韩顺靖", "韩正", "韩羽辉", "韩呈恬", "韩沂涵", "韩立童", "韩闰乐", "韩敬淋", "韩妙绚", "韩叶耕", "韩瑾桓", "韩逸", "韩家宁", "韩凤仿", "韩瀚涛", "韩焱", "郑淇燮", "郑恩骏", "郑美翔", "郑泽江", "郑真谣", "郑萱", "郑傲云", "郑小涵", "郑跃", "郑红珊", "郑延明", "郑心云", "郑毛羲", "郑阔溪", "郑世涵", "郑茹芙", "郑龙怡", "郑春蕊", "郑思慧", "郑瑞华", "郑咏静", "郑锡达", "郑若入", "郑文宇", "郑虹婷", "郑作淇", "郑奕宾", "郑菁", "郑麒永", "郑永晗", "郑圣轩", "郑痍熹", "郑俊英", "郑媛瑾", "郑模彤", "郑法右", "郑建骋", "郑涵雅", "郑朝萍", "郑正", "郑娟苒", "郑斐龙", "郑城恩", "郑思予", "郑群弈", "郑鹏虎", "郑宜谦", "郑木智", "郑大修", "郑增澳", "郑菲", "郑锦萍", "郑思婷", "郑宝霞", "郑妤", "郑奕东", "郑绍明", "郑德", "郑滔菱", "郑正涵", "郑佳子", "郑智静", "郑怡适", "郑帅蕙", "郑迅", "郑小昊", "郑语师", "郑东瑶", "郑榕", "郑伟涵", "郑得宇", "郑东璇", "郑惠骏", "郑长璁", "郑柔", "郑晟娜", "郑艳豪", "郑彭", "郑三渝", "郑济平", "郑少几", "郑科生", "郑路昕", "郑赢", "郑羽茗", "郑沛硕", "郑恺", "郑艺池", "郑卉", "郑忠兰", "郑易晗", "郑玉馨", "郑占玲", "姚铖雨", "姚学英", "姚森", "姚楠蘅", "姚玉婷", "姚霭", "姚莉滨", "姚景诗", "姚福红", "姚柯", "姚小淼", "姚意瑞", "姚蓓平", "姚凯闻", "姚剑彤", "姚钺", "姚学燕", "姚应华", "姚国畅", "姚安烁", "姚晨臻", "姚荣富", "姚世文", "姚泽玉", "姚茂驿", "姚可瑜", "姚志捷", "姚盛", "姚奕哲", "姚洪轩", "姚婧", "姚曦伊", "姚全伟", "姚韵军", "姚咏哲", "姚蒙佟", "姚楚芳", "姚超", "姚元", "姚博众", "姚一琳", "姚青林", "姚跃甲", "姚乐", "姚园宇", "姚耀修", "姚银源", "姚铭元", "姚慧彦", "姚晋恬", "姚稚辉", "姚赫秦", "姚志哲", "姚明蓉", "姚睿龙", "姚月灏", "姚火", "姚鑫哲", "姚耀鹏", "姚银琪", "姚宛吉", "姚熙", "姚蔓洁", "姚思鹏", "姚松雯", "姚建雄", "姚沅祥", "姚宽", "姚庆芸", "姚云卿", "姚璇", "姚静晗", "姚楚茜", "姚柯诚", "姚锦玉", "姚流", "姚拓", "姚基宝", "姚薇", "姚子莹", "姚红萌", "姚澜梅", "姚惠平", "姚馨", "姚晟伟", "姚占耀", "姚婕", "姚歆雄", "姚佳翩", "姚明玲", "姚启晓", "姚柯辉", "姚含", "唐红怡", "唐忠淼", "唐春扬", "唐熠滨", "唐乔东", "唐弈威", "唐丹平", "唐明华", "唐铎婉", "唐芷民", "唐伍骏", "唐睿琳", "唐泠怡", "唐素磊", "唐语恒", "唐书辰", "唐杰杰", "唐亭涛", "唐溥林", "唐瀚飞", "唐玉然", "唐子松", "唐丽璐", "唐梦玲", "唐静", "唐贤雪", "唐一鑫", "唐文", "唐锋", "唐龙芳", "唐宛虹", "唐彤萍", "唐彦梅", "唐继梅", "唐上", "唐小玲", "唐擎伟", "唐伊蔓", "唐疏春", "唐桐", "唐晶平", "唐乃丽", "唐苗", "唐语曦", "唐奕发", "唐元星", "唐玉蒙", "唐思莹", "唐香丞", "唐祥泳", "唐舍文", "唐明辉", "唐琦乐", "唐秀品", "唐逸桔", "唐莫", "唐锦栓", "唐朔涵", "唐慕荣", "唐钻予", "唐庆伟", "唐延浩", "唐睿妍", "唐欣芝", "唐卫文", "唐子琪", "唐遥", "唐邵岩", "唐睿屏", "唐世华", "唐景娇", "唐骏根", "唐倩侠", "唐宗文", "唐乙薇", "唐静林", "唐星", "唐馨文", "唐嘉莹", "唐雯缓", "唐艳衡", "唐晓微", "唐婧杰", "唐涵新", "唐小波", "唐大榕", "唐笑福", "唐瑶", "唐小林", "唐金琴", "唐淳斌", "唐栩浩", "唐哲", "姜毅航", "姜名漪", "姜泽翔", "姜露瑜", "姜逸雯", "姜然彤", "姜泽凡", "姜兰", "姜为生", "姜博柳", "姜豪静", "姜煜", "姜胤智", "姜艺君", "姜建吉", "姜翔玲", "姜祉财", "姜筱辉", "姜军欣", "姜享祖", "姜翰生", "姜军雄", "姜纾明", "姜炫伟", "姜艺轩", "姜儒楠", "姜学涟", "姜励", "姜煊", "姜元铭", "姜康花", "姜德贤", "姜子鑫", "姜土洋", "姜冠腾", "姜建懿", "姜泽明", "姜宇基", "姜姿元", "姜保嘉", "姜强锋", "姜嘉珍", "姜奕", "姜晓源", "姜渝", "姜缦萌", "姜芳", "姜当甜", "姜人林", "姜垠琪", "姜佳茜", "姜秋菲", "姜紫淞", "姜惠翔", "姜建雅", "姜晓宁", "姜曾薇", "姜荣英", "姜鹂", "姜友岑", "姜竞岩", "姜丽坡", "姜芹", "姜京雨", "姜照", "姜国钦", "姜婀祁", "姜竞丽", "姜与煌", "姜烜", "姜希学", "姜七杰", "姜新柳", "姜尚聪", "姜荧平", "姜怡栩", "姜晨霖", "姜秭安", "姜淼", "姜宝睿", "姜宸源", "姜毛婷", "姜鸿晨", "姜砚", "姜璨芬", "姜政芬", "姜释媛", "姜嘉睫", "姜来立", "姜敦桦", "姜子健", "姜江", "姜悛洁", "辛庄敏", "辛艾兰", "辛瑞峰", "辛巍汶", "辛纪森", "辛振珂", "辛峰", "辛乐乜", "辛清童", "辛飞康", "辛家荣", "辛睿成", "辛思英", "辛和欣", "辛程焕", "辛骅文", "辛君琪", "辛凌", "辛昭妍", "辛旭文", "辛睿", "辛致婷", "辛思峰", "辛垄宇", "辛绘琛", "辛禾腾", "辛芷阳", "辛俊", "辛清予", "辛晔", "辛珏", "辛文华", "辛加鸢", "辛程", "辛嘉博", "辛嘉铭", "辛忠惠", "辛彦诺", "辛冬", "齐业栋", "齐晓涛", "齐显换", "齐小桦", "齐一佳", "齐曾心", "齐聿旺", "齐惠君", "齐麓桐", "齐棚", "齐远", "齐桂超", "齐胺晨", "齐晋硕", "齐雨轩", "齐翰平", "齐梓臻", "齐笑娴", "齐涵丰", "齐骢潍", "齐海风", "齐泽城", "齐雨军", "齐子平", "齐俊晔", "齐锦远", "齐靖心", "齐明萱", "齐亮萱", "齐汉轩", "齐凯涵", "齐法墨", "齐瑾奕", "齐淑铭", "齐晓源", "齐钰泫", "齐律如", "齐佑田", "齐千", "齐海燕", "齐子涵", "齐海阳", "齐曜茜", "齐艺君", "齐佳", "齐康聆", "齐诗珂", "齐朝东", "齐可洋", "齐泽呈", "齐昱", "齐雨鑫", "齐炳枫", "齐笛", "齐睫", "齐昌昙", "齐秀", "齐甘骊", "齐一毅", "齐思炜", "齐胤昊", "齐坤琪", "齐艾霖", "齐子铧", "齐千勤", "齐理洋", "齐帆", "齐婕", "齐紫诚", "齐洵宜", "齐丹荧", "齐阿熙", "齐若涓", "齐家雯", "齐晖", "齐正展", "齐雨翔", "", "夏宇朵", "夏媛建", "夏宇戈", "夏丰婷", "夏骏宸", "夏子祥", "夏文芹", "夏子胜", "夏玺帅", "夏明潭", "夏新甫", "夏千华", "夏慧震", "夏瑞茹", "夏星升", "夏若霞", "夏玉城", "夏健玲", "夏玉芸", "夏海夫", "夏鸿阳", "夏振婷", "夏诚涵", "夏思林", "夏水轩", "夏默", "夏嘉运", "夏聪烨", "夏雅庭", "夏澄林", "夏冬良", "夏巴菱", "夏叶林", "夏树亚", "夏海鹰", "夏慧霖", "夏太金", "夏雨晨", "夏橹涵", "夏可辛", "夏裕谨", "夏子慧", "夏勋怡", "夏泓馨", "夏惠茹", "夏明谊", "夏子芸", "夏江识", "夏葱阳", "夏薇", "夏一格", "夏云悦", "夏云玉", "夏新荣", "夏悠", "夏永木", "夏谈芳", "夏兰岳", "夏凡", "夏海昶", "夏乐霖", "夏田翔", "夏双方", "夏宇妮", "夏李彬", "夏予清", "夏建晨", "夏鸿林", "夏东", "夏相炅", "夏鲜", "夏寒", "夏瑗非", "夏与", "夏至", "夏文铭", "夏梓乾", "夏元素", "夏靖霖", "夏烨", "夏子平", "夏佳琳", "夏沛红", "夏果", "夏芷莉", "夏万蓓", "夏春豫", "夏煜滢", "夏钰舒", "夏黎月", "夏永强", "夏明润", "夏丞宇"}

	users := make([]interface{}, 0)

	for i, name := range names {
		nn := fmt.Sprintf("%s", name)
		py := tools.Pin.Py(nn)
		pinyin := tools.Pin.Pinyin(nn)
		birth := fmt.Sprintf("19%d-%02d-%02d", 70+rand.Intn(30), 1+rand.Intn(12), 1+rand.Intn(28))
		sex := "1"
		if rand.Intn(10)%2 == 0 {
			sex = "2"
		}

		users = append(users, beans.User{
			Bean: beans.Bean{
				Id:     tools.PtrOfUUID(),
				Status: tools.PtrOfString("1"),
			},
			Dept:   tools.PtrOfString(deptIds[rand.Intn(len(deptIds))]),
			Type:   tools.PtrOfString("1"),
			Name:   &nn,
			Py:     &py,
			Pinyin: &pinyin,
			Avatar: nil,
			IdCard: tools.PtrOfString(fmt.Sprintf("330411%s%04d", strings.ReplaceAll(birth, "-", ""), i)),
			Birth:  tools.PtrOfInt64(tools.Time.MustParseYMD(birth).Unix()),
			Sex:    &sex,
			Sn:     tools.PtrOfString(fmt.Sprintf("X%04d", i)),
			Mobile: tools.PtrOfString(fmt.Sprintf("15%d%d%d%d%d%d%d%d%d", rand.Intn(10), rand.Intn(10), rand.Intn(10), rand.Intn(10), rand.Intn(10), rand.Intn(10), rand.Intn(10), rand.Intn(10), rand.Intn(10))),
		})
	}

	db.Insert(users...)
}

func InitTestRole(db *xorm.Engine) {
	roles := []map[string]interface{}{
		{"id": "1", "name": "管理员", "desc": "", "status": "1"},
		{"id": "3", "name": "经理", "desc": "", "status": "1"},
		{"id": "5", "name": "董事", "desc": "", "status": "1"},
		{"id": "6", "name": "巡查人员", "desc": "", "status": "1"},
		{"id": "7", "name": "测试组", "desc": "", "status": "1"},
		{"id": "8", "name": "研发组", "desc": "", "status": "1"},
		{"id": "9", "name": "外派员", "desc": "", "status": "1"},
	}

	for _, r := range roles {
		db.Table(beans.Role{}).Insert(r)
	}
}
