import grpc
from fastapi import FastAPI
from google.protobuf.json_format import MessageToDict
from sqlalchemy import text

from proto.cosmetic import cosmetic_pb2, cosmetic_pb2_grpc
from database import BaseRepository

def get_grpc_cosmetics():
    with grpc.insecure_channel("go-grpc-server:9000") as channel:
        # stub을 생성해줍니다.
        stub = cosmetic_pb2_grpc.CosmeticServiceStub(channel)
        # 요청을 보내고 결과를 받는데, 서버에서 지정한 메서드에 요청시 사용할 proto 메시지 형식으로 요청을 전송합니다.
        response = stub.ListCosmetics(cosmetic_pb2.ListCosmeticsRequest())

    return MessageToDict(response, including_default_value_fields=True)["data"]


app = FastAPI()


@app.get("/internal/cosmetics")
def get_cosmetics():
    return get_grpc_cosmetics()


@app.get("/")
def get_test():
    db_client1 = BaseRepository.get_connection_db1()
    sql = text('select * from Cosmetic')
    results = db_client1.execute(sql)
    results = [row[0] for row in results]

    db_client2 = BaseRepository.get_connection_db2()
    sql = text('select * from User')
    results2 = db_client1.execute(sql)
    results2 = [row[1] for row in results]

    return results
