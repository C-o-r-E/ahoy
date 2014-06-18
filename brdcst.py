import socket

bcast_ip = "255.255.255.255"
bcast_port = 2050
bcast_info = (bcast_ip, bcast_port)


print "attempting to send udp broadcast message"

s = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
s.setsockopt(socket.SOL_SOCKET, socket.SO_BROADCAST, 1)

s.sendto("this is a test", bcast_info)

print "done"
