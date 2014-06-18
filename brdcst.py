import socket

bcast_ip = "255.255.255.255"
bcast_port = 2050
bcast_info = (bcast_ip, bcast_port)


print "attempting to send udp broadcast message"

s = socket.socket(socket.af_INET, socket.SOCK_DGRAM)

s.sendto("this is a test", bcast_info)

print "done"
