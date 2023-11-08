package Math;

// Embeded Imports
using "@/StrongMath_Wrapper.f" as Strong;
using "@/Math_Area.f" as Area;
using "@/Math_Volumes.f" as Volumes;
using "@/Math_Vars.f" as Vars;

fn CalculateFibonacci(int n) -> int {

    if (n < 0) {
        std.print("Incorrect input", std.LineSeperator);
    } else if n == 0 {
        return 0;
    } else if (n == 1 or n == 2) {
        return 1;
    } else {
        return Calculate(n-1) + Calculate(n-2);
    }
}

fn IntToFloat(int a) -> float { return a * 1.0; }