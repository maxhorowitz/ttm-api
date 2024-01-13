# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

import quant_pb2 as quant__pb2


class QuantServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.RollingOrdinaryLeastSquares = channel.stream_unary(
                '/quant_service.QuantService/RollingOrdinaryLeastSquares',
                request_serializer=quant__pb2.RollingOrdinaryLeastSquaresDatapoint.SerializeToString,
                response_deserializer=quant__pb2.RollingOrdinaryLeastSquaresResult.FromString,
                )


class QuantServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def RollingOrdinaryLeastSquares(self, request_iterator, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_QuantServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'RollingOrdinaryLeastSquares': grpc.stream_unary_rpc_method_handler(
                    servicer.RollingOrdinaryLeastSquares,
                    request_deserializer=quant__pb2.RollingOrdinaryLeastSquaresDatapoint.FromString,
                    response_serializer=quant__pb2.RollingOrdinaryLeastSquaresResult.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'quant_service.QuantService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class QuantService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def RollingOrdinaryLeastSquares(request_iterator,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.stream_unary(request_iterator, target, '/quant_service.QuantService/RollingOrdinaryLeastSquares',
            quant__pb2.RollingOrdinaryLeastSquaresDatapoint.SerializeToString,
            quant__pb2.RollingOrdinaryLeastSquaresResult.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
