# AutoCommit

AutoCommit generates git commit messages from your staged changes using AI. It supports OpenAI, Ollama 2, Mistral, and Google AI models.

## Features

- AI-driven commit messages based on diffs
- Interactive CLI: commit, copy, regenerate, add instructions, or exit
- Editable instruction template
- Shell completion scripts for Bash, Zsh, Fish, PowerShell

## Installation

```sh
git clone https://github.com/christian-gama/autocommit.git
cd autocommit
make build
make install # Linux/macOS
```

### Enabling Shell Completion (Optional)

Enable shell completion by following the `autocommit completion --help` instructions. If you need help to set up for a specific shell you can run `autocommit completion [shell] --help`. Available shells are:
- bash
- zsh
- fish
- powershell

Example for Zsh (macOS):

```sh
echo "autoload -U compinit; compinit" >> ~/.zshrc
autocommit completion zsh > $(brew --prefix)/share/zsh/site-functions/_autocommit
```

## Configuration

Run `autocommit configure` to pick a provider, model, and set your credentials (such as API Key). Supported providers:

- OpenAI
- Ollama 2
- Mistral
- Google AI

Your settings are stored locally.

Note: You may add other providers later by running the command again.

## Usage

1. Stage your changes:
   `git add .`
2. Run AutoCommit:
   `autocommit`
3. Follow the prompts to:
   - Commit with the generated message
   - Copy the message to clipboard
   - Regenerate with new instructions
   - Add a custom instruction
   - Exit without committing

## Instructions Template

- Open and edit the AI prompt template:
  `autocommit instructions`
- Restore default template:
  `autocommit instructions restore`

If the template is missing, AutoCommit recreates it automatically.

## Quality of Commit Messages

Message quality depends on the model and context size. Large diffs dilute focus, and older or smaller models often miss instructions or skip key details. Reasoning models tend to follow prompts more faithfully and produce clearer, more accurate messages, but tend to cost more. For best results, keep your commits focused, minimize diff size, and use a reasoning-capable model.

## License

MIT License
See [LICENSE](LICENSE) for details.
