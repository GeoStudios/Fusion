package Math_Area;

// Embeded Imports
using "@/Math_Vars";

fn GetVolumeOfCubeF(float side) -> float { return side ** 3; }
fn GetVolumeOfCubeI(int side) -> int { return side ** 3; }

fn GetVolumeOfParallelPipedF(float length, float width, float height) -> float { return length * width * height; }
fn GetVolumeOfParallelPipedI(int length, int width, int height) -> int { return length * width * height; }

fn GetVolumeOfRectangularPrismF(float base float height) -> float { return base * height; }
fn GetVolumeOfRectangularPrismI(int base, int height) -> int { return base * height; }

fn GetVolumeOfCylinderF(float radius, float height) -> float { return ((radius ** 2) * Math_Vars.PI) * height; }
fn GetVolumeOfCylinderI(int radius, int height) -> int { return ((radius ** 2) * Math_Vars.PI) * height; }

fn GetVolumeOfConeF(float base float height) -> float { return 1/3 * (base * height); }
fn GetVolumeOfConeI(int base, int height) -> int { return 1/3 * (base * height); }

fn GetVolumeOfPyramidF(float base float height) -> float { return 1/3 * (base * height); }
fn GetVolumeOfPyramidI(int base, int height) -> int { return 1/3 * (base * height); }


fn GetVolumeOfSphereF(float radius, float height) -> float { return ((radius ** 2) * Math_Vars.PI) * 4/3; }
fn GetVolumeOfSphereI(int radius, int height) -> int { return ((radius ** 2) * Math_Vars.PI) * 4/3; }