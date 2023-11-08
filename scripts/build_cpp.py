import os

cpp_files, header_files, object_files = [], [], []

for root, paths, files in os.walk("src/cpp"):
    for file in files:
        if file.endswith(".cc") or file.endswith(".cpp") or file.endswith(".c"):
            cpp_files.append(os.path.abspath(os.path.join(root, file)))
            print(os.path.abspath(os.path.join(root, file)))
        if file.endswith(".hh") or file.endswith(".hpp") or file.endswith(".h"):
            header_files.append(os.path.abspath(os.path.join(root, file)))
            print(os.path.abspath(os.path.join(root, file)))

for file in cpp_files: 
    header = ""
    name = file.replace(".cc", "").replace(".cpp", "").replace(".c", "")
    for hFile in header_files:
        if hFile.__contains__(name):
            header = hFile
            break
    print("g++ -c -g "+file+" "+header+" -o "+name+".o")
    os.system("g++ -c -g "+file+" "+header+" -o "+name+".o")

for root, paths, files in os.walk("cpp_src"):
    for file in files:
        if file.endswith(".o"):
            print(os.path.abspath(os.path.join(root, file)))
            object_files.append(os.path.abspath(os.path.join(root, file)))

os.system("g++ -g "+" ".join(object_files)+" -o bin/cpp.exe")

for x in object_files:
    os.remove(x)

os.chdir("bin")
