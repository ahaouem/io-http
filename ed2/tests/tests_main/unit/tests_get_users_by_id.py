import pytest
from unittest.mock import patch
from main import app as flask_app  # ensure correct import based on your application structure
from main import NOT_FOUND, OK  # make sure to import these constants from your main application file

@pytest.fixture
def client():
    with flask_app.test_client() as testing_client:
        with flask_app.app_context():
            yield testing_client

# Assuming the controller is in a file named controllers.py and the class is UserController
@patch('controllers.UserController.get_user')
def test_get_user_not_found(mock_get_user, client):
    mock_get_user.return_value = None
    response = client.get("/users/1")
    assert response.status_code == NOT_FOUND

@patch('controllers.UserController.get_user')
def test_get_user_found(mock_get_user, client):
    mock_user = {'id': 1, 'name': 'John Doe'}
    mock_get_user.return_value = mock_user
    response = client.get("/users/1")
    assert response.status_code == OK
    assert response.data == str(mock_user).encode()

