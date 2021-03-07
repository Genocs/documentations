FROM python:3-alpine
 
WORKDIR /app
 
RUN echo "Flask==1.1.2" > requirements.txt
RUN pip install -r requirements.txt
COPY foo.py .
 
EXPOSE 5000
 
CMD ["python", "foo.py"]