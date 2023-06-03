import os
import subprocess
import sys
import random


print( "Current PID:", os.getpid() )
programs = []
outfiles = []
errfiles = []
NUMSERVERS = int(sys.argv[1])
MODCONS = 10000
starting_port = 8081
ports = []
port_keys = []
for i in range(NUMSERVERS):
    port_keys.append((random.randint(1,10000), starting_port+i))
    outfiles.append(open(f"out/out{port_keys[i][1]}.txt", "w"))
    errfiles.append(open(f"err/err{port_keys[i][1]}.txt", "w"))
port_keys.sort()

with open('pids.txt', 'w') as pid_file, \
     open('ports.txt', 'w') as port_file, \
     open('serverIds.txt', 'w') as sid_file:
    pid_file.write(str(os.getpid()) + '\n')  
    for i in range(NUMSERVERS):
        cmd_args = [f"{port_keys[(i+j)%NUMSERVERS][0], port_keys[(i+j)%NUMSERVERS][1]}" for j in range(-1,2)]
        sub = subprocess.Popen(f"go run server/main.go {cmd_args}", shell=True, stdout=outfiles[i], stderr=errfiles[i])
        pid_file.write(str(sub.pid) + '\n')
        port_file.write(str(port_keys[i][1]) + '\n')
        sid_file.write(str(port_keys[i][0]) + '\n')
        programs.append(sub)
        
print()

for i in range(len(programs)):
    print(f"started server --- sid: {port_keys[i][0]} port: {port_keys[i][1]} pid: {programs[i].pid}")

print()
    
while True:
    print(">>", end=' ')
    cmd = input()
    if cmd == "kill":
        break

print("\n", "-"*50, "\n")
print("Initiate kill sequence ...\n")
for i in range(len(programs)):
    programs[i].kill()
    print(f"killing: {programs[i].pid}")



# ports hash and mod into 10,000, so that server id's dont colide

# request
#     hash key with same has function and mode same value 10k
#     server id must always be larger than key, unless you loop around in which case  
#         later implement th euse of predecessors to find the node faster for the unroll case
#     for all servers find the 


#     client will just send everything to arbitrary server, could be same everytime or not\
    
# server
# is it sus for a distributed hash table
# each node needs to store its predecessor and successor
# everything will be mod 10k rn and 