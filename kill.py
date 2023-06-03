import os
import subprocess

pid_file = open('pids.txt', 'r')
pids = pid_file.readlines()

for pid in pids:
    os.system("kill " + pid)