import os
from signal import SIGKILL

with open('pids.txt', 'r') as pid_file:
    for pid in pid_file.read().splitlines():
        os.kill(int(pid), SIGKILL)