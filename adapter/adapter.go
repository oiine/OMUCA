package adapter

// 课程适配器
type OMUCAAdapter interface {
	// 课程名称 e.g 高等数学
	AdaptCourseName()string
	// 教室名称 e.g 信息楼 308
	AdaptCourseLocationName()string
	// 教师名称 e.g 小王
	AdaptCourseTeacherName()string
	// 开始节数 e.g 1 代表着 第一节课
	AdaptStartSection()int
	// 结束节数 e.g 4 代表着 第四节课
	AdaptEndSection()int
	// 持续的周数 e.g 1,2,3,4 代表着 第一、第二、第三、第四周
	AdaptWeeks()string
	// 周几课程 e.g 4 代表着周四的课程
	AdaptDayOfWeek()int
}