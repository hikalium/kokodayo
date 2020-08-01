default: kokodayo.bin

.FORCE :

kokodayo.bin : .FORCE
	go build -o $@

run : kokodayo.bin
	-killall kokodayo
	. ./env.sh && \
		MYSQL_PORT=$${MYSQL_PORT} \
		MYSQL_USER=$${MYSQL_USER} \
		./kokodayo.bin

run_prod : kokodayo.bin
	-killall kokodayo
	MYSQL_DATABASE=kokodayo_prod && \
	WWW_PORT=80 && \
		. ./env.sh && \
		MYSQL_PORT=$${MYSQL_PORT} \
		MYSQL_USER=$${MYSQL_USER} \
		MYSQL_DATABASE=$${MYSQL_DATABASE} \
		WWW_PORT=$${WWW_PORT} \
		./kokodayo.bin

test : kokodayo.bin
	-killall kokodayo
	./db_init_test.sh
	go test -v main/kokodayo
	dredd
	-killall kokodayo.bin
	@echo "PASS"

ci :
	git add .
	circleci local execute
