FROM python:3.10

WORKDIR /app

COPY ./requirements.txt /app/requirements.txt

RUN pip install --no-cache-dir -r requirements.txt

COPY . /app

EXPOSE 8060

CMD python manage.py makemigrations && python manage.py migrate && python manage.py runserver 0.0.0.0:8060

