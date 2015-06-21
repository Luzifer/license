# Luzifer / license

This small tool is mainly a helper for me not to have to look up the license text of the license I want to use every time again. With the help of this tool one command is sufficient to put a `LICENSE` file into the current directory filled with a chosen license.

The license template is pre-filled with the current year and the user information extracted from your `~/.gitconfig` file.

## Usage

```bash
# license help
license is a tool to quickly put your code under a license

Usage:
  license [command]

Available Commands:
  list        Lists all available licenses
  show        Prints the pre-filled license template to your terminal
  write       Write the pre-filled license template to ./LICENSE
  version     prints the current version of license
  badge       Gives you a markdown sniplet to show the license in your README.md
  help        Help about any command

Flags:
  -h, --help=false: help for license


Use "license help [command]" for more information about a command.
```

## Currently included licenses

- [Apache 2.0](https://www.apache.org/licenses/LICENSE-2.0)
- [BSD 2-Clause](http://opensource.org/licenses/BSD-2-Clause)
- [BSD 3-Clause](http://opensource.org/licenses/BSD-3-Clause)
- [GPLv3](http://www.gnu.org/licenses/gpl-3.0.en.html)
- [MIT](http://opensource.org/licenses/mit-license.html)
- [CC0 1.0](http://creativecommons.org/publicdomain/zero/1.0/)

## Adding more licenses

If you like this tool and need a new license included just open a pull-request with a new file named `lic<shortcode>.go` and the license template. The format to add the license to the application you can find in every other template file.
