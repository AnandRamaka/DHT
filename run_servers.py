import os
import subprocess
import sys

print( "Current PID:", os.getpid() )
programs = []
outfiles = []
errfiles = []
NUMSERVERS = int(sys.argv[1])
for i in range(NUMSERVERS):
    outfiles.append(open(f"out/out{i}.txt", "w"))
for i in range(NUMSERVERS):
    errfiles.append(open(f"err/err{i}.txt", "w"))


with open('pids.txt', 'w') as f:
    f.write(str(os.getpid()) + '\n')
    for i in range(NUMSERVERS):
        sub = subprocess.Popen("go run server/main.go", shell=True, stdout=outfiles[i], stderr=errfiles[i])
        f.write(str(sub.pid) + '\n')
        programs.append(sub)
        
for whole in programs:
    print(whole.pid)


cmd = input()

print(cmd)
if cmd == 'kill':
    for i in range(len(programs)):
        programs[i].kill()
