#!/usr/bin/env python3

import subprocess

def rename():
    """Rename all source files from pageXX.go to XX.go."""
    for i in range(1, 31):
        i = str(i)
        if int(i) < 10:
            subprocess.call(["git", "mv", "page0" + i + ".go", "0" + i + ".go"])
        else:
            subprocess.call(["git", "mv", "page" + i + ".go", i + ".go"])

if __name__ == "__main__":
    rename()
