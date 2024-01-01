// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var proto_service_pb = require('../proto/service_pb.js');

function serialize_app_service_BacktestRequest(arg) {
  if (!(arg instanceof proto_service_pb.BacktestRequest)) {
    throw new Error('Expected argument of type app_service.BacktestRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_app_service_BacktestRequest(buffer_arg) {
  return proto_service_pb.BacktestRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_app_service_BacktestResponse(arg) {
  if (!(arg instanceof proto_service_pb.BacktestResponse)) {
    throw new Error('Expected argument of type app_service.BacktestResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_app_service_BacktestResponse(buffer_arg) {
  return proto_service_pb.BacktestResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_app_service_PairwiseAnalysisRequest(arg) {
  if (!(arg instanceof proto_service_pb.PairwiseAnalysisRequest)) {
    throw new Error('Expected argument of type app_service.PairwiseAnalysisRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_app_service_PairwiseAnalysisRequest(buffer_arg) {
  return proto_service_pb.PairwiseAnalysisRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_app_service_PairwiseAnalysisResponse(arg) {
  if (!(arg instanceof proto_service_pb.PairwiseAnalysisResponse)) {
    throw new Error('Expected argument of type app_service.PairwiseAnalysisResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_app_service_PairwiseAnalysisResponse(buffer_arg) {
  return proto_service_pb.PairwiseAnalysisResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var AppServiceService = exports.AppServiceService = {
  pairwiseAnalysis: {
    path: '/app_service.AppService/PairwiseAnalysis',
    requestStream: false,
    responseStream: false,
    requestType: proto_service_pb.PairwiseAnalysisRequest,
    responseType: proto_service_pb.PairwiseAnalysisResponse,
    requestSerialize: serialize_app_service_PairwiseAnalysisRequest,
    requestDeserialize: deserialize_app_service_PairwiseAnalysisRequest,
    responseSerialize: serialize_app_service_PairwiseAnalysisResponse,
    responseDeserialize: deserialize_app_service_PairwiseAnalysisResponse,
  },
  backtest: {
    path: '/app_service.AppService/Backtest',
    requestStream: false,
    responseStream: false,
    requestType: proto_service_pb.BacktestRequest,
    responseType: proto_service_pb.BacktestResponse,
    requestSerialize: serialize_app_service_BacktestRequest,
    requestDeserialize: deserialize_app_service_BacktestRequest,
    responseSerialize: serialize_app_service_BacktestResponse,
    responseDeserialize: deserialize_app_service_BacktestResponse,
  },
};

exports.AppServiceClient = grpc.makeGenericClientConstructor(AppServiceService);
