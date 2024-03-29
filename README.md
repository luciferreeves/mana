<!-- # mana
Interpreted Toy Programming Language written in Go -->

# Mana

Mana is a toy programming language written in Go. It is a dynamically typed, interpreted language with a C-like syntax. 

> *Note*: The language is still in development and is not yet usable. The documentation below is a work in progress and is subject to change.

## Development Roadmap

| Implementation | Status | Specification | Example | Tests |
| --- | --- | --- | --- | --- |
| `LetStatement` | ✔️ | Let Statements are used to declare variables | `let x = 5;` | ✔️ |
| `ReturnStatement` | ✔️ | Return Statements are used to return values from functions | `return 5;` | ✔️ |
| `ExpressionStatement` | ✔️ | Expression Statements are used to evaluate expressions | `5 + 5;` | ✔️ |
| `IdentifierExpression` | ✔️ | Identifier Expressions are used to reference variables | `x` | ✔️ |
| `IntegerLiteralExpression` | ✔️ | Integer Literal Expressions are used to represent integer values | `5` | ✔️ |
| `PrefixExpression` | ✔️ | Prefix Expressions are used to represent prefix operators | `!true` | ✔️ |
| `InfixExpression` | ✔️ | Infix Expressions are used to represent infix operators | `5 + 5` | ✔️ |
| `BooleanLiteralExpression` | ✔️ | Boolean Literal Expressions are used to represent boolean values | `true` | ✔️ |
| `IfExpression` | ✔️ | If Expressions are used to represent conditional statements | `if (true) { return 5; }` | ✔️ |
| `BlockStatement` | ✔️ | Block Statements are used to represent blocks of code | `{ let x = 5; return x; }` | ✔️ |
| `FunctionLiteralExpression` | ✔️ | Function Literal Expressions are used to represent function definitions | `fn(x) { return x; }` | ✔️ |
| `CallExpression` | ✔️ | Call Expressions are used to call functions | `add(5, 5)` | ✔️ |
| `StringLiteralExpression` | NYI | String Literal Expressions are used to represent string values | `"Hello, World!"` | NYI |
| `ArrayLiteralExpression` | NYI | Array Literal Expressions are used to represent array values | `[1, 2, 3]` | NYI |
| `IndexExpression` | NYI | Index Expressions are used to index into arrays | `myArray[0]` | NYI |
| `HashLiteralExpression` | NYI | Hash Literal Expressions are used to represent hash values | `{"key": "value"}` | NYI |

\**NYI = Not Yet Implemented*

## REPL

Mana ships with a Read-Eval-Print-Loop (REPL) that can be used to evaluate Mana code. The REPL can be started by running the `mana` execultable. Note that you will need to [build the project](#building-the-project) before you can run the REPL.

```bash
/path/to/mana
```

Mana will then start the REPL and you can start typing Mana code. The REPL will evaluate the code and print the result.

```text
Hello <username>! Welcome to Mana REPL!

███╗░░░███╗░█████╗░███╗░░██╗░█████╗░
████╗░████║██╔══██╗████╗░██║██╔══██╗
██╔████╔██║███████║██╔██╗██║███████║
██║╚██╔╝██║██╔══██║██║╚████║██╔══██║
██║░╚═╝░██║██║░░██║██║░╚███║██║░░██║
╚═╝░░░░░╚═╝╚═╝░░╚═╝╚═╝░░╚══╝╚═╝░░╚═╝

>>> let x = 5;
```

## Syntax

Mana has a C-like syntax. The following is an example of a simple program written in Mana:

```rust
let x = 5;   // declare a variable named x and assign it the value 5
let y = 10;  // declare a variable named y and assign it the value 10

// this is a function that adds two numbers together
let add = fn(x, y) {
    return x + y;
};

// this is an alternative way to create a function
fn subtract(x, y) {
    return x - y;
}

// this will either add x and y if x is less than y, or return the difference between x and y
let result = if (x < y) { 
    add(x, y) 
} else { 
    subtract(x, y) 
}; // result = 15
```

## Types

Mana is a dynamically typed language. This means that the type of a variable is determined at runtime. The following are the types that Mana supports:

| Type | Description | Example |
| --- | --- | --- |
| `Integer` | A 64-bit signed integer | `5` |
| `Boolean` | A boolean value | `true` |

## Operators

Mana supports the following operators:

| Operator | Description | Example |
| --- | --- | --- |
| `+` | Addition | `5 + 5` |
| `-` | Subtraction | `5 - 5` |
| `*` | Multiplication | `5 * 5` |
| `/` | Division | `5 / 5` |
| `!` | Logical NOT | `!true` |
| `<` | Less Than | `5 < 5` |
| `>` | Greater Than | `5 > 5` |
| `==` | Equal To | `5 == 5` |
| `!=` | Not Equal To | `5 != 5` |

## Variables

Variables are declared using the `let` keyword. The variable name is followed by an equals sign and an expression. The expression is evaluated and the result is assigned to the variable. 

```rust
let x = 5;
```

## Conditionals

Mana supports If-Else conditionals. An `IfExpression` in Mana is composed of two parts: the condition and the consequence. The condition is an expression that evaluates to a boolean value. The consequence is a `BlockStatement` that is executed if the condition evaluates to `true`. The consequence is optional. If the condition evaluates to `false` and there is no consequence, then the `IfExpression` evaluates to `null`. If there is a consequence, then the `IfExpression` evaluates to the value of the last statement in the consequence.

```rust
if (x < y) {
    x + y;
} else {
    x - y;
}
```

## Functions

Mana supports first-class functions. This means that functions can be passed as arguments to other functions, returned from functions, and assigned to variables. The following is an example of a function definition in Mana.

Functions are defined using the `fn` keyword. The function name is followed by a list of parameters in parentheses. The function body is enclosed in curly braces. The function body is a `BlockStatement`, which means that it is a list of statements enclosed in curly braces. The last statement in the function body is the `return` statement, which is used to return a value from the function. Functions, themselves, are `ExpressionStatements`, which means that they evaluate to a value. The value that a function evaluates to is the value that is returned from the function.

```rust
fn add(x, y) {
    return x + y;
}
```

## Building the Project

To build the project, you will need to have Go installed on your machine. You can download Go from the [official website](https://golang.org/). Once you have Go installed, you can build the project by running the following command:

```bash
go build
```

This will create an executable file named `mana` in the root of the project directory.

## Running the Tests

To run the tests, you can use the `go test` command. This will run all of the tests in the project.

```bash
go test ./...
```

Or in a specific package:

```bash
go test ./lexer
```

If you want to run a specific test, you can use the `-run` flag to specify a regular expression that matches the test names.

```bash
go test -run TestLetStatement
```

Furthermore, you can use the `-v` flag to get verbose output from the tests.

```bash
go test -v ./...
```

## License
This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.
