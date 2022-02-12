FROM python:3-alpine

# http port exposed outside the container
EXPOSE 5000

# Set the Current Working Directory inside the container
WORKDIR /app

RUN echo "Flask==1.1.2" > requirements.txt
RUN pip install -r requirements.txt

# Copy the source code into the conatainer 
COPY app.py .

# Run the application
CMD ["python", "app.py"]