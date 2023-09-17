GOLANG = go
LDFLAGS = "-s -w"
all:
	sed -i "s/Commit_id \= \"\"/Commit_id \= \"`git log --oneline|head -1|cut -d " " -f 1`\"/" main.go
	$(GOLANG) mod download
	$(GOLANG) build -ldflags $(LDFLAGS) -o yoshinon main.go
format:
	gofmt -s -d -w `find|grep "\.go"`
