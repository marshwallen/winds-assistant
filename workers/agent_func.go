package workers

import (
	"winds-assistant/tools"
)

// ** Register Agent Tools
var ToolsFuncRegister = map[string]func(map[string]interface{}) (string){
	"get_win_event": getWinEvent,
	"get_file_tree": getFileTree,
	"get_sys_health": getSysHealth,
	"get_sys_process": getSysProcess,
	"get_sys_driver": getSysDriver,
	"get_bili_rcmd": getBiliRcmd,
}

// ** Register Agent Tools Prompt
var ToolsPromptRegister = []string{
	GET_WIN_EVENT_PROMPT,
	GET_FILE_TREE_PROMPT,
	GET_SYS_HEALTH_PROMPT,
	GET_SYS_PROCESS_PROMPT,
	GET_SYS_DRIVER_PROMPT,
	GET_BILI_RCMD_PROMPT,
}

// 在这里写 Agent Tools 的函数入口
// return type 必须为 string

// 查询 Windows 事件日志并返回结果
const GET_WIN_EVENT_PROMPT = `
工具 <get_win_event> 使用规则：
1. 如果用户提到了 <分析日志> 等类似的需求，你可以使用 <get_win_event> 工具来获取日志。
2. 此外，你还要解析用户需要分析的日志类型(Application, Security, System), 分析天数(StartTime, 为正数(默认1), 表示往前分析多少天)和最大事件数(MaxEvents, 默认50)。
3. 最后, 只返回如下类似的json内容, 除此之外不要说任何其他内容, 不要有多余的符号如 Markdown 代码块标识符, 无效换行和空白等:
{
	"tools": {
		"get_win_event": {
			"logName": "Application",
			"startTime": 1,
			"maxEvents": 50
		}
	}            
}`
func getWinEvent(q map[string]interface{}) string{
	logName, _ := q["logName"].(string)
    _s, _ := q["startTime"].(float64)
    _m, _ := q["maxEvents"].(float64)
    startTime := int(_s)
    maxEvents := int(_m)

	_o, _ := tools.QueryEvents(logName, startTime, maxEvents)
	output := "<get_win_event> 返回结果：" + _o
	return output
}

const GET_FILE_TREE_PROMPT = `
工具 <get_file_tree> 使用规则：
1. 如果用户提到了 <分析文件> 等类似的需求，你可以使用 <get_file_tree> 工具来获取文件树。
2. 此外，你还要解析用户想分析的盘符(如 C:/)。默认是 C:/。如有多个，请用字符串列表的形式表示。
3. 最后, 只返回如下类似的json内容, 除此之外不要说任何其他内容, 不要有多余的符号如 Markdown 代码块标识符, 无效换行和空白等:
{
	"tools": {
		"get_file_tree": {
			"disk": ["A:/", "B:/"]
		}
	}
}`
// 根据提供的磁盘列表获取文件树结构
func getFileTree(q map[string]interface{}) string{
	diskList, _ := q["disk"].([]interface{})

	output := "<get_file_tree> 返回结果："
	for _, d := range diskList {
		disk, _ := d.(string)
		_o, _ := tools.GetFileTree(disk, 3, 10, false)
		output += _o
	}
	return output
}

const GET_SYS_HEALTH_PROMPT = `
工具 <get_sys_health> 使用规则：
1. 如果用户提到了 <分析系统、硬件监控、CPU、GPU、内存、硬盘> 等类似的需求，你可以使用 <get_sys_health> 工具来获取系统信息。
2. 此外，你还要解析用户想分析的时间范围(Minutes, 为正数, 表示往前分析多少分钟)。默认是1, 表示分析最近一分钟。
3. 用户对每个指标返回一个列表, 分别代表每个指标在每个时间的值。默认每10秒记录一次数据。如果程序中断, 记录情况会断开。
4. 你需要结合多个指标对系统状态进行分析。此外，因为这些是时间序列，你需要额外地进行一些数学上的分析。
5. 最后, 只返回如下类似的json内容, 除此之外不要说任何其他内容, 不要有多余的符号如 Markdown 代码块标识符, 无效换行和空白等:
{
	"tools": {
		"get_sys_health": {                          
			"minutes": 1
		}
	}
}`
// 根据提供的参数获取系统健康数据
func getSysHealth(q map[string]interface{}) string{
	_m, _ := q["minutes"].(float64)
	minutes := int(_m)
	_o, _ := tools.GetSysHealthData(minutes)
	output := "<get_sys_health> 返回结果：" + _o
	return output
}

const GET_SYS_PROCESS_PROMPT = `
工具 <get_sys_process> 使用规则：
1. 如果用户提到了 <分析系统进程、查看运行情况> 等类似的需求，你可以使用 <get_sys_process> 工具来获取系统进程信息。
2. 你需要结合多个指标对系统状态进行分析, 如进程的CPU, 内存 ,执行路径 ,启动参数 ,运行状态 ,线程/句柄数 ,IO 统计等详细信息, 查找潜在问题。
3. 最后, 只返回如下类似的json内容, 除此之外不要说任何其他内容, 不要有多余的符号如 Markdown 代码块标识符, 无效换行和空白等:
{
	"tools": {
		"get_sys_process": {                          
			"enable": true
		}
	}
}`
// 根据传入的参数判断是否启用获取系统进程信息
func getSysProcess(q map[string]interface{}) string{
	enable, _ := q["enable"].(bool)
	if enable {
		// 获取系统进程信息
		_o := tools.GetSysProcessStr()
		output := "<get_sys_process> 返回结果：" + _o
		return output
	}
	return ""
}

const GET_SYS_DRIVER_PROMPT = `
工具 <get_sys_driver> 使用规则：
1. 如果用户提到了 <分析系统驱动状态> 等类似的需求，你可以使用 <get_sys_driver> 工具来获取系统进程信息。
2. 你的目的是排查系统中的驱动问题，如驱动程序信息，包括设备名称、制造商、驱动版本、状态和签名状态等。
3. 最后, 只返回如下类似的json内容, 除此之外不要说任何其他内容, 不要有多余的符号如 Markdown 代码块标识符, 无效换行和空白等:
{
	"tools": {
		"get_sys_driver": {						  
			"enable": true
		}
	}
}`

func getSysDriver(q map[string]interface{}) string{
	enable, _ := q["enable"].(bool)
	if enable {
		// 获取系统进程信息
		_o := tools.GetSysDriverStr()
		output := "<get_sys_driver> 返回结果：" + _o
		return output
	}
	return ""
}

const GET_BILI_RCMD_PROMPT = `
工具 <get_bili_rcmd> 使用规则：
1. 如果用户提到了 <B站视频推荐> 等类似的需求, 你可以使用 <get_bili_rcmd> 工具来获取系统进程信息。
2. 你的目的是根据用户的需求过滤得到的视频列表, 筛选出用户喜爱且高质量的视频给用户 (包含BV号、UP主、标题、视频长度、播放数、弹幕数、点赞数、推荐理由等重要信息)
3. 用户可能会给你一个cookie, 你可以通过在下面的json中加入该cookie, 从而获取用户喜欢的视频。如果没给, 就保持空字符串""。
4. rounds是需要获取几轮推荐。如果用户没有特别指出, 就保持默认的1。
5. 最后, 只返回如下类似的json内容, 除此之外不要说任何其他内容, 不要有多余的符号如 Markdown 代码块标识符, 无效换行和空白等:
{
	"tools": {
		"get_bili_rcmd": {						  
			"cookie": "xxx"
			"rounds": 1,
		}
	}
}`

func getBiliRcmd(q map[string]interface{}) string{
	cookie, _ := q["cookie"].(string)
	rounds, _ := q["rounds"].(float64)
	_o := tools.GetBiliRcmdStr(cookie, int(rounds))
	output := "<get_bili_rcmd> 返回结果：" + _o
	return output
}