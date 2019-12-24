package configmodel


type Target struct {
	Key   string
	Type  string
	Value string
}

type Operation struct {
	Key   string
	Type  string
	Value string
}

type Option struct {
	Cookie string
	Proxy  string
}
type Replace struct {
	Target string
	From   string
	To     string
}

type Transform struct {
	Target string
	From   string
	To     string
}

type After struct {
	Transform Transform
	Replace   Replace
}

type Before struct {
	Option  Option
	Replace Replace
}

type Action struct {
	Target    Target
	Operation Operation
	Before    Before
	After     After
	Return    string
}

type Actions struct {
	Action []Action
}
