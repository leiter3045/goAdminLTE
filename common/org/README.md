## beego page.go
封装的一个轻量级的分页类

## 特点
```
1.只需要在调用时填入几个参数，即可自己分页。生成跳转的<a>标签，前端只需要获取代码渲染即可

```
## 使用实例
```
model := org.Page{Request: c.Ctx.Request}
model.NewPage().SetParams(totalRow, c.pageRow)
model.SetConfig("prev", "上一页")
model.SetConfig("next", "下一页")
model.SetConfig("theme", " %HEADER%<li class='disabled'><a>%NOW_PAGE%/%TOTAL_PAGE% 页</a></li>%UP_PAGE%%FIRST%%LINK_PAGE%%DOWN_PAGE%")
str := model.Show()
c.Data["page_count"] = totalRow
c.Data["pages"] = str
```