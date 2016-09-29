FROM python:3.5-alpine

MAINTAINER Adrian Moreno <adrian.moreno@dell.com>

# Copy the requirements first to use cache
COPY requirements.txt /usr/src/app
WORKDIR /usr/src/app
RUN pip install -r requirements.txt

COPY . /usr/src/app
RUN python3 manage.py migrate --no-input &&\
    python3 manage.py collectstatic --no-input

EXPOSE 80

CMD ["gunicorn", "hero.wsgi", "--bind", "0.0.0.0:80", "--log-file", "-"]
