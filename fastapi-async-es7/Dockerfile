FROM python:3.9
WORKDIR /code
COPY ./requirements.txt /code/requirements.txt
COPY start-app.sh /code/start-app.sh
RUN pip install --no-cache-dir --upgrade -r /code/requirements.txt
COPY ./app /code/app
ENTRYPOINT ["sh", "/code/start-app.sh"]
