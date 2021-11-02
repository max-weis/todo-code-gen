package middleware

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/rs/zerolog"
	"io"
	"strconv"
	"time"
)

func AccessLogger(skipper middleware.Skipper) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if skipper(c) {
				return next(c)
			}
			req := c.Request()
			res := c.Response()
			start := time.Now()

			var err error
			if err = next(c); err != nil {
				c.Error(err)
			}
			stop := time.Now()

			id := req.Header.Get(echo.HeaderXRequestID)
			if id == "" {
				id = res.Header().Get(echo.HeaderXRequestID)
			}
			reqSize := req.Header.Get(echo.HeaderContentLength)
			if reqSize == "" {
				reqSize = "0"
			}

			log.Printf("%s %s [%v] %s %-7s %s %3d %s %s %13v %s %s",
				id,
				c.RealIP(),
				stop.Format(time.RFC3339),
				req.Host,
				req.Method,
				req.RequestURI,
				res.Status,
				reqSize,
				strconv.FormatInt(res.Size, 10),
				stop.Sub(start).String(),
				req.Referer(),
				req.UserAgent(),
			)
			return err
		}
	}
}

type ZerologAdapter struct {
	*zerolog.Logger
}

func ProvideZerologAdapter(logger zerolog.Logger) *ZerologAdapter {
	return &ZerologAdapter{Logger: &logger}
}

func toZerologLevel(level log.Lvl) zerolog.Level {
	switch level {
	case log.DEBUG:
		return zerolog.DebugLevel
	case log.INFO:
		return zerolog.InfoLevel
	case log.WARN:
		return zerolog.WarnLevel
	case log.ERROR:
		return zerolog.ErrorLevel
	}

	return zerolog.InfoLevel
}

func toEchoLevel(level zerolog.Level) log.Lvl {
	switch level {
	case zerolog.DebugLevel:
		return log.DEBUG
	case zerolog.InfoLevel:
		return log.INFO
	case zerolog.WarnLevel:
		return log.WARN
	case zerolog.ErrorLevel:
		return log.ERROR
	}

	return log.OFF
}

func (z *ZerologAdapter) Output() io.Writer {
	return z.Logger
}

func (z *ZerologAdapter) SetOutput(w io.Writer) {
	z.SetOutput(w)
}

func (z *ZerologAdapter) Prefix() string {
	return ""
}

func (z *ZerologAdapter) SetPrefix(p string) {
	// do nothing
}

func (z *ZerologAdapter) Level() log.Lvl {
	return toEchoLevel(z.Logger.GetLevel())
}

func (z *ZerologAdapter) SetLevel(v log.Lvl) {
	z.Logger.WithLevel(toZerologLevel(v))
}

func (z *ZerologAdapter) SetHeader(h string) {
	// do nothing
}

func (z *ZerologAdapter) Debug(i ...interface{}) {
	z.Logger.Debug().Msg(fmt.Sprint(i...))
}

func (z *ZerologAdapter) Debugf(format string, args ...interface{}) {
	z.Logger.Debug().Msgf(format, args...)
}

func (z *ZerologAdapter) Debugj(j log.JSON) {
	z.logJSON(z.Logger.Debug(), j)
}

func (z *ZerologAdapter) Info(i ...interface{}) {
	z.Logger.Info().Msg(fmt.Sprint(i...))
}

func (z *ZerologAdapter) Infof(format string, args ...interface{}) {
	z.Logger.Info().Msgf(fmt.Sprint(args...))
}

func (z *ZerologAdapter) Infoj(j log.JSON) {
	z.logJSON(z.Logger.Info(), j)
}

func (z *ZerologAdapter) Warn(i ...interface{}) {
	z.Logger.Warn().Msg(fmt.Sprint(i...))
}

func (z *ZerologAdapter) Warnf(format string, args ...interface{}) {
	z.Logger.Warn().Msgf(format, args...)
}

func (z *ZerologAdapter) Warnj(j log.JSON) {
	z.logJSON(z.Logger.Warn(), j)
}

func (z *ZerologAdapter) Error(i ...interface{}) {
	z.Logger.Error().Msg(fmt.Sprint(i...))
}

func (z *ZerologAdapter) Errorf(format string, i ...interface{}) {
	z.Logger.Error().Msgf(format, i...)
}

func (z *ZerologAdapter) Errorj(j log.JSON) {
	z.logJSON(z.Logger.Error(), j)
}

func (z *ZerologAdapter) Fatal(i ...interface{}) {
	z.Logger.Fatal().Msg(fmt.Sprint(i...))
}

func (z *ZerologAdapter) Fatalf(format string, i ...interface{}) {
	z.Logger.Fatal().Msgf(format, i...)
}

func (z *ZerologAdapter) Fatalj(j log.JSON) {
	z.logJSON(z.Logger.Fatal(), j)
}

func (z *ZerologAdapter) Panic(i ...interface{}) {
	z.Logger.Panic().Msg(fmt.Sprint(i...))
}

func (z *ZerologAdapter) Panicf(format string, i ...interface{}) {
	z.Logger.Panic().Msgf(format, i...)
}

func (z *ZerologAdapter) Panicj(j log.JSON) {
	z.logJSON(z.Logger.Panic(), j)
}

func (z *ZerologAdapter) Print(i ...interface{}) {
	z.Logger.WithLevel(zerolog.NoLevel).Str("level", "-").Msg(fmt.Sprint(i...))
}

func (z *ZerologAdapter) Printf(format string, i ...interface{}) {
	z.Logger.WithLevel(zerolog.NoLevel).Str("level", "-").Msgf(format, i...)
}

func (z *ZerologAdapter) Printj(j log.JSON) {
	z.logJSON(z.Logger.WithLevel(zerolog.NoLevel).Str("level", "-"), j)
}

func (z *ZerologAdapter) logJSON(event *zerolog.Event, j log.JSON) {
	for k, v := range j {
		event = event.Interface(k, v)
	}

	event.Msg("")
}
