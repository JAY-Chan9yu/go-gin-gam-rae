from sqlalchemy import create_engine
from sqlalchemy.orm import sessionmaker, scoped_session, declarative_base

from settings import config


Base = declarative_base()


class BaseRepository:
    _client = None
    _client2 = None
    _engine = None

    @classmethod
    def get_connection_db1(cls):
        if cls._client is None:
            engine = create_engine(
                f'mysql+pymysql://{config.DB_USERNAME}:{config.DB_PASSWORD}'
                f'@{config.DB_HOST}:33063/{config.DATABASE}'
            )  #echo="debug" )
            cls._client = scoped_session(sessionmaker(autocommit=False, autoflush=False, bind=engine))
        return cls._client

    @classmethod
    def get_connection_db2(cls):
        if cls._client2 is None:
            engine = create_engine(
                f'mysql+pymysql://{config.DB_USERNAME}:{config.DB_PASSWORD}'
                f'@{config.DB_HOST}:33064/{config.DATABASE}'
            )  #echo="debug" )
            cls._client2 = scoped_session(sessionmaker(autocommit=False, autoflush=False, bind=engine))
        return cls._client2

    @classmethod
    def get_engine(cls):
        """
        비동기로 db connection 을 처리하려면 scoped_session 을 사용하면 안된다.
        참고: https://blog.neonkid.xyz/266
        """
        if cls._engine is None:
            cls._engine = create_engine(
                f'mysql+pymysql://{config.DB_USERNAME}:{config.DB_PASSWORD}'
                f'@{config.DB_HOST}:{config.DB_PORT}/{config.DATABASE}'
            )  #echo="debug" )
            # cls._client = scoped_session(sessionmaker(autocommit=False, autoflush=False, bind=engine))
        return cls._engine
        # return cls._client
