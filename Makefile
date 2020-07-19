default: kokodayo.bin

.FORCE :

kokodayo.bin : .FORCE
	go build -o $@

run : kokodayo.bin
	. ./env.sh && \
		MYSQL_PORT=$${MYSQL_PORT} \
		MYSQL_USER=$${MYSQL_USER} \
		./kokodayo.bin

test : kokodayo.bin
	./db_init_test.sh
	go test -v main/kokodayo
	dredd
	@echo "PASS"

ci :
	git add .
	circleci local execute
