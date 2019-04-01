package aa

import (
	"fmt"
	"log"
	"runtime"
	"strings"
)

func (l *Log) caller() string {
	_, file, line, _ := runtime.Caller(2)
	a := strings.Split(file, "/")
	return fmt.Sprintf("%s:%d ", a[len(a) - 1], line)
}

func (l *Log) Printf(msg string, args ...interface{}) {
	log.Printf(l.caller() + msg + "\n", args...)
}

func (l *Log) Println(msg ...interface{}) {
	log.Println(l.caller(), fmt.Sprint(msg...))
}

// 紧急情况，需要立即通知技术人员。
func (l *Log) Emerg(msg string, args ...interface{}) {
	log.Printf("[emerg] " + l.caller() + msg + "\n", args...)
}

// 应该被立即改正的问题，如系统数据库被破坏，ISP连接丢失。
func (l *Log) Alert(msg string, args ...interface{}) {
	if l.priority > LogAlert {
		return
	}
	log.Printf("[alert] " + l.caller() + msg + "\n", args...)

}

// 重要情况，如硬盘错误，备用连接丢失
func (l *Log) Crit(msg string, args ...interface{}) {
	if l.priority > LogCrit {
		return
	}
	log.Printf("[crit] " + l.caller() + msg + "\n", args...)

}

// 错误，不是非常紧急，在一定时间内修复即可。
func (l *Log) Error(msg string, args ...interface{}) {
	if l.priority > LogErr {
		return
	}
	log.Printf("[error] " + l.caller() + msg + "\n", args...)
}

// 警告信息，不是错误，比如系统磁盘使用了85%等。
func (l *Log) Warn(msg string, args ...interface{}) {
	if l.priority > LogWarning {
		return
	}
	log.Printf("[warning] " + l.caller() + msg + "\n", args...)

}

// 不是错误情况，也不需要立即处理。
func (l *Log) Notice(msg string, args ...interface{}) {
	if l.priority > LogNotice {
		return
	}
	log.Printf("[warning] " + l.caller() + msg + "\n", args...)

}

// 情报信息，正常的系统消息，比如骚扰报告，带宽数据等，不需要处理。
func (l *Log) Info(msg string, args ...interface{}) {
	if l.priority > LogInfo {
		return
	}
	log.Printf("[info] " + l.caller() + msg + "\n", args...)
}

// 包含详细的开发情报的信息，通常只在调试一个程序时使用
func (l *Log) Debug(msg string, args ...interface{}) {
	if l.priority > LogDebug {
		return
	}
	log.Printf("[debug] " + l.caller() + msg + "\n", args...)
}
