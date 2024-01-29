GOLANG = go
LDFLAGS = "-s -w -X main.Commit_id=`git log --oneline|head -1|cut -d " " -f 1`"
all:
	$(GOLANG) mod download
	$(GOLANG) build -ldflags $(LDFLAGS) -o yoshinon main.go
format:
	gofmt -s -d -w `find|grep "\.go"`
