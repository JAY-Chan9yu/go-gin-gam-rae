# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/hello/hello.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from proto.google.api import annotations_pb2 as proto_dot_google_dot_api_dot_annotations__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x17proto/hello/hello.proto\x12\nhelloworld\x1a\"proto/google/api/annotations.proto\"\x1c\n\x0cHelloRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\"\x1d\n\nHelloReply\x12\x0f\n\x07message\x18\x01 \x01(\t2Z\n\x07Greeter\x12O\n\x08SayHello\x12\x18.helloworld.HelloRequest\x1a\x16.helloworld.HelloReply\"\x11\x82\xd3\xe4\x93\x02\x0b\"\x06/hello:\x01*BC\n\x1bio.grpc.examples.helloworldB\x0fHelloWorldProtoP\x01Z\x11/proto;helloworldb\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'proto.hello.hello_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\033io.grpc.examples.helloworldB\017HelloWorldProtoP\001Z\021/proto;helloworld'
  _GREETER.methods_by_name['SayHello']._options = None
  _GREETER.methods_by_name['SayHello']._serialized_options = b'\202\323\344\223\002\013\"\006/hello:\001*'
  _HELLOREQUEST._serialized_start=75
  _HELLOREQUEST._serialized_end=103
  _HELLOREPLY._serialized_start=105
  _HELLOREPLY._serialized_end=134
  _GREETER._serialized_start=136
  _GREETER._serialized_end=226
# @@protoc_insertion_point(module_scope)
