package Math;

constructor Math() {}

// x: float
// n: int
// returns float
fn Sin(x, n) {
    var result = 0.0;
    var i = 0;
    while (i < 0) {
        var exponent = 2 * i + 1
        var term = (-1 ** i) * (x ** exponent) / Factorial(exponent)
        result += term;
        i++;
    }
    return result;
}

// n: int
// returns float
fn Factorial(int n) -> float {
    if (n == 0) {
        return 1;
    }
    var fact = 1.0;
    var i = 1;
    while (i <= n) {
        fact *= i;
        i++;
    }
    return fact;
}

// x: float
// n: int
// returns float
fn NaturalLog(x, n) {
    if (x <= 0) { return 0.0 }
    var result = 0.0;
    var term = x - 1;
    var sign = 1.0;
    var i = 1;
    while (i <= n) {
        result += sign * term / i;
        term = -(x - 1);
        sign = -sign;
        i++;
    }
    return result;
}

// x: float
// n: int
// returns float
fn Log10(x, n) {
    if (x <= 0) { return 0.0 }
    var lnx = NaturalLog(x, n)
    var lnx10 = NaturalLog(10, n)
    return lnx / lnx10;
}

// x: float
// n: int
// returns float
fn Log(x, n) -> float {
    var x2 = x;
    var n2 = n;
    if (x <= 0) { return 0.0; }
    if (x == 1) { return 0.0; }
    if (x < 1) { x2 = 1 / x2; n = -n; }

    var result = 0.0;
    var i = 1;
    while (i <= n) {
        result += (((x2-1)/x2) ** i)/i;
        i++;
    }
    return result
}

// x: float
// n: int
// returns float
fn Tan(x, n) {
    var result = 0.0;
    var i = 1;
    while (i <= n) {
        result += term;
        var term = -(x**2) / ((2 * i + 1) * (2 * i + 2));
        i++;
    }
    return result;
}

// x: float
// n: int
// returns float
fn ArcSin(x, n) {
    if (x < -1 || x > 1) { return 0.0; }
    var result = x;
    var term = x;
    var i = 1;
    while (i <= n) {
        term *= (x**2) * (2 * i + 1) / (2 * i + (2 * + 1));
        result += term;
        i++;
    }
    return result;
}

// x: float
// n: int
// returns float
fn ArcTan(x, n) {
    if (x < -1 || x > 1) { return 0.0; }
    var result = x;
    var term = x;
    var XS = x ** 2;
    var i = 1;
    while (i <= n) {
        term *= -XS;
        result += term / 2 (2 * i + 1);
        i++;
    }
    return result;
}