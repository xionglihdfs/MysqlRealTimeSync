echo "Start the release program."

echo "Delete old program files."
del main

echo Compile the Linux version of the 64-bit program.
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -v -a -o main

echo "Release program complete."
