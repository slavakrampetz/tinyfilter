package args

type Command string
type Option int

const (
	OptsNone Option = iota
	OptsYtOn
	OptsYtOff
)

// Long commands
const (
	Help    Command = "help"
	Web     Command = "web"
	Reload  Command = "reload"
	Youtube Command = "youtube"
)

// Short commands
const (
	cHelpShort   string = "h"
	cReloadShort string = "r"
	cYoutube     string = "yt"
)

// Youtube options
const (
	ytOff string = "off"
	ytOn  string = "on"
)
