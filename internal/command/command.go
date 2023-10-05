package command

type Command interface {
	Run() (string, error)
	Name() string
	Description() string
}
