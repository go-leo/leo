package logx

import (
	kitlog "github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/go-kit/log/syslog"
	"github.com/go-kit/log/term"
	"io"
	"time"
)

type factory func(w io.Writer) kitlog.Logger

type options struct {
	factory      factory
	level        level.Value
	levelAllowed level.Option
	sync         func(w io.Writer) io.Writer
	syslog       func(factory) factory
	color        func(factory) factory
	withs        []func(factory) factory
	prefixes     []func(factory) factory
	suffixes     []func(factory) factory
}

type Option func(o *options)

// Level sets the log level.
// see: level.DebugValue(), level.InfoValue(), level.WarnValue(), level.ErrorValue()
func Level(l level.Value) Option {
	return func(o *options) {
		o.level = l
		o.levelAllowed = level.Allow(l)
	}
}

// JSON sets the log format to JSON.
func JSON() Option {
	return func(o *options) {
		o.factory = kitlog.NewJSONLogger
	}
}

// Logfmt sets the log format to logfmt.
func Logfmt() Option {
	return func(o *options) {
		o.factory = kitlog.NewLogfmtLogger
	}
}

// Sync sets the log writer to sync.
func Sync() Option {
	return func(o *options) {
		o.sync = kitlog.NewSyncWriter
	}
}

// With sets the keyvals to log. keyvals is appended to the existing context.
func With(keyvals ...interface{}) Option {
	return func(o *options) {
		o.withs = append(o.withs, func(f factory) factory {
			return func(writer io.Writer) kitlog.Logger {
				return kitlog.With(f(writer), keyvals...)
			}
		})
	}
}

// WithPrefix sets the keyvals to log, keyvals is prepended to the existing context.
func WithPrefix(keyvals ...interface{}) Option {
	return func(o *options) {
		o.prefixes = append(o.prefixes, func(f factory) factory {
			return func(writer io.Writer) kitlog.Logger {
				return kitlog.WithPrefix(f(writer), keyvals...)
			}
		})
	}
}

// WithSuffix sets the keyvals to log, keyvals is appended to the existing context.
func WithSuffix(keyvals ...interface{}) Option {
	return func(o *options) {
		o.suffixes = append(o.withs, func(f factory) factory {
			return func(writer io.Writer) kitlog.Logger {
				return kitlog.WithSuffix(f(writer), keyvals...)
			}
		})
	}
}

// Caller sets the caller to log.
func Caller(depth int) Option {
	return func(o *options) {
		With("caller", kitlog.Caller(depth+5))(o)
	}
}

// Timestamp sets the timestamp to log.
func Timestamp() Option {
	return func(o *options) {
		With("ts", kitlog.TimestampFormat(time.Now, time.DateTime))(o)
	}
}

// Syslog sets the syslog to log.
func Syslog(w syslog.SyslogWriter) Option {
	return func(o *options) {
		o.syslog = func(w syslog.SyslogWriter) func(factory) factory {
			return func(f factory) factory {
				return func(writer io.Writer) kitlog.Logger {
					return syslog.NewSyslogLogger(w, f)
				}
			}
		}(w)
	}
}

// Color sets the color to log.
func Color(color func(keyvals ...interface{}) term.FgBgColor) Option {
	return func(o *options) {
		o.color = func(color func(keyvals ...interface{}) term.FgBgColor) func(factory) factory {
			return func(f factory) factory {
				return func(w io.Writer) kitlog.Logger {
					return term.NewLogger(w, f, color)
				}
			}
		}(color)
	}
}

func New(w io.Writer, opts ...Option) kitlog.Logger {
	o := options{
		factory:      kitlog.NewLogfmtLogger,
		level:        level.DebugValue(),
		levelAllowed: level.AllowDebug(),
		sync:         func(w io.Writer) io.Writer { return w },
		syslog: func(f factory) factory {
			return func(writer io.Writer) kitlog.Logger {
				return f(writer)
			}
		},
		color: func(f factory) factory {
			return func(writer io.Writer) kitlog.Logger {
				return f(writer)
			}
		},
		withs:    nil,
		prefixes: nil,
		suffixes: nil,
	}
	for _, opt := range opts {
		opt(&o)
	}
	w = o.sync(w)
	create := o.factory
	for _, with := range o.withs {
		create = with(create)
	}
	for _, with := range o.prefixes {
		create = with(create)
	}
	for _, with := range o.suffixes {
		create = with(create)
	}
	create = o.color(create)
	create = o.syslog(create)
	logger := create(w)
	logger = level.NewInjector(logger, o.level)
	logger = level.NewFilter(logger, o.levelAllowed)
	return logger
}
