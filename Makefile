#pitop   
pitop is a terminal based activity monitor for Raspberry Pi 4  

build-32 : 
	@echo " ---- BUILD 32 bits ---- "
	go build . 

build-64 : 
	@echo " ---- BUILD 64 bits ---- "
	go build . 

install :
	@echo " ---- INSTALL ---- "

