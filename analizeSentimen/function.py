import functions_framework
import pickle
from google.cloud import storage
import sklearn

@functions_framework.http
def hello_http(request):
    """HTTP Cloud Function.
    Args:
        request (flask.Request): The request object.
        <https://flask.palletsprojects.com/en/1.1.x/api/#incoming-request-data>
    Returns:
        The response text, or any set of values that can be turned into a
        Response object using `make_response`
        <https://flask.palletsprojects.com/en/1.1.x/api/#flask.make_response>.
    """
    request_json = request.get_json(silent=True)
    request_args = request.args
    if request_json and 'sentiments' in request_json:
        # Download the classifier model from Cloud Storage
        storage_client = storage.Client()
        bucket = storage_client.get_bucket('trensentimen_bucket-1')
        blob_classifier = bucket.blob('modelNB6.pickle')
        blob_classifier.download_to_filename('/tmp/modelNB6.pickle')
        blob_vectorizer = bucket.blob('vectorizer.pickle')
        blob_vectorizer.download_to_filename('/tmp/vectorizer.pickle')
        # Load the classifier model from the downloaded file
        with open('/tmp/modelNB6.pickle', 'rb') as model_file:
            serverless_classifier = pickle.load(model_file)
        with open('/tmp/vectorizer.pickle', 'rb') as file:
            serverless_vectorizer = pickle.load(file)
        # Extract the sentiments from the request payload
        sentiments = request_json['sentiments']
        # Konversi teks menjadi vektor menggunakan CountVectorizer
        new_text_vectorized = serverless_vectorizer.transform(sentiments)
        # Make predictions using the loaded classifier model
        dataSentiments = serverless_classifier.predict(new_text_vectorized)
        response_data = {
            "message": True,
            "data": dataSentiments
        }
        response_data['data'] = response_data['data'].tolist()
    else:
        response_data = {
            "message": False,
            "data": []
        }
    return response_data
