A simple example of writing a CLI using Golang using the Cobra module.

Example output:

C02VQ1S6HTDD:golang_cli_tutorial nsiddiq$ ./main

hooray my root command works

C02VQ1S6HTDD:golang_cli_tutorial nsiddiq$ ./main sum

No arguments given. Expected 1 or more integers.

Usage:

  golang_cli_tutorial sum [flags]


Flags:

  -h, --help   help for sum


Global Flags:

      --config string   config file (default is $HOME/.golang_cli_tutorial.yaml)

C02VQ1S6HTDD:golang_cli_tutorial nsiddiq$ ./main sum a b

ERROR: Conversion of value a to integer failed.

C02VQ1S6HTDD:golang_cli_tutorial nsiddiq$ ./main sum 1 1

The sum of your integers is: 2

C02VQ1S6HTDD:golang_cli_tutorial nsiddiq$ ./main sum 1 1 1 2

The sum of your integers is: 5
