# Python 3.3.3 and 2.7.6
# python helloworld_python.py

from threading import Thread
from threading import Lock

chillout = Lock()

sum = 0

def add():
	global sum
	for j in range(0,1000000):
		with chillout:
			sum += 1


def subtract():
	global sum
	for j in range(0,1000001):
		with chillout:
			sum -= 1


def main():
	addThread = Thread(target = add, args = (),)
	addThread.start()

	subtractThread = Thread(target = subtract, args = (),)
	subtractThread.start()
	
	addThread.join()
	subtractThread.join()

	print(sum)

main()

