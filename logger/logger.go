package logger

import(
)

type Logger interface {
  Info(tag string, log string) error
  Debug(tag string, log string) error
  Error(tag string, log string) error
}
