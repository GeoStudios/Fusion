package main;

class test : Object {
    main test() {

    }
}

constructor hello_world() -> {
    print("Hello World")
    return 0;
}

fn char(string[] args) -> int {
    var jim: string = "Jimmy";
    var tom: string = "Tommy";
    
    print(jim + "does not like" + tom);
    return 0;

/*
    16 ** 2
    16 *^ 2  = 4
    16 *^ 3
    16 *^ 4
*/  

// ** = Power of. For instance 2 ** 3 = 8
// *^ = Sqroot of. For instance 16 *^ 2= 4
/*  
    16 ** 2 = 256
    16 *^ 2 = 4

    16 *^69 = 4

    16 ** 69 = 
*/

/*
*/

}

fn test(string[] args) -> int {}

public class SinCalculation {
    public static void main(String[] args) {
        double x = 0.5; // Angle in radians
        int n = 10; // Number of terms in the series
        double result = sin(x, n);
        System.out.println("sin(" + x + ") ≈ " + result);
    }

    public static double sin(double x, int n) {
        double result = 0;
        for (int i = 0; i < n; i++) {
            int exponent = 2 * i + 1;
            double term = Math.pow(-1, i) * Math.pow(x, exponent) / factorial(exponent);
            result += term;
        }
        return result;
    }

    public static double factorial(int n) {
        if (n == 0) {
            return 1;
        }
        double fact = 1.0;
        for (int i = 1; i <= n; i++) {
            fact *= i;
        }
        return fact;
    }
}

import java.util.Scanner;

public class CustomLog {

    // Define a method to calculate the natural logarithm
    public static double customLog(double x, int n) {
        if (x <= 0) {
            throw new IllegalArgumentException("Input must be greater than 0.");
        }

        if (x == 1) {
            return 0.0;
        }

        if (x < 1) {
            x = 1 / x;
            n = -n;
        }

        double result = 0.0;
        for (int i = 1; i <= n; i++) {
            result += Math.pow((x - 1) / x, i) / i;
        }

        return result;
    }

    public static void main(String[] args) {
        double x = 15.0; // The value for which you want to calculate the logarithm
        int n = 100; // Number of iterations for accuracy

        double result = customLog(x, n);
        System.out.println("Log of " + x + " is approximately: " + result);
    }
}

import java.util.Scanner;

public class TangentCalculator {
    public static void main(String[] args) {
        double angleInRadians = 0.5; // Angle in radians
        int n = 10; // Number of terms in the series
        double result = tangent(angleInRadians, n);
        System.out.println("tan(" + angleInRadians + ") ≈ " + result);
    }

    public static double tangent(double x, int n) {
        double result = 0;
        double term = x;
        for (int i = 1; i < n; i++) {
            result += term;
            term *= -(x * x) / ((2 * i + 1) * (2 * i + 2));
        }
        return result;
    }
}

public class ArcsineCalculator {
    public static void main(String[] args) {
        double x = 0.5; // Input value between -1 and 1
        int n = 10; // Number of terms in the series
        double result = arcsin(x, n);
        System.out.println("arcsin(" + x + ") ≈ " + result + " radians");
    }

    public static double arcsin(double x, int n) {
        if (x < -1 || x > 1) {
            throw new IllegalArgumentException("Input value must be between -1 and 1.");
        }

        double result = x;
        double term = x;
        for (int i = 1; i < n; i++) {
            term *= (x * x) * (2 * i - 1) / (2 * i * (2 * i + 1));
            result += term;
        }

        return result;
    }
}

public class ArccosineCalculator {
    public static void main(String[] args) {
        double x = 0.5; // Input value between -1 and 1
        int n = 10; // Number of terms in the series
        double result = arccos(x, n);
        System.out.println("arccos(" + x + ") ≈ " + result + " radians");
    }

    public static double arccos(double x, int n) {
        if (x < -1 || x > 1) {
            throw new IllegalArgumentException("Input value must be between -1 and 1.");
        }

        double result = Math.PI / 2 - x;
        double term = x;
        for (int i = 1; i < n; i++) {
            term *= (x * x) * (2 * i - 1) / (2 * i * (2 * i + 1));
            result -= term;
        }

        return result;
    }
}

public class ArctangentCalculator {
    public static void main(String[] args) {
        double x = 0.5; // Input value
        int n = 10; // Number of terms in the series
        double result = arctan(x, n);
        System.out.println("arctan(" + x + ") ≈ " + result + " radians");
    }

    public static double arctan(double x, int n) {
        if (x < -1 || x > 1) {
            throw new IllegalArgumentException("Input value should be within the range of -1 to 1.");
        }

        double result = x;
        double term = x;
        double xSquared = x * x;
        for (int i = 1; i < n; i++) {
            term *= -xSquared;
            result += term / (2 * i + 1);
        }

        return result;
    }
}

public class ExponentialCalculator {
    public static void main(String[] args) {
        double x = 2.0; // Input value
        int n = 10; // Number of terms in the series
        double result = exponential(x, n);
        System.out.println("exp(" + x + ") ≈ " + result);
    }

    public static double exponential(double x, int n) {
        double result = 1.0;
        double term = 1.0;

        for (int i = 1; i < n; i++) {
            term *= x / i;
            result += term;
        }

        return result;
    }
}

public class NaturalLogCalculator {
    public static void main(String[] args) {
        double x = 0.5; // Input value (should be greater than 0)
        int n = 10; // Number of terms in the series
        double result = naturalLog(1 + x, n); // Use ln(1 + x) for the series
        System.out.println("ln(1 + " + x + ") ≈ " + result);
    }

    public static double naturalLog(double x, int n) {
        if (x <= 0) {
            throw new IllegalArgumentException("Input value should be greater than 0.");
        }

        double result = 0.0;
        double term = x - 1; // ln(1 + x) starts with x - 1
        double sign = 1.0;

        for (int i = 1; i <= n; i++) {
            result += sign * term / i;
            term *= -(x - 1); // Update the term
            sign = -sign; // Toggle the sign
        }

        return result;
    }
}

public class Log10Calculator {
    public static void main(String[] args) {
        double x = 100.0; // Input value (should be greater than 0)
        int n = 10; // Number of terms in the series
        double result = log10(x, n);
        System.out.println("log10(" + x + ") ≈ " + result);
    }

    public static double log10(double x, int n) {
        if (x <= 0) {
            throw new IllegalArgumentException("Input value should be greater than 0.");
        }

        double lnX = naturalLog(x, n);
        double ln10 = naturalLog(10, n);

        return lnX / ln10;
    }

    public static double naturalLog(double x, int n) {
        if (x <= 0) {
            throw new IllegalArgumentException("Input value should be greater than 0.");
        }

        double result = 0.0;
        double term = (x - 1) / (x + 1);

        for (int i = 1; i <= n; i++) {
            double currentTerm = Math.pow(term, 2 * i - 1) / (2 * i - 1);
            result += currentTerm;
        }

        return 2 * result;
    }
}

public class ExponentiationCalculator {
    public static void main(String[] args) {
        double base = 2.0; // Base
        int exponent = 3; // Exponent
        double result = power(base, exponent);
        System.out.println(base + " raised to the power of " + exponent + " is: " + result);
    }

    public static double power(double base, int exponent) {
        if (exponent < 0) {
            throw new IllegalArgumentException("Exponent should be a non-negative integer.");
        }
        
        double result = 1.0;

        for (int i = 0; i < exponent; i++) {
            result *= base;
        }

        return result;
    }
}


public class AbsoluteValueCalculator {
    public static void main(String[] args) {
        double number = -5.25; // Input number
        double absValue = absoluteValue(number);
        System.out.println("The absolute value of " + number + " is: " + absValue);
    }

    public static double absoluteValue(double number) {
        if (number < 0) {
            return -number;
        } else {
            return number;
        }
    }
}

public class CeilingFunctionCalculator {
    public static void main(String[] args) {
        double number = 5.25; // Input number
        double ceilValue = ceilingFunction(number);
        System.out.println("The ceiling of " + number + " is: " + ceilValue);
    }

    public static double ceilingFunction(double number) {
        if (number == Math.floor(number)) {
            return number; // It's already an integer
        } else if (number > 0) {
            return Math.floor(number) + 1;
        } else {
            return Math.floor(number);
        }
    }
}

public class FloorFunctionCalculator {
    public static void main(String[] args) {
        double number = 5.75; // Input number
        double floorValue = floorFunction(number);
        System.out.println("The floor of " + number + " is: " + floorValue);
    }

    public static double floorFunction(double number) {
        return Math.floor(number);
    }
}

public class RoundingFunctionCalculator {
    public static void main(String[] args) {
        double number = 5.75; // Input number
        long roundedValue = roundingFunction(number);
        System.out.println("The rounded value of " + number + " is: " + roundedValue);
    }

    public static long roundingFunction(double number) {
        return Math.round(number);
    }
}
