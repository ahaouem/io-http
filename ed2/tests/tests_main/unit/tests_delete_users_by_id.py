import pytest
from unittest.mock import patch
from main import app as flask_app 
from main import OK, NOT_FOUND  

@pytest.fixture
def client():
    with flask_app.test_client() as testing_client:
        with flask_app.app_context():
            yield testing_client

@patch('controllers.UserController.remove_user')
def test_delete_user_success(mock_remove_user, client):
    mock_remove_user.return_value = True
    response = client.delete("/users/1")
    assert response.status_code == OK

@patch('controllers.UserController.remove_user')
def test_delete_user_not_found(mock_remove_user, client):
    mock_remove_user.return_value = False
    response = client.delete("/users/9999")
    assert response.status_code == NOT_FOUND
