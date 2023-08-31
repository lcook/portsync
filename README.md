## ⌛ portsync

![](https://github.com/lcook/portsync/actions/workflows/build.yaml/badge.svg)

Command-line utility tailored for FreeBSD, focused on management of package
updates, version tracking and streamlined commits of updated packages.

_Caveat emptor: Still a work in progress, not all cases are covered and you
are likely to encounter bugs. Little is done in terms of error-checking as primarily
we run a set of Makefiles for a bulk of the work. Note that static variables, e.g.,
build SHA hashes are not updated, there is still a level of manual work required
when updating packages albeit greatly reduced._

![Command-line demonstration](./demo.gif)

[Features](#features) | [Get started](#get-started) | [Usage](#usage) | [License](#license)

## Features

- Fetch and display the latest package versions from [portscout](https://portscout.freebsd.org/).
- Apply new updates to your local ports tree based.
- Streamlined commit process for updated packages.
- Conveniently build and test any changes made.

## Get started

1. Ensure you have Go installed on your system (minimum of 1.21). 
2. Clone the repository.
3. Build and install.
```sh
git clone https://github.com/lcook/portsync
cd portsync && make build install
```

#### Dependencies

We utilize a plethora of external tools used in helping us achieve a lot
of the functionality provided. Make sure that you have `portfmt` and `modules2tuple`
installed.
```sh
pkg install portfmt modules2tuple
```

#### Configuration file

There is an [included example](.portsync.example) configuration. You most only
need to worry about updating the `maintainer` field with your email, and setting
`base` (your ports tree) correctly. Copy this file to your $HOME directory.

Alternatively, you can specify a custom location for your configuration file by
passing `-c` to the utility. Likewise with the values in the configuration you
may pass them to the utility. See `portsync -h` for the appropriate flags.

## Usage

The general workflow is pretty straightforward.

Display any potential package updates for a given maintainer
```sh
portsync fetch -m "email@example.com"
```

If no maintainer is explicitly passed, it is populated by default with `ports@FreeBSD.org`
(ports with no maintainer prescribed), falling back to values found in the
configuration file as mentioned above. Similarly with the directory containing
your ports tree, defaulting to `/usr/ports`.

Display any potential package updates for a given maintainer, only processing
the packages updates for `foo/bar` and `bar/baz` (if any applicable)
```sh
portsync fetch -m "email@example.com" -o foo/bar -o bar/baz
```

Update and commit the package `foo/bar`
```sh
portsync update -o foo/bar --commit # Use -g for a shorthand to --commit
```

Build package `foo/bar`
```sh
portsync build -o foo/bar
```

Run package `foo/bar`'s test suite (if any applicable)
```sh
portsync test -o foo/bar
```

Having the ability to define a custom format specifier can be helpful when
scripting. For example this neat one-liner, providing a nice terminal interface
of available package updates through `fzf`
```sh
portsync fetch -f "%o" | fzf --multi \
  --header "Select port(s) to update" --preview "pkg rquery -r FreeBSD '%e' {}" \
  --preview-window=up | xargs portsync update -o
```

Or even be able to output updates in a particular file format, such as JSON
```sh
portsync fetch -f '{"origin": "%o", "current": "%v", "latest": "%l"}'
```

Happy hacking!

## License

[BSD 2-Clause](LICENSE)
