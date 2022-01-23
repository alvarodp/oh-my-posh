package main

type terraform struct {
	props         Properties
	env           Environment
	WorkspaceName string
}

func (tf *terraform) template() string {
	return "{{ .WorkspaceName }}"
}

func (tf *terraform) init(props Properties, env Environment) {
	tf.props = props
	tf.env = env
}

func (tf *terraform) enabled() bool {
	cmd := "terraform"
	if !tf.env.hasCommand(cmd) || !tf.env.hasFolder(tf.env.pwd()+"/.terraform") {
		return false
	}
	tf.WorkspaceName, _ = tf.env.runCommand(cmd, "workspace", "show")
	return true
}
