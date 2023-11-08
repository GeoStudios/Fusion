package math_tes;

using "#/std";
using "@/StrongMath_Wrapper.f" as "Math";

// args: Array
// returns int
fn main() -> void {
    std.println(Math.Sin(69));
    std.println(Math.Sinh(69));
    std.println(Math.ASin(69));
    std.println(Math.ASinh(69));
    std.print(std.LineSeparator);
    std.println(Math.Cosine(69));
    std.println(Math.Cosineh(69));
    std.println(Math.ACosine(69));
    std.println(Math.ACosineh(69));
    std.print(std.LineSeparator);
    std.println(Math.Tanh(69));
    std.println(Math.Tan(69));
    std.println(Math.ATan(69));
    std.println(Math.ATanh(69));
    std.println(Math.ATan2(69, 69));
    return;
}
main();