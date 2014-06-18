import socket

local_ip = "127.0.0.1"
bcast_port = 2050
sock_info = (local_ip, bcast port)

s = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
s.bind(sock_info)

print "listening for connections..."

while True:
    data, addr = s.recvfrom(1024)
    print "%s: [%s]" % addr, data
