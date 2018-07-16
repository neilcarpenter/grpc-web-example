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

const doReverse = text =>
    text
        .split('')
        .reverse()
        .join('');

const sendEnvoyMetadata = call => {
    // Ensure conformity with PROTOCOL-WEB, see
	// https://github.com/grpc/grpc/blob/master/doc/PROTOCOL-WEB.md#protocol-differences-vs-grpc-over-http2
    let metaData = new grpc.Metadata();
    metaData.add('accept', 'application/grpc-web-text');
    call.sendMetadata(metaData);
}

const reverse = (call, cb) => {
    sendEnvoyMetadata(call);   

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

    const message = doReverse(call.request.message);

    return cb(null, { message });
};

const serverStreamingReverse = call => {
    sendEnvoyMetadata(call);

    if (call.request.message === 'error') {
        call.emit('error', {
           code: grpc.status.INTERNAL,
           message: 'Example error message #2'
        });
        return;
    }

    const message = doReverse(call.request.message);
    const messageResponses = [];

    for (let i = 0; i < call.request.message_count; i++) {
        messageResponses.push(new Promise(resolve => {
            setTimeout(() => {
                call.write({ message });
                resolve();
            }, (call.request.message_interval * i));
        }))
    }

    Promise.all(messageResponses, call.end);
};

const getServer = () => {
    var server = new grpc.Server();
    server.addService(reverseService, { reverse, serverStreamingReverse });
    return server;
};

const server = getServer();
server.bind('0.0.0.0:9092', grpc.ServerCredentials.createInsecure());
server.start();
