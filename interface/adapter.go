package _interface

// 课程适配器
type OMUCAAdapter interface {
	// 课程名称 e.g 高等数学
	AdaptCourseName(v interface{})string
	// 教室名称 e.g 信息楼 308
	AdaptCourseLocationName(v interface{})string
	// 教师名称 e.g 小王
	AdaptCourseTeacherName(v interface{})string
	// 开始节数 e.g 1 代表着 第一节课
	AdaptStartSection(v interface{})int
	// 结束节数 e.g 4 代表着 第四节课
	AdaptEndSection(v interface{})int
	// 持续的周数 e.g [1,2,3,4] 代表着 第一、第二、第三、第四周
	AdaptWeeks(v interface{})[]int
	// 周几课程 e.g 4 代表着周四的课程
	AdaptDayOfWeek(v interface{})int
}