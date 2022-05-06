import * as $protobuf from "protobufjs";
/** Namespace blob. */
export namespace blob {

    /** Namespace v1. */
    namespace v1 {

        /** Properties of a CreateBlobRequest. */
        interface ICreateBlobRequest {

            /** CreateBlobRequest accountId */
            accountId?: (string|null);

            /** CreateBlobRequest uploadUrlTimeoutSeconds */
            uploadUrlTimeoutSeconds?: (number|null);
        }

        /** Represents a CreateBlobRequest. */
        class CreateBlobRequest implements ICreateBlobRequest {

            /**
             * Constructs a new CreateBlobRequest.
             * @param [properties] Properties to set
             */
            constructor(properties?: blob.v1.ICreateBlobRequest);

            /** CreateBlobRequest accountId. */
            public accountId: string;

            /** CreateBlobRequest uploadUrlTimeoutSeconds. */
            public uploadUrlTimeoutSeconds: number;

            /**
             * Creates a CreateBlobRequest message from a plain object. Also converts values to their respective internal types.
             * @param object Plain object
             * @returns CreateBlobRequest
             */
            public static fromObject(object: { [k: string]: any }): blob.v1.CreateBlobRequest;

            /**
             * Creates a plain object from a CreateBlobRequest message. Also converts values to other types if specified.
             * @param message CreateBlobRequest
             * @param [options] Conversion options
             * @returns Plain object
             */
            public static toObject(message: blob.v1.CreateBlobRequest, options?: $protobuf.IConversionOptions): { [k: string]: any };

            /**
             * Converts this CreateBlobRequest to JSON.
             * @returns JSON object
             */
            public toJSON(): { [k: string]: any };
        }

        /** Properties of a CreateBlobResponse. */
        interface ICreateBlobResponse {

            /** CreateBlobResponse id */
            id?: (string|null);

            /** CreateBlobResponse uploadUrl */
            uploadUrl?: (string|null);
        }

        /** Represents a CreateBlobResponse. */
        class CreateBlobResponse implements ICreateBlobResponse {

            /**
             * Constructs a new CreateBlobResponse.
             * @param [properties] Properties to set
             */
            constructor(properties?: blob.v1.ICreateBlobResponse);

            /** CreateBlobResponse id. */
            public id: string;

            /** CreateBlobResponse uploadUrl. */
            public uploadUrl: string;

            /**
             * Creates a CreateBlobResponse message from a plain object. Also converts values to their respective internal types.
             * @param object Plain object
             * @returns CreateBlobResponse
             */
            public static fromObject(object: { [k: string]: any }): blob.v1.CreateBlobResponse;

            /**
             * Creates a plain object from a CreateBlobResponse message. Also converts values to other types if specified.
             * @param message CreateBlobResponse
             * @param [options] Conversion options
             * @returns Plain object
             */
            public static toObject(message: blob.v1.CreateBlobResponse, options?: $protobuf.IConversionOptions): { [k: string]: any };

            /**
             * Converts this CreateBlobResponse to JSON.
             * @returns JSON object
             */
            public toJSON(): { [k: string]: any };
        }

        /** Properties of a GetBlobRequest. */
        interface IGetBlobRequest {

            /** GetBlobRequest id */
            id?: (string|null);
        }

        /** Represents a GetBlobRequest. */
        class GetBlobRequest implements IGetBlobRequest {

            /**
             * Constructs a new GetBlobRequest.
             * @param [properties] Properties to set
             */
            constructor(properties?: blob.v1.IGetBlobRequest);

            /** GetBlobRequest id. */
            public id: string;

            /**
             * Creates a GetBlobRequest message from a plain object. Also converts values to their respective internal types.
             * @param object Plain object
             * @returns GetBlobRequest
             */
            public static fromObject(object: { [k: string]: any }): blob.v1.GetBlobRequest;

            /**
             * Creates a plain object from a GetBlobRequest message. Also converts values to other types if specified.
             * @param message GetBlobRequest
             * @param [options] Conversion options
             * @returns Plain object
             */
            public static toObject(message: blob.v1.GetBlobRequest, options?: $protobuf.IConversionOptions): { [k: string]: any };

            /**
             * Converts this GetBlobRequest to JSON.
             * @returns JSON object
             */
            public toJSON(): { [k: string]: any };
        }

        /** Properties of a GetBlobResponse. */
        interface IGetBlobResponse {

            /** GetBlobResponse data */
            data?: (Uint8Array|null);
        }

        /** Represents a GetBlobResponse. */
        class GetBlobResponse implements IGetBlobResponse {

            /**
             * Constructs a new GetBlobResponse.
             * @param [properties] Properties to set
             */
            constructor(properties?: blob.v1.IGetBlobResponse);

            /** GetBlobResponse data. */
            public data: Uint8Array;

            /**
             * Creates a GetBlobResponse message from a plain object. Also converts values to their respective internal types.
             * @param object Plain object
             * @returns GetBlobResponse
             */
            public static fromObject(object: { [k: string]: any }): blob.v1.GetBlobResponse;

            /**
             * Creates a plain object from a GetBlobResponse message. Also converts values to other types if specified.
             * @param message GetBlobResponse
             * @param [options] Conversion options
             * @returns Plain object
             */
            public static toObject(message: blob.v1.GetBlobResponse, options?: $protobuf.IConversionOptions): { [k: string]: any };

            /**
             * Converts this GetBlobResponse to JSON.
             * @returns JSON object
             */
            public toJSON(): { [k: string]: any };
        }

        /** Properties of a GetBlobURLRequest. */
        interface IGetBlobURLRequest {

            /** GetBlobURLRequest id */
            id?: (string|null);

            /** GetBlobURLRequest timeoutSec */
            timeoutSec?: (number|null);
        }

        /** Represents a GetBlobURLRequest. */
        class GetBlobURLRequest implements IGetBlobURLRequest {

            /**
             * Constructs a new GetBlobURLRequest.
             * @param [properties] Properties to set
             */
            constructor(properties?: blob.v1.IGetBlobURLRequest);

            /** GetBlobURLRequest id. */
            public id: string;

            /** GetBlobURLRequest timeoutSec. */
            public timeoutSec: number;

            /**
             * Creates a GetBlobURLRequest message from a plain object. Also converts values to their respective internal types.
             * @param object Plain object
             * @returns GetBlobURLRequest
             */
            public static fromObject(object: { [k: string]: any }): blob.v1.GetBlobURLRequest;

            /**
             * Creates a plain object from a GetBlobURLRequest message. Also converts values to other types if specified.
             * @param message GetBlobURLRequest
             * @param [options] Conversion options
             * @returns Plain object
             */
            public static toObject(message: blob.v1.GetBlobURLRequest, options?: $protobuf.IConversionOptions): { [k: string]: any };

            /**
             * Converts this GetBlobURLRequest to JSON.
             * @returns JSON object
             */
            public toJSON(): { [k: string]: any };
        }

        /** Properties of a GetBlobURLResponse. */
        interface IGetBlobURLResponse {

            /** GetBlobURLResponse url */
            url?: (string|null);
        }

        /** Represents a GetBlobURLResponse. */
        class GetBlobURLResponse implements IGetBlobURLResponse {

            /**
             * Constructs a new GetBlobURLResponse.
             * @param [properties] Properties to set
             */
            constructor(properties?: blob.v1.IGetBlobURLResponse);

            /** GetBlobURLResponse url. */
            public url: string;

            /**
             * Creates a GetBlobURLResponse message from a plain object. Also converts values to their respective internal types.
             * @param object Plain object
             * @returns GetBlobURLResponse
             */
            public static fromObject(object: { [k: string]: any }): blob.v1.GetBlobURLResponse;

            /**
             * Creates a plain object from a GetBlobURLResponse message. Also converts values to other types if specified.
             * @param message GetBlobURLResponse
             * @param [options] Conversion options
             * @returns Plain object
             */
            public static toObject(message: blob.v1.GetBlobURLResponse, options?: $protobuf.IConversionOptions): { [k: string]: any };

            /**
             * Converts this GetBlobURLResponse to JSON.
             * @returns JSON object
             */
            public toJSON(): { [k: string]: any };
        }

        /** Represents a BlobService */
        class BlobService extends $protobuf.rpc.Service {

            /**
             * Constructs a new BlobService service.
             * @param rpcImpl RPC implementation
             * @param [requestDelimited=false] Whether requests are length-delimited
             * @param [responseDelimited=false] Whether responses are length-delimited
             */
            constructor(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean);

            /**
             * Calls CreateBlob.
             * @param request CreateBlobRequest message or plain object
             * @param callback Node-style callback called with the error, if any, and CreateBlobResponse
             */
            public createBlob(request: blob.v1.ICreateBlobRequest, callback: blob.v1.BlobService.CreateBlobCallback): void;

            /**
             * Calls CreateBlob.
             * @param request CreateBlobRequest message or plain object
             * @returns Promise
             */
            public createBlob(request: blob.v1.ICreateBlobRequest): Promise<blob.v1.CreateBlobResponse>;

            /**
             * Calls GetBlob.
             * @param request GetBlobRequest message or plain object
             * @param callback Node-style callback called with the error, if any, and GetBlobResponse
             */
            public getBlob(request: blob.v1.IGetBlobRequest, callback: blob.v1.BlobService.GetBlobCallback): void;

            /**
             * Calls GetBlob.
             * @param request GetBlobRequest message or plain object
             * @returns Promise
             */
            public getBlob(request: blob.v1.IGetBlobRequest): Promise<blob.v1.GetBlobResponse>;

            /**
             * Calls GetBlobURL.
             * @param request GetBlobURLRequest message or plain object
             * @param callback Node-style callback called with the error, if any, and GetBlobURLResponse
             */
            public getBlobURL(request: blob.v1.IGetBlobURLRequest, callback: blob.v1.BlobService.GetBlobURLCallback): void;

            /**
             * Calls GetBlobURL.
             * @param request GetBlobURLRequest message or plain object
             * @returns Promise
             */
            public getBlobURL(request: blob.v1.IGetBlobURLRequest): Promise<blob.v1.GetBlobURLResponse>;
        }

        namespace BlobService {

            /**
             * Callback as used by {@link blob.v1.BlobService#createBlob}.
             * @param error Error, if any
             * @param [response] CreateBlobResponse
             */
            type CreateBlobCallback = (error: (Error|null), response?: blob.v1.CreateBlobResponse) => void;

            /**
             * Callback as used by {@link blob.v1.BlobService#getBlob}.
             * @param error Error, if any
             * @param [response] GetBlobResponse
             */
            type GetBlobCallback = (error: (Error|null), response?: blob.v1.GetBlobResponse) => void;

            /**
             * Callback as used by {@link blob.v1.BlobService#getBlobURL}.
             * @param error Error, if any
             * @param [response] GetBlobURLResponse
             */
            type GetBlobURLCallback = (error: (Error|null), response?: blob.v1.GetBlobURLResponse) => void;
        }
    }
}
