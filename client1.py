#!/usr/bin/python3

import socket
import sys

server_address = (sys.argv[1], int(sys.argv[2]))
skt = socket.socket(family=socket.AF_INET, type=socket.SOCK_DGRAM)
skt.sendto(str.encode(sys.argv[3]), server_address)

buffer_size = 1024
reply = skt.recvfrom(buffer_size)
skt.close()

print("Reply from server: {}".format(reply[0]))

