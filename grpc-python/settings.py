from pydantic import BaseSettings, Field


class BaseConfig(BaseSettings):
    ALLOW_ORIGINS: list = Field(env="ALLOW_ORIGINS", default=['*'])

    DB_HOST: str = Field(env="DB_HOST")
    DB_PORT: int = Field(env="DB_PORT")
    DB_USERNAME: str = Field(env="DB_USERNAME")
    DB_PASSWORD: str = Field(env="DB_PASSWORD")
    DATABASE: str = Field(env="DATABASE")

    class Config:
        env_file = ".env"
        env_file_encoding = 'utf-8'


config = BaseConfig(_env_file=f'./.env')
