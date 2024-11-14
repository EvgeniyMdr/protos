### generate proto code
```
protoc -I proto proto/comments/comments.proto --go_out=./gen/go --go-grpc_out=./gen/go --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative
```

```
protoc -I proto proto/posts/posts.proto --go_out=./gen/go --go-grpc_out=./gen/go --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative
```
 