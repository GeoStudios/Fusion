package Math_Area;

// Embeded Imports
using "@/Math_Vars";

fn GetAreaOfSquareF(float length) -> float { return length ** 2; }
fn GetAreaOfSquareI(int length) -> int { return length ** 2; }

fn GetAreaOfRectangleF(float width, float height) -> float { return width * height; }
fn GetAreaOfRectangleI(int width, int height) -> int { return width * height; }

fn GetAreaOfTriangleF(float base, float height) -> float { return (base * height) / 2; }
fn GetAreaOfTriangleI(int base, int height) -> int { return (base * height) / 2; }

fn GetAreaOfRhombusF(float LargeSide, float SmallSide, float height) -> float { return ((LargeDiagonal * SmallDiagonal) / 2) * height; }
fn GetAreaOfRhombusI(int LargeSide, int SmallSide, int height) -> int { return ((LargeDiagonal * SmallDiagonal) / 2) * height; }

fn GetAreaOfRegularPolygonF(float Perimeter, float Apothem) -> float { return (Perimeter / 2) * Apothem; }
fn GetAreaOfRegularPolygonI(int Perimeter, int Apothem) -> int { return (Perimeter / 2) * Apothem; }

fn GetAreaOfCircleF(float Radius) -> float { return (Radius ** 2) * Math_Vars.PI; }
fn GetAreaOfCircleI(int Radius) -> int { return (Radius ** 2) * Math_Vars.PI; }

fn GetAreaOfConeLateralF(float Radius, float SlantHeight) -> float { return (Radius * Math_Vars.PI) * SlantHeight; }
fn GetAreaOfConeLateralI(float Radius, int SlantHeight) -> int { return (Radius * Math_Vars.PI) * SlantHeight; }

fn GetAreaOfSphereSurfaceF(float Radius, float SlantHeight) -> float { return ((Radius ** 2) * Math_Vars.PI) * 4; }
fn GetAreaOfSphereSurfaceI(float Radius, int SlantHeight) -> int { return ((Radius ** 2) * Math_Vars.PI) * 4; }