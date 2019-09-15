# z-table

## 依赖项 
1. xlsx
2. script-loader
3. file-saver
4. store.data.common.units
5. store.data.common.dicts
6. 新建工单(tkey=$ticket) - 依赖 map-position - 待完善
7. 新建工单(tkey=$ticket) - 依赖 user-select - 待完善
8. 新建工单(tkey=$ticket) - 依赖 $apolloProvider.clients.iot - 待完善

## 1. 属性

|名称|说明|类型|默认值|
|:-|:-|:-|:-|
|columns|支持iview原生属性所有功能，扩展功能见下表| Array|必填|
|isPage|是否需要分页|Boolean| true |
|total|数据总条数|Integer|isPage为true时有效|
|data|iview组件原生属性|Array|必填|
|loadData| 点击查询/修改页码等需要重新渲染数据的时候会调用，传入值为查询参数,fn为导出Excel所有数据的时候，需要isPage==true的时候fn({data:[],total:0}) isPage==false的时候 fn([]),并且数据不渲染到data中，如图![图片](./imgs/table_loaddata.png)  | Function(query,fn) |必填|
|filters| 条件过滤| Array | 为空的时候，不显示过滤框 |

###  columns 扩展

#### 1. tkey 

* 写法1：o_area.name 可以获取到对象中的对象的属性，不管有多少层，没有获取到，返回“ - ”
* 写法2：[]sensors.unit == 'WD';value 可以过滤到“温度传感器”（也可以写为“o_unit.name”），并且获取到value值，“[]...{Array}.{field} == '{值}';{要取的对象值}'”

```json
{
    "id":"5d4569effe00210001e9908e",
    "area":"5d3bf277383d9c0001232180",
    "type":"well",
    "time":1566390306,
    "o_area":{
        "name":"台州"
    },
    "sensors":[{
        "id":"5d4569effe00210001e9908f",
        "name":"氨氮",
        "time":1566390306,
        "unit":"NH",
        "o_unit":{
            "name":"氨氮"
        },
        "value":0
    },{
        "id":"5d4569effe00210001e99090",
        "name":"温度",
        "time":1566390306,
        "unit":"WD",
        "value":0,
        "o_unit":{
            "name":"氨氮"
        },
    }]
}
```

* 写法3：$ticket 新建工单(完成未测试)，组件内部设定了一些通用组件，可以直接拿来用

| 类型 | 说明 | 备注 |
|:- |:- |:- |
| ticketType | 工单类型的id或者名称 | *必填 |
| ticketOid | 工单接口oid的值 | *必填 |
| ticketOType | 工单接口的otype | *必填 |

```json
[
    { "title": "建单", "tkey":"$ticket", "ticketType":"工单类型","ticketOid":"o_station.id","ticketOType":"station"},
]
```

#### 2. tformat

输入类型
1. String 内容见如下表格说明

| 类型 | 说明 |
|:- |:- |
| unit | 用户格式化传感器的单位值单位等，前提是store的common中必须有unit数组 |
| time | 格式化时间戳，time:YMDHms [Y:年,M:月,D:日,H:小时,m:分钟,s:秒]，顺序不可以颠倒，格式化后的时间格式固定为“2019-09-05 09:38:26” |
| dict | 通过字典数据格式化值，例如“dict:S0001” |
| images | 将"a,b,c"格式化为["a","b","c"] |

```json
[
    { "title": "区域1", "width": "80", "tkey":"o_station.o_area.name","tformat":"unit:F1"},
    { "title": "区域2", "width": "80", "tkey":"o_station.o_area.name1","tformat":"dict:D0001"},
    { "title": "区域3", "width": "80", "tkey":"o_station.o_area.name2","tformat":"time:YMD"},
    { "title": "区域3", "width": "80", "tkey":"o_station.o_area.name2","tformat":"images"},
]
```

2. Function(row,value)


#### 3. tstyle
输入类型

| 对象类型 | 说明 |
|:- |:- |
| String(组件Tag)| 组件名，并且配值，例如：“Tag:success” /^Tag(:[0-9a-zA-Z]+)?$/ |
| String(html元素a)| 组件名，并且配值，例如：“a:success” /^a(:[0-9a-zA-Z]+)?$/ |
| String(html元素img)| 未实现 /^a(:[0-9a-zA-Z]+)?$/ |
| String(颜色值) | 列颜色值，可以为颜色名称，例如：“red”，可以为颜色值，例如：“#000000” /^[a-z]+$|^#[0-9a-fA-F]{3}$|^#[0-9a-fA-F]{6}$/ |
| Object | 不同的值（可以为显示值，也可以为实际值）对应不同的颜色，例如 {"1,故障,关":"red","2,开":"#00aa00"} |
| Function | 未设计 |
| Array | 未设计 |

tstyle的优先级为数组的顺序
```json 
[
    
    { "title": "区域", "width": "80", "tkey":"o_station.o_area.name","tstyle":"Tag:success"},
    { "title": "区域", "width": "80", "tkey":"o_station.o_area.name","tstyle":"a:https%3a%2f%2fwww.baidu.com"},
    { "title": "区域", "width": "80", "tkey":"o_station.o_area.name","tstyle":"img:80*80"},
    { "title": "区域", "width": "80", "tkey":"o_station.o_area.name","tstyle":"red"},
    { "title": "区域", "width": "80", "tkey":"o_station.o_area.name","tstyle":"#000"},
    { "title": "区域", "width": "80", "tkey":"o_station.o_area.name","tstyle":"#ffffff"},
    { "title": "区域", "width": "80", "tkey":"o_station.o_area.name","tstyle":{
        "1,故障,关":"red",
        "2,开":"#00aa00"
    }},
    { "title": "区域", "width": "80", "tkey":"o_station.o_area.name","tstyle":{
        "Tag:1,故障,关":"success",
        "Tag:2,开":"error"
    }},
]
```

#### 4. tfilters
输入类型 String，简写方式需要配置tformat，并且tformat的数据类型为String 

| 对象类型 | 说明 |
|:- |:- |
| dict | 通过字典数据过滤，例如 "tfilters":"dict:S0001;type"  其中type表示过滤的字段名称，简写方式 "tfilters":"type",若"tfilters":"[]type",则表示返回的过滤项为数组 |
| unit | 通过字典数据过滤，例如"tfilters":"unit:F3;unit" 其中第二个unit表示过滤的字段名称，简写方式 "tfilters":"unit",若"tfilters":"[]unit",则表示返回的过滤项为数组 |

#### 5. tsort
输入类型 String
```json 
[
    { "title": "区域", "width": "80", "tkey":"o_station.o_area.name","tsort":"order:id"},
    { "title": "区域", "width": "80", "tkey":"o_station.o_area.name","tsort":"[]order:id"},
]
```


#### 6. texcelName

整个columns中，只要配置了一个excelName并且 excelName!= undefined ,就会在Excel中多一行，显示字段名称，即texcelName，没有的取到texcelName则取key

#### 7. tclick 
输入类型 Function

点击整个单元格出发响应事件


### filters 条件过滤

#### 1. type 字段

| 类型 | 说明 |
|:- |:- |
| string | 手动输入字符串，主要用户模糊查询 |
| number | 手动输入数字，主要用户模糊查询 |
| dict | 选项列表dict:S0001 |
| select | 选项列表,list字段不能为空 |
| time | 支持的类型有["date","daterange","datetime","datetimerange","year","month"]，可以time:date方式来选择 |
| area | 未实现 |
| user | 未实现 |
| station | 未实现 |

数据格式：

```json
[
    {"title":"用户名称","type":"input:text","field":"name","placeholder":"请输入用户名称模糊查询"},
    {"title":"年龄","type":"input:number","field":"age"},
    {"title":"站点类型","type":"dict:S0002","field":"keyword"},
    {"title":"用户类型","type":"select","field":"user_type","list":[
        {"label":"微信用户","value":"1"},
        {"label":"支付宝用户","value":"2"},
    ]},
    { "type": "time", "title": "排水时间", "field": "start" },
    { "type": "time:daterange", "title": "排水时间", "field": "start,end" },
]
```

#### 2. field字段
在loadData的时候，返回的字段，有2种形式
1. "field":"name" 返回的值为对象
2. "field":"[]name" 返回的值为数组

#### 3. list 字段
数据格式

```json
{
    "list":[
        {"label":"微信用户","value":"1"},
        {"label":"支付宝用户","value":"2"},
    ]
}
```

#### 4. placeholder 字段(可空，有默认值)



## 2. 事件

| 事件名称 | 说明 | 备注 |
|:-|:-|:- |
| delete-batch | 批量删除确认的时候回调 | 勾选的数组 |

## 3. slot

| 名称 | 说明 |
|:-|:-|
| filter-left | 搜索框左边 |
| filter-right | 搜索框右边 |
| btn-left | 工具栏左边 |
| btn-center | 工具栏中间 |
| btn-right | 工具栏右边 |
| btn-group-left | 按钮组左边 |
| btn-group-right | 按钮组右边 |
| footer | 表格底部 |
