# inventory

1. command to generate protobuf files
#protoc --go_out=. --go-grpc_out=. inventorymanagement/inventory.proto

 2. initilizing go mod

 #go mod init example.com/go-inventory-grpc




 3. generate protobuf
 protoc ./internal/proto/inventory.proto --go_out=./internal/endpoint --go_opt=paths=source_relative --go-grpc_out=./internal/endpoint --go-grpc_opt=paths=source_relative ./internal/proto/inventory.proto


 4. Now, let’s get the module/ library needed for this project.
#go get entgo.io/ent/cmd/ent
#go get -u github.com/gorilla/mux

5. Initialize the Entity
#go run entgo.io/ent/cmd/ent init {{User}}
#User is table name

6. after adding field in the ent file use this command to generate ent again with added column or field

#go generate ./ent


protoc -I={{.INPUT_PROTOS}}

--go_out={{.OUTPUT_DIRECTORY}}

mkdir -p {{.OUTPUT_DIRECTORY}}
...

generate: golang:
        summary: Compile and generate protocol buffer golang definitions
    silent: true
    vars:
        INPUT_PROTOS: "((default "." . INPUT PROTOS})•
        UTPUT _DIRECTORY: "((default "." •OUTPUT _DIRECTORY))
cmds:
        - echo Compiling protobufs and generating definitions in Go...
        - echo - Input directory ((. INPUT_PROTOS))
        - echo - output directory ((.OUTPUT DIRECTORY))
        - mkdir - ((.OUTPUT_DIRECTORY))
            for f in ({. INPUT PROTOS)}/*. proto?
                echo - generating: $f
                protoc -I=({. INPUT _PROTOS)) fI
                    --gO_Out=((.OUTPUT DIRECTORY))
                    --go_opI-pathsasource_relative\
                    --g0-grpc_out=((. OUTPUT _DIRECTORY))
                    --go-grpc_opt=paths=source_relative\
                    --experimentaL_allow_proto3_optional
            done

generate: golang:
        summary: Compile and generate protocol buffer golang definitions
    silent: true
    vars:
        INPUT_PROTOS: "((default "." .inventory/inventorymanagement/internal/proto})•
        UTPUT _DIRECTORY: "((default "." •inventory/inventorymanagement/internal/endpoint))
cmds:
        - echo Compiling protobufs and generating definitions in Go...
        - echo - Input directory ((.inventory/inventorymanagement/internal/proto))
        - echo - output directory ((.inventory/inventorymanagement/internal/endpoint))
        - mkdir - ((.inventory/inventorymanagement/internal/endpoint))
            for f in ({. inventory/inventorymanagement/internal/proto)}/*. proto?
                echo - generating: $f
                protoc -I=./inventory/inventorymanagement/internal/proto/inventory.proto
                    --gO_Out=./inventory/inventorymanagement/internal/endpoint
                    --go_opI-pathsasource_relative\
                    --g0-grpc_out=./inventory/inventorymanagement/internal/endpoint
                    --go-grpc_opt=paths=source_relative\
                    --experimentaL_allow_proto3_optional
            done



protoc ./internal/proto/inventory.proto --go_out=./internal/endpoint --go_opt=paths=source_relative --go-grpc_out=./internal/endpoint --go-grpc_opt=paths=source_relative 

#protoc -I=/Users/hisilamanandhar/Documents/study/inventory/inventorymanagement/internal/proto/inventory.proto --go_out=/Users/hisilamanandhar/Documents/study/inventory/inventorymanagement/internal/endpoint --go-grpc_out=/Users/hisilamanandhar/Documents/study/inventory/inventorymanagement/internal/endpoint /Users/hisilamanandhar/Documents/study/inventory/inventorymanagement/internal/proto/inventory.proto

#protoc -I=/Users/hisilamanandhar/Documents/study/inventory/inventorymanagement/internal/proto/inventory.proto --go_out=./endpoint --go-grpc_out=./endpoint /Users/hisilamanandhar/Documents/study/inventory/inventorymanagement/internal/proto/inventory.proto

protoc --go_out=. --go-grpc_out=. inventorymanagement/inventory.proto




protoc -I=/Users/hisilamanandhar/Documents/study/inventory/inventorymanagement/internal/proto/inventory.proto --go_out=/Users/hisilamanandhar/Documents/study/inventory/inventorymanagement/internal/endpoint --go-grpc_out=/Users/hisilamanandhar/Documents/study/inventory/inventorymanagement/internal/endpoint /Users/hisilamanandhar/Documents/study/inventory/inventorymanagement/internal/proto/inventory.proto