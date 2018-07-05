const PROTO_PATH = __dirname + '/../proto/reverse/reverse.proto';
const grpc = require('grpc');
const protoLoader = require('@grpc/proto-loader');

const packageDefinition = protoLoader.loadSync(
    PROTO_PATH,
    {
        keepCase: true,
        longs: String,
        enums: String,
        defaults: true,
        oneofs: true
    }
);
const protoDescriptor = grpc.loadPackageDefinition(packageDefinition);
const ReverseService = protoDescriptor.grpc.web.reverse.ReverseService;

module.exports = () => new ReverseService('localhost:9092', grpc.credentials.createInsecure());
