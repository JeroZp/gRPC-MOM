package com.miempresa.inventoryservice;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.42.1)",
    comments = "Source: inventoryservice.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class InventoryServiceGrpc {

  private InventoryServiceGrpc() {}

  public static final String SERVICE_NAME = "inventoryservice.InventoryService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<com.miempresa.inventoryservice.InventoryServiceProto.AddProductRequest,
      com.miempresa.inventoryservice.InventoryServiceProto.AddProductResponse> getAddProductMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "AddProduct",
      requestType = com.miempresa.inventoryservice.InventoryServiceProto.AddProductRequest.class,
      responseType = com.miempresa.inventoryservice.InventoryServiceProto.AddProductResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.miempresa.inventoryservice.InventoryServiceProto.AddProductRequest,
      com.miempresa.inventoryservice.InventoryServiceProto.AddProductResponse> getAddProductMethod() {
    io.grpc.MethodDescriptor<com.miempresa.inventoryservice.InventoryServiceProto.AddProductRequest, com.miempresa.inventoryservice.InventoryServiceProto.AddProductResponse> getAddProductMethod;
    if ((getAddProductMethod = InventoryServiceGrpc.getAddProductMethod) == null) {
      synchronized (InventoryServiceGrpc.class) {
        if ((getAddProductMethod = InventoryServiceGrpc.getAddProductMethod) == null) {
          InventoryServiceGrpc.getAddProductMethod = getAddProductMethod =
              io.grpc.MethodDescriptor.<com.miempresa.inventoryservice.InventoryServiceProto.AddProductRequest, com.miempresa.inventoryservice.InventoryServiceProto.AddProductResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "AddProduct"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.miempresa.inventoryservice.InventoryServiceProto.AddProductRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.miempresa.inventoryservice.InventoryServiceProto.AddProductResponse.getDefaultInstance()))
              .setSchemaDescriptor(new InventoryServiceMethodDescriptorSupplier("AddProduct"))
              .build();
        }
      }
    }
    return getAddProductMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.miempresa.inventoryservice.InventoryServiceProto.GetProductRequest,
      com.miempresa.inventoryservice.InventoryServiceProto.GetProductResponse> getGetProductMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetProduct",
      requestType = com.miempresa.inventoryservice.InventoryServiceProto.GetProductRequest.class,
      responseType = com.miempresa.inventoryservice.InventoryServiceProto.GetProductResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.miempresa.inventoryservice.InventoryServiceProto.GetProductRequest,
      com.miempresa.inventoryservice.InventoryServiceProto.GetProductResponse> getGetProductMethod() {
    io.grpc.MethodDescriptor<com.miempresa.inventoryservice.InventoryServiceProto.GetProductRequest, com.miempresa.inventoryservice.InventoryServiceProto.GetProductResponse> getGetProductMethod;
    if ((getGetProductMethod = InventoryServiceGrpc.getGetProductMethod) == null) {
      synchronized (InventoryServiceGrpc.class) {
        if ((getGetProductMethod = InventoryServiceGrpc.getGetProductMethod) == null) {
          InventoryServiceGrpc.getGetProductMethod = getGetProductMethod =
              io.grpc.MethodDescriptor.<com.miempresa.inventoryservice.InventoryServiceProto.GetProductRequest, com.miempresa.inventoryservice.InventoryServiceProto.GetProductResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetProduct"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.miempresa.inventoryservice.InventoryServiceProto.GetProductRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.miempresa.inventoryservice.InventoryServiceProto.GetProductResponse.getDefaultInstance()))
              .setSchemaDescriptor(new InventoryServiceMethodDescriptorSupplier("GetProduct"))
              .build();
        }
      }
    }
    return getGetProductMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.miempresa.inventoryservice.InventoryServiceProto.UpdateProductRequest,
      com.miempresa.inventoryservice.InventoryServiceProto.UpdateProductResponse> getUpdateProductMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "UpdateProduct",
      requestType = com.miempresa.inventoryservice.InventoryServiceProto.UpdateProductRequest.class,
      responseType = com.miempresa.inventoryservice.InventoryServiceProto.UpdateProductResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.miempresa.inventoryservice.InventoryServiceProto.UpdateProductRequest,
      com.miempresa.inventoryservice.InventoryServiceProto.UpdateProductResponse> getUpdateProductMethod() {
    io.grpc.MethodDescriptor<com.miempresa.inventoryservice.InventoryServiceProto.UpdateProductRequest, com.miempresa.inventoryservice.InventoryServiceProto.UpdateProductResponse> getUpdateProductMethod;
    if ((getUpdateProductMethod = InventoryServiceGrpc.getUpdateProductMethod) == null) {
      synchronized (InventoryServiceGrpc.class) {
        if ((getUpdateProductMethod = InventoryServiceGrpc.getUpdateProductMethod) == null) {
          InventoryServiceGrpc.getUpdateProductMethod = getUpdateProductMethod =
              io.grpc.MethodDescriptor.<com.miempresa.inventoryservice.InventoryServiceProto.UpdateProductRequest, com.miempresa.inventoryservice.InventoryServiceProto.UpdateProductResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "UpdateProduct"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.miempresa.inventoryservice.InventoryServiceProto.UpdateProductRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.miempresa.inventoryservice.InventoryServiceProto.UpdateProductResponse.getDefaultInstance()))
              .setSchemaDescriptor(new InventoryServiceMethodDescriptorSupplier("UpdateProduct"))
              .build();
        }
      }
    }
    return getUpdateProductMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static InventoryServiceStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<InventoryServiceStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<InventoryServiceStub>() {
        @java.lang.Override
        public InventoryServiceStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new InventoryServiceStub(channel, callOptions);
        }
      };
    return InventoryServiceStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static InventoryServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<InventoryServiceBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<InventoryServiceBlockingStub>() {
        @java.lang.Override
        public InventoryServiceBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new InventoryServiceBlockingStub(channel, callOptions);
        }
      };
    return InventoryServiceBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static InventoryServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<InventoryServiceFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<InventoryServiceFutureStub>() {
        @java.lang.Override
        public InventoryServiceFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new InventoryServiceFutureStub(channel, callOptions);
        }
      };
    return InventoryServiceFutureStub.newStub(factory, channel);
  }

  /**
   */
  public static abstract class InventoryServiceImplBase implements io.grpc.BindableService {

    /**
     */
    public void addProduct(com.miempresa.inventoryservice.InventoryServiceProto.AddProductRequest request,
        io.grpc.stub.StreamObserver<com.miempresa.inventoryservice.InventoryServiceProto.AddProductResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getAddProductMethod(), responseObserver);
    }

    /**
     */
    public void getProduct(com.miempresa.inventoryservice.InventoryServiceProto.GetProductRequest request,
        io.grpc.stub.StreamObserver<com.miempresa.inventoryservice.InventoryServiceProto.GetProductResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetProductMethod(), responseObserver);
    }

    /**
     */
    public void updateProduct(com.miempresa.inventoryservice.InventoryServiceProto.UpdateProductRequest request,
        io.grpc.stub.StreamObserver<com.miempresa.inventoryservice.InventoryServiceProto.UpdateProductResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getUpdateProductMethod(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getAddProductMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                com.miempresa.inventoryservice.InventoryServiceProto.AddProductRequest,
                com.miempresa.inventoryservice.InventoryServiceProto.AddProductResponse>(
                  this, METHODID_ADD_PRODUCT)))
          .addMethod(
            getGetProductMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                com.miempresa.inventoryservice.InventoryServiceProto.GetProductRequest,
                com.miempresa.inventoryservice.InventoryServiceProto.GetProductResponse>(
                  this, METHODID_GET_PRODUCT)))
          .addMethod(
            getUpdateProductMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                com.miempresa.inventoryservice.InventoryServiceProto.UpdateProductRequest,
                com.miempresa.inventoryservice.InventoryServiceProto.UpdateProductResponse>(
                  this, METHODID_UPDATE_PRODUCT)))
          .build();
    }
  }

  /**
   */
  public static final class InventoryServiceStub extends io.grpc.stub.AbstractAsyncStub<InventoryServiceStub> {
    private InventoryServiceStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected InventoryServiceStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new InventoryServiceStub(channel, callOptions);
    }

    /**
     */
    public void addProduct(com.miempresa.inventoryservice.InventoryServiceProto.AddProductRequest request,
        io.grpc.stub.StreamObserver<com.miempresa.inventoryservice.InventoryServiceProto.AddProductResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getAddProductMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getProduct(com.miempresa.inventoryservice.InventoryServiceProto.GetProductRequest request,
        io.grpc.stub.StreamObserver<com.miempresa.inventoryservice.InventoryServiceProto.GetProductResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetProductMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void updateProduct(com.miempresa.inventoryservice.InventoryServiceProto.UpdateProductRequest request,
        io.grpc.stub.StreamObserver<com.miempresa.inventoryservice.InventoryServiceProto.UpdateProductResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getUpdateProductMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   */
  public static final class InventoryServiceBlockingStub extends io.grpc.stub.AbstractBlockingStub<InventoryServiceBlockingStub> {
    private InventoryServiceBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected InventoryServiceBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new InventoryServiceBlockingStub(channel, callOptions);
    }

    /**
     */
    public com.miempresa.inventoryservice.InventoryServiceProto.AddProductResponse addProduct(com.miempresa.inventoryservice.InventoryServiceProto.AddProductRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getAddProductMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.miempresa.inventoryservice.InventoryServiceProto.GetProductResponse getProduct(com.miempresa.inventoryservice.InventoryServiceProto.GetProductRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetProductMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.miempresa.inventoryservice.InventoryServiceProto.UpdateProductResponse updateProduct(com.miempresa.inventoryservice.InventoryServiceProto.UpdateProductRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getUpdateProductMethod(), getCallOptions(), request);
    }
  }

  /**
   */
  public static final class InventoryServiceFutureStub extends io.grpc.stub.AbstractFutureStub<InventoryServiceFutureStub> {
    private InventoryServiceFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected InventoryServiceFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new InventoryServiceFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.miempresa.inventoryservice.InventoryServiceProto.AddProductResponse> addProduct(
        com.miempresa.inventoryservice.InventoryServiceProto.AddProductRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getAddProductMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.miempresa.inventoryservice.InventoryServiceProto.GetProductResponse> getProduct(
        com.miempresa.inventoryservice.InventoryServiceProto.GetProductRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetProductMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.miempresa.inventoryservice.InventoryServiceProto.UpdateProductResponse> updateProduct(
        com.miempresa.inventoryservice.InventoryServiceProto.UpdateProductRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getUpdateProductMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_ADD_PRODUCT = 0;
  private static final int METHODID_GET_PRODUCT = 1;
  private static final int METHODID_UPDATE_PRODUCT = 2;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final InventoryServiceImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(InventoryServiceImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_ADD_PRODUCT:
          serviceImpl.addProduct((com.miempresa.inventoryservice.InventoryServiceProto.AddProductRequest) request,
              (io.grpc.stub.StreamObserver<com.miempresa.inventoryservice.InventoryServiceProto.AddProductResponse>) responseObserver);
          break;
        case METHODID_GET_PRODUCT:
          serviceImpl.getProduct((com.miempresa.inventoryservice.InventoryServiceProto.GetProductRequest) request,
              (io.grpc.stub.StreamObserver<com.miempresa.inventoryservice.InventoryServiceProto.GetProductResponse>) responseObserver);
          break;
        case METHODID_UPDATE_PRODUCT:
          serviceImpl.updateProduct((com.miempresa.inventoryservice.InventoryServiceProto.UpdateProductRequest) request,
              (io.grpc.stub.StreamObserver<com.miempresa.inventoryservice.InventoryServiceProto.UpdateProductResponse>) responseObserver);
          break;
        default:
          throw new AssertionError();
      }
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public io.grpc.stub.StreamObserver<Req> invoke(
        io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        default:
          throw new AssertionError();
      }
    }
  }

  private static abstract class InventoryServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    InventoryServiceBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return com.miempresa.inventoryservice.InventoryServiceProto.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("InventoryService");
    }
  }

  private static final class InventoryServiceFileDescriptorSupplier
      extends InventoryServiceBaseDescriptorSupplier {
    InventoryServiceFileDescriptorSupplier() {}
  }

  private static final class InventoryServiceMethodDescriptorSupplier
      extends InventoryServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    InventoryServiceMethodDescriptorSupplier(String methodName) {
      this.methodName = methodName;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.MethodDescriptor getMethodDescriptor() {
      return getServiceDescriptor().findMethodByName(methodName);
    }
  }

  private static volatile io.grpc.ServiceDescriptor serviceDescriptor;

  public static io.grpc.ServiceDescriptor getServiceDescriptor() {
    io.grpc.ServiceDescriptor result = serviceDescriptor;
    if (result == null) {
      synchronized (InventoryServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new InventoryServiceFileDescriptorSupplier())
              .addMethod(getAddProductMethod())
              .addMethod(getGetProductMethod())
              .addMethod(getUpdateProductMethod())
              .build();
        }
      }
    }
    return result;
  }
}
