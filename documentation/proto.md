# Working with Protocol Buffers

[Buf](https://buf.build/docs/introduction) is used for protocol buffer compilation.

## Generating the Protobuf and gRPC Stubs

Run the following command:

```bash
buf generate proto
```

## Sharing Schemas in a Monorepo

I have a separate project which describes [sharing Proto schemas across projects in a monorepo](https://github.com/wcygan/buf-polyglot-example). 

Essentially, the schemas reside in a folder, `/generated`, which contains the generated code for each language. Each project then pulls in the schemas from this folder.