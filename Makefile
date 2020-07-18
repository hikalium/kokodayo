default: kokodayo.bin

.FORCE :

kokodayo.bin : .FORCE
	go build -o $@

run : kokodayo.bin
	MYSQL_PORT=`docker port kokodayo-mysql 3306/tcp | cut -d ':' -f 2` \
			   ./kokodayo.bin

test : kokodayo.bin
	./db_init_test.sh
	go test -v main/kokodayo
	dredd
	@echo "PASS"

ci :
	git add .
	circleci local execute
