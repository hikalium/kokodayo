kokodayo.bin : main.go
	go build -o $@

run : kokodayo.bin
	./kokodayo.bin
