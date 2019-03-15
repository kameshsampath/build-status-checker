package types

//KbscOptions provides the options for the running kbsc
type KbscOptions struct {
	KubeConfig string
	LogLevel   string
}

//PollOptions - the options for knative build poll function
type PollOptions struct {
	BuildName string
	Namespace string
}
