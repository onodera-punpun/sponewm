/*
spone uses SPoneWM's IPC mechanism to send commands and output the response
to stdout. spone can also be used to view documentation on all available
commands.

Usage:
	spone [flags] [command]

Note that 'command' MUST be specified as a single argument. So that this is
illegal:

	spone AddWorkspace "new-workspace-name"

Use this instead:

	spone 'AddWorkspace "new-workspace-name"'

The flags are:
	-f file
		When set, commands will be read from the specified file. The file
		may contain multiple commands, where each is one its own line and
		enclosed with '(' and ')'.
		If "-" is used, commands will be read from stdin.
	--list
		List all commands and the names of each parameter for each command.
	--list-types
		The same as '--list', but includes the types of each parameter.
	--list-usage
		The same as '--list-types', but includes a description of the command.
	--usage command-name
		Get the usage information for the command named "command-name".
	--poll milliseconds
		When milliseconds is greater than 0, the given commands will be
		executed at the specified interval.
*/
package main
