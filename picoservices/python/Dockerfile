FROM python:3-alpine
 
WORKDIR /app
 
RUN echo "Flask==1.1.2" > requirements.txt
RUN pip install -r requirements.txt
COPY app.py .
 
EXPOSE 5000
 
CMD ["python", "app.py"]