package msgbuf

type Msgbuf struct {
	msg            []string
	overflow       bool
	length, ri, wi int
}

func NewMsgbuf(length int) *Msgbuf {
	return &Msgbuf{length: length, msg: make([]string, length)}
}

func (m *Msgbuf) Add(msg string) {
	m.msg[m.wi] = msg
	if m.overflow && m.wi == m.ri {
		m.ri++
	}
	m.wi++
	if m.wi == m.length {
		m.wi = 0
		m.overflow = true
	}

	if m.ri == m.length {
		m.ri = 0
	}
}

func (m *Msgbuf) Get(n int) []string {
	b := make([]string, n)
	for i := 0; i < n; i++ {
		b[i] = m.msg[(i+m.ri)%m.length]
	}
	return b
}
