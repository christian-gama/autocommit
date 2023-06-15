package git

func MakeDiffCommand() DiffCommand {
	return NewDiffCommand(MakeExecutor())
}

func MakeCommitCommand() CommitCommand {
	return NewCommitCommand(MakeExecutor())
}

func MakeExecutor() Executor {
	return NewExecutor()
}
