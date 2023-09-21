# brane

> Capture thoughts. Ask questions. Simplify your life.

`brane` is a very simple and intuitive command-line application for
quickly logging daily thoughts and seeking insights. Seamlessly jot down
ideas and leverage the power of OpenAI to query your notes.

## Why `brane` exists

As someone with
[ADHD](https://en.wikipedia.org/wiki/Attention_deficit_hyperactivity_disorder),
I forget things a lot. Most of the times I can't remember what happened
during my day, which is a big problem in therapy sessions. I wanted to
remember simple things like when I woke up or went to sleep, if I ate
lunch or not, or all those random ideas that I have throughout the day.

So, I built `brane`, a tool to write down quick notes during the day.
And with the OpenAI integration, I can ask questions and get insights
about my life, like what did I do this week, how many times I went to
the gym, or what did I eat last night.

`brane` made my therapy sessions more productive and richer in details.
I hope it can help others too.

## Features

- **Instant logging:** Log your thoughts in a jiffy. Simply type and let
  Brane handle the formatting.
- **OpenAI integration:** Don't just store notes; ask questions and get
  insights. Dive deep into your notes with the power of AI.
- **Simple structure:** Say goodbye to messy note files, databases, or
  proprietary file formats. With `brane`, every day gets its own
  Markdown file, neatly named and easy to navigate.

## Installation

### From source

First install the dependencies:

- Go 1.21 or above.
- make.
- [scdoc](https://git.sr.ht/~sircmpwn/scdoc).

Then compile and install:

```bash
make
sudo make install
```

## Usage

```bash
$ brane --help
NAME:
   brane - log and query your thoughts with the power of AI

USAGE:
   brane [global options] command [command options] [arguments...]

VERSION:
   0.1.0

COMMANDS:
   log, l  log a new thought
   ask, a  ask the AI questions about your thoughts

GLOBAL OPTIONS:
   --directory value, -d value  the directory where note files are stored (default: "/home/james/.local/share/brane") [$BRANE_DIRECTORY]
   --model value, -m value      the OpenAI model to use (default: "gpt-3.5-turbo-16k") [$BRANE_MODEL]
   --key value, -k value        the OpenAI API key to use [$BRANE_KEY]
   --help, -h                   show help
   --version, -v                print the version
```

See _brane(1)_ after installing for more information.

## Contributing

Anyone can help make `brane` better. Send patches on the [mailing
list](https://lists.sr.ht/~jamesponddotco/brane-devel) and report bugs
on the [issue tracker](https://todo.sr.ht/~jamesponddotco/brane).

You must sign-off your work using `git commit --signoff`. Follow the
[Linux kernel developer's certificate of
origin](https://www.kernel.org/doc/html/latest/process/submitting-patches.html#sign-your-work-the-developer-s-certificate-of-origin)
for more details.

All contributions are made under [the GPL-2.0 license](LICENSE.md).

## Resources

The following resources are available:

- [Support and general discussions](https://lists.sr.ht/~jamesponddotco/brane-discuss).
- [Patches and development related questions](https://lists.sr.ht/~jamesponddotco/brane-devel).
- [Instructions on how to prepare patches](https://git-send-email.io/).
- [Feature requests and bug reports](https://todo.sr.ht/~jamesponddotco/brane).

---

Released under the [GPL-2.0 license](LICENSE.md).
