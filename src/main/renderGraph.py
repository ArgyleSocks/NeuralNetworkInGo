import pygame, simplejson, math,sys
xmargin=500
ymargin=50
xOff=0
yOff=0
oldMouseX=0
oldMouseY=0
sensitivity=-1
# margin=50
squareSize=10
pygame.init()
file=open("drawBuffer.json","r")
screen = pygame.display.set_mode((640,1200),pygame.RESIZABLE)
pygame.display.update()
first=True
while(True):
	for et in pygame.event.get():
		if et.type == pygame.KEYDOWN:
			if (et.key == pygame.K_ESCAPE) or (et.type == pygame.QUIT):
				print "In Here"
				pygame.quit()
				sys.exit()
	pos=pygame.mouse.get_pos()
	if pygame.mouse.get_pressed()[0]:
		motionx=oldMouseX-pos[0]
		motiony=oldMouseY-pos[1]
		xOff+=motionx*sensitivity
		yOff+=motiony*sensitivity
	oldMouseX=pos[0]
	oldMouseY=pos[1]
	pygame.event.pump()
	screen.fill([225,225,225])
	file.seek(0)
	contents=file.read()
	print(contents)
	if contents!="" and "end" in contents:
		sep='end'
		contents=contents.split('end',1)[0].replace("end","")
		j=simplejson.loads(str(contents))
		if first==True:
			first=False
			pygame.display.set_mode((len(j)*(xmargin+squareSize),len(j[0])*(ymargin+squareSize)),pygame.RESIZABLE)
		for layer in range(len(j)):
			for node in range(len(j[layer])):
				# pygame.draw.rect(screen,(int(j[layer][node]["RefInputSum"]*225),0,0),(layer*50,node*50,20,20),0)
				if j[layer][node]["Weights"]!=None:
					for nextnode in range(len(j[layer][node]["Weights"])):
						print("it's happening")
						color=0
						color2=0
						# if j[layer][node]["WeightsChange"][nextnode]<0:
						# 	color=abs(int(j[layer][node]["WeightsChange"][nextnode]*225))
						# else:
						# 	color=int(j[layer][node]["WeightsChange"][nextnode]*225)
						if j[layer][node]["Weights"][nextnode]<0:
							color2=abs(int(j[layer][node]["Weights"][nextnode]))
						else:
							color2=int(j[layer][node]["Weights"][nextnode])
						# sigTransform1=1/math.pow(math.e,-color)
						sigTransform2=1/math.pow(math.e,-color2)*50
						# pygame.draw.line(screen,sigTransform1,(layer*50,node*50),((layer+1)*50,nextnode*50),4)
						pygame.draw.line(screen,(sigTransform2,0,0),(layer*(xmargin+squareSize)+squareSize+xOff,(node)*(ymargin+squareSize)+squareSize/2+yOff),((layer+1)*(xmargin+squareSize)+xOff,nextnode*(ymargin+squareSize)+squareSize/2+yOff),2)
				sigTransform=1/math.pow(math.e,-(j[layer][node]["RefInputSum"][2]))*50
				print(sigTransform)
				pygame.draw.rect(screen,(int(sigTransform),0,0),(layer*(xmargin+squareSize)+xOff,node*(ymargin+squareSize)+yOff,squareSize,squareSize),0)
		pygame.display.flip()
