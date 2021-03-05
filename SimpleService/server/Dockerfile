FROM python:3.7-alpine

RUN mkdir /app
WORKDIR /app
ADD . /app/
RUN pip install -r requirements.txt

ENV PYTHONUNBUFFERED=1

EXPOSE 12235
CMD ["python3", "server.py"]