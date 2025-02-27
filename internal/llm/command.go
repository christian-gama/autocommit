package llm

// chatCommandImpl is an implementation of ChatCommand.
type chatCommandImpl struct {
	chat Chat
}

// Execute implements the ChatCommand interface.
func (c *chatCommandImpl) Execute(config Config, system *System, input string) (string, error) {
	response, err := c.chat.Response(config, system, input)
	if err != nil {
		return "", err
	}
	return response, nil
}

// NewChatCommand creates a new instance of ChatCommand.
func NewChatCommand(chat Chat) ChatCommand {
	return &chatCommandImpl{
		chat: chat,
	}
}

// verifyConfigCommandImpl is an implementation of VerifyConfigCommand.
type verifyConfigCommandImpl struct {
	repo ConfigRepo
}

// Execute Implements the VerifyConfigCommand interface.
func (v *verifyConfigCommandImpl) Execute(
	getConfigsFn func() (Config, error),
) (config Config, err error) {
	ok := v.repo.Exists()
	if !ok {
		config, err = getConfigsFn()
		if err != nil {
			return nil, err
		}

		if err := v.repo.SaveConfig(config); err != nil {
			return nil, err
		}
	} else {
		config, err = v.repo.GetConfig()
		if err != nil {
			return nil, err
		}
	}

	return config, err
}

// NewVerifyConfigCommand creates a new instance of VerifyConfigCommand.
func NewVerifyConfigCommand(repo ConfigRepo) VerifyConfigCommand {
	return &verifyConfigCommandImpl{
		repo: repo,
	}
}

// resetConfigCommandImpl is an implementation of ResetConfigCommand.
type resetConfigCommandImpl struct {
	repo ConfigRepo
}

// Execute Implements the ResetConfigCommand interface.
func (r *resetConfigCommandImpl) Execute() error {
	if !r.repo.Exists() {
		return nil
	}

	return r.repo.DeleteConfig()
}

// NewResetConfigCommand creates a new instance of ResetConfigCommand.
func NewResetConfigCommand(repo ConfigRepo) ResetConfigCommand {
	return &resetConfigCommandImpl{
		repo: repo,
	}
}
