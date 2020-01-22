package status

// Status ...
type Status interface {
	GetStatus() string
}

type status struct {
	text string
}

// Install auto-status
func (s *status) GetStatus() string {
	s.text = "Nice day! :)"
	return "Status: " + s.text
}

func NewStatus() Status {
	return &status{}
}
