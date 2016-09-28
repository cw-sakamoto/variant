package step

import "text/template"

type ExecutionContext interface {
	GenerateAutoenv() (map[string]string, error)
	Caller() []Caller
	Key() Key
	Vars() map[string]interface{}
	CreateFuncMap() template.FuncMap
	ProjectName() string
	Autoenv() bool
	Autodir() bool
	Interactive() bool
	RunAnotherFlow(key string) (string, error)
}
