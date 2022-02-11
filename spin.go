package spunda

import "sync"

// Spinner types.
var (
	Box1    = `⠋⠙⠹⠸⠼⠴⠦⠧⠇⠏`
	Box2    = `⠋⠙⠚⠞⠖⠦⠴⠲⠳⠓`
	Box3    = `⠄⠆⠇⠋⠙⠸⠰⠠⠰⠸⠙⠋⠇⠆`
	Box4    = `⠋⠙⠚⠒⠂⠂⠒⠲⠴⠦⠖⠒⠐⠐⠒⠓⠋`
	Box5    = `⠁⠉⠙⠚⠒⠂⠂⠒⠲⠴⠤⠄⠄⠤⠴⠲⠒⠂⠂⠒⠚⠙⠉⠁`
	Box6    = `⠈⠉⠋⠓⠒⠐⠐⠒⠖⠦⠤⠠⠠⠤⠦⠖⠒⠐⠐⠒⠓⠋⠉⠈`
	Box7    = `⠁⠁⠉⠙⠚⠒⠂⠂⠒⠲⠴⠤⠄⠄⠤⠠⠠⠤⠦⠖⠒⠐⠐⠒⠓⠋⠉⠈⠈`
	Spin1   = `|/-\`
	Spin2   = `◴◷◶◵`
	Spin3   = `◰◳◲◱`
	Spin4   = `◐◓◑◒`
	Spin5   = `▉▊▋▌▍▎▏▎▍▌▋▊▉`
	Spin6   = `▌▄▐▀`
	Spin7   = `╫╪`
	Spin8   = `■□▪▫`
	Spin9   = `←↑→↓`
	Loadbar1 = `[] [=] [==] [===] [====] [=====] [======] [=======] [========] [=========] [==========]`
	Loadbar2 = `█▒▒▒▒▒▒▒▒▒ ███▒▒▒▒▒▒▒ █████▒▒▒▒▒ ███████▒▒▒ ██████████`
	Loadbar3 = `[ ] [=> ] [===> ] [=====> ] [======> ] [========> ] [==========> ] [============> ] [==============> ] [================> ] [==================> ] [===================>]`
	Default = Loadbar1
)

// Spinner is exactly what you think it is.
type Spinner2 struct {
	mu     sync.Mutex
	frames []rune
	length int
	pos    int
}

// New returns a spinner initialized with Default frames.
func New2() *Spinner2 {
	s := &Spinner{}
	s.Set(Default)
	return s
}

// Set frames to the given string which must not use spaces.
func (s *Spinner2) Set2(frames string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.frames = []rune(frames)
	s.length = len(s.frames)
}

// Current returns the current rune in the sequence.
func (s *Spinner2) Current2() string {
	s.mu.Lock()
	defer s.mu.Unlock()
	r := s.frames[s.pos%s.length]
	return string(r)
}

// Next returns the next rune in the sequence.
func (s *Spinner2) Next2() string {
	s.mu.Lock()
	defer s.mu.Unlock()
	r := s.frames[s.pos%s.length]
	s.pos++
	return string(r)
}

// Reset the spinner to its initial frame.
func (s *Spinner2) Reset2() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.pos = 0
}
