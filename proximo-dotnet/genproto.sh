
~/.nuget/packages/grpc.tools/1.0.0/tools/windows_x64/protoc --csharp_out=. --grpc_out=. --proto_path=../proto ../proto/*.proto --plugin=protoc-gen-grpc=packages/grpc.tools/1.0.0/tools/windows_x64/grpc_csharp_plugin.exe