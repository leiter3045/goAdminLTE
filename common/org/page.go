package org

import (
	"fmt"
	"math"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type Page struct {
	Request     *http.Request
	FirstRow    int
	ListRows   	int
	TotalRows   int
	TotalPages  int
	RollPage    int
	LastSuffix  bool
	Config		map[string]string
	Page   		int
	Url  		string
	NowPage 	int
}

func (c *Page) NewPage() *Page {
	config := make(map[string]string, 0)
	config["current_tag"] = "a"
	config["current_class"] = "active"
	config["header"] = "<li class='disabled'><a>共 %TOTAL_ROW% 条记录</a></li>"
	config["prev"] = "<<"
	config["next"] = ">>"
	config["first"] = "1..."
	config["last"] = "...%TOTAL_PAGE%"
	config["theme"] = "%FIRST% %UP_PAGE% %LINK_PAGE% %DOWN_PAGE% %END%"
	c.Config = config
	return c
}

func (c *Page) SetParams(totalRows int, listRows int) {
	/* 基础设置 */
	c.TotalRows  = totalRows //设置总记录数
	c.ListRows   = listRows  //设置每页显示行数
	c.LastSuffix = true
	c.RollPage 	 = 10
	c.Page 		 = 1
	c.NowPage    = 1 //empty($_GET[$this->p]) ? 1 : intval($_GET[$this->p]);
	c.FirstRow   = c.ListRows * (c.NowPage - 1)
}

func (c *Page) SetConfig(key string, value string)  {
	c.Config[key] = value
}

/**
 * build
 */
func (c Page) url(page int) string {
	link, _ := url.ParseRequestURI(c.Request.RequestURI)
	values := link.Query()
	if page == 1 {
		values.Del("page")
	} else {
		values.Set("page", strconv.Itoa(page))
	}
	link.RawQuery = values.Encode()
	return link.String()
}

/**
 * 压缩数组
 * @return string
 */
func (c Page) zipArray(a1, a2 []string) []string {
	r := make([]string, 2*len(a1))
	for i, e := range a1 {
		r[i*2] = e
		r[i*2+1] = a2[i]
	}
	return r
}

/**
 * 上一页
 * @return string
 */
func (c Page) hasPrev() (link string) {
	upRow  := c.NowPage - 1
	if upRow > 0 {
		link = "<li><a class='prev' href=" + c.url(upRow) + ">" + c.Config["prev"] + "</a></li>"
	} else {
		link = "<li class='disabled'><a class='prev' href='javascript:void(0);'>" + c.Config["prev"] + "</a></li>"
	}
	return link
}

/**
 * 下一页
 * @return string
 */
func (c Page) hasNext() (link string) {
	downRow  := c.NowPage + 1
	if downRow <= c.TotalPages {
		link = "<li><a class='next' href=" + c.url(downRow) + ">" + c.Config["next"] + "</a></li>"
	} else {
		link = "<li class='disabled'><a class='next' href='javascript:void(0);'>" + c.Config["next"] + "</a></li>"
	}
	return link
}

/**
 * 第一页
 * @return string
 */
func (c Page) pageLinkFirst(nowCoolPage float64) (link string) {
	floatFirst := float64(c.NowPage) - nowCoolPage
	if c.TotalPages > c.RollPage && floatFirst >= 1 {
		link = "<li><a class='first' href=" + c.url(1) + ">" + c.Config["first"] + "</a></li>"
	}
	return link
}

/**
 * 最后一页
 * @return string
 */
func (c Page) pageLinkPrev(nowCoolPage float64) (link string) {
	floatEnd := float64(c.NowPage) + nowCoolPage
	if c.TotalPages > c.RollPage && int(floatEnd) < c.TotalPages {
		link = "<li><a class='end' href=" + c.url(c.TotalPages) + ">" + c.Config["last"] + "</a></li>"
	}
	return link
}

func (c Page) pageLink(nowCoolPage float64, nowCoolPageCeil int) (link string) {
	var pageInt int
	nowpage := int(float64(c.NowPage) + nowCoolPage - 1)
	for i := 1; i <= c.RollPage; i++ {
		if float64(c.NowPage) - nowCoolPage <= 0 {
			pageInt = i;
		} else if  int(nowpage) >= c.TotalPages {
			pageInt = c.TotalPages - c.RollPage + i
		} else {
			pageInt = c.NowPage - nowCoolPageCeil + i
		}
		if pageInt > 0 && pageInt != c.NowPage {
			if pageInt <= c.TotalPages {
				link += "<li><a class='num' href=" + c.url(pageInt) + ">" + strconv.Itoa(pageInt) + "</a></li>"
			}else{
				break;
			}
		}else{
			if pageInt > 0 && c.TotalPages != 1 {
				link += "<li class=" + c.Config["current_class"] + "><" + c.Config["current_tag"] + " class=" + c.Config["current_class"] + ">" + strconv.Itoa(pageInt) + "</" + c.Config["current_tag"] + "></li>"
			}
		}
	}
	return link
}

func (c *Page) Show() string {
	if 0 == c.TotalRows {
		return ""
	};
	c.Page, _ = strconv.Atoi(c.Request.Form.Get("page"))
	link, _ := url.ParseRequestURI(c.Request.RequestURI)
	values := link.Query()
	page := values.Get("page")
	/* 生成URL */
	if c.Request.Form == nil {
		c.Request.ParseForm()
	}
	c.NowPage, _ = strconv.Atoi(page)
	c.Url = c.url(c.NowPage)
	c.TotalPages = int(math.Ceil(float64(c.TotalRows) / float64(c.ListRows)))
	/* 计算分页信息 */
	if c.TotalPages != 0 && c.NowPage > c.TotalPages {
		c.NowPage = c.TotalPages
	} else if c.NowPage == 0 {
		c.NowPage = 1
	}
	/* 计算分页零时变量 */
	nowCoolPage := float64(c.RollPage / 2)
	nowCoolPageCeil := int(math.Ceil(nowCoolPage))
	if c.LastSuffix {
		c.Config["last"] = strconv.Itoa(c.TotalPages)
	}
	fmt.Print(c.NowPage)
	//上一页
	upPage := c.hasPrev()
	//下一页
	downPage := c.hasNext()
	//第一页
	theFirst := c.pageLinkFirst(nowCoolPage)
	//最后一页
	theEnd := c.pageLinkPrev(nowCoolPage)
	//数字连接
	linkPage := c.pageLink(nowCoolPage, nowCoolPageCeil)
	//替换分页内容
	c.Config["header"] = strings.Replace(c.Config["header"], "%TOTAL_ROW%", strconv.Itoa(c.TotalRows), 1)
	pageArr := []string{"%HEADER%", "%NOW_PAGE%", "%UP_PAGE%", "%DOWN_PAGE%", "%FIRST%", "%LINK_PAGE%", "%END%", "%TOTAL_PAGE%"}
	RepArr  := []string{c.Config["header"], strconv.Itoa(c.NowPage), upPage, downPage, theFirst, linkPage, theEnd, strconv.Itoa(c.TotalPages)}
	pageStr := strings.NewReplacer(c.zipArray(pageArr, RepArr)...).Replace(c.Config["theme"])
	return "<ul class='pagination'>" + pageStr + "</ul>"
}



