FROM python:3.12-slim


WORKDIR /app

COPY . /app

RUN pip install uv

RUN uv pip install --system  -r requirements.txt

EXPOSE 8000