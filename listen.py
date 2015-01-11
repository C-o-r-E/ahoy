import socket
import time

local_ip = "0.0.0.0"
bcast_port = 2050
sock_info = (local_ip, bcast_port)

s = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
s.bind(sock_info)

start = time.time()
print "listening for connections..."

while True:
    data, addr = s.recvfrom(1024)
    now = time.time()
    delta = int(now - start)
    print "t={} {}: [{}]".format(delta, addr, data)
