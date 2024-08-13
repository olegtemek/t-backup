CMD_ROOT=cmd/main.go


run:
	go run $(CMD_ROOT) --name=myVaults --originalPath=~/Developer/t-backup/test-backup --driver=local --backupPath=~/Developer/t-backup --numberOfBackups=2
