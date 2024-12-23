EXERCISE_NAME=quiz
BUILD_DIR=build

build_bin:
	@echo "Building exercise ${EXERCISE_NAME}"
	env GOOS=linux CGO_ENABLED=0 go build -o ${BUILD_DIR}/${EXERCISE_NAME} ./src
	@echo "Done!"
