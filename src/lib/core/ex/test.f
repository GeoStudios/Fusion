package test;

// Embeded Imports
using "@/Math_Area.f" as Area;
using "@/LinkedList.f";

// Native Imports
using "#/std";

fn main() -> void {
    HashMap list = LinkedList.Construct();
    std.print(LinkedList.String(list), std.LineSeparator);
    return null;
}

main();