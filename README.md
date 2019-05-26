A simple example of writing a CLI using Golang using the Cobra module: https://github.com/spf13/cobra

Example output:


C02VQ1S6HTDD:golang_cli_tutorial nsiddiq$ go build

C02VQ1S6HTDD:golang_cli_tutorial nsiddiq$ ls

LICENSE			golang_cli_tutorial	main.go
README.md		lib
cmd			main

C02VQ1S6HTDD:golang_cli_tutorial nsiddiq$ ./golang_cli_tutorial

hooray my root command works

C02VQ1S6HTDD:golang_cli_tutorial nsiddiq$ ./golang_cli_tutorial sun

Error: unknown command "sun" for "golang_cli_tutorial"


Did you mean this?

	sum


Run 'golang_cli_tutorial --help' for usage.

unknown command "sun" for "golang_cli_tutorial"


Did you mean this?

	sum


C02VQ1S6HTDD:golang_cli_tutorial nsiddiq$ ./golang_cli_tutorial sum

No arguments given. Expected 1 or more integers.

Usage:

  golang_cli_tutorial sum [flags]


Flags:

  -h, --help   help for sum


Global Flags:

      --config string   config file (default is $HOME/.golang_cli_tutorial.yaml)

C02VQ1S6HTDD:golang_cli_tutorial nsiddiq$ ./golang_cli_tutorial sum a b

ERROR: Conversion of value a to integer failed.

C02VQ1S6HTDD:golang_cli_tutorial nsiddiq$ ./golang_cli_tutorial sum 1 1

The sum of your integers is: 2

C02VQ1S6HTDD:golang_cli_tutorial nsiddiq$ ./golang_cli_tutorial sum 1 1 1 2

The sum of your integers is: 5
