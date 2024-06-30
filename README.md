# frauCLI

## Overview
frauCLI is a command-line interface for managing recurrent expenses.

## Building the Project
To build the project, you need to have `make` installed on your system. Run the following command in the root directory of the project:

```

make build

```


This will compile the project and generate the necessary binaries.

## Running the Project
Once the project is built, you can run the CLI commands. For example, to add a recurrent expense, use the following command:

```


./bin/frau add-recurrent-expense --name='Netflix' --amount=9.99 --frequency='Monthly' --start-date='2021-01-01' --end-date='2021-12-31'

``` 


This command will add a recurrent expense with the specified details.

This `README.md` provides instructions on how to build the project using a Makefile and how to run the CLI commands. It also includes an example output for adding a recurrent expense.

