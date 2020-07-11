default: kokodayo.bin

.FORCE :

kokodayo.bin : .FORCE
	go build -o $@

run : kokodayo.bin
	./kokodayo.bin

test : 
	go test -v main/kokodayo
