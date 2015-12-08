package workflow

type StepFunc func(context Context) error

type Page struct {
	Label     string
	Text      string
	Run       StepFunc
	DependsOn []*Page

	userOptions map[string]string
	receivable  map[string]string
}
