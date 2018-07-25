"""
=================
Multiple subplots
=================

Simple demo with multiple subplots.
"""
import numpy as np
import matplotlib.pyplot as plt
import simplejson
from multiprocessing import Process, Pipe


length=0
arr=[0]*50
def readup(conn):
	file=open("costBuffer.json","r")
	global length
	print("running")
	while(True):
		#print("reading")
		file.seek(0)
		contents=file.read()
		print(contents)
		if contents!="" and "end" in contents:
			print("different story")
			sep='end'
			contents=contents.split('end',1)[0].replace("end","")
			j=simplejson.loads(str(contents))
			for i in j:
				length+=1
				arr[length-1]=i
			length=0
			print("sending")
			conn.send(arr)
			print("sent")
			print(arr)
parent,child=Pipe()
p=Process(target=readup, args=(child,))
print("sigh")
p.start()
print("hmm")
#p.join()
print("ok")
import numpy as np
import matplotlib.pyplot as plt
import matplotlib.animation as animation


def data_gen(t=0):
    global parent
    cnt=0
    while True:
        t += 1
        cnt+=1
        if t>=50:
            t=0
        print(arr,t)
        yield cnt, parent.recv()[t]


def init():
    ax.set_ylim(-1.1, 1.1)
    ax.set_xlim(0, 10)
    del xdata[:]
    del ydata[:]
    line.set_data(xdata, ydata)
    return line,

fig, ax = plt.subplots()
line, = ax.plot([], [], lw=2)
ax.grid()
xdata, ydata = [], []


def run(data):
    # update the data
    t, y = data
    xdata.append(t)
    ydata.append(y)
    xmin, xmax = ax.get_xlim()
    ax.set_xlim(0, 500)
    ax.set_ylim(0,20)
    #plt.cla()
    fig.canvas.draw()
    #plt.gcf().clear()
    line.set_data(xdata, ydata)

    return line,

ani = animation.FuncAnimation(fig, run, data_gen, blit=False, interval=10,
                              repeat=False, init_func=init)
plt.show()
