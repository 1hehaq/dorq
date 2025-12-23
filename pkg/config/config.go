package config

type Config struct {
	Query               string
	Limit               int
	JsonOutput          bool
	ShowHelp            bool
	ShowVersion         bool
	Update              bool
	Sources             string
	Timeout             int
	RequestProxy        string
	ForwardProxy        string
	ShowBrowser         bool
	AntibotTimeout      int
	Debug               bool
	Verbose             bool
	VVerbose            bool
	ResubmitWithoutSubs bool
	Output              string
}
