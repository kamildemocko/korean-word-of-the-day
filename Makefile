BUILD_BINARY=.\bin\korean-word-of-a-day.exe
GO_BINARY=korean-word-of-a-day.exe

build:
	@echo start build
	@go build -o ${BUILD_BINARY} .\cmd\app
	@echo - done

start: build
	@echo start run
	@start /B ${BUILD_BINARY} &
	@echo - done

stop:
	@echo killing proces
	@taskkill /IM ${GO_BINARY} /F
	@echo - done
