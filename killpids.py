import os
from signal import SIGKILL
counter = 0
with open('pids.txt', 'r') as pid_file:
    for pid in pid_file.read().splitlines():
        try:
            os.kill(int(pid), SIGKILL)
            counter+=1
        except:
            continue
print(f"killed {counter} processes")