package common

import (
    "context"
    "fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
    "fyne.io/fyne/v2"
)

// 解析 /api/chat
type LLMMessage struct {
    Role 		string `json:"role"`
    Content 	string `json:"content"`
}
type Chan struct {
    Message 	LLMMessage  `json:"message"`
    Done     	bool   		`json:"done"`
}
// 解析/api/tags
type Model struct {
    Name       string `json:"name"`
    ModifiedAt string `json:"modified_at"`
    Size       int64  `json:"size"`
    Digest     string `json:"digest"`
}
// 模型列表
type ModelListResponse struct {
    Models []Model `json:"models"`
}

// 用户设置
type Settings struct {
    URL         string              // API 地址
    Token       string              // API 密钥
    Model       string              // 当前选中的模型
    ModelList   []string            // 模型列表
    CancelFunc  context.CancelFunc  // 是否停止对话
    DialogID    string              // 对话 ID
    EnableAgent bool                // 是否启用Agent调用系统能力
}

// 配置文件解析
type LLMConfig struct {
    Backend struct {
        Ollama struct {
            URL      string `yaml:"url"`
            Token    string `yaml:"token"`
        } `yaml:"ollama"`
    } `yaml:"backend"`
}

type Widgets struct {
    Window      fyne.Window
	MainSplit 	*container.Split
	ChatDisplay *widget.Label
	ChatScroll 	*container.Scroll
	InputEntry 	*widget.Entry
}

// 构造 Get-WinEvent 查询体
type EventQuery struct {
    LogName   string        `json:"logName"`     // 日志类型 (Application, Security, System, etc.)
    StartTime  int          `json:"startTime"`   // 起始时间 (正数，代表往前推多少天)
    MaxEvents  int          `json:"maxEvents"`   // 最大事件数
}

type FileTreeQuery struct {
    Disk        []string    `json:"disk"`
}

type SysHealthQuery struct {
    Minutes     int         `json:"minutes"`
}