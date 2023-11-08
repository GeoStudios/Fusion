package StrongMath_Wrapper;

// Native Imports
using "#/StrongMath" as "NMath";

fn Sin(int a) -> float { return NMath.Sin(a); }
fn Sinh(int a) -> float { return NMath.Sinh(a); }
fn ASin(int a) -> float { return NMath.ASin(a); }
fn ASinh(int a) -> float { return NMath.ASinh(a); }

fn Cosine(int a) -> float { return NMath.Cosine(a); }
fn Cosineh(int a) -> float { return NMath.Cosineh(a); }
fn ACosine(int a) -> float { return NMath.ACosine(a); }
fn ACosineh(int a) -> float { return NMath.ACosineh(a); }

fn Tan(int a) -> float { return NMath.Tan(a); }
fn Tanh(int a) -> float { return NMath.Tanh(a); }

fn ATan(int a) -> float { return NMath.ATan(a); }
fn ATanh(int a) -> float { return NMath.ATanh(a); }
fn ATan2(int a, int b) -> float { return NMath.ATan2(a, b); }