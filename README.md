<!-- # mana
Interpreted Toy Programming Language written in Go -->

# Mana

Mana is a toy programming language written in Go. It is a dynamically typed, interpreted language with a C-like syntax. 


> **Note**: The language is still in development and is not yet usable. The documentation below is a work in progress and is subject to change.

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
| `CallExpression` | NYI | Call Expressions are used to call functions | `add(5, 5)` | NYI |
| `StringLiteralExpression` | NYI | String Literal Expressions are used to represent string values | `"Hello, World!"` | NYI |
| `ArrayLiteralExpression` | NYI | Array Literal Expressions are used to represent array values | `[1, 2, 3]` | NYI |
| `IndexExpression` | NYI | Index Expressions are used to index into arrays | `myArray[0]` | NYI |
| `HashLiteralExpression` | NYI | Hash Literal Expressions are used to represent hash values | `{"key": "value"}` | NYI |

\**NYI = Not Yet Implemented*

## REPL

Mana will come with a REPL (Read-Eval-Print-Loop) that can be used to evaluate Mana code.

> **Note**: The REPL is still in development and is not yet usable. Currently, it can only parse and print the AST of the input code.

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
let result = if (x < y) { add(x, y) } else { subtract(x, y) };
```