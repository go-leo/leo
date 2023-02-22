package log

import (
	"log"
	"net/http"
	"os"
)

var DefaultLogger Logger

func init() {
	DefaultLogger = &defaultLogger{
		lvl:         Info(),
		debugLogger: log.New(os.Stdout, Debug().Name(), log.LstdFlags|log.Lshortfile),
		infoLogger:  log.New(os.Stdout, Info().Name(), log.LstdFlags|log.Lshortfile),
		warnLogger:  log.New(os.Stdout, Warn().Name(), log.LstdFlags|log.Lshortfile),
		errorLogger: log.New(os.Stdout, Error().Name(), log.LstdFlags|log.Lshortfile),
		panicLogger: log.New(os.Stdout, Panic().Name(), log.LstdFlags|log.Lshortfile),
		fatalLogger: log.New(os.Stdout, Fatal().Name(), log.LstdFlags|log.Lshortfile),
	}
}

type defaultLogger struct {
	lvl         Level
	debugLogger *log.Logger
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
	panicLogger *log.Logger
	fatalLogger *log.Logger
	field       []Field
}

func (l defaultLogger) SetLevel(lvl Level) {
	l.lvl = lvl
}

func (l defaultLogger) GetLevel() Level {
	return l.lvl
}

func (l defaultLogger) EnableDebug() {
	l.lvl = Debug
}

func (l defaultLogger) IsDebugEnabled() bool { return l.lvl <= Debug }

func (l defaultLogger) Debug(...any) {

}

func (l defaultLogger) Debugf(string, ...any) {}

func (l defaultLogger) DebugF(...Field) {}

func (l defaultLogger) EnableInfo() {}

func (l defaultLogger) IsInfoEnabled() bool { return false }

func (l defaultLogger) Info(...any) {}

func (l defaultLogger) Infof(string, ...any) {}

func (l defaultLogger) InfoF(...Field) {}

func (l defaultLogger) EnableWarn() {}

func (l defaultLogger) IsWarnEnabled() bool { return false }

func (l defaultLogger) Warn(...any) {}

func (l defaultLogger) Warnf(string, ...any) {}

func (l defaultLogger) WarnF(...Field) {}

func (l defaultLogger) EnableError() {}

func (l defaultLogger) IsErrorEnabled() bool { return false }

func (l defaultLogger) Error(...any) {}

func (l defaultLogger) Errorf(string, ...any) {}

func (l defaultLogger) ErrorF(...Field) {}

func (l defaultLogger) EnablePanic() {}

func (l defaultLogger) IsPanicEnabled() bool { return false }

func (l defaultLogger) Panic(...any) {}

func (l defaultLogger) Panicf(string, ...any) {}

func (l defaultLogger) PanicF(...Field) {}

func (l defaultLogger) EnableFatal() {}

func (l defaultLogger) IsFatalEnabled() bool { return false }

func (l defaultLogger) Fatal(...any) {}

func (l defaultLogger) Fatalf(string, ...any) {}

func (l defaultLogger) FatalF(...Field) {}

func (l defaultLogger) SkipCaller(int) Logger { return l }

func (l defaultLogger) With(...Field) Logger { return l }

func (l defaultLogger) Clone() Logger { return l }

func (l defaultLogger) ServeHTTP(http.ResponseWriter, *http.Request) {}
