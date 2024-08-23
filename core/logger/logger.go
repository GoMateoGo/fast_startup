package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	L *zap.SugaredLogger
)

// --------------------------------------------------------
// -初始化日志组件
func InitLogger() {
	logMode := zapcore.DebugLevel
	//if !viper.GetBool("mod.develop") {
	//	logMode = zapcore.InfoLevel
	//}
	core := zapcore.NewCore(getEncoder(), zapcore.NewMultiWriteSyncer(getWriteSyncer(), zapcore.AddSync(os.Stdout)), logMode)
	z := zap.New(core).Sugar()
	if z == nil {
		panic("日志组件初始化失败!")
	}
	log.Println("=================日志组件加载完毕=================")
	L = z
}

// --------------------------------------------------------
// -设置日志编码规范
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(t.Local().Format(time.DateTime))
	}
	return zapcore.NewJSONEncoder(encoderConfig)
}

// --------------------------------------------------------
// -获取日志写入方式
func getWriteSyncer() zapcore.WriteSyncer {
	sSeparator := string(filepath.Separator) //取分隔符
	sRootDir, _ := os.Getwd()                //取当前文件目录
	// 当前文件目录 + 分隔符 + log文件夹 + 分隔符 + 当前时间 + .txt
	sLogFilePath := sRootDir + sSeparator + "log" + sSeparator + time.Now().Format(time.DateOnly) + ".txt"
	fmt.Println(sLogFilePath)

	lumberjackSyncer := &lumberjack.Logger{
		Filename: sLogFilePath,
		MaxSize:  5, //日志文件最大尺寸(M),超限后自动分隔
		//MaxBackups: 100, //保留旧文件的最大个数
		MaxAge: 90, //保留旧文件的最大天数
	}

	return zapcore.AddSync(lumberjackSyncer)
}
