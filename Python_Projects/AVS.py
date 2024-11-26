import os
import subprocess

def open_VScode(path):
    if os.path.exists(path) and os.path.isdir(path):
        try:
            subprocess.run([r"C:\Users\user\AppData\Local\Programs\Microsoft VS Code\Code.exe",path],check=True)#subprocess.run runs the terminal or cmd and check is given as true since if code command is not found the it can raise the CallProcessError
        except FileNotFoundError:
            print("Vs code is not available") #here file is not found is used since it means that the vscode file is missing;
        except subprocess.CalledProcessError as e:
            print("Vs code failed to open")
        except Exception as e:
            print("Exception is "+ str(e))
    else:
        print("Invalid path")

if __name__ == "__main__":
    open_VScode("C:\\Users\\user\\Desktop\\Github\\Go_Programs")
