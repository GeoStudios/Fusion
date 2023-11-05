import os

os.system("clear")
os.chdir("src")
print("------------Compile-Binary-------------")
print("./bin/fusion.exe")
os.system("go get .")
os.system("go build -o ../bin/fusion.exe")
os.chdir("../bin")
# os.system("clear")
print("--------------Run-Binary---------------")
# print("Running bytecode.js in SimmerJs")
os.system("fusion.exe ../tests/math_tes.f -i -o")
os.chdir("../")
