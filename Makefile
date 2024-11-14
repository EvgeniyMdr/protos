# .PHONY для указания целей, которые не являются файлами
.PHONY: generate clean

# Цель для генерации структур Go и gRPC файлов из proto
generate:
	mkdir -p gen/go  # Создаем директорию для сгенерированных файлов
	protoc -I proto proto/comments/comments.proto --go_out=./gen/go --go-grpc_out=./gen/go --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative

# Очистка сгенерированных файлов
clean:
	rm -rf gen/go/*
