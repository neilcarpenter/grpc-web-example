const grpc = require('grpc');
const protoLoader = require('@grpc/proto-loader');

const PROTO_PATH = __dirname + '/../proto/reverse/reverse.proto';

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
const reverseService = protoDescriptor.grpc.web.reverse.ReverseService.service;

const reverse = (call, cb) => {
    // Ensure conformity with PROTOCOL-WEB, see
	// https://github.com/grpc/grpc/blob/master/doc/PROTOCOL-WEB.md#protocol-differences-vs-grpc-over-http2
    let metaData = new grpc.Metadata();
    metaData.add('accept', 'application/grpc-web-text');
    call.sendMetadata(metaData);

    if (!call.request.message) {
        cb({
           code: grpc.status.INVALID_ARGUMENT,
           message: 'Message not set'
        });
        return;
    }

    if (call.request.message === 'error') {
        cb({
           code: grpc.status.INTERNAL,
           message: 'Example error message'
        });
        return;
    }

    const message = call.request.message
        .split('')
        .reverse()
        .join('');

    return cb(null, { message });
};

const getServer = () => {
    var server = new grpc.Server();
    server.addService(reverseService, { reverse });
    return server;
};

const server = getServer();
server.bind('0.0.0.0:9092', grpc.ServerCredentials.createInsecure());
server.start();
