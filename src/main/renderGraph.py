import pygame, simplejson

pygame.init()
file=open("drawBuffer.json","r")
screen = pygame.display.set_mode((640,480))
pygame.display.update()
while(True):
	screen.fill([225,225,225])
	file.seek(0)
	contents=file.read()
	print(contents)
	if contents!="" and "end" in contents:
		sep='end'
		contents=contents.split('end',1)[0].replace("end","")
		j=simplejson.loads(str(contents))
		for layer in range(len(j)):
			for node in range(len(j[layer])):
				pygame.draw.rect(screen,(int(j[layer][node]["RefInputSum"]*225),0,0),(layer*50,node*50,20,20),0)
				if j[layer][node]["Weights"]!=None:
					for nextnode in range(len(j[layer][node]["Weights"])):
						print("it's happening")
						color=(0,0,0)
						color2=(0,0,0)
						# if j[layer][node]["WeightsChange"][nextnode]<0:
						# 	color=(0,0,abs(int(j[layer][node]["WeightsChange"][nextnode]*225)))
						# else:
						# 	color=(int(j[layer][node]["WeightsChange"][nextnode]*225),0,0)
						# if j[layer][node]["Weights"][nextnode]<0:
						# 	color2=(0,0,abs(int(j[layer][node]["Weights"][nextnode])))
						# else:
						# 	color2=(int(j[layer][node]["Weights"][nextnode]),0,0)
						# # pygame.draw.line(screen,color,(layer*50,node*50),((layer+1)*50,nextnode*50),4)
						# pygame.draw.line(screen,color2,(layer*50,(node)*50+20),((layer+1)*50,nextnode*50),1)
				pygame.draw.rect(screen,(int(j[layer][node]["RefInputSum"]*225),0,0),(layer*50,node*50,20,20),0)
		pygame.display.flip()
