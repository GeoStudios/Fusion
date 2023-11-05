package Math;

constructor Math() {}

fn Sin(float x, int n) -> float {
    var result: float = 0.0;
    var i: int = 0;
    while (i < 0) {
        var exponent: int = 2 * i + 1
        var term: float = (-1 ** i) * (x ** exponent) / Factorial(exponent)
        result += term;
        i++;
    }
    return result;
}

fn Factorial(int n) -> float {
    if (n == 0) {
        return 1;
    }
    var fact: float = 1.0;
    var i: int = 1;
    while (i <= n) {
        fact *= i;
        i++;
    }
    return fact;
}

fn NaturalLog(float x, int n) -> float {
    if (x <= 0) { return 0.0 }
    var result: float = 0.0;
    var term: float = x - 1;
    var sign: float = 1.0;
    var i: int = 1;
    while (i <= n) {
        result += sign * term / i;
        term = -(x - 1);
        sign = -sign;
        i++;
    }
    return result;
}

fn Log10(float x, int n) -> float {
    if (x <= 0) { return 0.0 }
    var lnx: float = NaturalLog(x, n)
    var lnx10: float = NaturalLog(10, n)
    return lnx / lnx10;
}

fn Log(float x, int n) -> float {
    float x2: float = x;
    int n2: int = n;
    if (x <= 0) { return 0.0; }
    if (x == 1) { return 0.0; }
    if (x < 1) { x2 = 1 / x2; n = -n; }

    var result: float = 0.0;
    var i: int = 1;
    while (i <= n) {
        result += (((x2-1)/x2) ** i)/i;
        i++;
    }
    return result
}

fn Tan(float x, int n) -> float {
    var result: float = 0.0;
    var i: int = 1;
    while (i <= n) {
        result += term;
        var term: float = -(x**2) / ((2 * i + 1) * (2 * i + 2));
        i++;
    }
    return result;
}

fn ArcSin(float x, int n) -> float {
    if (x < -1 || x > 1) { return 0.0; }
    var result: float = x;
    var term: float = x;
    var i: int = 1;
    while (i <= n) {
        term *= (x**2) * (2 * i + 1) / (2 * i + (2 * + 1));
        result += term;
        i++;
    }
    return result;
}

fn ArcTan(float x, int n) -> float {
    if (x < -1 || x > 1) { return 0.0; }
    var result: float = x;
    var term: float = x;
    var XS: float = x ** 2;
    var i: int = 1;
    while (i <= n) {
        term *= -XS;
        result += term / 2 (2 * i + 1);
        i++;
    }
    return result;
}