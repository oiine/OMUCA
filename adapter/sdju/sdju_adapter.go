package sdju

import (
	"github.com/oiine/OMUCA/adapter"
	"strconv"
	"strings"
)

type Adapter struct {
	adapter.OMUCAAdapter
	LocationName string `json:"cdmc"` // 上课地点名称
	Name string `json:"kcmc"` // 课程名称
	Day string `json:"xqjmc"` // 所在周几
	Sections string `json:"jcs"` // 开始节数 e.g 1-4
	TeacherName string `json:"xm"` // 任课教师姓名
	Weeks string `json:"zcd"`// 周数 18-19周
	SchoolArea string `json:"xqmc"` // 校区
}

// 适配课程名
func (a *Adapter)AdaptCourseName()string  {
	return a.Name
}

// 适配地点名
func (a *Adapter)AdaptCourseLocationName()string{
	return a.SchoolArea + a.LocationName
}

// 适配教师名
func (a *Adapter)AdaptCourseTeacherName()string{
	return a.TeacherName
}

// 适配开始节数
func (a *Adapter) AdaptStartSection()int {
	sectionsArr := strings.Split(a.Sections, "-")
	startSection, err := strconv.Atoi(sectionsArr[0])
	if err != nil {
		return 0
	}
	return startSection
}

// 适配结束节数
func (a *Adapter)AdaptEndSection()int{
	sectionsArr := strings.Split(a.Sections, "-")
	endSection, err := strconv.Atoi(sectionsArr[1])
	if err != nil{
		return 0
	}
	return endSection
}

// 适配持续的周数
func (a *Adapter)AdaptWeeks()string {
	// 设置上课的周数 1-8周 -> 1,2,3,4,5,6,7,8 或 2-8(双) -> 2,4,6,8
	var weeksArrWithStartToEnd []string
	var trimWeeks string
	// 先把周字去掉
	trimWeeks = strings.Trim(a.Weeks, "周")
	if strings.Contains(trimWeeks, "(双)") {
		// 双周处理
		trimWeeks = strings.Trim(trimWeeks, "(双)")
	}else if strings.Contains(trimWeeks, "(单)"){
		// 单周处理
		trimWeeks = strings.Trim(trimWeeks, "(单)")
	}else{
		return ""
	}
	// 分割字符 1-8 -> [1,8]
	weeksArr := strings.Split(trimWeeks, "-")
	if len(weeksArr) != 1{
		var (
			startWeek int
			endWeek int
		)
		startWeek, err := strconv.Atoi(weeksArr[0])
		if err != nil {
			return ""
		}
		endWeek, err = strconv.Atoi(weeksArr[1])
		if err != nil {
			return ""
		}
		// 填充周数
		for startWeek <= endWeek {
			weeksArrWithStartToEnd = append(weeksArrWithStartToEnd, strconv.Itoa(startWeek))
			if strings.Contains(a.Weeks, "(双)") || strings.Contains(a.Weeks, "(单)") {
				// 这里不判断单双周直接 append 是因为目标格式双周情况下就是偶数开头， 单周默认就是奇数开头
				startWeek +=2
			}else{
				// 其余情况全部 +1
				startWeek ++
			}
		}
	}else{
		// 数组长度为1，trimWeeks 可能是 1周 -> 1 -> [1]
		weeksArrWithStartToEnd = append(weeksArrWithStartToEnd, trimWeeks)
	}
	// 以 , 填充
	a.Weeks = strings.Join(weeksArrWithStartToEnd, ",")
	return a.Weeks
}

// 适配周几上课
func (a *Adapter)AdaptDayOfWeek()int {
	return map[string]int{"星期一": 1, "星期二": 2, "星期三": 3, "星期四": 4, "星期五": 5, "星期六": 6, "星期天": 7, "星期日": 7}[a.Day]
}



