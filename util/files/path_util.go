package files

import (
	"github.com/huwhy/commons/config"
	zaprotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap/zapcore"
	"os"
	"path"
	"strings"
	"time"
)

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func GetWriteSyncer(conf *config.Zap) (zapcore.WriteSyncer, error) {
	fileWriter, err := zaprotatelogs.New(
		path.Join(conf.Dir, "%Y-%m-%d.log"),
		//zaprotatelogs.WithLinkName("link_name"),
		zaprotatelogs.WithMaxAge(7*24*time.Hour),
		zaprotatelogs.WithRotationTime(24*time.Hour),
	)
	if conf.ShowConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
	}
	return zapcore.AddSync(fileWriter), err
}

func GetSuffix(filename string) string {
	return filename[strings.LastIndex(filename, "."):]
}

func MoveFile(src, target string) {
	os.Rename(src, target)
}
