# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from proto.cosmetic import cosmetic_pb2 as proto_dot_cosmetic_dot_cosmetic__pb2


class CosmeticServiceStub(object):
    """The greeting service definition.
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.deleteCosmetic = channel.unary_unary(
                '/cosmetic.CosmeticService/deleteCosmetic',
                request_serializer=proto_dot_cosmetic_dot_cosmetic__pb2.deleteCosmeticRequest.SerializeToString,
                response_deserializer=proto_dot_cosmetic_dot_cosmetic__pb2.deleteCosmeticReply.FromString,
                )
        self.updateCosmetic = channel.unary_unary(
                '/cosmetic.CosmeticService/updateCosmetic',
                request_serializer=proto_dot_cosmetic_dot_cosmetic__pb2.updateCosmeticRequest.SerializeToString,
                response_deserializer=proto_dot_cosmetic_dot_cosmetic__pb2.updateCosmeticReply.FromString,
                )
        self.createCosmetic = channel.unary_unary(
                '/cosmetic.CosmeticService/createCosmetic',
                request_serializer=proto_dot_cosmetic_dot_cosmetic__pb2.createCosmeticRequest.SerializeToString,
                response_deserializer=proto_dot_cosmetic_dot_cosmetic__pb2.createCosmeticResponse.FromString,
                )
        self.ListCosmetics = channel.unary_unary(
                '/cosmetic.CosmeticService/ListCosmetics',
                request_serializer=proto_dot_cosmetic_dot_cosmetic__pb2.ListCosmeticsRequest.SerializeToString,
                response_deserializer=proto_dot_cosmetic_dot_cosmetic__pb2.ListCosmeticsResponse.FromString,
                )


class CosmeticServiceServicer(object):
    """The greeting service definition.
    """

    def deleteCosmetic(self, request, context):
        """Sends a greeting
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def updateCosmetic(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def createCosmetic(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ListCosmetics(self, request, context):
        """Sends another greeting
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_CosmeticServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'deleteCosmetic': grpc.unary_unary_rpc_method_handler(
                    servicer.deleteCosmetic,
                    request_deserializer=proto_dot_cosmetic_dot_cosmetic__pb2.deleteCosmeticRequest.FromString,
                    response_serializer=proto_dot_cosmetic_dot_cosmetic__pb2.deleteCosmeticReply.SerializeToString,
            ),
            'updateCosmetic': grpc.unary_unary_rpc_method_handler(
                    servicer.updateCosmetic,
                    request_deserializer=proto_dot_cosmetic_dot_cosmetic__pb2.updateCosmeticRequest.FromString,
                    response_serializer=proto_dot_cosmetic_dot_cosmetic__pb2.updateCosmeticReply.SerializeToString,
            ),
            'createCosmetic': grpc.unary_unary_rpc_method_handler(
                    servicer.createCosmetic,
                    request_deserializer=proto_dot_cosmetic_dot_cosmetic__pb2.createCosmeticRequest.FromString,
                    response_serializer=proto_dot_cosmetic_dot_cosmetic__pb2.createCosmeticResponse.SerializeToString,
            ),
            'ListCosmetics': grpc.unary_unary_rpc_method_handler(
                    servicer.ListCosmetics,
                    request_deserializer=proto_dot_cosmetic_dot_cosmetic__pb2.ListCosmeticsRequest.FromString,
                    response_serializer=proto_dot_cosmetic_dot_cosmetic__pb2.ListCosmeticsResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'cosmetic.CosmeticService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class CosmeticService(object):
    """The greeting service definition.
    """

    @staticmethod
    def deleteCosmetic(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/cosmetic.CosmeticService/deleteCosmetic',
            proto_dot_cosmetic_dot_cosmetic__pb2.deleteCosmeticRequest.SerializeToString,
            proto_dot_cosmetic_dot_cosmetic__pb2.deleteCosmeticReply.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def updateCosmetic(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/cosmetic.CosmeticService/updateCosmetic',
            proto_dot_cosmetic_dot_cosmetic__pb2.updateCosmeticRequest.SerializeToString,
            proto_dot_cosmetic_dot_cosmetic__pb2.updateCosmeticReply.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def createCosmetic(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/cosmetic.CosmeticService/createCosmetic',
            proto_dot_cosmetic_dot_cosmetic__pb2.createCosmeticRequest.SerializeToString,
            proto_dot_cosmetic_dot_cosmetic__pb2.createCosmeticResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ListCosmetics(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/cosmetic.CosmeticService/ListCosmetics',
            proto_dot_cosmetic_dot_cosmetic__pb2.ListCosmeticsRequest.SerializeToString,
            proto_dot_cosmetic_dot_cosmetic__pb2.ListCosmeticsResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
