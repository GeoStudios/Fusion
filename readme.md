# Fusion

Example:
```cs
package main;

using "#/std";

fn main() -> int {
    int x = 4;
    int z = -2;

    int add = x + z;
    std.print(add)
    return 0;
}
main();
```

All .f files must have a `fn main(string[] args) -> int {}` so it can launch it if it is the main file
the package name must be the same as the file name