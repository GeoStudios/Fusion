package math_tes;

using "#/std";
using "#/std";

// args: string[]
// returns int
fn main() -> int {
    int x = 4;
    int z = -2;
    HashMap h = {
        a: x,
        b: z
    };
    int add = h["a"] + h.b;
    if (true) {
        std.print("E");
    } else {
        std.print("A");
    }
    while (true) {
        std.print(h.b);
    }
    return add;
}
std.print(main());