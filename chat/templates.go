package chat

const ShortTemplate = `
Objective: Generate a very brief commit messages based on the output of a 'git diff' command in the style of Conventional Commits in a very concise and brief manner, with no additional information.

UserInput: A string representing the output from a 'git diff' command.

Output: A string that represents a commit message. This message should adhere to the Conventional Commits specification. It should include a type (such as "feat", "fix", "chore", etc.), an optional scope (which could be a specific part of the codebase), and a succinct description of the change. The description should provide information on what the commit does, not what it's doing. It should be the shortest possible message that conveys the necessary information.

Structure of the output: 
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]

Allowed types:
feat: Represents the introduction of a new feature to the codebase.
chore: Examples include updating dependencies or tweaking build scripts.
fix: Is used to indicate that a particular change resolves an issue or bug in the codebase.
refactor: Represents a change that does not add new functionality or fix a bug. Instead, it restructures the code in some way to make it cleaner, more efficient, or more maintainable.
ci: Involves changes to CI configuration files and scripts such as a CI pipeline.
build: Is used for changes that affect the build system or external dependencies.
docs: This can include changes to comments, markdown files, or any other form of documentation, such as README or LICENSE.
test: This can include adding new tests, refactoring existing tests, or fixing failing tests.
perf: Represents changes made with the intent of improving performance.
`

const VerboseTemplate = `
Objective: Generate commit messages based on the output of a 'git diff' command in the style of Conventional Commits.

UserInput: A string representing the output from a 'git diff' command.

Output: A string that represents a commit message. This message should adhere to the Conventional Commits specification. It should include a type (such as "feat", "fix", "chore", etc.), an optional scope (which could be a specific part of the codebase), and a succinct description of the change. The description should provide information on what the commit does, not what it's doing. It should add any additional information, such as body or footer, if necessary. 

Structure of the output: 
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]

Allowed types:
feat: Represents the introduction of a new feature to the codebase.
chore: Examples include updating dependencies or tweaking build scripts.
fix: Is used to indicate that a particular change resolves an issue or bug in the codebase.
refactor: Represents a change that does not add new functionality or fix a bug. Instead, it restructures the code in some way to make it cleaner, more efficient, or more maintainable.
ci: Involves changes to CI configuration files and scripts such as a CI pipeline.
build: Is used for changes that affect the build system or external dependencies.
docs: This can include changes to comments, markdown files, or any other form of documentation, such as README or LICENSE.
test: This can include adding new tests, refactoring existing tests, or fixing failing tests.
perf: Represents changes made with the intent of improving performance.
`
