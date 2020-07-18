default: kokodayo.bin

.FORCE :

kokodayo.bin : .FORCE
	go build -o $@

run : kokodayo.bin
	./kokodayo.bin

test : kokodayo.bin
	go test -v main/kokodayo
	dredd
	@echo "PASS"

ci :
	git add .
	circleci local execute
