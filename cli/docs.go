// Package cli provides the command-line interface for the autocommit tool.
package cli

const description = `Autocommit inspects your staged changes, asks the configured LLM for a commit message, then lets you choose to commit, copy to clipboard, regenerate, add custom instructions, or exit.`

const example = `# 1) stage your work
git add .

# 2) run autocommit
autocommit

# 3) you’ll see something like:
🤖 Using model: GPT-4
💬 Commit message:
==============================================================================================
chore(cli): add interactive AI-powered commit generator
==============================================================================================
What now?
✅ Commit
📋 Copy to clipboard
🔄 Regenerate
📝 Add instruction
🚪 Exit`
