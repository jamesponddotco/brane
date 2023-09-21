brane(1)

# NAME

brane - log and query your thoughts with the power of AI

# SYNOPSIS

*brane* [OPTIONS] COMMAND [ARGUMENTS]

# DESCRIPTION

*brane* is a stupid simple and easy-to-use command-line application for logging
daily thoughts and seeking insights through OpenAI. By utilizing a simple and
intuitive interface, you can effortlessly record ideas and leverage the
capabilities of OpenAI to query your notes.

# OPTIONS

*-d*, *--directory*
	Specifies the directory where note files are stored. Defaults to
	_~/.local/share/brane_. Can also be set with the _BRANE_DIRECTORY_
	environment variable.

*-m*, *--model*
	Specifies the OpenAI model to use. Defaults to _gpt-3.5-turbo-16k_. Can also
	be set with the _BRANE_MODEL_ environment variable.

*-k*, *--key*
	Specifies the OpenAI API key to use. Mandatory for the *ask* command. Can be
	set with the _BRANE_KEY_ environment variable.

*-h*, *--help*
	Show help message and quit.

*-v*, *--version*
	Show version number and quit.

# COMMANDS

*l*, *log* [THOUGHT]
	Log a new thought.

*a*, *ask* [QUERY]
	Ask the AI questions about your thoughts.

# ENVIRONMENT

*BRANE_DIRECTORY*
	Specifies the directory where note files are stored.

*BRANE_MODEL*
	Specifies the OpenAI model to use.

*BRANE_KEY*
	Specifies the OpenAI API key to use.

# EXAMPLES

*1. Log a new thought*
	$ brane log "Had an amazing dinner at Joe's diner. The steak was delicious."

*2. Ask the AI questions about your thoughts*
	$ brane ask "What did I eat last night?"

# AUTHORS

Maintained by James Pond <james@cipher.host>.
