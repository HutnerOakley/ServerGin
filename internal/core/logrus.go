package core

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"sync"

	"github.com/sirupsen/logrus"
)

type MyLog struct {
	NoColor bool //控制输出颜色(写文件不带颜色)
}

// 颜色常量
const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

/*
日志配置格式化
*/
func (ml *MyLog) Format(entry *logrus.Entry) ([]byte, error) {
	// 根据不同的level展示不同的颜色
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = gray
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red
	default:
		levelColor = blue
	}
	var buffer *bytes.Buffer
	if entry.Buffer != nil {
		buffer = entry.Buffer
	} else {
		buffer = &bytes.Buffer{}
	}

	// 自定义日期格式
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	// 防止Caller为nil
	fileVal := ""
	funcVal := ""
	if entry.Caller != nil {
		// 自定义文件路径
		fileVal = fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
		funcVal = entry.Caller.Function
	}

	// 根据NoColor判断格式
	if ml.NoColor {
		// 写文件格式（极度简洁，无乱码）
		fmt.Fprintf(buffer, "[%s] [%s] %s %s %s\n", timestamp, entry.Level, fileVal, funcVal, entry.Message)
	} else {
		// 终端输出格式（带颜色）
		fmt.Fprintf(buffer, "[%s] \x1b[%dm[%s]\x1b[0m %s %s\n", timestamp, levelColor, entry.Level, fileVal, entry.Message)
	}
	return buffer.Bytes(), nil
}

/*
日志初始化
*/
func InitLogger() {
	// 设置日志级别
	logrus.SetLevel(logrus.InfoLevel)
	// 设置日志行号
	logrus.SetReportCaller(true)
	// 设置日志格式
	logrus.SetFormatter(&MyLog{NoColor: false})
	logrus.AddHook(&MyHook{
		LogPath:       "logs",
		fileFormatter: &MyLog{NoColor: true},
	})
}

/*
日志钩子
按天分片，错误日志单独存放
*/
type MyHook struct {
	file          *os.File   // 当前打开的日志文件
	errFile       *os.File   // 错误日志的文件
	date          string     // 当前日志的时间
	LogPath       string     // 日志的目录
	fileFormatter *MyLog     // 文件格式化器
	mu            sync.Mutex // 锁
}

/*
写到文件中的钩子
*/
func (hook *MyHook) Fire(entry *logrus.Entry) error {
	// 1. 写入到文件
	// 2.按时间分片
	// 3.把错误的日志单独存放
	hook.mu.Lock()
	defer hook.mu.Unlock()

	// 检查日期，触发轮转
	date := entry.Time.Format("2006-01-02")
	if hook.date != date {
		// 换时间
		err := hook.rotateFiles(date)
		if err != nil {
			return err
		}
		hook.date = date
	}

	// 复用Format函数，不手写Sprintf
	msgBytes, err := hook.fileFormatter.Format(entry)
	if err != nil {
		return err
	}
	// 写入普通日志
	if hook.file != nil {
		if _, err := hook.file.Write(msgBytes); err != nil {
			return fmt.Errorf("写入日志文件失败: %v", err)
		}
	}

	// 写入错误日志
	if entry.Level <= logrus.ErrorLevel && hook.errFile != nil {
		if _, err := hook.errFile.Write(msgBytes); err != nil {
			return fmt.Errorf("写入错误日志文件失败: %v", err)
		}
	}

	return nil
}

/*
切换日志文件，日志轮转
*/
func (hook *MyHook) rotateFiles(timer string) error {
	// 1.关闭文件后，要把指针置为nil
	if hook.file != nil {
		err := hook.file.Close()
		if err != nil {
			return fmt.Errorf("文件关闭失败: %v", err)
		}
		hook.file = nil
	}

	if hook.errFile != nil {
		err := hook.errFile.Close()
		if err != nil {
			return fmt.Errorf("文件关闭失败: %v", err)
		}
		hook.errFile = nil
	}

	// 2.创建新目录
	logDir := filepath.Join(hook.LogPath, timer)
	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		return fmt.Errorf("创建日志目录错误: %v", err)
	}

	// 3.打开日志文件
	logPath := filepath.Join(logDir, "info.log")
	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("打开info日志失败: %v", err)
	}
	hook.file = file

	errLogPath := filepath.Join(logDir, "error.log")
	errFile, err := os.OpenFile(errLogPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("打开err日志失败: %v", err)
	}
	hook.errFile = errFile

	return nil
}

func (hook *MyHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
