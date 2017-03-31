package utils

import (
	"github.com/Sirupsen/logrus"
	"time"
	"fmt"
	"strings"
	"go.uber.org/zap"
)


type customFormmater struct {}



var Logger *logrus.Entry
var ZapLogger *zap.Logger

func (_ customFormmater) Format(e *logrus.Entry) ([]byte, error){
	timeStr := time.Now().Format(time.RFC3339)

	res := fmt.Sprintf("[%s]%s, Msg: %s \n", strings.ToUpper(e.Level.String()), timeStr, e.Message)
	//e.Data


	return []byte(res), nil
}



func init() {

	//logrus.Formatter()

	Logger = logrus.WithFields(logrus.Fields{})
	logrus.SetFormatter(customFormmater{})
	//logrus.S

	//logrus.New().
	//logrus.SetFormatter()
	//logrus.

	logrus.SetLevel(logrus.DebugLevel)

	//l,_ := zap.NewProduction()

	//l.

	//config := zap.NewProductionConfig()
	//config.Level = /
	//config.Level = zap.NewAtomicLevel().SetLevel(zap.DebugLevel)

	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.DebugLevel)
	config := zap.NewProductionConfig()
	config.Level = atomicLevel
	//zap.New

	//zap.Option
	config.DisableCaller = true;
	config.DisableStacktrace = true;
	//config.
	ZapLogger, _ = config.Build()


	sl := ZapLogger.Sugar()

	ZapLogger.Info("ZAP info message")
	ZapLogger.Info("URL params", zap.String("key", "value"))
	sl.Debug("SUGAR debug")
	sl.Info("SUGAR info message")
	sl.Warn("SUGAR worn")
	sl.Error("SUGAR error message")
	sl.Infow("SUGAR Infow", "url", 5, "asd", true)


	//Logger.Debug("Init utils debug")
	//Logger.Info("Init utils info")
	//Logger.Warn("Init utils warn")
	Logger.Error("Init utils error")
}
