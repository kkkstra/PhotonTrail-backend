FROM python:3.12-slim

WORKDIR /app

COPY . /app

ENV PIP_INDEX_URL=https://pypi.tuna.tsinghua.edu.cn/simple
ENV PIP_TRUSTED_HOST=pypi.tuna.tsinghua.edu.cn

RUN pip install --no-cache-dir flask \
    alibabacloud_tea_openapi \
    alibabacloud_sts20150401 \
    alibabacloud_credentials

EXPOSE 8000

CMD ["python", "main.py"]