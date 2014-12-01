import socket
from uuid import getnode
from time import sleep

bcast_ip = "255.255.255.255"
bcast_port = 2050
bcast_info = (bcast_ip, bcast_port)

mac = getnode()

s = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
s.setsockopt(socket.SOL_SOCKET, socket.SO_BROADCAST, 1)

while(True):
    s.sendto("SolidXpress " + hex(mac), bcast_info)
    sleep(2)


print "done"
