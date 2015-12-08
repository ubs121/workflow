package workflow

type (
	Context interface{}

	Workflow struct {
		Start     *Page
		OnFailure FailureFunc
		Context   Context

		queue   []*Page
		inQueue map[*Page]bool
	
		recordsPerPage int
	}

)

func New() *Workflow {
	w := &Workflow{}
	w.inQueue = make(map[*Page]bool)
	return w
}

func (w *Workflow) Run() error {
	w.loadQueue(w.Start)
	for _, step := range w.queue {
		fmt.Printf("Running step: %s ", step.Label)
		if err := step.Run(w.Context); err != nil {
			if err := w.OnFailure(err, step, w.Context); err != nil {
				fmt.Println("FAILED")
				return err
			}
		}
		fmt.Println("COMPLETE")
	}
	return nil
}

func (w *Workflow) loadQueue(s *Page) {
	if s == nil {
		return
	}

	for _, step := range s.DependsOn {
		w.loadQueue(step)
	}

	if !w.inQueue[s] {
		w.inQueue[s] = true
		w.queue = append(w.queue, s)
	}
	return
}


func (scr *Screen) process(ses USSDSession) string {
	responseHandler := new(ResponseHandler)

	pages = []
	ses.recordsPerPage = 0

	// If no backend hit is required
	if len(scr.services) == 0 {
		responseHandler.handleStaticResponse(ses, s.successText)
		ses.getPage[0].userOptions = s.userOptions
	} else {
		for service : range scr.services {
			if !service.invoke() {
				// Load error message screen and show error message from back-end

			}
		}
	}

	page:=ses.getPage(0)
	ses.userOptions = page.userOptions
	return page.Text
}
