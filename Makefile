build: gene v56error v57error v80error

gene:
	go build -o gene github.com/ytnobody/gomysqlerror/generator

v56error:
	mkdir -p v56error
	./gene 5.6 > v56error/v56error.go

v57error:
	mkdir -p v57error
	./gene 5.7 > v57error/v57error.go

v80error:
	mkdir -p v80error
	./gene 8.0 > v80error/v80error.go

clean:
	rm -fr ./gene ./v*error